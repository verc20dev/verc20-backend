package _default

import (
	"ethsyncer/pkg/context"
	"ethsyncer/pkg/orm"
	"ethsyncer/pkg/protocol/common"
	"ethsyncer/util"
	log "github.com/sirupsen/logrus"
	"math/big"
	"strconv"
)

type Handler struct {
}

func (h *Handler) Process(tx orm.TxModel, ctx *context.IndexerContext) {
	op := common.GetProtocolOp(tx.Input, ctx)
	if op == common.Deploy {
		h.processDeploy(tx, ctx)
	} else if op == common.Mint {
		h.processMint(tx, ctx)
	} else if op == common.Transfer {
		h.processTransfer(tx, ctx)
	} else if op == common.List {
		h.processList(tx, ctx)
	}
}

func (h *Handler) processDeploy(tx orm.TxModel, ctx *context.IndexerContext) {
	data, err := FormDeployData(ctx)
	// invalid tx, ignore
	if err != nil {
		return
	}

	_, ok := ctx.TokenInfos[data.tick]
	// already exist, ignore
	if ok {
		return
	}

	decimal := common.MAX_DECIMAL
	if data.dec != "" {
		decimal, err = strconv.Atoi(data.dec)
		// invalid decimal, ignore
		if err != nil {
			return
		}
	}
	// invalid decimal, ignore
	if decimal > common.MAX_DECIMAL {
		return
	}

	tokenInfo := orm.TokenInfoModel{
		Name:        data.tick,
		Decimals:    decimal,
		TotalMinted: "0",
		TotalTxs:    1,
		StartBlock:  tx.BlockNum,
		CreatedBy:   tx.From,
		CreatedAt:   tx.Timestamp,
		CreationTx:  tx.Hash,
		MintedOutAt: 0,
	}

	if data.startBlock != "" {
		startBlock, err := strconv.Atoi(data.startBlock)
		if err == nil {
			if uint64(startBlock) > tx.BlockNum {
				tokenInfo.StartBlock = uint64(startBlock)
			}
		}
	}

	if data.duration != "" {
		duration, err := strconv.Atoi(data.duration)
		if err == nil {
			if uint64(duration) > 0 && uint64(duration) <= common.MAX_DURATION {
				tokenInfo.Duration = uint64(duration)
			}
		}
	}

	// check type
	if data.t == "fair" {
		// fair mode, totalSupply starts from 0
		tokenInfo.TotalSupply = "0"
		tokenInfo.Type = "fair"

		maxSupply, precision, err := util.NewDecimalFromString(data.max)
		// fair mode, max is optional, only process when it's valid
		if err == nil && maxSupply.Sign() > 0 && !maxSupply.IsOverflowUint64() && precision == decimal {
			tokenInfo.TotalSupply = data.max
		}


		if data.lim == "" {
			// invalid limit, ignore
			return
		}
		limit, precision, err := util.NewDecimalFromString(data.lim)
		// invalid limit, ignore
		if err != nil {
			return
		}
		// invalid limit, ignore
		if limit.Sign() < 0 || limit.IsOverflowUint64() || precision > decimal {
			return
		}
		tokenInfo.Limit = data.lim

		// check if duration is already set in tokenInfo
		if tokenInfo.Duration == 0 {
			log.Error("Duration should be set in fair mode. tx: ", tx.ID, " data: ", data)
			return
		}

	} else {
		tokenInfo.Type = "normal"

		maxSupply, precision, err := util.NewDecimalFromString(data.max)
		// invalid max, ignore
		if err != nil {
			return
		}

		// invalid maxSupply, ignore
		if maxSupply.Sign() < 0 || maxSupply.IsOverflowUint64() || precision > decimal {
			return
		}

		tokenInfo.TotalSupply = data.max

		if data.lim != "" {
			// invalid limit, ignore
			limit, precision, err := util.NewDecimalFromString(data.lim)
			if err != nil {
				return
			}

			// invalid limit, ignore
			if limit.Sign() < 0 || limit.IsOverflowUint64() || precision > decimal {
				return
			}

			tokenInfo.Limit = data.lim
		}
	}

	ctx.TokenInfos[data.tick] = tokenInfo
	ctx.Histories = append(ctx.Histories, orm.HistoryModel{
		TokenName: tokenInfo.Name,
		TxId:      tx.ID,
		Type:      "deploy",
	})
}

func (h *Handler) processMint(tx orm.TxModel, ctx *context.IndexerContext) {
	data, err := FormMintData(ctx)
	// invalid tx, ignore
	if err != nil {
		return
	}

	tokenInfo, ok := ctx.TokenInfos[data.tick]
	// mint a not existed token, ignore
	if !ok {
		return
	}

	// already minted out, ignore
	if tokenInfo.MintedOutAt != 0 {
		return
	}

	// exceed duration, ignore
	if tokenInfo.Duration != 0 {
		if tx.BlockNum > tokenInfo.StartBlock+tokenInfo.Duration {
			return
		}
	}

	if tokenInfo.StartBlock != 0 {
		if tx.BlockNum < tokenInfo.StartBlock {
			log.Error("Minted block number should be greater than start block number. tx: ", tx.ID, " data: ", data)
			return
		}
	}

	realAmt, precision, err := util.NewDecimalFromString(data.amt)
	// invalid mint amount, ignore
	if err != nil {
		return
	}

	// invalid precision, ignore
	if precision > tokenInfo.Decimals {
		return
	}

	totalMinted, _, _ := util.NewDecimalFromString(tokenInfo.TotalMinted)
	mintedOut := false

	if tokenInfo.Type == "fair" {

		limit, _, _ := util.NewDecimalFromString(tokenInfo.Limit)
		// invalid mint amount, ignore
		if realAmt.Sign() < 0 || realAmt.IsOverflowUint64() || realAmt.Cmp(limit) > 0 {
			return
		}

		// fair mode, when maxSupply is set, check if maxSupply is reached
		// treat 0 maxSupply as unlimited
		maxSupply, _, err := util.NewDecimalFromString(tokenInfo.TotalSupply)
		if err == nil && maxSupply.Float64() > 0 {
			remainAmount := maxSupply.Sub(totalMinted)
			if realAmt.Cmp(remainAmount) > 0 {
				// mint what's left
				realAmt = remainAmount
				mintedOut = true
			}
		}

	} else {
		// normal mode
		totalSupply, _, _ := util.NewDecimalFromString(tokenInfo.TotalSupply)
		remainAmount := totalSupply.Sub(totalMinted)

		if tokenInfo.Limit != "" {
			// ignore limit conversion error because it's already checked in deploy
			limit, _, _ := util.NewDecimalFromString(tokenInfo.Limit)
			// invalid mint amount, ignore
			if realAmt.Sign() < 0 || realAmt.IsOverflowUint64() || realAmt.Cmp(limit) > 0 {
				return
			}
		}

		if realAmt.Cmp(remainAmount) > 0 {
			// mint what's left
			realAmt = remainAmount
			mintedOut = true
		}

	}

	totalMinted = totalMinted.Add(realAmt)
	tokenInfo.TotalMinted = totalMinted.String()

	tokenInfo.TotalTxs += 1
	if mintedOut {
		tokenInfo.MintedOutAt = tx.Timestamp
	}

	delta, ok := ctx.TokenHolderBalanceDelta[tx.From][data.tick]
	if ok {
		ctx.TokenHolderBalanceDelta[tx.From][data.tick] = delta.Add(delta, realAmt.Value)
	} else {
		if _, ok := ctx.TokenHolderBalanceDelta[tx.From]; !ok {
			ctx.TokenHolderBalanceDelta[tx.From] = make(map[string]*big.Int)
		}
		ctx.TokenHolderBalanceDelta[tx.From][data.tick] = realAmt.Value
	}

	ctx.TokenInfos[data.tick] = tokenInfo
	ctx.Histories = append(ctx.Histories, orm.HistoryModel{
		TokenName: tokenInfo.Name,
		TxId:      tx.ID,
		Type:      "mint",
		Amount:    realAmt.String(),
	})

}

func (h *Handler) processTransfer(tx orm.TxModel, ctx *context.IndexerContext) {
	data, err := FormTransferData(ctx)
	if err != nil {
		// invalid tx, ignore
		return
	}

	tokenInfo, ok := ctx.TokenInfos[data.tick]
	if !ok {
		// transfer a not existed token, ignore
		return
	}

	// TODO: refactor this to achieve less db query
	fromBalanceModel := orm.TokenHolderModel{}
	err = ctx.DbClient.
		Where("address = ? AND token_name = ?", tx.From, data.tick).
		First(&fromBalanceModel).
		Error
	if err != nil {
		// txFrom have no balance to transfer, ignore
		return
	}

	fromBalance, _, err := util.NewDecimalFromString(fromBalanceModel.Balance)
	if err != nil {
		// invalid balance, ignore
		return
	}
	transferAmt, precision, err := util.NewDecimalFromString(data.amt)
	if err != nil {
		// invalid transfer amount, ignore
		return
	}

	if precision > tokenInfo.Decimals {
		// invalid precision, ignore
		return
	}

	if fromBalance.Cmp(transferAmt) == -1 {
		// txFrom have no enough balance to transfer, ignore
		return
	}

	// init empty fields
	if _, ok := ctx.TokenHolderBalanceDelta[tx.From]; !ok {
		ctx.TokenHolderBalanceDelta[tx.From] = make(map[string]*big.Int)
	}
	if _, ok := ctx.TokenHolderBalanceDelta[tx.To]; !ok {
		ctx.TokenHolderBalanceDelta[tx.To] = make(map[string]*big.Int)
	}
	if _, ok := ctx.TokenHolderBalanceDelta[tx.From][data.tick]; !ok {
		ctx.TokenHolderBalanceDelta[tx.From][data.tick] = big.NewInt(0)
	}
	if _, ok := ctx.TokenHolderBalanceDelta[tx.To][data.tick]; !ok {
		ctx.TokenHolderBalanceDelta[tx.To][data.tick] = big.NewInt(0)
	}

	ctx.TokenHolderBalanceDelta[tx.From][data.tick].
		Sub(ctx.TokenHolderBalanceDelta[tx.From][data.tick], transferAmt.Value)
	ctx.TokenHolderBalanceDelta[tx.To][data.tick].
		Add(ctx.TokenHolderBalanceDelta[tx.To][data.tick], transferAmt.Value)

	tokenInfo.TotalTxs += 1
	ctx.TokenInfos[data.tick] = tokenInfo

	ctx.Histories = append(ctx.Histories, orm.HistoryModel{
		TokenName: tokenInfo.Name,
		TxId:      tx.ID,
		Type:      "transfer",
		Amount:    transferAmt.String(),
	})

}

// it's basically the same as processTransfer, but normally tx.To is market contract address
func (h *Handler) processList(tx orm.TxModel, ctx *context.IndexerContext) {
	// temporary disable index list, will do it from api level


	//data, err := FormTransferData(ctx)
	//if err != nil {
	//	// invalid tx, ignore
	//	return
	//}
	//
	//tokenInfo, ok := ctx.TokenInfos[data.tick]
	//if !ok {
	//	// transfer a not existed token, ignore
	//	return
	//}
	//
	//// TODO: refactor this to achieve less db query
	//fromBalanceModel := orm.TokenHolderModel{}
	//err = ctx.DbClient.
	//	Where("address = ? AND token_name = ?", tx.From, data.tick).
	//	First(&fromBalanceModel).
	//	Error
	//if err != nil {
	//	// txFrom have no balance to transfer, ignore
	//	return
	//}
	//
	//fromBalance, _, err := util.NewDecimalFromString(fromBalanceModel.Balance)
	//if err != nil {
	//	// invalid balance, ignore
	//	return
	//}
	//transferAmt, precision, err := util.NewDecimalFromString(data.amt)
	//if err != nil {
	//	// invalid transfer amount, ignore
	//	return
	//}
	//
	//if precision > tokenInfo.Decimals {
	//	// invalid precision, ignore
	//	return
	//}
	//
	//if fromBalance.Cmp(transferAmt) == -1 {
	//	// txFrom have no enough balance to transfer, ignore
	//	return
	//}
	//
	//// init empty fields
	//if _, ok := ctx.TokenHolderBalanceDelta[tx.From]; !ok {
	//	ctx.TokenHolderBalanceDelta[tx.From] = make(map[string]*big.Int)
	//}
	//if _, ok := ctx.TokenHolderBalanceDelta[tx.To]; !ok {
	//	ctx.TokenHolderBalanceDelta[tx.To] = make(map[string]*big.Int)
	//}
	//if _, ok := ctx.TokenHolderBalanceDelta[tx.From][data.tick]; !ok {
	//	ctx.TokenHolderBalanceDelta[tx.From][data.tick] = big.NewInt(0)
	//}
	//if _, ok := ctx.TokenHolderBalanceDelta[tx.To][data.tick]; !ok {
	//	ctx.TokenHolderBalanceDelta[tx.To][data.tick] = big.NewInt(0)
	//}
	//
	//ctx.TokenHolderBalanceDelta[tx.From][data.tick].
	//	Sub(ctx.TokenHolderBalanceDelta[tx.From][data.tick], transferAmt.Value)
	//ctx.TokenHolderBalanceDelta[tx.To][data.tick].
	//	Add(ctx.TokenHolderBalanceDelta[tx.To][data.tick], transferAmt.Value)
	//
	//tokenInfo.TotalTxs += 1
	//ctx.TokenInfos[data.tick] = tokenInfo
	//
	//ctx.Histories = append(ctx.Histories, orm.HistoryModel{
	//	TokenName: tokenInfo.Name,
	//	TxId:      tx.ID,
	//	Type:      "list",
	//	Amount:    transferAmt.String(),
	//})

	// todo: add list to database

}