package privatekey

import (
	"fmt"
	"time"

	"github.com/valkyrieworks/vault"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
	"github.com/valkyrieworks/kinds"
)

//
//
type ReprocessNotaryCustomer struct {
	following    *NotaryCustomer
	attempts int
	deadline time.Duration
}

//
//
func NewReprocessNotaryCustomer(sc *NotaryCustomer, attempts int, deadline time.Duration) *ReprocessNotaryCustomer {
	return &ReprocessNotaryCustomer{sc, attempts, deadline}
}

var _ kinds.PrivateRatifier = (*ReprocessNotaryCustomer)(nil)

func (sc *ReprocessNotaryCustomer) End() error {
	return sc.following.End()
}

func (sc *ReprocessNotaryCustomer) IsLinked() bool {
	return sc.following.IsLinked()
}

func (sc *ReprocessNotaryCustomer) WaitForLinkage(maximumWait time.Duration) error {
	return sc.following.WaitForLinkage(maximumWait)
}

//
//

func (sc *ReprocessNotaryCustomer) Ping() error {
	return sc.following.Ping()
}

func (sc *ReprocessNotaryCustomer) FetchPublicKey() (vault.PublicKey, error) {
	var (
		pk  vault.PublicKey
		err error
	)
	for i := 0; i < sc.attempts || sc.attempts == 0; i++ {
		pk, err = sc.following.FetchPublicKey()
		if err == nil {
			return pk, nil
		}
		//
		if _, ok := err.(*DistantNotaryFault); ok {
			return nil, err
		}
		time.Sleep(sc.deadline)
	}
	return nil, fmt.Errorf("REDACTED", err)
}

func (sc *ReprocessNotaryCustomer) AttestBallot(ledgerUID string, ballot *engineproto.Ballot) error {
	var err error
	for i := 0; i < sc.attempts || sc.attempts == 0; i++ {
		err = sc.following.AttestBallot(ledgerUID, ballot)
		if err == nil {
			return nil
		}
		//
		if _, ok := err.(*DistantNotaryFault); ok {
			return err
		}
		time.Sleep(sc.deadline)
	}
	return fmt.Errorf("REDACTED", err)
}

func (sc *ReprocessNotaryCustomer) AttestNomination(ledgerUID string, nomination *engineproto.Nomination) error {
	var err error
	for i := 0; i < sc.attempts || sc.attempts == 0; i++ {
		err = sc.following.AttestNomination(ledgerUID, nomination)
		if err == nil {
			return nil
		}
		//
		if _, ok := err.(*DistantNotaryFault); ok {
			return err
		}
		time.Sleep(sc.deadline)
	}
	return fmt.Errorf("REDACTED", err)
}
