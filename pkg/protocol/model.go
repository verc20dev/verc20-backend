package protocol

import (
	"encoding/json"
	"ethsyncer/pkg/context"
	"ethsyncer/pkg/orm"
	"ethsyncer/pkg/protocol/default"
	"strings"
)


type ProtocolType int

const (
	ERC20 ProtocolType = iota
	ERC20_FAIR
	UNKNOWN_PROTOCOL
)

type Handler interface {
	Process(tx orm.TxModel, ctx *context.IndexerContext)
}


// GetProtocolType returns the protocol type of the tx
// if the tx is not supported, return UNKNOWN_PROTOCOL
// if the tx is supported, return the protocol type and put unmarshalled tx input into ctx
func GetProtocolType(data string, ctx *context.IndexerContext) ProtocolType {
	var target map[string]interface{}
	err := json.Unmarshal([]byte(data), &target)
	if err != nil {
		return UNKNOWN_PROTOCOL
	}
	p, ok := target["p"]
	if !ok {
		return UNKNOWN_PROTOCOL
	}

	pStr, ok := p.(string)
	if !ok {
		return UNKNOWN_PROTOCOL
	}

	if ctx.TxInputUnmarshalled == nil {
		ctx.TxInputUnmarshalled = make(map[string]interface{})
	}
	ctx.TxInputUnmarshalled = target

	if strings.ToLower(pStr) == "erc-20" {
		if isFairMode(target) {
			return ERC20_FAIR
		}
		return ERC20
	}

	return UNKNOWN_PROTOCOL
}


func GetHandler(pType ProtocolType) Handler  {
	if pType == ERC20 {
		return &_default.Handler{}
	} else {
		return nil
	}
}

func isFairMode(data map[string]interface{}) bool {
	t, ok := data["t"]
	if !ok {
		return false
	}

	tStr, ok := t.(string)
	if !ok {
		return false
	}

	if strings.ToLower(tStr) != "fair" {
		return false
	}

	return true
}
