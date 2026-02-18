package coregrpc

import (
	"context"

	iface "github.com/valkyrieworks/iface/kinds"
	core "github.com/valkyrieworks/rpc/core"
	rpctypes "github.com/valkyrieworks/rpc/jsonrpc/kinds"
)

type multicastAPI struct {
	env *core.Context
}

func (baseapi *multicastAPI) Ping(context.Context, *QueryPing) (*AnswerPing, error) {
	//
	return &AnswerPing{}, nil
}

func (baseapi *multicastAPI) MulticastTransfer(_ context.Context, req *QueryMulticastTransfer) (*AnswerMulticastTransfer, error) {
	//
	//
	res, err := baseapi.env.MulticastTransferEndorse(&rpctypes.Context{}, req.Tx)
	if err != nil {
		return nil, err
	}

	return &AnswerMulticastTransfer{
		InspectTransfer: &iface.ReplyInspectTransfer{
			Code: res.InspectTransfer.Code,
			Data: res.InspectTransfer.Data,
			Log:  res.InspectTransfer.Log,
		},
		TransOutcome: &iface.InvokeTransferOutcome{
			Code: res.TransOutcome.Code,
			Data: res.TransOutcome.Data,
			Log:  res.TransOutcome.Log,
		},
	}, nil
}
