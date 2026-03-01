package verify

import (
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

func OriginPaper(
	moment time.Time,
	assessors []*kinds.Assessor,
	agreementParameters *kinds.AgreementSettings,
	successionUUID string,
) *kinds.OriginPaper {
	inaugurationAssessors := make([]kinds.OriginAssessor, len(assessors))

	for i := range assessors {
		inaugurationAssessors[i] = kinds.OriginAssessor{
			Potency:  assessors[i].BallotingPotency,
			PublicToken: assessors[i].PublicToken,
		}
	}

	return &kinds.OriginPaper{
		OriginMoment:     moment,
		PrimaryAltitude:   1,
		SuccessionUUID:         successionUUID,
		Assessors:      inaugurationAssessors,
		AgreementSettings: agreementParameters,
	}
}
