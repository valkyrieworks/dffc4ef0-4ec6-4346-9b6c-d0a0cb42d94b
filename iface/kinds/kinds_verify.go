package kinds_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	iface "github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/vault/merkle"
)

func VerifyDigestAndDemonstrateOutcomes(t *testing.T) {
	trs := []*iface.InvokeTransferOutcome{
		//
		{Code: 0, Data: nil},
		{Code: 0, Data: []byte{}},

		{Code: 0, Data: []byte("REDACTED")},
		{Code: 14, Data: nil},
		{Code: 14, Data: []byte("REDACTED")},
		{Code: 14, Data: []byte("REDACTED")},
	}

	//
	bz0, err := trs[0].Serialize()
	require.NoError(t, err)
	bz1, err := trs[1].Serialize()
	require.NoError(t, err)
	require.Equal(t, bz0, bz1)

	//
	rs, err := iface.SerializeTransferOutcomes(trs)
	require.NoError(t, err)
	origin := merkle.DigestFromOctetSegments(rs)
	assert.NotEmpty(t, origin)

	_, evidences := merkle.EvidencesFromOctetSegments(rs)
	for i, tr := range trs {
		bz, err := tr.Serialize()
		require.NoError(t, err)

		sound := evidences[i].Validate(origin, bz)
		assert.NoError(t, sound, "REDACTED", i)
	}
}

func VerifyDigestCertainAttributesSolely(t *testing.T) {
	tr1 := iface.InvokeTransferOutcome{
		Code:      1,
		Data:      []byte("REDACTED"),
		Log:       "REDACTED",
		Details:      "REDACTED",
		FuelDesired: 1000,
		FuelApplied:   1000,
		Events:    []iface.Event{},
		Codex: "REDACTED",
	}
	tr2 := iface.InvokeTransferOutcome{
		Code:      1,
		Data:      []byte("REDACTED"),
		Log:       "REDACTED",
		Details:      "REDACTED",
		FuelDesired: 1000,
		FuelApplied:   1000,
		Events:    []iface.Event{},
		Codex: "REDACTED",
	}
	r1, err := iface.SerializeTransferOutcomes([]*iface.InvokeTransferOutcome{&tr1})
	require.NoError(t, err)
	r2, err := iface.SerializeTransferOutcomes([]*iface.InvokeTransferOutcome{&tr2})
	require.NoError(t, err)
	require.Equal(t, merkle.DigestFromOctetSegments(r1), merkle.DigestFromOctetSegments(r2))
}
