package syncer

import (
	"errors"
	"ethsyncer/pkg/orm"
	"ethsyncer/pkg/web3"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

type Syncer struct {
	Web3Client *ethclient.Client
	DbClient   *gorm.DB
	FromBlock  uint64
	ToBlock    uint64
	StartTxIdx uint64
	DryRun     bool
	Condition  func(tx *types.Transaction) bool
}

func NewSyncer(
	web3Client *ethclient.Client,
	dbClient *gorm.DB,
	fromBlock uint64,
	toBlock uint64,
	startTxIdx uint64,
	dryRun bool,
	condition func(tx *types.Transaction) bool,
) *Syncer {
	return &Syncer{
		Web3Client: web3Client,
		DbClient:   dbClient,
		FromBlock:  fromBlock,
		ToBlock:    toBlock,
		StartTxIdx: startTxIdx,
		DryRun:     dryRun,
		Condition:  condition,
	}
}

func (s *Syncer) Sync() error {
	if s.FromBlock > s.ToBlock {
		log.Error("Invalid block range. from: ",
			s.FromBlock, "'s index ", s.StartTxIdx, ", to: ", s.ToBlock)
		return errors.New("invalid block range")
	}
	if s.DryRun {
		log.Info("Syncer is running in dry-run mode")
	}

	log.Info("Syncing from block ",
		s.FromBlock, "'s index ", s.StartTxIdx, " to block ", s.ToBlock)
	for i := s.FromBlock; i <= s.ToBlock; i++ {
		log.Info("Syncing block ", i)
		startTime := time.Now()

		txs, err := web3.GetTxsInBlock(s.Web3Client, i)
		if err != nil {
			return err
		}
		var txModels []orm.TxModel
		for idx, tx := range txs {
			log.Debug("processing tx ", tx.Hash().Hex(), " at index ", idx, " in block ", i)
			if uint64(idx) < s.StartTxIdx {
				log.Debug("Skip tx ", tx.Hash().Hex(), " at index ", idx)
				continue
			}

			if s.Condition(&tx) {
				log.Debug("Syncing tx ", tx.Hash().Hex(), " at index ", idx)
				txFrom, err := web3.GetTxFrom(&tx)
				if err != nil {
					return err
				}
				decodedData, err := web3.CborToJsonString(tx.Data())
				if err != nil {
					return err
				}
				txModel := orm.TxModel{
					Hash:      tx.Hash().Hex(),
					From:      txFrom,
					To:        tx.To().String(),
					BlockNum:  i,
					Index:     uint64(idx),
					Timestamp: uint64(tx.Time().Unix()),
					Input:     decodedData,
				}
				txModels = append(txModels, txModel)
			} else {
				log.Debug("Condition not match, skipping tx ", tx.Hash().Hex())
			}
		}
		if s.DryRun {
			log.Info("Dry-run mode, skipping saving txs to db")
			continue
		}

		err = s.DbClient.Transaction(func(dbTx *gorm.DB) error {
			if txModels != nil && len(txModels) > 0 {
				err = dbTx.Create(&txModels).Error
				if err != nil {
					return err
				}
			}

			err = dbTx.
				Model(&orm.StatusModel{}).
				Where("id = ?", 1).
				Update("synced_block_num", i).
				Error
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}

		log.Info("block ", i, " successfully synced. cost ",
			time.Since(startTime).Milliseconds(), " ms")
		log.Info("total valid txs in block ", i, ": ", len(txModels))
	}
	log.Info("Successfully synced blocks from ",
		s.FromBlock, " to ", s.ToBlock, " with start tx index ", s.StartTxIdx)
	return nil
}

type SyncManager struct {
	syncer *Syncer
}

func NewSyncManager(wc *ethclient.Client, dbc *gorm.DB) *SyncManager {
	latestBlockNum, err := web3.GetLatestBlock(wc)
	if err != nil {
		log.Fatal(err)
	}

	latestSyncedBlockNum, err := orm.GetLatestSyncedBlockNum(dbc)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("latest block: ", latestBlockNum)
	log.Info("synced block: ", latestSyncedBlockNum)
	return &SyncManager{
		syncer: NewSyncer(
			wc,
			dbc,
			latestSyncedBlockNum+1,
			latestBlockNum,
			0,
			false,
			web3.SniffIfCbor,
		),
	}
}

func (sm *SyncManager) SyncOnce() error {
	return sm.syncer.Sync()
}

func (sm *SyncManager) Start(interval int) {
	for {
		latestBlockNum, err := web3.GetLatestBlock(sm.syncer.Web3Client)
		if err != nil {
			log.Error(err)
			continue
		}

		latestSyncedBlockNum, err := orm.GetLatestSyncedBlockNum(sm.syncer.DbClient)
		if err != nil {
			log.Error(err)
			continue
		}

		if latestBlockNum == latestSyncedBlockNum {
			log.Info("Already synced, latest block: ", latestBlockNum)
			log.Info("waiting ", interval, "s for new blocks...")
			time.Sleep(time.Duration(interval) * time.Second)
			continue
		}

		sm.syncer.FromBlock = latestSyncedBlockNum + 1
		sm.syncer.ToBlock = latestBlockNum

		err = sm.SyncOnce()
		if err != nil {
			log.Error(err)
		}
	}
}
