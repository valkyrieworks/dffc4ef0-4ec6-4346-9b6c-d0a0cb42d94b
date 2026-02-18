package core

import (
	ctypes "github.com/valkyrieworks/rpc/core/kinds"
	rpctypes "github.com/valkyrieworks/rpc/jsonrpc/kinds"
)

//
func (env *Context) RiskyPurgeTxpool(*rpctypes.Context) (*ctypes.OutcomeRiskyPurgeTxpool, error) {
	env.Txpool.Purge()
	return &ctypes.OutcomeRiskyPurgeTxpool{}, nil
}
