package base

import (
	"context"

	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/octets"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/delegate"
	ktypes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/base/kinds"
	remoteifacetypes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/jsoniface/kinds"
)

//
//
func (env *Context) IfaceInquire(
	_ *remoteifacetypes.Env,
	route string,
	data octets.HexadecimalOctets,
	altitude int64,
	ascertain bool,
) (*ktypes.OutcomeIfaceInquire, error) {
	outcomeInquire, err := env.DelegateApplicationInquire.Inquire(context.TODO(), &iface.SolicitInquire{
		Route:   route,
		Data:   data,
		Altitude: altitude,
		Validate:  ascertain,
	})
	if err != nil {
		return nil, err
	}

	return &ktypes.OutcomeIfaceInquire{Reply: *outcomeInquire}, nil
}

//
//
func (env *Context) IfaceDetails(_ *remoteifacetypes.Env) (*ktypes.OutcomeIfaceDetails, error) {
	resultDetails, err := env.DelegateApplicationInquire.Details(context.TODO(), delegate.SolicitDetails)
	if err != nil {
		return nil, err
	}

	return &ktypes.OutcomeIfaceDetails{Reply: *resultDetails}, nil
}
