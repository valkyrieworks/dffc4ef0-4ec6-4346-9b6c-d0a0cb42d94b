package kinds_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/hashmap"
)

func VerifyDigestAlsoAscertainOutcomes(t *testing.T) {
	trs := []*iface.InvokeTransferOutcome{
		//
		{Cipher: 0, Data: nil},
		{Cipher: 0, Data: []byte{}},

		{Cipher: 0, Data: []byte("REDACTED")},
		{Cipher: 14, Data: nil},
		{Cipher: 14, Data: []byte("REDACTED")},
		{Cipher: 14, Data: []byte("REDACTED")},
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
	origin := hashmap.DigestOriginatingOctetSegments(rs)
	assert.NotEmpty(t, origin)

	_, attestations := hashmap.AttestationsOriginatingOctetSegments(rs)
	for i, tr := range trs {
		bz, err := tr.Serialize()
		require.NoError(t, err)

		sound := attestations[i].Validate(origin, bz)
		assert.NoError(t, sound, "REDACTED", i)
	}
}

func VerifyDigestCertainAreasSolely(t *testing.T) {
	tr1 := iface.InvokeTransferOutcome{
		Cipher:      1,
		Data:      []byte("REDACTED"),
		Log:       "REDACTED",
		Details:      "REDACTED",
		FuelDesired: 1000,
		FuelUtilized:   1000,
		Incidents:    []iface.Incident{},
		Codeset: "REDACTED",
	}
	tr2 := iface.InvokeTransferOutcome{
		Cipher:      1,
		Data:      []byte("REDACTED"),
		Log:       "REDACTED",
		Details:      "REDACTED",
		FuelDesired: 1000,
		FuelUtilized:   1000,
		Incidents:    []iface.Incident{},
		Codeset: "REDACTED",
	}
	r1, err := iface.SerializeTransferOutcomes([]*iface.InvokeTransferOutcome{&tr1})
	require.NoError(t, err)
	r2, err := iface.SerializeTransferOutcomes([]*iface.InvokeTransferOutcome{&tr2})
	require.NoError(t, err)
	require.Equal(t, hashmap.DigestOriginatingOctetSegments(r1), hashmap.DigestOriginatingOctetSegments(r2))
}
