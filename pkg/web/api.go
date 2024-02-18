package web

import (
	"errors"
	"ethsyncer/pkg/orm"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"math"
	"strconv"
)

type HolderAPI struct {
}

type TokenAPI struct {
}

type OrderAPI struct {
}

// GetStatus Get /status
func (api *TokenAPI) GetStatus(c *gin.Context) {
	db := orm.GetDbClient()
	var status orm.StatusModel
	err := db.
		First(&status).
		Error
	if err != nil {
		log.Error(err)
		c.JSON(500, gin.H{"status": "ERROR"})
		return
	}
	c.JSON(200, Status{
		LatestImportedBlockNumber: status.ImportedBlockNum,
		LatestSyncedBlockNumber:   status.SyncedBlockNum,
	})
}

// GetHolder Get /holders/:address
func (api *HolderAPI) GetHolder(c *gin.Context) {
	address := c.Param("address")
	if address == "" {
		c.JSON(400, gin.H{"status": "ERROR"})
		return
	}

	db := orm.GetDbClient()
	var holderModels []orm.TokenHolderModel
	tokenName, ok := c.GetQuery("tick")
	var err error
	if ok {
		err = db.
			Where("address = ?", address).
			Where("token_name = ?", tokenName).
			Find(&holderModels).
			Error
	} else {
		err = db.
			Where("address = ?", address).
			Find(&holderModels).
			Error
	}

	if err != nil {
		log.Error(err)
		c.JSON(500, gin.H{"status": "ERROR"})
		return
	}

	var tokenHoldings []TokenHolding
	for _, holderModel := range holderModels {
		tokenHoldings = append(
			tokenHoldings,
			TokenHolding{
				Name:    holderModel.TokenName,
				Balance: holderModel.Balance,
			})
	}

	c.JSON(200, HolderDetail{
		Tokens: tokenHoldings,
	})
}

// GetToken Get /tokens/:name
// get token
func (api *TokenAPI) GetToken(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(400, gin.H{"status": "ERROR"})
		return
	}

	db := orm.GetDbClient()
	var tokenModel orm.TokenInfoModel
	err := db.
		Where("name = ?", name).
		First(&tokenModel).
		Error
	if err != nil {
		log.Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{"status": "Token " + name + " not found"})
			return
		}
		c.JSON(500, gin.H{"status": "ERROR"})
		return
	}

	if tokenModel == (orm.TokenInfoModel{}) {
		c.JSON(404, gin.H{"status": "ERROR"})
		return
	}

	var holderCount int64
	err = db.
		Model(&orm.TokenHolderModel{}).
		Where("token_name = ?", name).
		Count(&holderCount).
		Error
	if err != nil {
		c.JSON(500, gin.H{"status": "ERROR"})
		return
	}

	c.JSON(200, tokenModelToDetail(tokenModel, holderCount))
}

// ListToken Get /tokens
// list token
func (api *TokenAPI) ListToken(c *gin.Context) {
	// get pagination params
	defaultSort := "created_at"
	defaultOrder := "asc"
	params := getPaginationParams(c)
	if params.sort == "" {
		params.sort = defaultSort
	}
	if params.order == "" {
		params.order = defaultOrder
	}

	weightOrderQuery := "sort_weight desc nulls last,"
	orderQuery := weightOrderQuery + " " + params.sort + " " + params.order + ", created_at asc"

	db := orm.GetDbClient()
	var tokenModels []orm.TokenInfoModel

	var totalRow int64
	err := db.
		Model(&orm.TokenInfoModel{}).
		Count(&totalRow).
		Error
	if err != nil {
		log.Error(err)
		c.JSON(500, gin.H{"status": "ERROR"})
		return
	}

	err = db.
		Offset(params.offset).
		Limit(params.limit).
		Order(orderQuery).
		Find(&tokenModels).
		Error
	if err != nil {
		log.Error(err)
		c.JSON(500, gin.H{"status": "ERROR"})
		return
	}

	holderCountMap := make(map[string]int64)
	type Result struct {
		Name    string
		Holders int64
	}
	var res []Result
	err = db.
		Model(&orm.TokenHolderModel{}).
		Select("token_name as name, count(distinct address) as holders").
		Group("token_name").
		Find(&res).
		Error
	if err != nil {
		log.Error(err)
		c.JSON(500, gin.H{"status": "ERROR"})
		return
	}
	for _, r := range res {
		holderCountMap[r.Name] = r.Holders
	}

	var tokenDetails []TokenDetail
	for _, tokenModel := range tokenModels {
		tokenDetails = append(
			tokenDetails,
			tokenModelToDetail(tokenModel, holderCountMap[tokenModel.Name]))
	}

	c.JSON(200, ListTokenResponse{
		Total: int64(math.Ceil(float64(totalRow) / float64(params.limit))),
		Data:  tokenDetails,
	})
}

// ListTokenHistories Get /tokens/:name/histories
// list token histories
func (api *TokenAPI) ListTokenHistories(c *gin.Context) {
	defaultSort := "time"
	defaultOrder := "desc"
	params := getPaginationParams(c)
	if params.sort == "" {
		params.sort = defaultSort
	}
	if params.order == "" {
		params.order = defaultOrder
	}

	orderQuery := params.sort + " " + params.order
	if params.sort == "time" {
		orderQuery = "tx_id " + params.order
	}

	name := c.Param("name")
	if name == "" {
		c.JSON(400, gin.H{"status": "ERROR"})
		return
	}

	db := orm.GetDbClient()
	var tokenModel orm.TokenInfoModel
	err := db.
		Where("name = ?", name).
		First(&tokenModel).
		Error
	if err != nil {
		log.Error(err)
		c.JSON(404, gin.H{"status": "token " + name + " not found"})
		return
	}

	var totalRow int64
	err = db.
		Model(&orm.HistoryModel{}).
		Where("token_name = ?", name).
		Count(&totalRow).
		Error
	if err != nil {
		log.Error(err)
		c.JSON(500, gin.H{"status": "ERROR"})
		return
	}

	var histories []TokenHistoryInfo

	err = db.
		Limit(params.limit).
		Offset(params.offset).
		Order(orderQuery).
		Model(&orm.HistoryModel{}).
		Select("histories.\"type\" as method, t.\"from\", t.\"to\", block_num, amount as quantity, timestamp as created_at, hash as creation_tx").
		Joins("left join transactions t on t.id = histories.tx_id").
		Where("token_name = ?", name).
		Scan(&histories).
		Error

	if err != nil {
		log.Error(err)
		c.JSON(500, gin.H{"status": "ERROR"})
		return
	}

	c.JSON(200, ListTokenHistoryResponse{
		Total: int64(math.Ceil(float64(totalRow) / float64(params.limit))),
		Data:  histories,
	})
}

// ListTokenHolders Get /tokens/:name/holders
// list token holders
func (api *TokenAPI) ListTokenHolders(c *gin.Context) {
	defaultSort := "balance"
	defaultOrder := "desc"
	params := getPaginationParams(c)
	if params.sort == "" {
		params.sort = defaultSort
	}
	if params.order == "" {
		params.order = defaultOrder
	}

	orderQuery := params.sort + " " + params.order
	if params.sort == "balance" {
		orderQuery = "cast(balance as decimal) " + params.order
	}

	name := c.Param("name")
	if name == "" {
		c.JSON(400, gin.H{"status": "ERROR"})
		return
	}

	db := orm.GetDbClient()
	var tokenModel orm.TokenInfoModel
	err := db.
		Where("name = ?", name).
		First(&tokenModel).
		Error
	if err != nil {
		log.Error(err)
		c.JSON(404, gin.H{"status": "token " + name + " not found"})
		return
	}

	var totalRow int64
	err = db.
		Model(&orm.TokenHolderModel{}).
		Where("token_name = ?", name).
		Count(&totalRow).
		Error
	if err != nil {
		log.Error(err)
		c.JSON(500, gin.H{"status": "ERROR"})
		return
	}

	var tokenHolderModels []orm.TokenHolderModel
	err = db.
		Where("token_name = ?", name).
		Offset(params.offset).
		Limit(params.limit).
		Order(orderQuery).
		Find(&tokenHolderModels).
		Error
	if err != nil {
		log.Error(err)
		c.JSON(500, gin.H{"status": "ERROR"})
		return
	}

	var tokenHolderInfos []TokenHolderInfo
	for idx, tokenHolderModel := range tokenHolderModels {
		tokenHolderInfos = append(
			tokenHolderInfos,
			tokenHolderToTokenHolderInfo(tokenHolderModel, idx))
	}

	c.JSON(200, ListTokenHolderResponse{
		Total: int64(math.Ceil(float64(totalRow) / float64(params.limit))),
		Data:  tokenHolderInfos,
	})
}


// ListHolderHistories Get /holders/:address/histories
func (api *HolderAPI) ListHolderHistories(c *gin.Context) {
	defaultSort := "time"
	defaultOrder := "desc"
	params := getPaginationParams(c)
	if params.sort == "" {
		params.sort = defaultSort
	}
	if params.order == "" {
		params.order = defaultOrder
	}

	orderQuery := params.sort + " " + params.order
	if params.sort == "time" {
		orderQuery = "tx_id " + params.order
	}

	address := c.Param("address")
	if address == "" {
		c.JSON(400, gin.H{"status": "ERROR"})
		return
	}

	db := orm.GetDbClient()
	var totalRow int64
	err := db.
		Model(&orm.HistoryModel{}).
		Joins("left join transactions t on t.id = histories.tx_id").
		Where("t.\"from\" = ? OR t.\"to\" = ?", address, address).
		Count(&totalRow).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(200, ListTokenHistoryResponse{
				Total: 0,
				Data:  nil,
			})
			return
		}
		c.JSON(500, gin.H{"status": "ERROR"})
		return
	}

	var histories []TokenHistoryInfo
	err = db.
		Limit(params.limit).
		Offset(params.offset).
		Order(orderQuery).
		Model(&orm.HistoryModel{}).
		Select("histories.\"type\" as method, t.\"from\", t.\"to\", block_num, amount as quantity, timestamp as created_at, hash as creation_tx, token_name as name").
		Joins("left join transactions t on t.id = histories.tx_id").
		Where("t.\"from\" = ? OR t.\"to\" = ?", address, address).
		Scan(&histories).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(200, ListTokenHistoryResponse{
				Total: 0,
				Data:  nil,
			})
			return
		}
		log.Error(err)
		c.JSON(500, gin.H{"status": "ERROR"})
		return
	}

	c.JSON(200, ListTokenHistoryResponse{
		Total: int64(math.Ceil(float64(totalRow) / float64(params.limit))),
		Data:  histories,
	})

}

type PaginationParams struct {
	limit  int
	offset int
	sort   string
	order  string
}

func getPaginationParams(c *gin.Context) PaginationParams {
	var params PaginationParams
	params.limit = 10
	params.offset = 0
	params.sort = ""
	params.order = ""

	if limit, ok := c.GetQuery("limit"); ok {
		l, err := strconv.Atoi(limit)
		if err == nil {
			params.limit = l
		}
	}
	if offset, ok := c.GetQuery("offset"); ok {
		o, err := strconv.Atoi(offset)
		if err == nil {
			params.offset = o
		}
	}
	if sort, ok := c.GetQuery("sort"); ok {
		params.sort = sort
	}
	if order, ok := c.GetQuery("order"); ok {
		params.order = order
	}

	return params
}

func tokenModelToDetail(tokenModel orm.TokenInfoModel, holders int64) TokenDetail {
	return TokenDetail{
		Name:         tokenModel.Name,
		TotalSupply:  tokenModel.TotalSupply,
		Decimals:     tokenModel.Decimals,
		Minted:       tokenModel.TotalMinted,
		Holders:      holders,
		Transactions: tokenModel.TotalTxs,
		CreationTx:   tokenModel.CreationTx,
		CreatedAt:    tokenModel.CreatedAt,
		CreatedBy:    tokenModel.CreatedBy,
		MintedOutAt:  tokenModel.MintedOutAt,
		Limit:        tokenModel.Limit,
		StartBlock:   tokenModel.StartBlock,
		Type:         tokenModel.Type,
		Duration:     tokenModel.Duration,
		IsVerified:   tokenModel.IsVerified,
		IsOfficial:   tokenModel.IsOfficial,
	}
}

func tokenHolderToTokenHolderInfo(tokenHolder orm.TokenHolderModel, idx int) TokenHolderInfo {
	return TokenHolderInfo{
		Rank:    int32(idx),
		Address: tokenHolder.Address,
		Balance: tokenHolder.Balance,
	}
}
