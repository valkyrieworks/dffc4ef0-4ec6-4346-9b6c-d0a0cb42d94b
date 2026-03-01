package kinds

import (
	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	cryptocode "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/serialization"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
)

//

//
//
var Temp2buffer = temp2buffer{}

type temp2buffer struct{}

func (temp2buffer) Heading(heading *Heading) commitchema.Heading {
	return commitchema.Heading{
		Edition: heading.Edition,
		SuccessionUUID: heading.SuccessionUUID,
		Altitude:  heading.Altitude,
		Moment:    heading.Moment,

		FinalLedgerUuid: heading.FinalLedgerUUID.TowardSchema(),

		FinalEndorseDigest: heading.FinalEndorseDigest,
		DataDigest:       heading.DataDigest,

		AssessorsDigest:     heading.AssessorsDigest,
		FollowingAssessorsDigest: heading.FollowingAssessorsDigest,
		AgreementDigest:      heading.AgreementDigest,
		PlatformDigest:            heading.PlatformDigest,
		FinalOutcomesDigest:    heading.FinalOutcomesDigest,

		ProofDigest:    heading.ProofDigest,
		NominatorLocation: heading.NominatorLocation,
	}
}

func (temp2buffer) Assessor(val *Assessor) iface.Assessor {
	return iface.Assessor{
		Location: val.PublicToken.Location(),
		Potency:   val.BallotingPotency,
	}
}

func (temp2buffer) LedgerUUID(ledgerUUID LedgerUUID) commitchema.LedgerUUID {
	return commitchema.LedgerUUID{
		Digest:          ledgerUUID.Digest,
		FragmentAssignHeading: Temp2buffer.FragmentAssignHeading(ledgerUUID.FragmentAssignHeading),
	}
}

func (temp2buffer) FragmentAssignHeading(heading FragmentAssignHeading) commitchema.FragmentAssignHeading {
	return commitchema.FragmentAssignHeading{
		Sum: heading.Sum,
		Digest:  heading.Digest,
	}
}

//
func (temp2buffer) AssessorRevise(val *Assessor) iface.AssessorRevise {
	pk, err := cryptocode.PublicTokenTowardSchema(val.PublicToken)
	if err != nil {
		panic(err)
	}
	return iface.AssessorRevise{
		PublicToken: pk,
		Potency:  val.BallotingPotency,
	}
}

//
func (temp2buffer) AssessorRevisions(values *AssessorAssign) []iface.AssessorRevise {
	assessors := make([]iface.AssessorRevise, values.Extent())
	for i, val := range values.Assessors {
		assessors[i] = Temp2buffer.AssessorRevise(val)
	}
	return assessors
}

//
func (temp2buffer) FreshAssessorRevise(publickey security.PublicToken, potency int64) iface.AssessorRevise {
	publickeyIface, err := cryptocode.PublicTokenTowardSchema(publickey)
	if err != nil {
		panic(err)
	}
	return iface.AssessorRevise{
		PublicToken: publickeyIface,
		Potency:  potency,
	}
}

//

//
//
var Buffer2temp = buffer2temp{}

type buffer2temp struct{}

func (buffer2temp) AssessorRevisions(values []iface.AssessorRevise) ([]*Assessor, error) {
	cometValues := make([]*Assessor, len(values))
	for i, v := range values {
		pub, err := cryptocode.PublicTokenOriginatingSchema(v.PublicToken)
		if err != nil {
			return nil, err
		}
		cometValues[i] = FreshAssessor(pub, v.Potency)
	}
	return cometValues, nil
}
