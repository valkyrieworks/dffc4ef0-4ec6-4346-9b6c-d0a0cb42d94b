package verify

import (
	"context"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

func Assessor(_ context.Context, ballotingPotency int64) (*kinds.Assessor, kinds.PrivateAssessor, error) {
	privateItem := kinds.FreshSimulatePRV()
	publicToken, err := privateItem.ObtainPublicToken()
	if err != nil {
		return nil, nil, err
	}

	val := kinds.FreshAssessor(publicToken, ballotingPotency)
	return val, privateItem, nil
}

func AssessorAssign(ctx context.Context, t *testing.T, countAssessors int, ballotingPotency int64) (*kinds.AssessorAssign, []kinds.PrivateAssessor) {
	var (
		validz           = make([]*kinds.Assessor, countAssessors)
		privateAssessors = make([]kinds.PrivateAssessor, countAssessors)
	)
	t.Helper()

	for i := 0; i < countAssessors; i++ {
		val, privateAssessor, err := Assessor(ctx, ballotingPotency)
		require.NoError(t, err)
		validz[i] = val
		privateAssessors[i] = privateAssessor
	}

	sort.Sort(kinds.PrivateAssessorsViaLocation(privateAssessors))

	return kinds.FreshAssessorAssign(validz), privateAssessors
}
