package privatevalue

import (
	"fmt"
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

//
//
type ReissueEndorserCustomer struct {
	following    *EndorserCustomer
	attempts int
	deadline time.Duration
}

//
//
func FreshReissueEndorserCustomer(sc *EndorserCustomer, attempts int, deadline time.Duration) *ReissueEndorserCustomer {
	return &ReissueEndorserCustomer{sc, attempts, deadline}
}

var _ kinds.PrivateAssessor = (*ReissueEndorserCustomer)(nil)

func (sc *ReissueEndorserCustomer) Shutdown() error {
	return sc.following.Shutdown()
}

func (sc *ReissueEndorserCustomer) EqualsAssociated() bool {
	return sc.following.EqualsAssociated()
}

func (sc *ReissueEndorserCustomer) PauseForeachLinkage(maximumPause time.Duration) error {
	return sc.following.PauseForeachLinkage(maximumPause)
}

//
//

func (sc *ReissueEndorserCustomer) Ping() error {
	return sc.following.Ping()
}

func (sc *ReissueEndorserCustomer) ObtainPublicToken() (security.PublicToken, error) {
	var (
		pk  security.PublicToken
		err error
	)
	for i := 0; i < sc.attempts || sc.attempts == 0; i++ {
		pk, err = sc.following.ObtainPublicToken()
		if err == nil {
			return pk, nil
		}
		//
		if _, ok := err.(*RemoteEndorserFailure); ok {
			return nil, err
		}
		time.Sleep(sc.deadline)
	}
	return nil, fmt.Errorf("REDACTED", err)
}

func (sc *ReissueEndorserCustomer) AttestBallot(successionUUID string, ballot *commitchema.Ballot) error {
	var err error
	for i := 0; i < sc.attempts || sc.attempts == 0; i++ {
		err = sc.following.AttestBallot(successionUUID, ballot)
		if err == nil {
			return nil
		}
		//
		if _, ok := err.(*RemoteEndorserFailure); ok {
			return err
		}
		time.Sleep(sc.deadline)
	}
	return fmt.Errorf("REDACTED", err)
}

func (sc *ReissueEndorserCustomer) AttestNomination(successionUUID string, nomination *commitchema.Nomination) error {
	var err error
	for i := 0; i < sc.attempts || sc.attempts == 0; i++ {
		err = sc.following.AttestNomination(successionUUID, nomination)
		if err == nil {
			return nil
		}
		//
		if _, ok := err.(*RemoteEndorserFailure); ok {
			return err
		}
		time.Sleep(sc.deadline)
	}
	return fmt.Errorf("REDACTED", err)
}
