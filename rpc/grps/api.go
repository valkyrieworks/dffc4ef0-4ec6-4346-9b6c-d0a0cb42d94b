package coregrpc

import (
	"context"

	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	base "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/base"
	remoteifacetypes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/jsoniface/kinds"
)

type multicastAPI struct {
	env *base.Context
}

func (baseiface *multicastAPI) Ping(context.Context, *SolicitPing) (*ReplyPing, error) {
	//
	return &ReplyPing{}, nil
}

func (baseiface *multicastAPI) MulticastTransfer(_ context.Context, req *SolicitMulticastTransfer) (*ReplyMulticastTransfer, error) {
	//
	//
	res, err := baseiface.env.MulticastTransferEndorse(&remoteifacetypes.Env{}, req.Tx)
	if err != nil {
		return nil, err
	}

	return &ReplyMulticastTransfer{
		InspectTransfer: &iface.ReplyInspectTransfer{
			Cipher: res.InspectTransfer.Cipher,
			Data: res.InspectTransfer.Data,
			Log:  res.InspectTransfer.Log,
		},
		TransferOutcome: &iface.InvokeTransferOutcome{
			Cipher: res.TransferOutcome.Cipher,
			Data: res.TransferOutcome.Data,
			Log:  res.TransferOutcome.Log,
		},
	}, nil
}
