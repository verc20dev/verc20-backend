package web

import (
	"errors"
	"ethsyncer/pkg/orm"
	protocolDefault "ethsyncer/pkg/protocol/default"
	"ethsyncer/pkg/web3"
	"ethsyncer/types"
	"ethsyncer/util"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"math"
	"math/big"
	"time"
)

type MarketAPI struct {
}

// ListOrders Get /orders
// get list of token orders
func (api *MarketAPI) ListOrders(c *gin.Context) {
	// get pagination params
	defaultSort := "unit_price"
	defaultOrder := "asc"
	params := getPaginationParams(c)
	if params.sort == "" {
		params.sort = defaultSort
	}
	if params.order == "" {
		params.order = defaultOrder
	}

	orderQuery := params.sort + " " + params.order

	nameFilter, ok := c.GetQuery("tick")
	if !ok || nameFilter == "" {
		log.Error("tick is required, req: ", c.Request)
		c.JSON(400, gin.H{"status": "Request is invalid"})
		return
	}

	typeFilter, ok := c.GetQuery("type")
	if !ok || typeFilter == "" {
		log.Error("type is required, req: ", c.Request)
		c.JSON(400, gin.H{"status": "Request is invalid"})
		return
	}

	if typeFilter != "ask" && typeFilter != "bid" {
		log.Error("type is invalid, req: ", c.Request)
		c.JSON(400, gin.H{"status": "Request is invalid"})
		return
	}

	sellFilter := typeFilter == "ask"

	ownerFilter, haveOwnerFilter := c.GetQuery("owner")

	db := orm.GetDbClient()
	var orderModels []orm.OrderModel

	var totalRow int64
	var err error
	if haveOwnerFilter {
		err = db.
			Model(&orm.OrderModel{}).
			Where(
				"token_name = ? AND maker = ? AND sell = ? AND executed != true AND canceled != true",
				nameFilter, ownerFilter, sellFilter,
			).
			Count(&totalRow).
			Error
	} else {
		err = db.
			Model(&orm.OrderModel{}).
			Where("token_name = ? AND sell = ? AND expiration_time > ? AND executed != true AND canceled != true",
				nameFilter, sellFilter, time.Now().Unix(),
			).
			Count(&totalRow).
			Error
	}

	if err != nil {
		log.Error(err)
		c.JSON(500, gin.H{"status": "ERROR"})
		return
	}

	if haveOwnerFilter {
		err = db.
			Where("token_name = ? AND maker = ?  AND sell = ? AND executed != true AND canceled != true",
				nameFilter, ownerFilter, sellFilter).
			Offset(params.offset).
			Limit(params.limit).
			Order(orderQuery).
			Find(&orderModels).
			Error
	} else {
		err = db.
			Where(
				"token_name = ? AND sell = ? AND expiration_time > ? AND executed != true AND canceled != true",
				nameFilter, sellFilter, time.Now().Unix(),
			).
			Offset(params.offset).
			Limit(params.limit).
			Order(orderQuery).
			Find(&orderModels).
			Error
	}

	if err != nil {
		log.Error(err)
		c.JSON(500, gin.H{"status": "ERROR"})
		return
	}

	var orderDetails []OrderDetail
	for _, orderModel := range orderModels {
		orderDetails = append(
			orderDetails,
			OrderDetail{
				ID:             orderModel.ID,
				Tick:           orderModel.TokenName,
				Quantity:       orderModel.Amount,
				UnitPrice:      orderModel.UnitPrice,
				Owner:          orderModel.Maker,
				Status:         getOrderStatus(orderModel),
				CreatedAt:      orderModel.CreationTime,
				ExpirationTime: orderModel.ExpirationTime,
				Taker:          orderModel.Taker,
				Input:          orderModel.Input,
				Signature:      orderModel.Signature,
				Sell:           orderModel.Sell,
			})
	}

	c.JSON(200, ListOrdersResponse{
		Total: int64(math.Ceil(float64(totalRow) / float64(params.limit))),
		Data:  orderDetails,
	})
}

// GetOrderDetail Get /orders/:id
// get order detail
func (api *MarketAPI) GetOrderDetail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"status": "ERROR"})
		return
	}

	db := orm.GetDbClient()
	var orderModel orm.OrderModel
	err := db.
		Where("id = ?", id).
		First(&orderModel).
		Error
	if err != nil {
		log.Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{"status": "Order " + id + " not found"})
			return
		}
		c.JSON(500, gin.H{"status": "ERROR"})
		return
	}

	c.JSON(200, OrderDetail{
		ID:             orderModel.ID,
		Tick:           orderModel.TokenName,
		Quantity:       orderModel.Amount,
		UnitPrice:      orderModel.UnitPrice,
		Owner:          orderModel.Maker,
		Status:         getOrderStatus(orderModel),
		CreatedAt:      orderModel.CreationTime,
		ExpirationTime: orderModel.ExpirationTime,
		Taker:          orderModel.Taker,
		Input:          orderModel.Input,
		Signature:      orderModel.Signature,
		Sell:           orderModel.Sell,
	})
}

// CreateOrder Post /orders
// create new order
func (api *MarketAPI) CreateOrder(c *gin.Context) {
	var createOrderRequest CreateOrderRequest
	err := c.ShouldBindJSON(&createOrderRequest)
	if err != nil {
		log.Error(err)
		c.JSON(400, gin.H{"status": "Request is invalid"})
		return
	}

	// check if token exists
	db := orm.GetDbClient()
	var tokenModel orm.TokenInfoModel
	err = db.
		Where("name = ?", createOrderRequest.Tick).
		First(&tokenModel).
		Error
	if err != nil {
		log.Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(400, gin.H{"status": "Token " + createOrderRequest.Tick + " not found"})
			return
		}
		c.JSON(500, gin.H{"status": "ERROR"})
		return
	}

	// TODO: check signature to prove that the order is created by the owner

	if *createOrderRequest.Sell {
		// check owner balance
		var ownerBalanceModel orm.TokenHolderModel
		err = db.
			Where("token_name = ? and address = ?", createOrderRequest.Tick, createOrderRequest.Owner).
			First(&ownerBalanceModel).
			Error
		if err != nil {
			log.Error(err)
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(400, gin.H{"status": "Owner " + createOrderRequest.Owner + " does not have " + createOrderRequest.Tick})
				return
			}
			c.JSON(500, gin.H{"status": "ERROR"})
			return
		}

		ownerBalance, _, err := util.NewDecimalFromString(ownerBalanceModel.Balance)
		if err != nil {
			log.Error(err)
			c.JSON(500, gin.H{"status": "ERROR"})
			return
		}
		orderQuantity, precision, err := util.NewDecimalFromString(createOrderRequest.Quantity)
		if err != nil {
			log.Error(err)
			c.JSON(400, gin.H{"status": "Quantity is invalid"})
			return
		}

		if precision > tokenModel.Decimals {
			log.Error("Quantity Precision is invalid, req: ", c.Request)
			c.JSON(400, gin.H{"status": "Quantity is invalid"})
			return
		}

		if ownerBalance.Cmp(orderQuantity) == -1 {
			log.Error("Owner does not have enough balance, req: ", c.Request)
			c.JSON(400, gin.H{"status": "Owner does not have enough balance"})
			return
		}

		// update owner balance
		ownerBalance = ownerBalance.Sub(orderQuantity)
		// if ownerBalance is 0, delete the record
		if ownerBalance.Cmp(util.NewDecimalFromBigInt(big.NewInt(0))) == 0 {
			err = db.
				Delete(&ownerBalanceModel).
				Error
			if err != nil {
				log.Error(err)
				c.JSON(500, gin.H{"status": "ERROR"})
				return
			}
		} else {
			ownerBalanceModel.Balance = ownerBalance.String()
			err = db.
				Model(&ownerBalanceModel).
				Where("token_name = ? and address = ?", createOrderRequest.Tick, createOrderRequest.Owner).
				Update("balance", ownerBalance.String()).
				Error
			if err != nil {
				log.Error(err)
				c.JSON(500, gin.H{"status": "ERROR"})
				return
			}
		}
	} else {
		// check Tx value matches the order
		// get tx receipt
		web3Client, err := web3.GetWeb3Client(viper.GetString("rpcUrl"))
		if err != nil {
			log.Error(err)
			c.JSON(500, gin.H{"status": "ERROR"})
			return
		}
		txReceipt, err := web3.GetTxReceipt(web3Client, createOrderRequest.Tx)
		if err != nil {
			log.Error(err)
			c.JSON(400, gin.H{"status": "tx invalid"})
			return
		}

		if txReceipt.Status == 0 {
			c.JSON(400, gin.H{"status": "tx failed"})
			return
		}

		tx, err := web3.GetTx(web3Client, createOrderRequest.Tx)
		if err != nil {
			log.Error(err)
			c.JSON(400, gin.H{"status": "tx invalid"})
			return
		}

		totalValue, _, err := util.NewDecimalFromString(tx.Value().String())
		if err != nil {
			log.Error(err)
			c.JSON(400, gin.H{"status": "tx value is invalid"})
			return
		}
		orderAmount, _, err := util.NewDecimalFromString(createOrderRequest.Quantity)
		if err != nil {
			log.Error(err)
			c.JSON(400, gin.H{"status": "Quantity is invalid"})
			return
		}
		orderUnitPrice, _, err := util.NewDecimalFromString(createOrderRequest.UnitPrice)
		if err != nil {
			log.Error(err)
			c.JSON(400, gin.H{"status": "UnitPrice is invalid"})
			return
		}
		orderValue := orderAmount.Mul(orderUnitPrice)

		if totalValue.Cmp(orderValue) == -1 {
			log.Error("Tx value does not match the order, req: ", c.Request)
			c.JSON(400, gin.H{"status": "tx value does not match the order"})
			return
		}
	}

	// save order to db
	orderModel := orm.OrderModel{
		TokenName:      createOrderRequest.Tick,
		Maker:          createOrderRequest.Owner,
		Amount:         createOrderRequest.Quantity,
		UnitPrice:      createOrderRequest.UnitPrice,
		Tx:             createOrderRequest.Tx,
		CreationTime:   createOrderRequest.CreationTime,
		ExpirationTime: createOrderRequest.ExpirationTime,
		Signature:      createOrderRequest.Signature,
		Input:          createOrderRequest.Input,
		Sell:           *createOrderRequest.Sell,
	}

	err = db.
		Create(&orderModel).
		Error
	if err != nil {
		log.Error(err)
		c.JSON(500, gin.H{"status": "ERROR"})
		return
	}

	// update token floor price
	currentFloorPrice, _, err := util.NewDecimalFromString(tokenModel.FloorPrice)
	orderPrice, _, err := util.NewDecimalFromString(orderModel.UnitPrice)

	if currentFloorPrice == nil || orderPrice.Cmp(currentFloorPrice) == -1 {
		tokenModel.FloorPrice = orderModel.UnitPrice
		err = db.Save(&tokenModel).Error
		if err != nil {
			log.Error(err)
			c.JSON(500, gin.H{"status": "ERROR"})
			return
		}
	}

	activityType := "Bid"
	if *createOrderRequest.Sell {
		activityType = "Ask"
	}

	// create trading activity record
	tradingActivity := orm.TradingActivityModel{
		TokenName: orderModel.TokenName,
		OrderId:   orderModel.ID,
		Type:      activityType,
		Amount:    orderModel.Amount,
		UnitPrice: orderModel.UnitPrice,
		CreatedAt: orderModel.CreationTime,
		From:      orderModel.Maker,
		To:        "",
		TxHash:    createOrderRequest.Tx,
	}
	err = db.
		Create(&tradingActivity).
		Error
	if err != nil {
		log.Error(err)
		c.JSON(500, gin.H{"status": "ERROR"})
		return
	}

	c.Status(201)
}

// ExecuteOrder Post /orders/:id/execute
// execute order
func (api *MarketAPI) ExecuteOrder(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"status": "ERROR"})
		return
	}

	var executeOrderRequest ExecuteOrderRequest
	err := c.ShouldBindJSON(&executeOrderRequest)
	if err != nil {
		log.Error(err)
		c.JSON(400, gin.H{"status": "Request is invalid"})
		return
	}

	tx := executeOrderRequest.Tx
	if tx == "" {
		c.JSON(400, gin.H{"status": "tx is required"})
		return
	}

	// get tx receipt
	web3Client, err := web3.GetWeb3Client(viper.GetString("rpcUrl"))
	if err != nil {
		log.Error(err)
		c.JSON(500, gin.H{"status": "ERROR"})
		return
	}
	txReceipt, err := web3.GetTxReceipt(web3Client, tx)
	if err != nil {
		log.Error(err)
		c.JSON(400, gin.H{"status": "tx invalid"})
		return
	}

	if txReceipt.Status == 0 {
		c.JSON(400, gin.H{"status": "tx failed"})
		return
	}

	contractAddress := ethCommon.HexToAddress(viper.GetString("marketAddress"))
	contract, err := types.NewVERC20Market(contractAddress, web3Client)
	if err != nil {
		log.Error(err)
		c.JSON(500, gin.H{"status": "ERROR"})
		return
	}

	var orderExecutedEvent *types.VERC20MarketVERC20OrderExecuted
	for _, txLog := range txReceipt.Logs {
		event, err := contract.ParseVERC20OrderExecuted(*txLog)
		if err != nil {
			continue
		}
		orderExecutedEvent = event
		break
	}

	if orderExecutedEvent == nil {
		log.Error("OrderExecuted event not found in tx receipt logs")
		c.JSON(400, gin.H{"status": "tx invalid"})
		return
	}

	// check if order exists

	db := orm.GetDbClient()
	var orderModel orm.OrderModel
	err = db.
		Where("id = ?", id).
		First(&orderModel).
		Error
	if err != nil {
		log.Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{"status": "Order " + id + " not found"})
			return
		}
		c.JSON(500, gin.H{"status": "ERROR"})
		return
	}

	if orderModel.Executed {
		c.JSON(400, gin.H{"status": "Order " + id + " is already executed"})
		return
	}

	// check if event match order
	if orderModel.TokenName != orderExecutedEvent.Tick {
		log.Error("Order ", id, " and event mismatch")
		c.JSON(400, gin.H{"status": "tx invalid"})
		return
	}

	if util.HexToLowercase(orderModel.Tx) !=
		util.HexToLowercase(util.ByteArrayToHex(orderExecutedEvent.ListId[:])) {
		log.Error("Order ", id, " and event listId mismatch")
		c.JSON(400, gin.H{"status": "tx invalid"})
		return
	}

	// update order
	protocolDefault.ProcessOrderExecuted(orderExecutedEvent)
	c.Status(200)
}

// FreezeOrder Post /orders/:id/freeze
// freeze order
func (api *MarketAPI) FreezeOrder(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"status": "ERROR"})
		return
	}

	var freezeOrderRequest FreezeOrderRequest
	err := c.ShouldBindJSON(&freezeOrderRequest)
	if err != nil {
		log.Error(err)
		c.JSON(400, gin.H{"status": "Request is invalid"})
		return
	}

	// TODO: check signature to prove that address really sent the request

	// check if order exists
	db := orm.GetDbClient()
	var orderModel orm.OrderModel
	err = db.
		Where("id = ?", id).
		First(&orderModel).
		Error
	if err != nil {
		log.Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{"status": "Order " + id + " not found"})
			return
		}
		c.JSON(500, gin.H{"status": "ERROR"})
		return
	}

	if orderModel.Sell {
		c.JSON(400, gin.H{"status": "Order " + id + " is a ask order"})
		return
	}

	if orderModel.Executed {
		c.JSON(400, gin.H{"status": "Order " + id + " is already executed"})
		return
	}

	if orderModel.Canceled {
		c.JSON(400, gin.H{"status": "Order " + id + " is already canceled"})
		return
	}

	// freeze user balance
	var userBalanceModel orm.TokenHolderModel
	err = db.
		Where("token_name = ? and address = ?", orderModel.TokenName, freezeOrderRequest.Address).
		First(&userBalanceModel).
		Error
	if err != nil {
		log.Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(400, gin.H{"status": "User " + freezeOrderRequest.Address + " does not have " + orderModel.TokenName})
			return
		}
		c.JSON(500, gin.H{"status": "ERROR"})
		return
	}

	userBalance, _, err := util.NewDecimalFromString(userBalanceModel.Balance)
	if err != nil {
		log.Error(err)
		c.JSON(500, gin.H{"status": "ERROR"})
		return
	}
	orderAmount, _, err := util.NewDecimalFromString(orderModel.Amount)
	if err != nil {
		log.Error(err)
		c.JSON(500, gin.H{"status": "ERROR"})
		return
	}

	if userBalance.Cmp(orderAmount) == -1 {
		log.Error("User does not have enough balance, req: ", c.Request)
		c.JSON(400, gin.H{"status": "Insufficient balance to take the bid order"})
		return
	}

	userBalance = userBalance.Sub(orderAmount)
	// if userBalance is 0, delete the record
	if userBalance.Cmp(util.NewDecimalFromBigInt(big.NewInt(0))) == 0 {
		err = db.
			Delete(&userBalanceModel).
			Error
		if err != nil {
			log.Error(err)
			c.JSON(500, gin.H{"status": "ERROR"})
			return
		}
	} else {
		userBalanceModel.Balance = userBalance.String()
		err = db.
			Model(&userBalanceModel).
			Where("token_name = ? and address = ?", orderModel.TokenName, freezeOrderRequest.Address).
			Update("balance", userBalance.String()).
			Error
		if err != nil {
			log.Error(err)
			c.JSON(500, gin.H{"status": "ERROR"})
			return
		}
	}

	c.Status(200)
}

// CancelOrder Post /orders/:id/cancel
// cancel order
func (api *MarketAPI) CancelOrder(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"status": "ERROR"})
		return
	}

	var cancelOrderRequest CancelOrderRequest
	err := c.ShouldBindJSON(&cancelOrderRequest)
	if err != nil {
		log.Error(err)
		c.JSON(400, gin.H{"status": "Request is invalid"})
		return
	}

	db := orm.GetDbClient()
	var orderModel orm.OrderModel
	err = db.
		Where("id = ?", id).
		First(&orderModel).
		Error
	if err != nil {
		log.Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{"status": "Order " + id + " not found"})
			return
		}
		c.JSON(500, gin.H{"status": "ERROR"})
		return
	}

	if orderModel.Canceled {
		c.JSON(400, gin.H{"status": "Order " + id + " is already canceled"})
		return
	}

	tx := cancelOrderRequest.Tx
	if tx == "" {
		c.JSON(400, gin.H{"status": "tx is required"})
		return
	}

	// get tx receipt
	web3Client, err := web3.GetWeb3Client(viper.GetString("rpcUrl"))
	if err != nil {
		log.Error(err)
		c.JSON(500, gin.H{"status": "ERROR"})
		return
	}
	txReceipt, err := web3.GetTxReceipt(web3Client, tx)
	if err != nil {
		log.Error(err)
		c.JSON(400, gin.H{"status": "tx invalid"})
		return
	}

	if txReceipt.Status == 0 {
		c.JSON(400, gin.H{"status": "tx failed"})
		return
	}

	contractAddress := ethCommon.HexToAddress(viper.GetString("marketAddress"))
	contract, err := types.NewVERC20Market(contractAddress, web3Client)
	if err != nil {
		log.Error(err)
		c.JSON(500, gin.H{"status": "ERROR"})
		return
	}

	var orderCanceledEvent *types.VERC20MarketVERC20OrderCanceled
	for _, txLog := range txReceipt.Logs {
		event, err := contract.ParseVERC20OrderCanceled(*txLog)
		if err != nil {
			continue
		}
		orderCanceledEvent = event
		break
	}

	if orderCanceledEvent == nil {
		log.Error("OrderCanceled event not found in tx receipt logs")
		c.JSON(400, gin.H{"status": "tx invalid"})
		return
	}

	if util.HexToLowercase(orderModel.Tx) !=
		util.HexToLowercase(util.ByteArrayToHex(orderCanceledEvent.ListId[:])) {
		log.Error("Order ", id, " and event listId mismatch")
		c.JSON(400, gin.H{"status": "tx invalid"})
	}

	// update order
	protocolDefault.ProcessOrderCanceled(orderCanceledEvent)

	c.Status(204)
}

// ListTradingActivities Get /market/activities
// list trading activities
func (api *MarketAPI) ListTradingActivities(c *gin.Context) {
	tokenName, ok := c.GetQuery("tick")
	if !ok || tokenName == "" {
		c.JSON(400, gin.H{"status": "ERROR"})
		return
	}

	// get pagination params
	defaultSort := "created_at"
	defaultOrder := "desc"
	params := getPaginationParams(c)
	if params.sort == "" {
		params.sort = defaultSort
	}
	if params.order == "" {
		params.order = defaultOrder
	}

	orderQuery := params.sort + " " + params.order

	// check if token exists
	db := orm.GetDbClient()
	var tokenModel orm.TokenInfoModel
	err := db.
		Where("name = ?", tokenName).
		First(&tokenModel).
		Error
	if err != nil {
		log.Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(400, gin.H{"status": "Token " + tokenName + " not found"})
			return
		}
		c.JSON(500, gin.H{"status": "ERROR"})
		return
	}

	var totalRow int64
	err = db.
		Model(&orm.TradingActivityModel{}).
		Where("token_name = ?", tokenName).
		Count(&totalRow).
		Error
	if err != nil {
		log.Error(err)
		c.JSON(500, gin.H{"status": "ERROR"})
		return
	}

	var tradingActivityModels []orm.TradingActivityModel
	err = db.
		Where("token_name = ?", tokenName).
		Offset(params.offset).
		Limit(params.limit).
		Order(orderQuery).
		Find(&tradingActivityModels).
		Error
	if err != nil {
		log.Error(err)
		c.JSON(500, gin.H{"status": "ERROR"})
		return
	}

	var tradingActivities []TradingActivity
	for _, tradingActivityModel := range tradingActivityModels {
		tradingActivities = append(
			tradingActivities,
			TradingActivity{
				ID:        tradingActivityModel.ID,
				Tick:      tradingActivityModel.TokenName,
				Type:      tradingActivityModel.Type,
				Quantity:  tradingActivityModel.Amount,
				UnitPrice: tradingActivityModel.UnitPrice,
				CreatedAt: tradingActivityModel.CreatedAt,
				From:      tradingActivityModel.From,
				To:        tradingActivityModel.To,
				Tx:        tradingActivityModel.TxHash,
			})
	}

	c.JSON(200, ListTradingActivitiesResponse{
		Total: int64(math.Ceil(float64(totalRow) / float64(params.limit))),
		Data:  tradingActivities,
	})

}

// ListMarketTokens Get /market/tokens
// list market tokens
func (api *MarketAPI) ListMarketTokens(c *gin.Context) {
	// get pagination params
	defaultSort := "total_volume"
	defaultOrder := "desc"
	params := getPaginationParams(c)
	if params.sort == "" {
		params.sort = defaultSort
	}
	if params.order == "" {
		params.order = defaultOrder
	}

	if trendingFilter, ok := c.GetQuery("trending"); ok && trendingFilter == "true" {
		params.sort = "daily_volume"
		params.order = "desc"
	}

	nameFilter := ""
	if tick, ok := c.GetQuery("tick"); ok && tick != "" {
		nameFilter = tick
	}

	weightOrderQuery := "sort_weight desc nulls last,"
	orderQuery := weightOrderQuery + " " + params.sort + " " + params.order + ", created_at asc"

	db := orm.GetDbClient()
	var tokenModels []orm.TokenInfoModel

	var totalRow int64
	if nameFilter != "" {
		err := db.
			Model(&orm.TokenInfoModel{}).
			Where("name like ?", "%"+nameFilter+"%").
			Count(&totalRow).
			Error
		if err != nil {
			log.Error(err)
			c.JSON(500, gin.H{"status": "ERROR"})
			return
		}
	} else {
		err := db.
			Model(&orm.TokenInfoModel{}).
			Count(&totalRow).
			Error
		if err != nil {
			log.Error(err)
			c.JSON(500, gin.H{"status": "ERROR"})
			return
		}
	}

	if nameFilter != "" {
		result := db.
			Where("name like ?", "%"+nameFilter+"%").
			Offset(params.offset).
			Limit(params.limit).
			Order(orderQuery).
			Find(&tokenModels)

		err := result.Error
		if err != nil {
			log.Error(err)
			c.JSON(500, gin.H{"status": "ERROR"})
			return
		}
	} else {
		result := db.
			Offset(params.offset).
			Limit(params.limit).
			Order(orderQuery).
			Find(&tokenModels)

		err := result.Error
		if err != nil {
			log.Error(err)
			c.JSON(500, gin.H{"status": "ERROR"})
			return
		}
	}

	tokenNames := util.Map(tokenModels, func(tokenModel orm.TokenInfoModel) string {
		return tokenModel.Name
	})

	holderCountMap := make(map[string]int64)
	type Result struct {
		Name    string
		Holders int64
	}
	var res []Result
	err := db.
		Model(&orm.TokenHolderModel{}).
		Select("token_name as name, count(distinct address) as holders").
		Where("token_name in ?", tokenNames).
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

	type DailyVolumeResult struct {
		TokenName   string
		DailyVolume string
	}
	var dailyVolumeRes []DailyVolumeResult
	err = db.
		Model(&orm.OrderModel{}).
		Select("token_name, sum(cast(orders.unit_price as decimal) * cast(amount as decimal)) as daily_volume").
		Where(
			"executed = true AND executed_at between ? and ? AND token_name IN ?",
			time.Now().AddDate(0, 0, -1).Unix(),
			time.Now().Unix(), tokenNames).
		Group("token_name").
		Find(&dailyVolumeRes).
		Error
	if err != nil {
		log.Error(err)
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(500, gin.H{"status": "ERROR"})
			return
		}
	}
	var dailyVolumeMap = make(map[string]string)
	for _, r := range dailyVolumeRes {
		dailyVolumeMap[r.TokenName] = r.DailyVolume
	}

	var marketTokenDetails []MarketTokenDetail
	for _, tokenModel := range tokenModels {
		marketTokenDetails = append(
			marketTokenDetails,
			MarketTokenDetail{
				Tick:        tokenModel.Name,
				TotalSupply: tokenModel.TotalSupply,
				Holders:     holderCountMap[tokenModel.Name],
				FloorPrice:  tokenModel.FloorPrice,
				TotalVolume: tokenModel.TotalVolume,
				DailyVolume: dailyVolumeMap[tokenModel.Name],
				IsOfficial:  tokenModel.IsOfficial,
				IsVerified:  tokenModel.IsVerified,
			})
	}

	c.JSON(200, ListMarketTokensResponse{
		Total: int64(math.Ceil(float64(totalRow) / float64(params.limit))),
		Data:  marketTokenDetails,
	})

}

// GetMarketTokenDetail Get /market/tokens/:name
// get market token detail
func (api *MarketAPI) GetMarketTokenDetail(c *gin.Context) {
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

	type Result struct {
		Name    string
		Holders int64
	}
	var res []Result
	err = db.
		Model(&orm.TokenHolderModel{}).
		Select("token_name as name, count(distinct address) as holders").
		Where("token_name = ?", name).
		Group("token_name").
		Find(&res).
		Error
	if err != nil {
		log.Error(err)
		c.JSON(500, gin.H{"status": "ERROR"})
		return
	}

	holderCount := int64(0)
	if len(res) > 0 {
		holderCount = res[0].Holders
	}

	type DailyVolumeResult struct {
		TokenName   string
		DailyVolume string
	}
	var dailyVolumeRes []DailyVolumeResult
	err = db.
		Model(&orm.OrderModel{}).
		Select("token_name, sum(cast(orders.unit_price as decimal) * cast(amount as decimal)) as daily_volume").
		Where(
			"executed = true AND executed_at between ? and ? AND token_name = ?",
			time.Now().AddDate(0, 0, -1).Unix(),
			time.Now().Unix(), name).
		Group("token_name").
		Find(&dailyVolumeRes).
		Error
	if err != nil {
		log.Error(err)
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(500, gin.H{"status": "ERROR"})
			return
		}
	}
	var dailyVolumeMap = make(map[string]string)
	for _, r := range dailyVolumeRes {
		dailyVolumeMap[r.TokenName] = r.DailyVolume
	}

	c.JSON(200, MarketTokenDetail{
		Tick:        tokenModel.Name,
		TotalSupply: tokenModel.TotalSupply,
		Holders:     holderCount,
		FloorPrice:  tokenModel.FloorPrice,
		TotalVolume: tokenModel.TotalVolume,
		DailyVolume: dailyVolumeMap[tokenModel.Name],
		IsOfficial:  tokenModel.IsOfficial,
		IsVerified:  tokenModel.IsVerified,
	})
}

func getOrderStatus(orderModel orm.OrderModel) string {
	if orderModel.Taker != "" {
		return "filled"
	}
	if orderModel.Canceled {
		return "canceled"
	}
	now := uint64(time.Now().Unix())
	endTime := orderModel.ExpirationTime
	if now > endTime {
		return "expired"
	}

	return "active"
}
