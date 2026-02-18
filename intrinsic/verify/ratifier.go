package verify

import (
	"context"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/kinds"
)

func Ratifier(_ context.Context, pollingEnergy int64) (*kinds.Ratifier, kinds.PrivateRatifier, error) {
	privateValue := kinds.NewEmulatePV()
	publicKey, err := privateValue.FetchPublicKey()
	if err != nil {
		return nil, nil, err
	}

	val := kinds.NewRatifier(publicKey, pollingEnergy)
	return val, privateValue, nil
}

func RatifierAssign(ctx context.Context, t *testing.T, countRatifiers int, pollingEnergy int64) (*kinds.RatifierAssign, []kinds.PrivateRatifier) {
	var (
		valz           = make([]*kinds.Ratifier, countRatifiers)
		privateRatifiers = make([]kinds.PrivateRatifier, countRatifiers)
	)
	t.Helper()

	for i := 0; i < countRatifiers; i++ {
		val, privateRatifier, err := Ratifier(ctx, pollingEnergy)
		require.NoError(t, err)
		valz[i] = val
		privateRatifiers[i] = privateRatifier
	}

	sort.Sort(kinds.PrivateRatifiersByLocation(privateRatifiers))

	return kinds.NewRatifierCollection(valz), privateRatifiers
}
