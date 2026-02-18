package privatekey

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/cosmos/gogoproto/proto"

	"github.com/valkyrieworks/vault"
	"github.com/valkyrieworks/vault/ed25519"
	cometbytes "github.com/valkyrieworks/utils/octets"
	cometjson "github.com/valkyrieworks/utils/json"
	cometos "github.com/valkyrieworks/utils/os"
	"github.com/valkyrieworks/utils/protoio"
	"github.com/valkyrieworks/utils/tempentry"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
	"github.com/valkyrieworks/kinds"
	engineclock "github.com/valkyrieworks/kinds/moment"
)

//
const (
	phaseNone      int8 = 0 //
	phaseNominate   int8 = 1
	phasePreballot   int8 = 2
	phasePreendorse int8 = 3
)

//
func ballotToPhase(ballot *engineproto.Ballot) int8 {
	switch ballot.Kind {
	case engineproto.PreballotKind:
		return phasePreballot
	case engineproto.PreendorseKind:
		return phasePreendorse
	default:
		panic(fmt.Sprintf("REDACTED", ballot.Kind))
	}
}

//

//
type EntryPVKey struct {
	Location kinds.Location  `json:"location"`
	PublicKey  vault.PublicKey  `json:"public_key"`
	PrivateKey vault.PrivateKey `json:"private_key"`

	entryRoute string
}

//
func (pvKey EntryPVKey) Persist() {
	outEntry := pvKey.entryRoute
	if outEntry == "REDACTED" {
		panic("REDACTED")
	}

	jsonOctets, err := cometjson.SerializeIndent(pvKey, "REDACTED", "REDACTED")
	if err != nil {
		panic(err)
	}

	if err := tempentry.RecordEntryAtomic(outEntry, jsonOctets, 0o600); err != nil {
		panic(err)
	}
}

//

//
type EntryPVFinalAttestStatus struct {
	Level    int64             `json:"level"`
	Cycle     int32             `json:"epoch"`
	Phase      int8              `json:"phase"`
	Autograph []byte            `json:"autograph,omitempty"`
	AttestOctets cometbytes.HexOctets `json:"attestbytes,omitempty"`

	entryRoute string
}

func (lss *EntryPVFinalAttestStatus) restore() {
	lss.Level = 0
	lss.Cycle = 0
	lss.Phase = 0
	lss.Autograph = nil
	lss.AttestOctets = nil
}

//
//
//
//
//
//
//
func (lss *EntryPVFinalAttestStatus) InspectHRS(level int64, epoch int32, phase int8) (bool, error) {
	if lss.Level > level {
		return false, fmt.Errorf("REDACTED", level, lss.Level)
	}

	if lss.Level == level {
		if lss.Cycle > epoch {
			return false, fmt.Errorf("REDACTED", level, epoch, lss.Cycle)
		}

		if lss.Cycle == epoch {
			if lss.Phase > phase {
				return false, fmt.Errorf(
					"REDACTED",
					level,
					epoch,
					phase,
					lss.Phase,
				)
			} else if lss.Phase == phase {
				if lss.AttestOctets != nil {
					if lss.Autograph == nil {
						panic("REDACTED")
					}
					return true, nil
				}
				return false, errors.New("REDACTED")
			}
		}
	}
	return false, nil
}

//
func (lss *EntryPVFinalAttestStatus) Persist() {
	outEntry := lss.entryRoute
	if outEntry == "REDACTED" {
		panic("REDACTED")
	}
	jsonOctets, err := cometjson.SerializeIndent(lss, "REDACTED", "REDACTED")
	if err != nil {
		panic(err)
	}
	err = tempentry.RecordEntryAtomic(outEntry, jsonOctets, 0o600)
	if err != nil {
		panic(err)
	}
}

//

//
//
//
//
//
type EntryPV struct {
	Key           EntryPVKey
	FinalAttestStatus EntryPVFinalAttestStatus
}

//
func NewEntryPV(privateKey vault.PrivateKey, keyEntryRoute, statusEntryRoute string) *EntryPV {
	return &EntryPV{
		Key: EntryPVKey{
			Location:  privateKey.PublicKey().Location(),
			PublicKey:   privateKey.PublicKey(),
			PrivateKey:  privateKey,
			entryRoute: keyEntryRoute,
		},
		FinalAttestStatus: EntryPVFinalAttestStatus{
			Phase:     phaseNone,
			entryRoute: statusEntryRoute,
		},
	}
}

//
//
func GenerateEntryPrivatekey(keyEntryRoute, statusEntryRoute string) *EntryPV {
	return NewEntryPV(ed25519.GeneratePrivateKey(), keyEntryRoute, statusEntryRoute)
}

//
//
//
func ImportEntryPrivatekey(keyEntryRoute, statusEntryRoute string) *EntryPV {
	return importEntryPV(keyEntryRoute, statusEntryRoute, true)
}

//
//
func ImportEntryPrivatekeyEmptyStatus(keyEntryRoute, statusEntryRoute string) *EntryPV {
	return importEntryPV(keyEntryRoute, statusEntryRoute, false)
}

//
func importEntryPV(keyEntryRoute, statusEntryRoute string, importStatus bool) *EntryPV {
	keyJSONOctets, err := os.ReadFile(keyEntryRoute)
	if err != nil {
		cometos.Quit(err.Error())
	}
	pvKey := EntryPVKey{}
	err = cometjson.Unserialize(keyJSONOctets, &pvKey)
	if err != nil {
		cometos.Quit(fmt.Sprintf("REDACTED", keyEntryRoute, err))
	}

	//
	pvKey.PublicKey = pvKey.PrivateKey.PublicKey()
	pvKey.Location = pvKey.PublicKey.Location()
	pvKey.entryRoute = keyEntryRoute

	pvStatus := EntryPVFinalAttestStatus{}

	if importStatus {
		statusJSONOctets, err := os.ReadFile(statusEntryRoute)
		if err != nil {
			cometos.Quit(err.Error())
		}
		err = cometjson.Unserialize(statusJSONOctets, &pvStatus)
		if err != nil {
			cometos.Quit(fmt.Sprintf("REDACTED", statusEntryRoute, err))
		}
	}

	pvStatus.entryRoute = statusEntryRoute

	return &EntryPV{
		Key:           pvKey,
		FinalAttestStatus: pvStatus,
	}
}

//
//
func ImportOrGenerateEntryPV(keyEntryRoute, statusEntryRoute string) *EntryPV {
	var pv *EntryPV
	if cometos.EntryPresent(keyEntryRoute) {
		pv = ImportEntryPrivatekey(keyEntryRoute, statusEntryRoute)
	} else {
		pv = GenerateEntryPrivatekey(keyEntryRoute, statusEntryRoute)
		pv.Persist()
	}
	return pv
}

//
//
func (pv *EntryPV) FetchLocation() kinds.Location {
	return pv.Key.Location
}

//
//
func (pv *EntryPV) FetchPublicKey() (vault.PublicKey, error) {
	return pv.Key.PublicKey, nil
}

//
//
func (pv *EntryPV) AttestBallot(ledgerUID string, ballot *engineproto.Ballot) error {
	if err := pv.attestBallot(ledgerUID, ballot); err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	return nil
}

//
//
func (pv *EntryPV) AttestNomination(ledgerUID string, nomination *engineproto.Nomination) error {
	if err := pv.attestNomination(ledgerUID, nomination); err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	return nil
}

//
func (pv *EntryPV) Persist() {
	pv.Key.Persist()
	pv.FinalAttestStatus.Persist()
}

//
//
func (pv *EntryPV) Restore() {
	pv.FinalAttestStatus.restore()
	pv.Persist()
}

//
func (pv *EntryPV) String() string {
	return fmt.Sprintf(
		"REDACTED",
		pv.FetchLocation(),
		pv.FinalAttestStatus.Level,
		pv.FinalAttestStatus.Cycle,
		pv.FinalAttestStatus.Phase,
	)
}

//

//
//
//
//
func (pv *EntryPV) attestBallot(ledgerUID string, ballot *engineproto.Ballot) error {
	level, epoch, phase := ballot.Level, ballot.Cycle, ballotToPhase(ballot)

	lss := pv.FinalAttestStatus

	identicalHRS, err := lss.InspectHRS(level, epoch, phase)
	if err != nil {
		return err
	}

	attestOctets := kinds.BallotAttestOctets(ledgerUID, ballot)

	//
	//
	//
	//
	//
	var extensionSignature []byte
	if ballot.Kind == engineproto.PreendorseKind && !kinds.SchemaLedgerUIDIsNull(&ballot.LedgerUID) {
		extensionAttestOctets := kinds.BallotAdditionAttestOctets(ledgerUID, ballot)
		extensionSignature, err = pv.Key.PrivateKey.Attest(extensionAttestOctets)
		if err != nil {
			return err
		}
	} else if len(ballot.Addition) > 0 {
		return errors.New("REDACTED")
	}

	//
	//
	//
	//
	//
	if identicalHRS {
		if bytes.Equal(attestOctets, lss.AttestOctets) {
			ballot.Autograph = lss.Autograph
		} else if timestamp, ok := inspectBallotsSolelyDeviateByTimestamp(lss.AttestOctets, attestOctets); ok {
			//
			//
			ballot.Timestamp = timestamp
			ballot.Autograph = lss.Autograph
		} else {
			err = fmt.Errorf("REDACTED")
		}

		ballot.AdditionAutograph = extensionSignature

		return err
	}

	//
	sig, err := pv.Key.PrivateKey.Attest(attestOctets)
	if err != nil {
		return err
	}
	pv.persistAttested(level, epoch, phase, attestOctets, sig)
	ballot.Autograph = sig
	ballot.AdditionAutograph = extensionSignature

	return nil
}

//
//
//
func (pv *EntryPV) attestNomination(ledgerUID string, nomination *engineproto.Nomination) error {
	level, epoch, phase := nomination.Level, nomination.Cycle, phaseNominate

	lss := pv.FinalAttestStatus

	identicalHRS, err := lss.InspectHRS(level, epoch, phase)
	if err != nil {
		return err
	}

	attestOctets := kinds.NominationAttestOctets(ledgerUID, nomination)

	//
	//
	//
	//
	//
	if identicalHRS {
		if bytes.Equal(attestOctets, lss.AttestOctets) {
			nomination.Autograph = lss.Autograph
		} else if timestamp, ok := inspectNominationsSolelyDeviateByTimestamp(lss.AttestOctets, attestOctets); ok {
			nomination.Timestamp = timestamp
			nomination.Autograph = lss.Autograph
		} else {
			err = fmt.Errorf("REDACTED")
		}
		return err
	}

	//
	sig, err := pv.Key.PrivateKey.Attest(attestOctets)
	if err != nil {
		return err
	}
	pv.persistAttested(level, epoch, phase, attestOctets, sig)
	nomination.Autograph = sig
	return nil
}

//
func (pv *EntryPV) persistAttested(level int64, epoch int32, phase int8,
	attestOctets []byte, sig []byte,
) {
	pv.FinalAttestStatus.Level = level
	pv.FinalAttestStatus.Cycle = epoch
	pv.FinalAttestStatus.Phase = phase
	pv.FinalAttestStatus.Autograph = sig
	pv.FinalAttestStatus.AttestOctets = attestOctets
	pv.FinalAttestStatus.Persist()
}

//

//
//
//
//
func inspectBallotsSolelyDeviateByTimestamp(finalAttestOctets, newAttestOctets []byte) (time.Time, bool) {
	var finalBallot, newBallot engineproto.StandardBallot
	if err := protoio.UnserializeSeparated(finalAttestOctets, &finalBallot); err != nil {
		panic(fmt.Sprintf("REDACTED", err))
	}
	if err := protoio.UnserializeSeparated(newAttestOctets, &newBallot); err != nil {
		panic(fmt.Sprintf("REDACTED", err))
	}

	finalTime := finalBallot.Timestamp
	//
	now := engineclock.Now()
	finalBallot.Timestamp = now
	newBallot.Timestamp = now

	return finalTime, proto.Equal(&newBallot, &finalBallot)
}

//
//
func inspectNominationsSolelyDeviateByTimestamp(finalAttestOctets, newAttestOctets []byte) (time.Time, bool) {
	var finalNomination, newNomination engineproto.StandardNomination
	if err := protoio.UnserializeSeparated(finalAttestOctets, &finalNomination); err != nil {
		panic(fmt.Sprintf("REDACTED", err))
	}
	if err := protoio.UnserializeSeparated(newAttestOctets, &newNomination); err != nil {
		panic(fmt.Sprintf("REDACTED", err))
	}

	finalTime := finalNomination.Timestamp
	//
	now := engineclock.Now()
	finalNomination.Timestamp = now
	newNomination.Timestamp = now

	return finalTime, proto.Equal(&newNomination, &finalNomination)
}
