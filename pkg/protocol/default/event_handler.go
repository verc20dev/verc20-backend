package _default

import (
	"errors"
	"ethsyncer/pkg/orm"
	"ethsyncer/types"
	"ethsyncer/util"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func ProcessOrderExecuted(event *types.VERC20MarketVERC20OrderExecuted) {
	dbc := orm.GetDbClient()

	activityType := "Bid Executed"
	if event.Sell {
		activityType = "Ask Executed"
	}

	from := event.Taker.String()
	to := event.Maker.String()
	if event.Sell {
		from = event.Maker.String()
		to = event.Taker.String()
	}

	activity := orm.TradingActivityModel{
		TokenName: event.Tick,
		Type:      activityType,
		Amount:    event.Amount.String(),
		UnitPrice: event.Price.String(),
		CreatedAt: event.Timestamp,
		From:      from,
		To:        to,
		TxHash:    event.Raw.TxHash.String(),
	}

	dbc.Create(&activity)

	tokenRecipient := orm.TokenHolderModel{}

	// if sell, taker is the recipient
	// if buy, maker is the recipient
	tokenRecipientAddress := event.Maker.String()
	if event.Sell {
		tokenRecipientAddress = event.Taker.String()
	}

	queryTakerResult := dbc.Where(
		"address = ? AND token_name = ?",
		tokenRecipientAddress, event.Tick).
		First(&tokenRecipient)
	if queryTakerResult.Error != nil {
		if errors.Is(queryTakerResult.Error, gorm.ErrRecordNotFound) {
			tokenRecipient = orm.TokenHolderModel{
				TokenName: event.Tick,
				Address:   event.Taker.String(),
				Balance:   event.Amount.String(),
			}
			dbc.Create(&tokenRecipient)
		} else {
			log.Error(queryTakerResult.Error)
		}

	} else {
		takerBalance, _, _ := util.NewDecimalFromString(tokenRecipient.Balance)
		amount, _, _ := util.NewDecimalFromString(event.Amount.String())
		tokenRecipient.Balance = takerBalance.Add(amount).String()
		dbc.Save(&tokenRecipient)
	}

	// update token volume
	token := orm.TokenInfoModel{}
	queryTokenResult := dbc.Where("name = ?", event.Tick).First(&token)
	if queryTokenResult.Error != nil {
		log.Error(queryTokenResult.Error)
	} else {
		totalVolume, _, _ := util.NewDecimalFromString(token.TotalVolume)

		amount, _, _ := util.NewDecimalFromString(event.Amount.String())
		price, _, _ := util.NewDecimalFromString(event.Price.String())

		orderVolume := amount.Mul(price)

		if totalVolume == nil {
			token.TotalVolume = orderVolume.String()
		} else {
			token.TotalVolume = totalVolume.Add(orderVolume).String()
		}

		dbc.Save(&token)
	}

	// update order status
	order := orm.OrderModel{}
	txHash := util.ByteArrayToHex(event.ListId[:])
	if util.HexIsUppercase(txHash) {
		txHash = util.HexToLowercase(txHash)
	}
	queryOrderResult := dbc.Where("LOWER(tx) = ?", txHash).First(&order)
	if queryOrderResult.Error != nil {
		log.Error(queryOrderResult.Error)
	} else {
		order.Executed = true
		order.ExecutedAt = event.Timestamp
		dbc.Save(&order)
	}

	floorOrder := orm.OrderModel{}
	queryFloorOrderResult := dbc.Where(
		"token_name = ? AND canceled != true AND executed != true", event.Tick).
		Order("unit_price asc").First(&floorOrder)
	if queryFloorOrderResult.Error != nil {
		// did not find any floor order so set floor to latest order price
		log.Error(queryFloorOrderResult.Error)
		token.FloorPrice = event.Price.String()
		dbc.Save(&token)
	} else {
		token.FloorPrice = floorOrder.UnitPrice
		dbc.Save(&token)
	}

	// insert transaction
	transaction := orm.TxModel{
		Hash:      event.Raw.TxHash.String(),
		From:      from,
		To:        to,
		BlockNum:  event.Raw.BlockNumber,
		Index:     uint64(event.Raw.Index),
		Timestamp: event.Timestamp,
		Input:     "",
	}

	dbc.Create(&transaction)

	// update history
	history := orm.HistoryModel{
		TokenName: event.Tick,
		TxId:      transaction.ID,
		Type:      "transfer",
		Amount:    event.Amount.String(),
	}

	dbc.Create(&history)
}

func ProcessOrderCanceled(event *types.VERC20MarketVERC20OrderCanceled) {
	dbc := orm.GetDbClient()

	order := orm.OrderModel{}
	txHash := util.ByteArrayToHex(event.ListId[:])
	txHash = util.HexToLowercase(txHash)

	queryOrderResult := dbc.Where("LOWER(tx) = ?", txHash).First(&order)
	if queryOrderResult.Error != nil {
		log.Error(queryOrderResult.Error)
	} else {
		order.Canceled = true
		order.CanceledAt = event.Timestamp
		dbc.Save(&order)
	}

	activity := orm.TradingActivityModel{
		TokenName: order.TokenName,
		Type:      "Cancel",
		Amount:    order.Amount,
		UnitPrice: order.UnitPrice,
		CreatedAt: event.Timestamp,
		From:      order.Maker,
		TxHash:    event.Raw.TxHash.String(),
	}

	dbc.Create(&activity)

	// when sell order cancel, return token to maker
	// when buy order cancel, eth already returned to maker, no need to do anything
	if event.Sell {
		// update user balance
		maker := orm.TokenHolderModel{}
		queryMakerResult := dbc.Where("address = ? AND token_name = ?",
			order.Maker, order.TokenName).First(&maker)
		if queryMakerResult.Error != nil {
			log.Error(queryMakerResult.Error)
			// if maker does not exist, create one
			if errors.Is(queryMakerResult.Error, gorm.ErrRecordNotFound) {
				maker = orm.TokenHolderModel{
					TokenName: order.TokenName,
					Address:   order.Maker,
					Balance:   order.Amount,
				}
				dbc.Create(&maker)
			}
		} else {
			makerBalance, _, _ := util.NewDecimalFromString(maker.Balance)
			amount, _, _ := util.NewDecimalFromString(order.Amount)
			maker.Balance = makerBalance.Add(amount).String()
			dbc.Save(&maker)
		}
	}

	token := orm.TokenInfoModel{}
	queryTokenResult := dbc.Where("name = ?", order.TokenName).First(&token)
	if queryTokenResult.Error != nil {
		log.Error(queryTokenResult.Error)
		return
	}

	// update floor price
	floorOrder := orm.OrderModel{}
	queryFloorOrderResult := dbc.Where(
		"token_name = ? AND canceled != true AND executed != true", order.TokenName).
		Order("unit_price asc").First(&floorOrder)
	if queryFloorOrderResult.Error != nil {
		// did not find any floor order so set floor to latest executed order price
		log.Error(queryFloorOrderResult.Error)
		executedOrder := orm.OrderModel{}
		queryExecutedOrderResult := dbc.Where(
			"token_name = ? AND executed = true", order.TokenName).
			Order("executed_at desc").First(&executedOrder)
		if queryExecutedOrderResult.Error != nil {
			// did not find any executed order so set floor to 0
			log.Error(queryExecutedOrderResult.Error)
			token.FloorPrice = ""
			dbc.Save(&token)
		} else {
			token.FloorPrice = executedOrder.UnitPrice
			dbc.Save(&token)
		}
	} else {
		token.FloorPrice = floorOrder.UnitPrice
		dbc.Save(&token)
	}

}
