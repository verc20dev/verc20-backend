package _default

import (
	"errors"
	"ethsyncer/pkg/context"
)

type DeployData struct {
	p    string
	op   string
	t    string
	tick string
	max  string
	lim  string
	dec  string

	startBlock string
	duration   string
}

type MintOrTransferData struct {
	p    string
	op   string
	tick string
	amt  string
}

func FormDeployData(ctx *context.IndexerContext) (*DeployData, error) {
	res := DeployData{}

	if ctx.TxInputUnmarshalled == nil {
		return nil, errors.New("invalid tx input")
	}

	p, ok := ctx.TxInputUnmarshalled["p"]
	if !ok {
		return nil, errors.New("invalid tx input, there is no p field")
	}
	pStr, ok := p.(string)
	if !ok {
		return nil, errors.New("invalid tx input, the p field is not string")
	}
	res.p = pStr

	op, ok := ctx.TxInputUnmarshalled["op"]
	if !ok {
		return nil, errors.New("invalid tx input, there is no op field")
	}
	opStr, ok := op.(string)
	if !ok {
		return nil, errors.New("invalid tx input, the op field is not string")
	}
	res.op = opStr

	tick, ok := ctx.TxInputUnmarshalled["tick"]
	if !ok {
		return nil, errors.New("invalid tx input, there is no tick field")
	}
	tickStr, ok := tick.(string)
	if !ok {
		return nil, errors.New("invalid tx input, the tick field is not string")
	}
	res.tick = tickStr

	// t is optional
	// if t is not exist, then it's normal deploy
	t, ok := ctx.TxInputUnmarshalled["t"]
	if ok {
		tStr, ok := t.(string)
		if ok {
			res.t = tStr
		}
	}

	// m is optional in some cases, validate it later
	m, ok := ctx.TxInputUnmarshalled["max"]
	if ok {
		mStr, ok := m.(string)
		if ok {
			res.max = mStr
		}
	}


	// lim is optional in some cases, validate it later
	lim, ok := ctx.TxInputUnmarshalled["lim"]
	if ok {
		limStr, ok := lim.(string)
		if ok {
			res.lim = limStr
		}
	}

	// startBlock is optional in some cases, validate it later
	startBlock, ok := ctx.TxInputUnmarshalled["startBlock"]
	if ok {
		startBlockStr, ok := startBlock.(string)
		if ok {
			res.startBlock = startBlockStr
		}
	}

	// duration is optional in some cases, validate it later
	duration, ok := ctx.TxInputUnmarshalled["duration"]
	if ok {
		durationStr, ok := duration.(string)
		if ok {
			res.duration = durationStr
		}
	}

	// dec is optional
	dec, ok := ctx.TxInputUnmarshalled["dec"]
	if ok {
		decStr, ok := dec.(string)
		if !ok {
			return nil, errors.New("invalid tx input, the dec field is not string")
		}
		res.dec = decStr
	}

	if err := validateDeployData(res); err != nil {
		return nil, err
	} else {
		return &res, nil
	}
}

func validateDeployData(data DeployData) error {
	err := errors.New("invalid format")
	if data.p == "" {
		return err
	}
	if data.op != "deploy" {
		return err
	}
	if data.tick == "" {
		return err
	}

	// when t is not exist, then it's normal deploy, max is required
	if data.t == "fair" {
		// when t is fair, max is optional
		if data.max != "" {
			return err
		}
		// lim is required
		if data.lim == "" {
			return err
		}
	} else {
		// just treat it as normal deploy
		if data.max == "" {
			return err
		}
	}

	return nil
}

func FormMintData(ctx *context.IndexerContext) (*MintOrTransferData, error) {
	res := MintOrTransferData{}

	if ctx.TxInputUnmarshalled == nil {
		return nil, errors.New("invalid tx input")
	}

	p, ok := ctx.TxInputUnmarshalled["p"]
	if !ok {
		return nil, errors.New("invalid tx input, there is no p field")
	}
	pStr, ok := p.(string)
	if !ok {
		return nil, errors.New("invalid tx input, the p field is not string")
	}
	res.p = pStr

	op, ok := ctx.TxInputUnmarshalled["op"]
	if !ok {
		return nil, errors.New("invalid tx input, there is no op field")
	}
	opStr, ok := op.(string)
	if !ok {
		return nil, errors.New("invalid tx input, the op field is not string")
	}
	res.op = opStr

	tick, ok := ctx.TxInputUnmarshalled["tick"]
	if !ok {
		return nil, errors.New("invalid tx input, there is no tick field")
	}
	tickStr, ok := tick.(string)
	if !ok {
		return nil, errors.New("invalid tx input, the tick field is not string")
	}
	res.tick = tickStr

	amt, ok := ctx.TxInputUnmarshalled["amt"]
	if !ok {
		return nil, errors.New("invalid tx input, there is no amt field")
	}
	amtStr, ok := amt.(string)
	if !ok {
		return nil, errors.New("invalid tx input, the amt field is not string")
	}
	res.amt = amtStr

	if err := validateMintData(res); err != nil {
		return nil, err
	} else {
		return &res, nil
	}
}

func validateMintData(data MintOrTransferData) error {
	err := errors.New("invalid format")
	if data.p == "" {
		return err
	}
	if data.op != "mint" {
		return err
	}
	if data.tick == "" {
		return err
	}
	if data.amt == "" {
		return err
	}
	return nil
}

func FormTransferData(ctx *context.IndexerContext) (*MintOrTransferData, error) {
	res := MintOrTransferData{}

	if ctx.TxInputUnmarshalled == nil {
		return nil, errors.New("invalid tx input")
	}

	p, ok := ctx.TxInputUnmarshalled["p"]
	if !ok {
		return nil, errors.New("invalid tx input, there is no p field")
	}
	pStr, ok := p.(string)
	if !ok {
		return nil, errors.New("invalid tx input, the p field is not string")
	}
	res.p = pStr

	op, ok := ctx.TxInputUnmarshalled["op"]
	if !ok {
		return nil, errors.New("invalid tx input, there is no op field")
	}
	opStr, ok := op.(string)
	if !ok {
		return nil, errors.New("invalid tx input, the op field is not string")
	}
	res.op = opStr

	tick, ok := ctx.TxInputUnmarshalled["tick"]
	if !ok {
		return nil, errors.New("invalid tx input, there is no tick field")
	}
	tickStr, ok := tick.(string)
	if !ok {
		return nil, errors.New("invalid tx input, the tick field is not string")
	}
	res.tick = tickStr

	amt, ok := ctx.TxInputUnmarshalled["amt"]
	if !ok {
		return nil, errors.New("invalid tx input, there is no amt field")
	}
	amtStr, ok := amt.(string)
	if !ok {
		return nil, errors.New("invalid tx input, the amt field is not string")
	}
	res.amt = amtStr

	if err := validateTransferData(res); err != nil {
		return nil, err
	} else {
		return &res, nil
	}
}

func validateTransferData(data MintOrTransferData) error {
	err := errors.New("invalid format")
	if data.p == "" {
		return err
	}
	if data.op != "transfer" {
		return err
	}
	if data.tick == "" {
		return err
	}
	if data.amt == "" {
		return err
	}
	return nil
}
