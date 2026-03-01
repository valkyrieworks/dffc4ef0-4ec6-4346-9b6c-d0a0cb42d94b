package base

import (
	"context"
	"errors"
	"fmt"
	"time"

	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	txpooll "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/txpool"
	ktypes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/base/kinds"
	remoteifacetypes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/jsoniface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

var FaultGatewayTerminatedObtainingAscend = errors.New("REDACTED")

//
//

//
//
//
func (env *Context) MulticastTransferAsyncronous(_ *remoteifacetypes.Env, tx kinds.Tx) (*ktypes.OutcomeMulticastTransfer, error) {
	if env.TxpoolHandler.AwaitChronize() {
		return nil, FaultGatewayTerminatedObtainingAscend
	}
	err := env.Txpool.InspectTransfer(tx, nil, txpooll.TransferDetails{})
	if err != nil {
		return nil, err
	}
	return &ktypes.OutcomeMulticastTransfer{Digest: tx.Digest()}, nil
}

//
//
//
func (env *Context) MulticastTransferChronize(ctx *remoteifacetypes.Env, tx kinds.Tx) (*ktypes.OutcomeMulticastTransfer, error) {
	if env.TxpoolHandler.AwaitChronize() {
		return nil, FaultGatewayTerminatedObtainingAscend
	}

	resultChnl := make(chan *iface.ReplyInspectTransfer, 1)
	err := env.Txpool.InspectTransfer(tx, func(res *iface.ReplyInspectTransfer) {
		select {
		case <-ctx.Env().Done():
		case resultChnl <- res:
		}
	}, txpooll.TransferDetails{})
	if err != nil {
		return nil, err
	}

	select {
	case <-ctx.Env().Done():
		return nil, fmt.Errorf("REDACTED", ctx.Env().Err())
	case res := <-resultChnl:
		return &ktypes.OutcomeMulticastTransfer{
			Cipher:      res.Cipher,
			Data:      res.Data,
			Log:       res.Log,
			Codeset: res.Codeset,
			Digest:      tx.Digest(),
		}, nil
	}
}

//
//
func (env *Context) MulticastTransferEndorse(ctx *remoteifacetypes.Env, tx kinds.Tx) (*ktypes.OutcomeMulticastTransferEndorse, error) {
	if env.TxpoolHandler.AwaitChronize() {
		return nil, FaultGatewayTerminatedObtainingAscend
	}

	listener := ctx.DistantLocation()

	if env.IncidentChannel.CountCustomers() >= env.Settings.MaximumListeningCustomers {
		return nil, fmt.Errorf("REDACTED", env.Settings.MaximumListeningCustomers)
	} else if env.IncidentChannel.CountCustomerFeeds(listener) >= env.Settings.MaximumFeedsEveryCustomer {
		return nil, fmt.Errorf("REDACTED", env.Settings.MaximumFeedsEveryCustomer)
	}

	//
	underContext, abort := context.WithTimeout(ctx.Env(), ListenDeadline)
	defer abort()
	q := kinds.IncidentInquireTransferForeach(tx)
	transferUnder, err := env.IncidentChannel.Listen(underContext, listener, q)
	if err != nil {
		err = fmt.Errorf("REDACTED", err)
		env.Tracer.Failure("REDACTED", "REDACTED", err)
		return nil, err
	}
	defer func() {
		if err := env.IncidentChannel.Unlisten(context.Background(), listener, q); err != nil {
			env.Tracer.Failure("REDACTED", "REDACTED", err)
		}
	}()

	//
	inspectTransferResultChnl := make(chan *iface.ReplyInspectTransfer, 1)
	err = env.Txpool.InspectTransfer(tx, func(res *iface.ReplyInspectTransfer) {
		select {
		case <-ctx.Env().Done():
		case inspectTransferResultChnl <- res:
		}
	}, txpooll.TransferDetails{})
	if err != nil {
		env.Tracer.Failure("REDACTED", "REDACTED", err)
		return nil, fmt.Errorf("REDACTED", err)
	}
	select {
	case <-ctx.Env().Done():
		return nil, fmt.Errorf("REDACTED", ctx.Env().Err())
	case inspectTransferResult := <-inspectTransferResultChnl:
		if inspectTransferResult.Cipher != iface.CipherKindOKAY {
			return &ktypes.OutcomeMulticastTransferEndorse{
				InspectTransfer:  *inspectTransferResult,
				TransferOutcome: iface.InvokeTransferOutcome{},
				Digest:     tx.Digest(),
			}, nil
		}

		//
		select {
		case msg := <-transferUnder.Out(): //
			transferOutcomeIncident := msg.Data().(kinds.IncidentDataTransfer)
			return &ktypes.OutcomeMulticastTransferEndorse{
				InspectTransfer:  *inspectTransferResult,
				TransferOutcome: transferOutcomeIncident.Outcome,
				Digest:     tx.Digest(),
				Altitude:   transferOutcomeIncident.Altitude,
			}, nil
		case <-transferUnder.Aborted():
			var rationale string
			if transferUnder.Err() == nil {
				rationale = "REDACTED"
			} else {
				rationale = transferUnder.Err().Error()
			}
			err = fmt.Errorf("REDACTED", rationale)
			env.Tracer.Failure("REDACTED", "REDACTED", err)
			return &ktypes.OutcomeMulticastTransferEndorse{
				InspectTransfer:  *inspectTransferResult,
				TransferOutcome: iface.InvokeTransferOutcome{},
				Digest:     tx.Digest(),
			}, err
		case <-time.After(env.Settings.DeadlineMulticastTransferEndorse):
			err = errors.New("REDACTED")
			env.Tracer.Failure("REDACTED", "REDACTED", err)
			return &ktypes.OutcomeMulticastTransferEndorse{
				InspectTransfer:  *inspectTransferResult,
				TransferOutcome: iface.InvokeTransferOutcome{},
				Digest:     tx.Digest(),
			}, err
		}
	}
}

//
//
//
func (env *Context) PendingTrans(_ *remoteifacetypes.Env, thresholdReference *int) (*ktypes.OutcomePendingTrans, error) {
	//
	threshold := env.certifyEveryScreen(thresholdReference)

	txs := env.Txpool.HarvestMaximumTrans(threshold)
	return &ktypes.OutcomePendingTrans{
		Tally:      len(txs),
		Sum:      env.Txpool.Extent(),
		SumOctets: env.Txpool.ExtentOctets(),
		Txs:        txs,
	}, nil
}

//
//
func (env *Context) CountPendingTrans(*remoteifacetypes.Env) (*ktypes.OutcomePendingTrans, error) {
	return &ktypes.OutcomePendingTrans{
		Tally:      env.Txpool.Extent(),
		Sum:      env.Txpool.Extent(),
		SumOctets: env.Txpool.ExtentOctets(),
	}, nil
}

//
//
//
func (env *Context) InspectTransfer(_ *remoteifacetypes.Env, tx kinds.Tx) (*ktypes.OutcomeInspectTransfer, error) {
	res, err := env.DelegateApplicationTxpool.InspectTransfer(context.TODO(), &iface.SolicitInspectTransfer{Tx: tx})
	if err != nil {
		return nil, err
	}
	return &ktypes.OutcomeInspectTransfer{ReplyInspectTransfer: *res}, nil
}
