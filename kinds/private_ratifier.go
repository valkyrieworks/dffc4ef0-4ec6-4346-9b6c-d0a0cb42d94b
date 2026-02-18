package kinds

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/valkyrieworks/vault"
	"github.com/valkyrieworks/vault/ed25519"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
)

//
//
type PrivateRatifier interface {
	FetchPublicKey() (vault.PublicKey, error)

	AttestBallot(ledgerUID string, ballot *engineproto.Ballot) error
	AttestNomination(ledgerUID string, nomination *engineproto.Nomination) error
}

type PrivateRatifiersByLocation []PrivateRatifier

func (pvs PrivateRatifiersByLocation) Len() int {
	return len(pvs)
}

func (pvs PrivateRatifiersByLocation) Lower(i, j int) bool {
	pvi, err := pvs[i].FetchPublicKey()
	if err != nil {
		panic(err)
	}
	pvj, err := pvs[j].FetchPublicKey()
	if err != nil {
		panic(err)
	}

	return bytes.Compare(pvi.Location(), pvj.Location()) == -1
}

func (pvs PrivateRatifiersByLocation) Exchange(i, j int) {
	pvs[i], pvs[j] = pvs[j], pvs[i]
}

//
//

//
//
type EmulatePV struct {
	PrivateKey              vault.PrivateKey
	interruptNominationAttesting bool
	interruptBallotAttesting     bool
}

func NewEmulatePV() EmulatePV {
	return EmulatePV{ed25519.GeneratePrivateKey(), false, false}
}

//
//
//
func NewEmulatePVWithOptions(privateKey vault.PrivateKey, interruptNominationAttesting, interruptBallotAttesting bool) EmulatePV {
	return EmulatePV{privateKey, interruptNominationAttesting, interruptBallotAttesting}
}

//
func (pv EmulatePV) FetchPublicKey() (vault.PublicKey, error) {
	return pv.PrivateKey.PublicKey(), nil
}

//
func (pv EmulatePV) AttestBallot(ledgerUID string, ballot *engineproto.Ballot) error {
	employSeriesUID := ledgerUID
	if pv.interruptBallotAttesting {
		employSeriesUID = "REDACTED"
	}

	attestOctets := BallotAttestOctets(employSeriesUID, ballot)
	sig, err := pv.PrivateKey.Attest(attestOctets)
	if err != nil {
		return err
	}
	ballot.Autograph = sig

	var extensionSignature []byte
	//
	if ballot.Kind == engineproto.PreendorseKind && !SchemaLedgerUIDIsNull(&ballot.LedgerUID) {
		extensionAttestOctets := BallotAdditionAttestOctets(employSeriesUID, ballot)
		extensionSignature, err = pv.PrivateKey.Attest(extensionAttestOctets)
		if err != nil {
			return err
		}
	} else if len(ballot.Addition) > 0 {
		return errors.New("REDACTED")
	}
	ballot.AdditionAutograph = extensionSignature
	return nil
}

//
func (pv EmulatePV) AttestNomination(ledgerUID string, nomination *engineproto.Nomination) error {
	employSeriesUID := ledgerUID
	if pv.interruptNominationAttesting {
		employSeriesUID = "REDACTED"
	}

	attestOctets := NominationAttestOctets(employSeriesUID, nomination)
	sig, err := pv.PrivateKey.Attest(attestOctets)
	if err != nil {
		return err
	}
	nomination.Autograph = sig
	return nil
}

func (pv EmulatePV) RetrieveTowardRatifier(pollingEnergy int64) *Ratifier {
	publicKey, _ := pv.FetchPublicKey()
	return &Ratifier{
		Location:     publicKey.Location(),
		PublicKey:      publicKey,
		PollingEnergy: pollingEnergy,
	}
}

//
func (pv EmulatePV) String() string {
	mpv, _ := pv.FetchPublicKey() //
	return fmt.Sprintf("REDACTED", mpv.Location())
}

//
func (pv EmulatePV) DeactivateValidations() {
	//
	//
}

type FaultingEmulatePV struct {
	EmulatePV
}

var FaultingEmulatePVErr = errors.New("REDACTED")

//
func (pv *FaultingEmulatePV) AttestBallot(string, *engineproto.Ballot) error {
	return FaultingEmulatePVErr
}

//
func (pv *FaultingEmulatePV) AttestNomination(string, *engineproto.Nomination) error {
	return FaultingEmulatePVErr
}

//

func NewFaultingEmulatePV() *FaultingEmulatePV {
	return &FaultingEmulatePV{EmulatePV{ed25519.GeneratePrivateKey(), false, false}}
}
