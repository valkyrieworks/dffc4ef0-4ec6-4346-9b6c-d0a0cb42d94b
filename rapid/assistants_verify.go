package rapid_test

import (
	"time"

	"github.com/valkyrieworks/vault"
	"github.com/valkyrieworks/vault/ed25519"
	"github.com/valkyrieworks/vault/comethash"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
	cometrelease "github.com/valkyrieworks/schema/consensuscore/release"
	"github.com/valkyrieworks/kinds"
	engineclock "github.com/valkyrieworks/kinds/moment"
	"github.com/valkyrieworks/release"
)

//
//
//
//
//
//
//
type privateKeys []vault.PrivateKey

//
func generatePrivateKeys(n int) privateKeys {
	res := make(privateKeys, n)
	for i := range res {
		res[i] = ed25519.GeneratePrivateKey()
	}
	return res
}

//
//
//
//
//
//
//

//
func (pkz privateKeys) Expand(n int) privateKeys {
	surplus := generatePrivateKeys(n)
	return append(pkz, surplus...)
}

//
//
//
//
//
//
//
//

//
//
//
//
//

//
//
//
//
func (pkz privateKeys) ToRatifiers(init, inc int64) *kinds.RatifierAssign {
	res := make([]*kinds.Ratifier, len(pkz))
	for i, k := range pkz {
		res[i] = kinds.NewRatifier(k.PublicKey(), init+int64(i)*inc)
	}
	return kinds.NewRatifierCollection(res)
}

//
func (pkz privateKeys) attestHeading(heading *kinds.Heading, valueCollection *kinds.RatifierAssign, initial, final int) *kinds.Endorse {
	endorseAutographs := make([]kinds.EndorseSignature, len(pkz))
	for i := 0; i < len(pkz); i++ {
		endorseAutographs[i] = kinds.NewEndorseSignatureMissing()
	}

	ledgerUID := kinds.LedgerUID{
		Digest:          heading.Digest(),
		SegmentAssignHeading: kinds.SegmentAssignHeading{Sum: 1, Digest: vault.CRandomOctets(32)},
	}

	//
	for i := initial; i < final && i < len(pkz); i++ {
		ballot := createBallot(heading, valueCollection, pkz[i], ledgerUID)
		endorseAutographs[ballot.RatifierOrdinal] = ballot.EndorseSignature()
	}

	return &kinds.Endorse{
		Level:     heading.Level,
		Cycle:      1,
		LedgerUID:    ledgerUID,
		Endorsements: endorseAutographs,
	}
}

func createBallot(heading *kinds.Heading, ratifierset *kinds.RatifierAssign,
	key vault.PrivateKey, ledgerUID kinds.LedgerUID,
) *kinds.Ballot {
	address := key.PublicKey().Location()
	idx, _ := ratifierset.FetchByLocation(address)
	ballot := &kinds.Ballot{
		RatifierLocation: address,
		RatifierOrdinal:   idx,
		Level:           heading.Level,
		Cycle:            1,
		Timestamp:        engineclock.Now(),
		Kind:             engineproto.PreendorseKind,
		LedgerUID:          ledgerUID,
	}

	v := ballot.ToSchema()
	//
	attestOctets := kinds.BallotAttestOctets(heading.LedgerUID, v)
	sig, err := key.Attest(attestOctets)
	if err != nil {
		panic(err)
	}
	ballot.Autograph = sig

	extensionAttestOctets := kinds.BallotAdditionAttestOctets(heading.LedgerUID, v)
	extensionSignature, err := key.Attest(extensionAttestOctets)
	if err != nil {
		panic(err)
	}
	ballot.AdditionAutograph = extensionSignature

	return ballot
}

func generateHeading(ledgerUID string, level int64, byteTime time.Time, txs kinds.Txs,
	ratifierset, followingRatifierset *kinds.RatifierAssign, applicationDigest, constDigest, outputDigest []byte,
) *kinds.Heading {
	return &kinds.Heading{
		Release: cometrelease.Agreement{Ledger: release.LedgerProtocol, App: 0},
		LedgerUID: ledgerUID,
		Level:  level,
		Time:    byteTime,
		//
		//
		RatifiersDigest:     ratifierset.Digest(),
		FollowingRatifiersDigest: followingRatifierset.Digest(),
		DataDigest:           txs.Digest(),
		ApplicationDigest:            applicationDigest,
		AgreementDigest:      constDigest,
		FinalOutcomesDigest:    outputDigest,
		RecommenderLocation:    ratifierset.Ratifiers[0].Location,
	}
}

//
func (pkz privateKeys) GenerateAttestedHeading(ledgerUID string, level int64, byteTime time.Time, txs kinds.Txs,
	ratifierset, followingRatifierset *kinds.RatifierAssign, applicationDigest, constDigest, outputDigest []byte, initial, final int,
) *kinds.AttestedHeading {
	heading := generateHeading(ledgerUID, level, byteTime, txs, ratifierset, followingRatifierset, applicationDigest, constDigest, outputDigest)
	return &kinds.AttestedHeading{
		Heading: heading,
		Endorse: pkz.attestHeading(heading, ratifierset, initial, final),
	}
}

//
func (pkz privateKeys) GenerateAttestedHeadingFinalLedgerUID(ledgerUID string, level int64, byteTime time.Time, txs kinds.Txs,
	ratifierset, followingRatifierset *kinds.RatifierAssign, applicationDigest, constDigest, outputDigest []byte, initial, final int,
	finalLedgerUID kinds.LedgerUID,
) *kinds.AttestedHeading {
	heading := generateHeading(ledgerUID, level, byteTime, txs, ratifierset, followingRatifierset, applicationDigest, constDigest, outputDigest)
	heading.FinalLedgerUID = finalLedgerUID
	return &kinds.AttestedHeading{
		Heading: heading,
		Endorse: pkz.attestHeading(heading, ratifierset, initial, final),
	}
}

func (pkz privateKeys) AlterKeys(variance int) privateKeys {
	newKeys := pkz[variance:]
	return newKeys.Expand(variance)
}

//
//
//
func generateEmulateMemberWithKeys(
	ledgerUID string,
	ledgerVolume int64,
	valueVolume int,
	valueDeviation float32,
	byteTime time.Time) (
	map[int64]*kinds.AttestedHeading,
	map[int64]*kinds.RatifierAssign,
	map[int64]privateKeys,
) {
	var (
		headings         = make(map[int64]*kinds.AttestedHeading, ledgerVolume)
		ratifierset          = make(map[int64]*kinds.RatifierAssign, ledgerVolume+1)
		keyindex          = make(map[int64]privateKeys, ledgerVolume+1)
		keys            = generatePrivateKeys(valueVolume)
		sumDeviation  = valueDeviation
		valueDeviationInteger int
		newKeys         privateKeys
	)

	valueDeviationInteger = int(sumDeviation)
	sumDeviation = -float32(valueDeviationInteger)
	newKeys = keys.AlterKeys(valueDeviationInteger)
	keyindex[1] = keys
	keyindex[2] = newKeys

	//
	finalHeading := keys.GenerateAttestedHeading(ledgerUID, 1, byteTime.Add(1*time.Minute), nil,
		keys.ToRatifiers(2, 0), newKeys.ToRatifiers(2, 0), digest("REDACTED"), digest("REDACTED"),
		digest("REDACTED"), 0, len(keys))
	ongoingHeading := finalHeading
	headings[1] = ongoingHeading
	ratifierset[1] = keys.ToRatifiers(2, 0)
	keys = newKeys

	for level := int64(2); level <= ledgerVolume; level++ {
		sumDeviation += valueDeviation
		valueDeviationInteger = int(sumDeviation)
		sumDeviation = -float32(valueDeviationInteger)
		newKeys = keys.AlterKeys(valueDeviationInteger)
		ongoingHeading = keys.GenerateAttestedHeadingFinalLedgerUID(ledgerUID, level, byteTime.Add(time.Duration(level)*time.Minute),
			nil,
			keys.ToRatifiers(2, 0), newKeys.ToRatifiers(2, 0), digest("REDACTED"), digest("REDACTED"),
			digest("REDACTED"), 0, len(keys), kinds.LedgerUID{Digest: finalHeading.Digest()})
		headings[level] = ongoingHeading
		ratifierset[level] = keys.ToRatifiers(2, 0)
		finalHeading = ongoingHeading
		keys = newKeys
		keyindex[level+1] = keys
	}

	return headings, ratifierset, keyindex
}

func generateEmulateMember(
	ledgerUID string,
	ledgerVolume int64,
	valueVolume int,
	valueDeviation float32,
	byteTime time.Time) (
	string,
	map[int64]*kinds.AttestedHeading,
	map[int64]*kinds.RatifierAssign,
) {
	headings, ratifierset, _ := generateEmulateMemberWithKeys(ledgerUID, ledgerVolume, valueVolume, valueDeviation, byteTime)
	return ledgerUID, headings, ratifierset
}

func digest(s string) []byte {
	return comethash.Sum([]byte(s))
}
