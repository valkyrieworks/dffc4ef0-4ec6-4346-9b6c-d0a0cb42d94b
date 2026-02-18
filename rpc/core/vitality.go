package core

import (
	ctypes "github.com/valkyrieworks/rpc/core/kinds"
	rpctypes "github.com/valkyrieworks/rpc/jsonrpc/kinds"
)

//
//
//
func (env *Context) Vitality(*rpctypes.Context) (*ctypes.OutcomeVitality, error) {
	return &ctypes.OutcomeVitality{}, nil
}
