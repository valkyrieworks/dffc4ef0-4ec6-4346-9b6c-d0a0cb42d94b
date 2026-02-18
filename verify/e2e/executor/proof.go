package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/valkyrieworks/vault"
	"github.com/valkyrieworks/vault/comethash"
	"github.com/valkyrieworks/intrinsic/verify"
	cometjson "github.com/valkyrieworks/utils/json"
	"github.com/valkyrieworks/privatekey"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
	cometrelease "github.com/valkyrieworks/schema/consensuscore/release"
	e2e "github.com/valkyrieworks/verify/e2e/pkg"
	"github.com/valkyrieworks/kinds"
	"github.com/valkyrieworks/release"
)

//
const rapidCustomerProofProportion = 4

//
//
//
//
func InsertProof(ctx context.Context, r *rand.Rand, verifychain *e2e.Verifychain, quantity int) error {
	//
	var objectiveMember *e2e.Member

	for _, idx := range r.Perm(len(verifychain.Instances)) {
		objectiveMember = verifychain.Instances[idx]

		if objectiveMember.Style == e2e.StyleOrigin || objectiveMember.Style == e2e.StyleRapid {
			objectiveMember = nil
			continue
		}

		break
	}

	if objectiveMember == nil {
		return errors.New("REDACTED")
	}

	tracer.Details(fmt.Sprintf("REDACTED", objectiveMember.Label, quantity))

	customer, err := objectiveMember.Customer()
	if err != nil {
		return err
	}

	//
	ledgerOutput, err := customer.Ledger(ctx, nil)
	if err != nil {
		return err
	}
	proofLevel := ledgerOutput.Ledger.Level
	waitLevel := ledgerOutput.Ledger.Level + 3

	nRatifiers := 100
	valueOutput, err := customer.Ratifiers(ctx, &proofLevel, nil, &nRatifiers)
	if err != nil {
		return err
	}

	valueCollection, err := kinds.RatifierCollectionFromPresentRatifiers(valueOutput.Ratifiers)
	if err != nil {
		return err
	}

	//
	privateValues, err := fetchPrivateRatifierKeys(verifychain)
	if err != nil {
		return err
	}

	//
	//
	_, err = waitForMember(ctx, objectiveMember, waitLevel, time.Minute)
	if err != nil {
		return err
	}

	var ev kinds.Proof
	for i := 0; i < quantity; i++ {
		soundEvt := true
		if i%rapidCustomerProofProportion == 0 {
			soundEvt = i%(rapidCustomerProofProportion*2) != 0 //
			ev, err = composeRapidCustomerAssaultProof(
				ctx, privateValues, proofLevel, valueCollection, verifychain.Label, ledgerOutput.Ledger.Time, soundEvt,
			)
		} else {
			var dve *kinds.ReplicatedBallotProof
			dve, err = composeReplicatedBallotProof(
				privateValues, proofLevel, valueCollection, verifychain.Label, ledgerOutput.Ledger.Time,
			)
			if dve.BallotA.Level < verifychain.BallotPluginsActivateLevel {
				dve.BallotA.Addition = nil
				dve.BallotA.AdditionAutograph = nil
				dve.BallotBYTE.Addition = nil
				dve.BallotBYTE.AdditionAutograph = nil
			}
			ev = dve
		}
		if err != nil {
			return err
		}

		_, err := customer.MulticastProof(ctx, ev)
		if !soundEvt {
			//
			//
			quantity++
		}
		if soundEvt != (err == nil) {
			if err == nil {
				return errors.New("REDACTED")
			}
			return err
		}
		time.Sleep(5 * time.Second / time.Duration(quantity))
	}

	//
	//
	_, err = waitForMember(ctx, objectiveMember, ledgerOutput.Ledger.Level+2, 30*time.Second)
	if err != nil {
		return err
	}

	tracer.Details(fmt.Sprintf("REDACTED", ledgerOutput.Ledger.Level+2))

	return nil
}

func fetchPrivateRatifierKeys(verifychain *e2e.Verifychain) ([]kinds.EmulatePV, error) {
	privateValues := []kinds.EmulatePV{}

	for _, member := range verifychain.Instances {
		if member.Style == e2e.StyleRatifier {
			privateKeyRoute := filepath.Join(verifychain.Dir, member.Label, PrivatekeyKeyEntry)
			privateKey, err := fetchPrivateKey(privateKeyRoute)
			if err != nil {
				return nil, err
			}
			//
			//
			privateValues = append(privateValues, kinds.NewEmulatePVWithOptions(privateKey, false, false))
		}
	}

	return privateValues, nil
}

//
//
func composeRapidCustomerAssaultProof(
	ctx context.Context,
	privateValues []kinds.EmulatePV,
	level int64,
	values *kinds.RatifierAssign,
	ledgerUID string,
	evtTime time.Time,
	soundProof bool,
) (*kinds.RapidCustomerAssaultProof, error) {
	//
	falsifiedLevel := level + 2
	falsifiedTime := evtTime.Add(1 * time.Second)
	heading := createHeadingArbitrary(ledgerUID, falsifiedLevel)
	heading.Time = falsifiedTime

	//
	//
	pv, clashingValues, err := transformRatifierCollection(ctx, privateValues, values, !soundProof)
	if err != nil {
		return nil, err
	}

	heading.RatifiersDigest = clashingValues.Digest()

	//
	ledgerUID := createLedgerUID(heading.Digest(), 1000, []byte("REDACTED"))
	ballotCollection := kinds.NewBallotCollection(ledgerUID, falsifiedLevel, 0, engineproto.AttestedMessageKind(2), clashingValues)
	endorse, err := verify.CreateEndorseFromBallotCollection(ledgerUID, ballotCollection, pv, falsifiedTime)
	if err != nil {
		return nil, err
	}

	//
	if !soundProof {
		endorse.Endorsements[len(endorse.Endorsements)-1].Autograph[0]++
	}

	ev := &kinds.RapidCustomerAssaultProof{
		ClashingLedger: &kinds.RapidLedger{
			AttestedHeading: &kinds.AttestedHeading{
				Heading: heading,
				Endorse: endorse,
			},
			RatifierAssign: clashingValues,
		},
		SharedLevel:     level,
		SumPollingEnergy: values.SumPollingEnergy(),
		Timestamp:        evtTime,
	}
	ev.FaultyRatifiers = ev.FetchFaultyRatifiers(values, &kinds.AttestedHeading{
		Heading: createHeadingArbitrary(ledgerUID, falsifiedLevel),
	})
	return ev, nil
}

//
//
func composeReplicatedBallotProof(
	privateValues []kinds.EmulatePV,
	level int64,
	values *kinds.RatifierAssign,
	ledgerUID string,
	moment time.Time,
) (*kinds.ReplicatedBallotProof, error) {
	privateValue, valueIdx, err := fetchArbitraryRatifierOrdinal(privateValues, values)
	if err != nil {
		return nil, err
	}
	ballotA, err := kinds.CreateBallot(privateValue, ledgerUID, valueIdx, level, 0, 2, createArbitraryLedgerUID(), moment)
	if err != nil {
		return nil, err
	}
	ballotBYTE, err := kinds.CreateBallot(privateValue, ledgerUID, valueIdx, level, 0, 2, createArbitraryLedgerUID(), moment)
	if err != nil {
		return nil, err
	}
	ev, err := kinds.NewReplicatedBallotProof(ballotA, ballotBYTE, moment, values)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	return ev, nil
}

//
//
func fetchArbitraryRatifierOrdinal(privateValues []kinds.EmulatePV, values *kinds.RatifierAssign) (kinds.EmulatePV, int32, error) {
	for _, idx := range rand.Perm(len(privateValues)) {
		pv := privateValues[idx]
		valueIdx, _ := values.FetchByLocation(pv.PrivateKey.PublicKey().Location())
		if valueIdx >= 0 {
			return pv, valueIdx, nil
		}
	}
	return kinds.EmulatePV{}, -1, errors.New("REDACTED")
}

func fetchPrivateKey(keyEntryRoute string) (vault.PrivateKey, error) {
	keyJSONOctets, err := os.ReadFile(keyEntryRoute)
	if err != nil {
		return nil, err
	}
	pvKey := privatekey.EntryPVKey{}
	err = cometjson.Unserialize(keyJSONOctets, &pvKey)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", keyEntryRoute, err)
	}

	return pvKey.PrivateKey, nil
}

func createHeadingArbitrary(ledgerUID string, level int64) *kinds.Heading {
	return &kinds.Heading{
		Release:            cometrelease.Agreement{Ledger: release.LedgerProtocol, App: 1},
		LedgerUID:            ledgerUID,
		Level:             level,
		Time:               time.Now(),
		FinalLedgerUID:        createLedgerUID([]byte("REDACTED"), 1000, []byte("REDACTED")),
		FinalEndorseDigest:     vault.CRandomOctets(comethash.Volume),
		DataDigest:           vault.CRandomOctets(comethash.Volume),
		RatifiersDigest:     vault.CRandomOctets(comethash.Volume),
		FollowingRatifiersDigest: vault.CRandomOctets(comethash.Volume),
		AgreementDigest:      vault.CRandomOctets(comethash.Volume),
		ApplicationDigest:            vault.CRandomOctets(comethash.Volume),
		FinalOutcomesDigest:    vault.CRandomOctets(comethash.Volume),
		ProofDigest:       vault.CRandomOctets(comethash.Volume),
		RecommenderLocation:    vault.CRandomOctets(vault.LocationVolume),
	}
}

func createArbitraryLedgerUID() kinds.LedgerUID {
	return createLedgerUID(vault.CRandomOctets(comethash.Volume), 100, vault.CRandomOctets(comethash.Volume))
}

func createLedgerUID(digest []byte, sectionCollectionVolume uint32, sectionCollectionDigest []byte) kinds.LedgerUID {
	var (
		h   = make([]byte, comethash.Volume)
		psH = make([]byte, comethash.Volume)
	)
	copy(h, digest)
	copy(psH, sectionCollectionDigest)
	return kinds.LedgerUID{
		Digest: h,
		SegmentAssignHeading: kinds.SegmentAssignHeading{
			Sum: sectionCollectionVolume,
			Digest:  psH,
		},
	}
}

func transformRatifierCollection(
	ctx context.Context,
	privateValues []kinds.EmulatePV,
	values *kinds.RatifierAssign,
	nop bool,
) ([]kinds.PrivateRatifier, *kinds.RatifierAssign, error) {
	newValue, newPrivateValue, err := verify.Ratifier(ctx, 10)
	if err != nil {
		return nil, nil, err
	}

	var newValues *kinds.RatifierAssign
	if nop {
		newValues = kinds.NewRatifierCollection(values.Clone().Ratifiers)
	} else {
		if values.Volume() > 2 {
			newValues = kinds.NewRatifierCollection(append(values.Clone().Ratifiers[:values.Volume()-1], newValue))
		} else {
			newValues = kinds.NewRatifierCollection(append(values.Clone().Ratifiers, newValue))
		}
	}

	//
	pv := make([]kinds.PrivateRatifier, newValues.Volume())
	for idx, val := range newValues.Ratifiers {
		located := false
		for _, p := range append(privateValues, newPrivateValue.(kinds.EmulatePV)) {
			if bytes.Equal(p.PrivateKey.PublicKey().Location(), val.Location) {
				pv[idx] = p
				located = true
				break
			}
		}
		if !located {
			return nil, nil, fmt.Errorf("REDACTED", val.Location)
		}
	}

	return pv, newValues, nil
}
