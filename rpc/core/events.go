package core

import (
	"context"
	"errors"
	"fmt"
	"time"

	cometbroadcast "github.com/valkyrieworks/utils/broadcast"
	cmtinquire "github.com/valkyrieworks/utils/broadcast/inquire"
	ctypes "github.com/valkyrieworks/rpc/core/kinds"
	rpctypes "github.com/valkyrieworks/rpc/jsonrpc/kinds"
)

const (
	//
	//
	maximumInquireExtent = 512
)

//
//
func (env *Context) Enrol(ctx *rpctypes.Context, inquire string) (*ctypes.OutcomeEnrol, error) {
	address := ctx.DistantAddress()

	countAgents := env.EventBus.CountAgents()
	countEnrollments := env.EventBus.CountCustomerRegistrations(address)

	switch {
	case countAgents >= env.Settings.MaximumEnrollmentAgents:
		return nil, fmt.Errorf(
			"REDACTED",
			env.Settings.MaximumEnrollmentAgents,
		)
	case countEnrollments >= env.Settings.MaximumRegistrationsEachCustomer:
		return nil, fmt.Errorf(
			"REDACTED",
			env.Settings.MaximumRegistrationsEachCustomer,
		)
	case len(inquire) > maximumInquireExtent:
		return nil, errors.New("REDACTED")
	}

	env.Tracer.Details("REDACTED", "REDACTED", address, "REDACTED", inquire)

	q, err := cmtinquire.New(inquire)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	subtractCtx, revoke := context.WithTimeout(ctx.Context(), EnrolDeadline)
	defer revoke()

	sub, err := env.EventBus.Enrol(subtractCtx, address, q, env.Settings.EnrollmentBufferVolume)
	if err != nil {
		return nil, err
	}

	endIfGradual := env.Settings.EndOnGradualCustomer

	//
	enrollmentUID := ctx.JSONRequest.ID
	go func() {
		for {
			select {
			case msg := <-sub.Out():
				var (
					outcomeEvent = &ctypes.OutcomeEvent{Inquire: inquire, Data: msg.Data(), Events: msg.Events()}
					reply        = rpctypes.NewRPCSuccessReply(enrollmentUID, outcomeEvent)
				)
				recordCtx, revoke := context.WithTimeout(context.Background(), 10*time.Second)
				defer revoke()
				if err := ctx.WSLink.RecordRPCReply(recordCtx, reply); err != nil {
					env.Tracer.Details("REDACTED",
						"REDACTED", address, "REDACTED", enrollmentUID, "REDACTED", err)

					if endIfGradual {
						var (
							err  = errors.New("REDACTED")
							reply = rpctypes.RPCHostFault(enrollmentUID, err)
						)
						if !ctx.WSLink.AttemptRecordRPCReply(reply) {
							env.Tracer.Details("REDACTED",
								"REDACTED", address, "REDACTED", enrollmentUID, "REDACTED", err)
						}
						return
					}
				}
			case <-sub.Revoked():
				if sub.Err() != cometbroadcast.ErrDeactivated {
					var cause string
					if sub.Err() == nil {
						cause = "REDACTED"
					} else {
						cause = sub.Err().Error()
					}
					var (
						err  = fmt.Errorf("REDACTED", cause)
						reply = rpctypes.RPCHostFault(enrollmentUID, err)
					)
					if !ctx.WSLink.AttemptRecordRPCReply(reply) {
						env.Tracer.Details("REDACTED",
							"REDACTED", address, "REDACTED", enrollmentUID, "REDACTED", err)
					}
				}
				return
			}
		}
	}()

	return &ctypes.OutcomeEnrol{}, nil
}

//
//
func (env *Context) Deenroll(ctx *rpctypes.Context, inquire string) (*ctypes.OutcomeDeenroll, error) {
	address := ctx.DistantAddress()
	env.Tracer.Details("REDACTED", "REDACTED", address, "REDACTED", inquire)
	q, err := cmtinquire.New(inquire)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}
	err = env.EventBus.Deenroll(context.Background(), address, q)
	if err != nil {
		return nil, err
	}
	return &ctypes.OutcomeDeenroll{}, nil
}

//
//
func (env *Context) DeenrollAll(ctx *rpctypes.Context) (*ctypes.OutcomeDeenroll, error) {
	address := ctx.DistantAddress()
	env.Tracer.Details("REDACTED", "REDACTED", address)
	err := env.EventBus.DeenrollAll(context.Background(), address)
	if err != nil {
		return nil, err
	}
	return &ctypes.OutcomeDeenroll{}, nil
}
