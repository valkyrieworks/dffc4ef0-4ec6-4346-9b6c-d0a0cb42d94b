package privatevalue

import (
	"fmt"
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	strongminderrors "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds/faults"

	cryptocode "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/serialization"
	privatevalueschema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/privatevalue"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

//
//
type EndorserCustomer struct {
	gateway *EndorserObserverGateway
	successionUUID  string
}

var _ kinds.PrivateAssessor = (*EndorserCustomer)(nil)

//
//
func FreshEndorserCustomer(gateway *EndorserObserverGateway, successionUUID string) (*EndorserCustomer, error) {
	if !gateway.EqualsActive() {
		if err := gateway.Initiate(); err != nil {
			return nil, fmt.Errorf("REDACTED", err)
		}
	}

	return &EndorserCustomer{gateway: gateway, successionUUID: successionUUID}, nil
}

//
func (sc *EndorserCustomer) Shutdown() error {
	return sc.gateway.Shutdown()
}

//
func (sc *EndorserCustomer) EqualsAssociated() bool {
	return sc.gateway.EqualsAssociated()
}

//
func (sc *EndorserCustomer) PauseForeachLinkage(maximumPause time.Duration) error {
	return sc.gateway.PauseForeachLinkage(maximumPause)
}

//
//

//
func (sc *EndorserCustomer) Ping() error {
	reply, err := sc.gateway.TransmitSolicit(shouldEncloseSignal(&privatevalueschema.PingSolicit{}))
	if err != nil {
		sc.gateway.Tracer.Failure("REDACTED", "REDACTED", err)
		return nil
	}

	pb := reply.ObtainPingReply()
	if pb == nil {
		return err
	}

	return nil
}

//
//
func (sc *EndorserCustomer) ObtainPublicToken() (security.PublicToken, error) {
	reply, err := sc.gateway.TransmitSolicit(shouldEncloseSignal(&privatevalueschema.PublicTokenSolicit{SuccessionUuid: sc.successionUUID}))
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	reply := reply.ObtainPublicTokenReply()
	if reply == nil {
		return nil, strongminderrors.FaultMandatoryAttribute{Attribute: "REDACTED"}
	}
	if reply.Failure != nil {
		return nil, &RemoteEndorserFailure{Cipher: int(reply.Failure.Cipher), Characterization: reply.Failure.Characterization}
	}

	pk, err := cryptocode.PublicTokenOriginatingSchema(reply.PublicToken)
	if err != nil {
		return nil, err
	}

	return pk, nil
}

//
func (sc *EndorserCustomer) AttestBallot(successionUUID string, ballot *commitchema.Ballot) error {
	reply, err := sc.gateway.TransmitSolicit(shouldEncloseSignal(&privatevalueschema.AttestBallotSolicit{Ballot: ballot, SuccessionUuid: successionUUID}))
	if err != nil {
		return err
	}

	reply := reply.ObtainNotatedBallotReply()
	if reply == nil {
		return strongminderrors.FaultMandatoryAttribute{Attribute: "REDACTED"}
	}
	if reply.Failure != nil {
		return &RemoteEndorserFailure{Cipher: int(reply.Failure.Cipher), Characterization: reply.Failure.Characterization}
	}

	*ballot = reply.Ballot

	return nil
}

//
func (sc *EndorserCustomer) AttestNomination(successionUUID string, nomination *commitchema.Nomination) error {
	reply, err := sc.gateway.TransmitSolicit(shouldEncloseSignal(
		&privatevalueschema.AttestNominationSolicit{Nomination: nomination, SuccessionUuid: successionUUID},
	))
	if err != nil {
		return err
	}

	reply := reply.ObtainNotatedNominationReply()
	if reply == nil {
		return strongminderrors.FaultMandatoryAttribute{Attribute: "REDACTED"}
	}
	if reply.Failure != nil {
		return &RemoteEndorserFailure{Cipher: int(reply.Failure.Cipher), Characterization: reply.Failure.Characterization}
	}

	*nomination = reply.Nomination

	return nil
}
