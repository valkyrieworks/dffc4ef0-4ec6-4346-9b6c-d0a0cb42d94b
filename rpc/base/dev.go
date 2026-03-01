package base

import (
	ktypes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/base/kinds"
	remoteifacetypes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/jsoniface/kinds"
)

//
func (env *Context) InsecurePurgeTxpool(*remoteifacetypes.Env) (*ktypes.OutcomeInsecurePurgeTxpool, error) {
	env.Txpool.Purge()
	return &ktypes.OutcomeInsecurePurgeTxpool{}, nil
}
