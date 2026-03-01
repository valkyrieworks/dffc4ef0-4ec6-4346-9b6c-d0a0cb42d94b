package base

import (
	"context"
	"errors"
	"fmt"
	"time"

	tendermintpubsub "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/broadcastlisten"
	tendermintinquire "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/broadcastlisten/inquire"
	ktypes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/base/kinds"
	remoteifacetypes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/jsoniface/kinds"
)

const (
	//
	//
	maximumInquireMagnitude = 512
)

//
//
func (env *Context) Listen(ctx *remoteifacetypes.Env, inquire string) (*ktypes.OutcomeListen, error) {
	location := ctx.DistantLocation()

	countCustomers := env.IncidentChannel.CountCustomers()
	countSubscriptions := env.IncidentChannel.CountCustomerFeeds(location)

	switch {
	case countCustomers >= env.Settings.MaximumListeningCustomers:
		return nil, fmt.Errorf(
			"REDACTED",
			env.Settings.MaximumListeningCustomers,
		)
	case countSubscriptions >= env.Settings.MaximumFeedsEveryCustomer:
		return nil, fmt.Errorf(
			"REDACTED",
			env.Settings.MaximumFeedsEveryCustomer,
		)
	case len(inquire) > maximumInquireMagnitude:
		return nil, errors.New("REDACTED")
	}

	env.Tracer.Details("REDACTED", "REDACTED", location, "REDACTED", inquire)

	q, err := tendermintinquire.New(inquire)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	underContext, abort := context.WithTimeout(ctx.Env(), ListenDeadline)
	defer abort()

	sub, err := env.IncidentChannel.Listen(underContext, location, q, env.Settings.ListeningReserveExtent)
	if err != nil {
		return nil, err
	}

	shutdownConditionalGradual := env.Settings.ShutdownUponGradualCustomer

	//
	listeningUUID := ctx.JSNRequest.ID
	go func() {
		for {
			select {
			case msg := <-sub.Out():
				var (
					outcomeIncident = &ktypes.OutcomeIncident{Inquire: inquire, Data: msg.Data(), Incidents: msg.Incidents()}
					reply        = remoteifacetypes.FreshRemoteTriumphReply(listeningUUID, outcomeIncident)
				)
				persistContext, abort := context.WithTimeout(context.Background(), 10*time.Second)
				defer abort()
				if err := ctx.SocketLink.PersistRemoteReply(persistContext, reply); err != nil {
					env.Tracer.Details("REDACTED",
						"REDACTED", location, "REDACTED", listeningUUID, "REDACTED", err)

					if shutdownConditionalGradual {
						var (
							err  = errors.New("REDACTED")
							reply = remoteifacetypes.RemoteDaemonFailure(listeningUUID, err)
						)
						if !ctx.SocketLink.AttemptPersistRemoteReply(reply) {
							env.Tracer.Details("REDACTED",
								"REDACTED", location, "REDACTED", listeningUUID, "REDACTED", err)
						}
						return
					}
				}
			case <-sub.Aborted():
				if sub.Err() != tendermintpubsub.FaultUnlistened {
					var rationale string
					if sub.Err() == nil {
						rationale = "REDACTED"
					} else {
						rationale = sub.Err().Error()
					}
					var (
						err  = fmt.Errorf("REDACTED", rationale)
						reply = remoteifacetypes.RemoteDaemonFailure(listeningUUID, err)
					)
					if !ctx.SocketLink.AttemptPersistRemoteReply(reply) {
						env.Tracer.Details("REDACTED",
							"REDACTED", location, "REDACTED", listeningUUID, "REDACTED", err)
					}
				}
				return
			}
		}
	}()

	return &ktypes.OutcomeListen{}, nil
}

//
//
func (env *Context) Unlisten(ctx *remoteifacetypes.Env, inquire string) (*ktypes.OutcomeUnlisten, error) {
	location := ctx.DistantLocation()
	env.Tracer.Details("REDACTED", "REDACTED", location, "REDACTED", inquire)
	q, err := tendermintinquire.New(inquire)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}
	err = env.IncidentChannel.Unlisten(context.Background(), location, q)
	if err != nil {
		return nil, err
	}
	return &ktypes.OutcomeUnlisten{}, nil
}

//
//
func (env *Context) UnlistenEvery(ctx *remoteifacetypes.Env) (*ktypes.OutcomeUnlisten, error) {
	location := ctx.DistantLocation()
	env.Tracer.Details("REDACTED", "REDACTED", location)
	err := env.IncidentChannel.UnlistenEvery(context.Background(), location)
	if err != nil {
		return nil, err
	}
	return &ktypes.OutcomeUnlisten{}, nil
}
