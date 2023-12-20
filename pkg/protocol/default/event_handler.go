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

	activity := orm.TradingActivityModel{
		TokenName: event.Tick,
		Type:      "Sale",
		Amount:    event.Amount.String(),
		UnitPrice: event.Price.String(),
		CreatedAt: event.Timestamp,
		From:      event.Seller.String(),
		To:        event.Taker.String(),
		TxHash:    event.Raw.TxHash.String(),
	}

	dbc.Create(&activity)

	// update taker's balance
	taker := orm.TokenHolderModel{}
	// check if taker exists, if not create one, else update balance
	queryTakerResult := dbc.Where("address = ?", event.Taker.String()).First(&taker)
	if queryTakerResult.Error != nil {
		if errors.Is(queryTakerResult.Error, gorm.ErrRecordNotFound) {
			taker = orm.TokenHolderModel{
				TokenName: event.Tick,
				Address:   event.Taker.String(),
				Balance:   event.Amount.String(),
			}
			dbc.Create(&taker)
		} else {
			log.Error(queryTakerResult.Error)
		}

	} else {
		takerBalance, _, _ := util.NewDecimalFromString(taker.Balance)
		amount, _, _ := util.NewDecimalFromString(event.Amount.String())
		taker.Balance = takerBalance.Add(amount).String()
		dbc.Save(&taker)
	}

	// update token volume, daily volume
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

	// update user balance
	maker := orm.TokenHolderModel{}
	queryMakerResult := dbc.Where("address = ?", order.Maker).First(&maker)
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
