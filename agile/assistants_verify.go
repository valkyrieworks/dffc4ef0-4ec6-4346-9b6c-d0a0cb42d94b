package agile_test

import (
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/edwards25519"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/tenderminthash"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
	strongmindedition "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/edition"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
	committime "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds/moment"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/edition"
)

//
//
//
//
//
//
//
type privateTokens []security.PrivateToken

//
func producePrivateTokens(n int) privateTokens {
	res := make(privateTokens, n)
	for i := range res {
		res[i] = edwards25519.ProducePrivateToken()
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
func (pkz privateTokens) Broaden(n int) privateTokens {
	surplus := producePrivateTokens(n)
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
func (pkz privateTokens) TowardAssessors(initialize, inc int64) *kinds.AssessorAssign {
	res := make([]*kinds.Assessor, len(pkz))
	for i, k := range pkz {
		res[i] = kinds.FreshAssessor(k.PublicToken(), initialize+int64(i)*inc)
	}
	return kinds.FreshAssessorAssign(res)
}

//
func (pkz privateTokens) attestHeadline(heading *kinds.Heading, itemAssign *kinds.AssessorAssign, initial, final int) *kinds.Endorse {
	endorseSignatures := make([]kinds.EndorseSignature, len(pkz))
	for i := 0; i < len(pkz); i++ {
		endorseSignatures[i] = kinds.FreshEndorseSignatureMissing()
	}

	ledgerUUID := kinds.LedgerUUID{
		Digest:          heading.Digest(),
		FragmentAssignHeading: kinds.FragmentAssignHeading{Sum: 1, Digest: security.CHARArbitraryOctets(32)},
	}

	//
	for i := initial; i < final && i < len(pkz); i++ {
		ballot := createBallot(heading, itemAssign, pkz[i], ledgerUUID)
		endorseSignatures[ballot.AssessorOrdinal] = ballot.EndorseSignature()
	}

	return &kinds.Endorse{
		Altitude:     heading.Altitude,
		Iteration:      1,
		LedgerUUID:    ledgerUUID,
		Notations: endorseSignatures,
	}
}

func createBallot(heading *kinds.Heading, validset *kinds.AssessorAssign,
	key security.PrivateToken, ledgerUUID kinds.LedgerUUID,
) *kinds.Ballot {
	location := key.PublicToken().Location()
	idx, _ := validset.ObtainViaLocation(location)
	ballot := &kinds.Ballot{
		AssessorLocation: location,
		AssessorOrdinal:   idx,
		Altitude:           heading.Altitude,
		Iteration:            1,
		Timestamp:        committime.Now(),
		Kind:             commitchema.PreendorseKind,
		LedgerUUID:          ledgerUUID,
	}

	v := ballot.TowardSchema()
	//
	attestOctets := kinds.BallotAttestOctets(heading.SuccessionUUID, v)
	sig, err := key.Attest(attestOctets)
	if err != nil {
		panic(err)
	}
	ballot.Notation = sig

	addnAttestOctets := kinds.BallotAdditionAttestOctets(heading.SuccessionUUID, v)
	addnSignature, err := key.Attest(addnAttestOctets)
	if err != nil {
		panic(err)
	}
	ballot.AdditionNotation = addnSignature

	return ballot
}

func produceHeadline(successionUUID string, altitude int64, byteMoment time.Time, txs kinds.Txs,
	validset, followingValidset *kinds.AssessorAssign, platformDigest, consensusDigest, resultDigest []byte,
) *kinds.Heading {
	return &kinds.Heading{
		Edition: strongmindedition.Agreement{Ledger: edition.LedgerScheme, App: 0},
		SuccessionUUID: successionUUID,
		Altitude:  altitude,
		Moment:    byteMoment,
		//
		//
		AssessorsDigest:     validset.Digest(),
		FollowingAssessorsDigest: followingValidset.Digest(),
		DataDigest:           txs.Digest(),
		PlatformDigest:            platformDigest,
		AgreementDigest:      consensusDigest,
		FinalOutcomesDigest:    resultDigest,
		NominatorLocation:    validset.Assessors[0].Location,
	}
}

//
func (pkz privateTokens) ProduceNotatedHeadline(successionUUID string, altitude int64, byteMoment time.Time, txs kinds.Txs,
	validset, followingValidset *kinds.AssessorAssign, platformDigest, consensusDigest, resultDigest []byte, initial, final int,
) *kinds.NotatedHeading {
	heading := produceHeadline(successionUUID, altitude, byteMoment, txs, validset, followingValidset, platformDigest, consensusDigest, resultDigest)
	return &kinds.NotatedHeading{
		Heading: heading,
		Endorse: pkz.attestHeadline(heading, validset, initial, final),
	}
}

//
func (pkz privateTokens) ProduceNotatedHeadlineFinalLedgerUUID(successionUUID string, altitude int64, byteMoment time.Time, txs kinds.Txs,
	validset, followingValidset *kinds.AssessorAssign, platformDigest, consensusDigest, resultDigest []byte, initial, final int,
	finalLedgerUUID kinds.LedgerUUID,
) *kinds.NotatedHeading {
	heading := produceHeadline(successionUUID, altitude, byteMoment, txs, validset, followingValidset, platformDigest, consensusDigest, resultDigest)
	heading.FinalLedgerUUID = finalLedgerUUID
	return &kinds.NotatedHeading{
		Heading: heading,
		Endorse: pkz.attestHeadline(heading, validset, initial, final),
	}
}

func (pkz privateTokens) AlterationTokens(variation int) privateTokens {
	freshTokens := pkz[variation:]
	return freshTokens.Broaden(variation)
}

//
//
//
func produceSimulatePeerUsingTokens(
	successionUUID string,
	ledgerExtent int64,
	itemExtent int,
	itemDeviation float32,
	byteMoment time.Time) (
	map[int64]*kinds.NotatedHeading,
	map[int64]*kinds.AssessorAssign,
	map[int64]privateTokens,
) {
	var (
		headings         = make(map[int64]*kinds.NotatedHeading, ledgerExtent)
		validset          = make(map[int64]*kinds.AssessorAssign, ledgerExtent+1)
		tokenmap          = make(map[int64]privateTokens, ledgerExtent+1)
		tokens            = producePrivateTokens(itemExtent)
		sumDeviation  = itemDeviation
		itemDeviationInteger int
		freshTokens         privateTokens
	)

	itemDeviationInteger = int(sumDeviation)
	sumDeviation = -float32(itemDeviationInteger)
	freshTokens = tokens.AlterationTokens(itemDeviationInteger)
	tokenmap[1] = tokens
	tokenmap[2] = freshTokens

	//
	finalHeadline := tokens.ProduceNotatedHeadline(successionUUID, 1, byteMoment.Add(1*time.Minute), nil,
		tokens.TowardAssessors(2, 0), freshTokens.TowardAssessors(2, 0), digest("REDACTED"), digest("REDACTED"),
		digest("REDACTED"), 0, len(tokens))
	prevailingHeadline := finalHeadline
	headings[1] = prevailingHeadline
	validset[1] = tokens.TowardAssessors(2, 0)
	tokens = freshTokens

	for altitude := int64(2); altitude <= ledgerExtent; altitude++ {
		sumDeviation += itemDeviation
		itemDeviationInteger = int(sumDeviation)
		sumDeviation = -float32(itemDeviationInteger)
		freshTokens = tokens.AlterationTokens(itemDeviationInteger)
		prevailingHeadline = tokens.ProduceNotatedHeadlineFinalLedgerUUID(successionUUID, altitude, byteMoment.Add(time.Duration(altitude)*time.Minute),
			nil,
			tokens.TowardAssessors(2, 0), freshTokens.TowardAssessors(2, 0), digest("REDACTED"), digest("REDACTED"),
			digest("REDACTED"), 0, len(tokens), kinds.LedgerUUID{Digest: finalHeadline.Digest()})
		headings[altitude] = prevailingHeadline
		validset[altitude] = tokens.TowardAssessors(2, 0)
		finalHeadline = prevailingHeadline
		tokens = freshTokens
		tokenmap[altitude+1] = tokens
	}

	return headings, validset, tokenmap
}

func produceSimulatePeer(
	successionUUID string,
	ledgerExtent int64,
	itemExtent int,
	itemDeviation float32,
	byteMoment time.Time) (
	string,
	map[int64]*kinds.NotatedHeading,
	map[int64]*kinds.AssessorAssign,
) {
	headings, validset, _ := produceSimulatePeerUsingTokens(successionUUID, ledgerExtent, itemExtent, itemDeviation, byteMoment)
	return successionUUID, headings, validset
}

func digest(s string) []byte {
	return tenderminthash.Sum([]byte(s))
}
