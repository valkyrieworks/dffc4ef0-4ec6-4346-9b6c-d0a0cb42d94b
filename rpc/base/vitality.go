package base

import (
	ktypes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/base/kinds"
	remoteifacetypes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/jsoniface/kinds"
)

//
//
//
func (env *Context) Vitality(*remoteifacetypes.Env) (*ktypes.OutcomeVitality, error) {
	return &ktypes.OutcomeVitality{}, nil
}
