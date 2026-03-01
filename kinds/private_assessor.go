package kinds

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/edwards25519"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
)

//
//
type PrivateAssessor interface {
	ObtainPublicToken() (security.PublicToken, error)

	AttestBallot(successionUUID string, ballot *commitchema.Ballot) error
	AttestNomination(successionUUID string, nomination *commitchema.Nomination) error
}

type PrivateAssessorsViaLocation []PrivateAssessor

func (pvs PrivateAssessorsViaLocation) Len() int {
	return len(pvs)
}

func (pvs PrivateAssessorsViaLocation) Inferior(i, j int) bool {
	pvi, err := pvs[i].ObtainPublicToken()
	if err != nil {
		panic(err)
	}
	pvj, err := pvs[j].ObtainPublicToken()
	if err != nil {
		panic(err)
	}

	return bytes.Compare(pvi.Location(), pvj.Location()) == -1
}

func (pvs PrivateAssessorsViaLocation) Exchange(i, j int) {
	pvs[i], pvs[j] = pvs[j], pvs[i]
}

//
//

//
//
type SimulatePRV struct {
	PrivateToken              security.PrivateToken
	breachNominationNotating bool
	breachBallotNotating     bool
}

func FreshSimulatePRV() SimulatePRV {
	return SimulatePRV{edwards25519.ProducePrivateToken(), false, false}
}

//
//
//
func FreshSimulatePRVUsingParameters(privateToken security.PrivateToken, breachNominationNotating, breachBallotNotating bool) SimulatePRV {
	return SimulatePRV{privateToken, breachNominationNotating, breachBallotNotating}
}

//
func (pv SimulatePRV) ObtainPublicToken() (security.PublicToken, error) {
	return pv.PrivateToken.PublicToken(), nil
}

//
func (pv SimulatePRV) AttestBallot(successionUUID string, ballot *commitchema.Ballot) error {
	utilizeSuccessionUUID := successionUUID
	if pv.breachBallotNotating {
		utilizeSuccessionUUID = "REDACTED"
	}

	attestOctets := BallotAttestOctets(utilizeSuccessionUUID, ballot)
	sig, err := pv.PrivateToken.Attest(attestOctets)
	if err != nil {
		return err
	}
	ballot.Notation = sig

	var addnSignature []byte
	//
	if ballot.Kind == commitchema.PreendorseKind && !SchemaLedgerUUIDEqualsVoid(&ballot.LedgerUUID) {
		addnAttestOctets := BallotAdditionAttestOctets(utilizeSuccessionUUID, ballot)
		addnSignature, err = pv.PrivateToken.Attest(addnAttestOctets)
		if err != nil {
			return err
		}
	} else if len(ballot.Addition) > 0 {
		return errors.New("REDACTED")
	}
	ballot.AdditionNotation = addnSignature
	return nil
}

//
func (pv SimulatePRV) AttestNomination(successionUUID string, nomination *commitchema.Nomination) error {
	utilizeSuccessionUUID := successionUUID
	if pv.breachNominationNotating {
		utilizeSuccessionUUID = "REDACTED"
	}

	attestOctets := NominationAttestOctets(utilizeSuccessionUUID, nomination)
	sig, err := pv.PrivateToken.Attest(attestOctets)
	if err != nil {
		return err
	}
	nomination.Notation = sig
	return nil
}

func (pv SimulatePRV) DeriveWithinAssessor(ballotingPotency int64) *Assessor {
	publicToken, _ := pv.ObtainPublicToken()
	return &Assessor{
		Location:     publicToken.Location(),
		PublicToken:      publicToken,
		BallotingPotency: ballotingPotency,
	}
}

//
func (pv SimulatePRV) Text() string {
	mpv, _ := pv.ObtainPublicToken() //
	return fmt.Sprintf("REDACTED", mpv.Location())
}

//
func (pv SimulatePRV) DeactivateVerifications() {
	//
	//
}

type FaultingSimulatePRV struct {
	SimulatePRV
}

var FaultingSimulatePRVFault = errors.New("REDACTED")

//
func (pv *FaultingSimulatePRV) AttestBallot(string, *commitchema.Ballot) error {
	return FaultingSimulatePRVFault
}

//
func (pv *FaultingSimulatePRV) AttestNomination(string, *commitchema.Nomination) error {
	return FaultingSimulatePRVFault
}

//

func FreshFaultingSimulatePRV() *FaultingSimulatePRV {
	return &FaultingSimulatePRV{SimulatePRV{edwards25519.ProducePrivateToken(), false, false}}
}
