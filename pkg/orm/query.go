package orm

import "gorm.io/gorm"

func GetLatestSyncedBlockNum(db *gorm.DB) (uint64, error) {
	var status StatusModel
	err := db.First(&status).Error
	if err != nil {
		return 0, err
	}
	return status.SyncedBlockNum, nil
}

func GetLastTxIdxInBlock(db *gorm.DB, blockNum uint64) (uint64, error) {
	var tx TxModel
	err := db.
		Where("block_num = ?", blockNum).
		Order("index desc").
		First(&tx).
		Error
	if err != nil {
		return 0, err
	}
	return tx.Index, nil
}

func GetTxsInBlock(db *gorm.DB, blockNum uint64) ([]TxModel, error) {
	var txs []TxModel
	err := db.
		Where("block_num = ?", blockNum).
		Order("index asc").
		Find(&txs).
		Error
	if err != nil {
		return nil, err
	}
	return txs, nil
}

func GetAllTokenInfo(db *gorm.DB) ([]TokenInfoModel, error) {
	var tokenInfos []TokenInfoModel
	err := db.Find(&tokenInfos).Error
	if err != nil {
		return nil, err
	}
	return tokenInfos, nil
}

func GetTokenHolderBalanceByTokenName(
	db *gorm.DB,
	tokenNames []string,
) ([]TokenHolderModel, error) {
	var tokenHolders []TokenHolderModel
	err := db.Where("token_name in ?", tokenNames).
		Find(&tokenHolders).
		Error
	if err != nil {
		return nil, err
	}
	return tokenHolders, nil
}

func GetLatestImportedBlockNum(db *gorm.DB) (uint64, error) {
	var status StatusModel
	err := db.First(&status).Error
	if err != nil {
		return 0, err
	}
	return status.ImportedBlockNum, nil
}

func GetLatestFetchedBlockNum(db *gorm.DB) (uint64, error) {
	var status StatusModel
	err := db.First(&status).Error
	if err != nil {
		return 0, err
	}
	return status.EventFetchedBlockNum, nil
}
