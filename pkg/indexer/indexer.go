package indexer

import (
	"ethsyncer/pkg/context"
	"ethsyncer/pkg/orm"
	"ethsyncer/pkg/protocol"
	"ethsyncer/util"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"math/big"
	"time"
)

type Indexer struct {
	DbClient  *gorm.DB
	FromBlock uint64
	ToBlock   uint64
	Context   context.IndexerContext
}

type Manager struct {
	Indexer *Indexer
}

func NewIndexer(
	dbClient *gorm.DB,
	fromBlock uint64,
	toBlock uint64,
) *Indexer {
	return &Indexer{
		DbClient:  dbClient,
		FromBlock: fromBlock,
		ToBlock:   toBlock,
	}
}

func NewManager(dbc *gorm.DB) *Manager {
	latestImportedBlockNum, err := orm.GetLatestImportedBlockNum(dbc)
	if err != nil {
		log.Fatal(err)
	}
	latestSyncedBlockNum, err := orm.GetLatestSyncedBlockNum(dbc)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("latest imported block num: ", latestImportedBlockNum)
	log.Info("latest synced block num: ", latestSyncedBlockNum)

	return &Manager{
		Indexer: NewIndexer(dbc, latestImportedBlockNum+1, latestSyncedBlockNum),
	}
}

func (manager *Manager) IndexOnce() error {
	return manager.Indexer.Index()
}

func (manager *Manager) Start(interval int) {
	for {
		latestImportedBlockNum, err := orm.GetLatestImportedBlockNum(manager.Indexer.DbClient)
		if err != nil {
			log.Fatal(err)
		}
		latestSyncedBlockNum, err := orm.GetLatestSyncedBlockNum(manager.Indexer.DbClient)
		if err != nil {
			log.Fatal(err)
		}

		manager.Indexer.FromBlock = latestImportedBlockNum + 1
		manager.Indexer.ToBlock = latestSyncedBlockNum

		if manager.Indexer.FromBlock > manager.Indexer.ToBlock {
			log.Info("Already indexed, latest block: ", latestImportedBlockNum)
			log.Info("waiting ", interval, "s for new blocks...")
			time.Sleep(time.Duration(interval) * time.Second)
			continue
		}

		err = manager.IndexOnce()
		if err != nil {
			log.Error(err)
		}

	}
}

func (indexer *Indexer) Index() error {

	if indexer.FromBlock > indexer.ToBlock {
		log.Info("already imported to latest synced block num: ", indexer.ToBlock)
		return nil
	}

	log.Info("Indexing from block ", indexer.FromBlock, " to block ", indexer.ToBlock)
	for j := indexer.FromBlock; j <= indexer.ToBlock; j++ {
		startTime := time.Now()
		log.Info("Indexing block ", j)
		txs, err := orm.GetTxsInBlock(indexer.DbClient, j)
		if err != nil {
			return err
		}
		err = indexer.initIndexerContext()
		if err != nil {
			return err
		}

		for _, tx := range txs {
			indexer.indexTx(tx, &indexer.Context)
		}

		err = indexer.dump(j)
		if err != nil {
			return err
		}
		log.Info("block ", j, " successfully imported. cost: ",
			time.Since(startTime).Milliseconds(), "ms")
	}

	return nil
}

func (indexer *Indexer) indexTx(tx orm.TxModel, ctx *context.IndexerContext) {
	pType := protocol.GetProtocolType(tx.Input, ctx)
	h := protocol.GetHandler(pType)
	if h == nil {
		return
	}
	h.Process(tx, ctx)
}

func (indexer *Indexer) initIndexerContext() error {
	tokenInfos, err := orm.GetAllTokenInfo(indexer.DbClient)
	if err != nil {
		return err
	}
	tokenInfoMap := util.GroupingBy(tokenInfos, func(t orm.TokenInfoModel) string {
		return t.Name
	})

	indexer.Context = context.IndexerContext{
		DbClient: indexer.DbClient,
		TokenInfos: util.MapValues(tokenInfoMap, func(t []orm.TokenInfoModel) orm.TokenInfoModel {
			if len(t) > 0 {
				return t[0]
			} else {
				return orm.TokenInfoModel{}
			}
		}),
		TokenHolderBalanceDelta: make(map[string]map[string]*big.Int),
		Histories:               make([]orm.HistoryModel, 0),
	}
	return nil
}

func (indexer *Indexer) dump(blockNum uint64) error {
	ZERO := big.NewInt(0)
	tokenNames := util.Keys(indexer.Context.TokenInfos)
	tokenBalances, err := orm.GetTokenHolderBalanceByTokenName(indexer.DbClient, tokenNames)
	if err != nil {
		return err
	}

	tokenBalancesMap := formTokenBalanceMap(tokenBalances)
	var newTokenBalances []orm.TokenHolderModel
	var deleteTokenBalanceIds []uint64
	var updateTokenBalances []orm.TokenHolderModel
	for address, v := range indexer.Context.TokenHolderBalanceDelta {
		for name, delta := range v {
			if _, ok := tokenBalancesMap[address]; !ok {
				tokenBalancesMap[address] = make(map[string]*big.Int)
			}
			if _, ok := tokenBalancesMap[address][name]; !ok {
				tokenBalancesMap[address][name] = big.NewInt(0)
			}

			savedBalance := util.NewDecimalFromBigInt(tokenBalancesMap[address][name])
			deltaD := util.NewDecimalFromBigInt(delta)
			result := savedBalance.Add(deltaD)

			if savedBalance.Sign() == 0 && result.Cmp(&util.Decimal{Value: ZERO}) > 0 {
				newTokenBalances = append(newTokenBalances, orm.TokenHolderModel{
					TokenName: name,
					Address:   address,
					Balance:   result.String(),
				})
			} else if savedBalance.Sign() == 0 && result.Cmp(&util.Decimal{Value: ZERO}) <= 0 {
				continue
			} else if savedBalance.Sign() > 0 && result.Cmp(&util.Decimal{Value: ZERO}) <= 0 {
				tokenBalance, exist := util.FindFirst(tokenBalances, func(m orm.TokenHolderModel) bool {
					return m.TokenName == name && m.Address == address
				})
				if !exist {
					continue
				}
				deleteTokenBalanceIds = append(deleteTokenBalanceIds, tokenBalance.ID)
			} else if savedBalance.Sign() > 0 && result.Cmp(&util.Decimal{Value: ZERO}) > 0 {
				tokenBalance, exist := util.FindFirst(tokenBalances, func(m orm.TokenHolderModel) bool {
					return m.TokenName == name && m.Address == address
				})
				if !exist {
					continue
				}
				tokenBalance.Balance = result.String()
				updateTokenBalances = append(updateTokenBalances, tokenBalance)
			}
		}
	}

	err = indexer.DbClient.Transaction(func(dbTx *gorm.DB) error {
		tokenInfos := util.Values(indexer.Context.TokenInfos)
		if tokenInfos != nil && len(tokenInfos) > 0 {
			err := dbTx.Clauses(clause.OnConflict{
				Columns: []clause.Column{{Name: "name"}},
				DoUpdates: clause.AssignmentColumns([]string{
					"total_minted",
					"total_txs",
					"minted_out_at",
				}),
			}).Create(tokenInfos).Error
			if err != nil {
				return err
			}
		}

		if len(newTokenBalances) > 0 || len(updateTokenBalances) > 0 {
			// merge newTokenBalances and updateTokenBalances
			upsertTokenBalances := make([]orm.TokenHolderModel, 0)
			upsertTokenBalances = append(upsertTokenBalances, newTokenBalances...)
			upsertTokenBalances = append(upsertTokenBalances, updateTokenBalances...)

			err = dbTx.Clauses(clause.OnConflict{
				Columns: []clause.Column{
					{Name: "token_name"},
					{Name: "address"},
				},
				DoUpdates: clause.AssignmentColumns([]string{
					"balance",
				}),
			}).Create(upsertTokenBalances).Error
			if err != nil {
				return err
			}
		}

		if len(deleteTokenBalanceIds) > 0 {
			err = dbTx.Delete(&orm.TokenHolderModel{}, deleteTokenBalanceIds).Error
			if err != nil {
				return err
			}
		}

		if len(indexer.Context.Histories) > 0 {
			err = dbTx.Create(indexer.Context.Histories).Error
			if err != nil {
				return err
			}
		}

		err = dbTx.
			Model(&orm.StatusModel{}).
			Where("id = ?", 1).
			Update("imported_block_num", blockNum).
			Error
		if err != nil {
			return err
		}

		return nil
	})

	return err
}

func formTokenBalanceMap(tokenBalances []orm.TokenHolderModel) map[string]map[string]*big.Int {
	// from list<{name, address, balance}> to map<address, map<name, balance>>
	tokenBalancesByAddress := util.GroupingBy(tokenBalances, func(t orm.TokenHolderModel) string {
		return t.Address
	})
	tokenBalancesMap := make(map[string]map[string][]orm.TokenHolderModel)
	for k, v := range tokenBalancesByAddress {
		tokenBalancesName := make(map[string][]orm.TokenHolderModel)
		tokenBalancesName = util.GroupingBy(v, func(t orm.TokenHolderModel) string {
			return t.TokenName
		})
		tokenBalancesMap[k] = tokenBalancesName
	}

	return util.MapValues(tokenBalancesMap, func(t map[string][]orm.TokenHolderModel) map[string]*big.Int {
		return util.MapValues(t, func(t []orm.TokenHolderModel) *big.Int {
			if len(t) == 0 {
				return big.NewInt(0)
			} else {
				balance, _, err := util.NewDecimalFromString(t[0].Balance)
				if err != nil {
					return big.NewInt(0)
				}
				return balance.Value
			}
		})
	})
}
