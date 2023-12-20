package common

import (
	"encoding/json"
	"ethsyncer/pkg/context"
)

type Op int

const (
	Deploy Op = iota
	Mint
	Transfer
	List
	UNKNOWN_OP
)

const MAX_DECIMAL = 18

// MAX_DURATION average block time is 12s, set to two days for now
const MAX_DURATION  = (60 * 60 * 24 * 2) / 12

// GetProtocolOp returns the protocol op of the tx
// if the tx is not supported, return UNKNOWN_OP
// if the tx is supported, return the protocol op and put unmarshalled tx input into ctx
// after calling this function, if Op not UNKNOWN_OP, then ctx.TxInputUnmarshalled is not nil
func GetProtocolOp(data string, ctx *context.IndexerContext) Op {
	if ctx.TxInputUnmarshalled == nil {
		var target map[string]interface{}
		err := json.Unmarshal([]byte(data), &target)
		if err != nil {
			return UNKNOWN_OP
		}
		ctx.TxInputUnmarshalled = target
	}

	op, ok := ctx.TxInputUnmarshalled["op"]
	if !ok {
		return UNKNOWN_OP
	}

	opStr, ok := op.(string)
	if !ok {
		return UNKNOWN_OP
	}

	if opStr == "deploy" {
		return Deploy
	} else if opStr == "mint" {
		return Mint
	} else if opStr == "transfer" {
		return Transfer
	} else if opStr == "list" {
		return List
	} else {
		return UNKNOWN_OP
	}
}
