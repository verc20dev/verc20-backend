package context

import (
	"ethsyncer/pkg/orm"
	"gorm.io/gorm"
	"math/big"
)

type IndexerContext struct {
	DbClient                *gorm.DB
	TokenInfos              map[string]orm.TokenInfoModel  // token_name -> info
	TokenHolderBalanceDelta map[string]map[string]*big.Int // address -> token_name -> balance
	Histories               []orm.HistoryModel
	TxInputUnmarshalled     map[string]interface{}
}
