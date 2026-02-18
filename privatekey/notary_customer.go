package privatekey

import (
	"fmt"
	"time"

	"github.com/valkyrieworks/vault"
	cometfaults "github.com/valkyrieworks/kinds/faults"

	cryptocode "github.com/valkyrieworks/vault/codec"
	privatekeyproto "github.com/valkyrieworks/schema/consensuscore/privatekey"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
	"github.com/valkyrieworks/kinds"
)

//
//
type NotaryCustomer struct {
	gateway *NotaryObserverTerminus
	ledgerUID  string
}

var _ kinds.PrivateRatifier = (*NotaryCustomer)(nil)

//
//
func NewNotaryCustomer(gateway *NotaryObserverTerminus, ledgerUID string) (*NotaryCustomer, error) {
	if !gateway.IsActive() {
		if err := gateway.Begin(); err != nil {
			return nil, fmt.Errorf("REDACTED", err)
		}
	}

	return &NotaryCustomer{gateway: gateway, ledgerUID: ledgerUID}, nil
}

//
func (sc *NotaryCustomer) End() error {
	return sc.gateway.End()
}

//
func (sc *NotaryCustomer) IsLinked() bool {
	return sc.gateway.IsLinked()
}

//
func (sc *NotaryCustomer) WaitForLinkage(maximumWait time.Duration) error {
	return sc.gateway.WaitForLinkage(maximumWait)
}

//
//

//
func (sc *NotaryCustomer) Ping() error {
	reply, err := sc.gateway.TransmitQuery(shouldEncloseMessage(&privatekeyproto.PingQuery{}))
	if err != nil {
		sc.gateway.Tracer.Fault("REDACTED", "REDACTED", err)
		return nil
	}

	pb := reply.FetchPingAnswer()
	if pb == nil {
		return err
	}

	return nil
}

//
//
func (sc *NotaryCustomer) FetchPublicKey() (vault.PublicKey, error) {
	reply, err := sc.gateway.TransmitQuery(shouldEncloseMessage(&privatekeyproto.PublicKeyQuery{SeriesUid: sc.ledgerUID}))
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	reply := reply.FetchPublicKeyAnswer()
	if reply == nil {
		return nil, cometfaults.ErrMandatoryField{Field: "REDACTED"}
	}
	if reply.Fault != nil {
		return nil, &DistantNotaryFault{Code: int(reply.Fault.Code), Summary: reply.Fault.Summary}
	}

	pk, err := cryptocode.PublicKeyFromSchema(reply.PublicKey)
	if err != nil {
		return nil, err
	}

	return pk, nil
}

//
func (sc *NotaryCustomer) AttestBallot(ledgerUID string, ballot *engineproto.Ballot) error {
	reply, err := sc.gateway.TransmitQuery(shouldEncloseMessage(&privatekeyproto.AttestBallotQuery{Ballot: ballot, SeriesUid: ledgerUID}))
	if err != nil {
		return err
	}

	reply := reply.FetchAttestedBallotAnswer()
	if reply == nil {
		return cometfaults.ErrMandatoryField{Field: "REDACTED"}
	}
	if reply.Fault != nil {
		return &DistantNotaryFault{Code: int(reply.Fault.Code), Summary: reply.Fault.Summary}
	}

	*ballot = reply.Ballot

	return nil
}

//
func (sc *NotaryCustomer) AttestNomination(ledgerUID string, nomination *engineproto.Nomination) error {
	reply, err := sc.gateway.TransmitQuery(shouldEncloseMessage(
		&privatekeyproto.AttestNominationQuery{Nomination: nomination, SeriesUid: ledgerUID},
	))
	if err != nil {
		return err
	}

	reply := reply.FetchAttestedNominationAnswer()
	if reply == nil {
		return cometfaults.ErrMandatoryField{Field: "REDACTED"}
	}
	if reply.Fault != nil {
		return &DistantNotaryFault{Code: int(reply.Fault.Code), Summary: reply.Fault.Summary}
	}

	*nomination = reply.Nomination

	return nil
}
