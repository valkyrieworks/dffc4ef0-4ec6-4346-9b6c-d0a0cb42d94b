package core

import (
	"context"

	iface "github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/utils/octets"
	"github.com/valkyrieworks/gateway"
	ctypes "github.com/valkyrieworks/rpc/core/kinds"
	rpctypes "github.com/valkyrieworks/rpc/jsonrpc/kinds"
)

//
//
func (env *Context) IfaceInquire(
	_ *rpctypes.Context,
	route string,
	data octets.HexOctets,
	level int64,
	demonstrate bool,
) (*ctypes.OutcomeIfaceInquire, error) {
	outcomeInquire, err := env.GatewayApplicationInquire.Inquire(context.TODO(), &iface.QueryInquire{
		Route:   route,
		Data:   data,
		Level: level,
		Demonstrate:  demonstrate,
	})
	if err != nil {
		return nil, err
	}

	return &ctypes.OutcomeIfaceInquire{Reply: *outcomeInquire}, nil
}

//
//
func (env *Context) IfaceDetails(_ *rpctypes.Context) (*ctypes.OutcomeIfaceDetails, error) {
	outputDetails, err := env.GatewayApplicationInquire.Details(context.TODO(), gateway.QueryDetails)
	if err != nil {
		return nil, err
	}

	return &ctypes.OutcomeIfaceDetails{Reply: *outputDetails}, nil
}
