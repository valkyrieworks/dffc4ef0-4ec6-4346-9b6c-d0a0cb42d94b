package core

import (
	"context"
	"errors"
	"fmt"
	"time"

	iface "github.com/valkyrieworks/iface/kinds"
	txpool "github.com/valkyrieworks/txpool"
	ctypes "github.com/valkyrieworks/rpc/core/kinds"
	rpctypes "github.com/valkyrieworks/rpc/jsonrpc/kinds"
	"github.com/valkyrieworks/kinds"
)

var ErrTerminusHaltedTrappingUp = errors.New("REDACTED")

//
//

//
//
//
func (env *Context) MulticastTransferAsync(_ *rpctypes.Context, tx kinds.Tx) (*ctypes.OutcomeMulticastTransfer, error) {
	if env.TxpoolHandler.WaitAlign() {
		return nil, ErrTerminusHaltedTrappingUp
	}
	err := env.Txpool.InspectTransfer(tx, nil, txpool.TransferDetails{})
	if err != nil {
		return nil, err
	}
	return &ctypes.OutcomeMulticastTransfer{Digest: tx.Digest()}, nil
}

//
//
//
func (env *Context) MulticastTransferAlign(ctx *rpctypes.Context, tx kinds.Tx) (*ctypes.OutcomeMulticastTransfer, error) {
	if env.TxpoolHandler.WaitAlign() {
		return nil, ErrTerminusHaltedTrappingUp
	}

	outputChan := make(chan *iface.ReplyInspectTransfer, 1)
	err := env.Txpool.InspectTransfer(tx, func(res *iface.ReplyInspectTransfer) {
		select {
		case <-ctx.Context().Done():
		case outputChan <- res:
		}
	}, txpool.TransferDetails{})
	if err != nil {
		return nil, err
	}

	select {
	case <-ctx.Context().Done():
		return nil, fmt.Errorf("REDACTED", ctx.Context().Err())
	case res := <-outputChan:
		return &ctypes.OutcomeMulticastTransfer{
			Code:      res.Code,
			Data:      res.Data,
			Log:       res.Log,
			Codex: res.Codex,
			Digest:      tx.Digest(),
		}, nil
	}
}

//
//
func (env *Context) MulticastTransferEndorse(ctx *rpctypes.Context, tx kinds.Tx) (*ctypes.OutcomeMulticastTransferEndorse, error) {
	if env.TxpoolHandler.WaitAlign() {
		return nil, ErrTerminusHaltedTrappingUp
	}

	enrollee := ctx.DistantAddress()

	if env.EventBus.CountAgents() >= env.Settings.MaximumEnrollmentAgents {
		return nil, fmt.Errorf("REDACTED", env.Settings.MaximumEnrollmentAgents)
	} else if env.EventBus.CountCustomerRegistrations(enrollee) >= env.Settings.MaximumRegistrationsEachCustomer {
		return nil, fmt.Errorf("REDACTED", env.Settings.MaximumRegistrationsEachCustomer)
	}

	//
	subtractCtx, revoke := context.WithTimeout(ctx.Context(), EnrolDeadline)
	defer revoke()
	q := kinds.EventInquireTransferFor(tx)
	transferSubtract, err := env.EventBus.Enrol(subtractCtx, enrollee, q)
	if err != nil {
		err = fmt.Errorf("REDACTED", err)
		env.Tracer.Fault("REDACTED", "REDACTED", err)
		return nil, err
	}
	defer func() {
		if err := env.EventBus.Deenroll(context.Background(), enrollee, q); err != nil {
			env.Tracer.Fault("REDACTED", "REDACTED", err)
		}
	}()

	//
	inspectTransferOutputChan := make(chan *iface.ReplyInspectTransfer, 1)
	err = env.Txpool.InspectTransfer(tx, func(res *iface.ReplyInspectTransfer) {
		select {
		case <-ctx.Context().Done():
		case inspectTransferOutputChan <- res:
		}
	}, txpool.TransferDetails{})
	if err != nil {
		env.Tracer.Fault("REDACTED", "REDACTED", err)
		return nil, fmt.Errorf("REDACTED", err)
	}
	select {
	case <-ctx.Context().Done():
		return nil, fmt.Errorf("REDACTED", ctx.Context().Err())
	case inspectTransferOutput := <-inspectTransferOutputChan:
		if inspectTransferOutput.Code != iface.CodeKindSuccess {
			return &ctypes.OutcomeMulticastTransferEndorse{
				InspectTransfer:  *inspectTransferOutput,
				TransOutcome: iface.InvokeTransferOutcome{},
				Digest:     tx.Digest(),
			}, nil
		}

		//
		select {
		case msg := <-transferSubtract.Out(): //
			transferOutcomeEvent := msg.Data().(kinds.EventDataTransfer)
			return &ctypes.OutcomeMulticastTransferEndorse{
				InspectTransfer:  *inspectTransferOutput,
				TransOutcome: transferOutcomeEvent.Outcome,
				Digest:     tx.Digest(),
				Level:   transferOutcomeEvent.Level,
			}, nil
		case <-transferSubtract.Revoked():
			var cause string
			if transferSubtract.Err() == nil {
				cause = "REDACTED"
			} else {
				cause = transferSubtract.Err().Error()
			}
			err = fmt.Errorf("REDACTED", cause)
			env.Tracer.Fault("REDACTED", "REDACTED", err)
			return &ctypes.OutcomeMulticastTransferEndorse{
				InspectTransfer:  *inspectTransferOutput,
				TransOutcome: iface.InvokeTransferOutcome{},
				Digest:     tx.Digest(),
			}, err
		case <-time.After(env.Settings.DeadlineMulticastTransEndorse):
			err = errors.New("REDACTED")
			env.Tracer.Fault("REDACTED", "REDACTED", err)
			return &ctypes.OutcomeMulticastTransferEndorse{
				InspectTransfer:  *inspectTransferOutput,
				TransOutcome: iface.InvokeTransferOutcome{},
				Digest:     tx.Digest(),
			}, err
		}
	}
}

//
//
//
func (env *Context) UnattestedTrans(_ *rpctypes.Context, ceilingPointer *int) (*ctypes.OutcomeUnattestedTrans, error) {
	//
	ceiling := env.certifyEachScreen(ceilingPointer)

	txs := env.Txpool.HarvestMaximumTrans(ceiling)
	return &ctypes.OutcomeUnattestedTrans{
		Number:      len(txs),
		Sum:      env.Txpool.Volume(),
		SumOctets: env.Txpool.VolumeOctets(),
		Txs:        txs,
	}, nil
}

//
//
func (env *Context) CountUnattestedTrans(*rpctypes.Context) (*ctypes.OutcomeUnattestedTrans, error) {
	return &ctypes.OutcomeUnattestedTrans{
		Number:      env.Txpool.Volume(),
		Sum:      env.Txpool.Volume(),
		SumOctets: env.Txpool.VolumeOctets(),
	}, nil
}

//
//
//
func (env *Context) InspectTransfer(_ *rpctypes.Context, tx kinds.Tx) (*ctypes.OutcomeInspectTransfer, error) {
	res, err := env.GatewayApplicationTxpool.InspectTransfer(context.TODO(), &iface.QueryInspectTransfer{Tx: tx})
	if err != nil {
		return nil, err
	}
	return &ctypes.OutcomeInspectTransfer{ReplyInspectTransfer: *res}, nil
}
