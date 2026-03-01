package p2p

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/edwards25519"
	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
)

func VerifyFetchEitherProducePeerToken(t *testing.T) {
	recordRoute := filepath.Join(os.TempDir(), commitrand.Str(12)+"REDACTED")

	peerToken, err := FetchEitherProducePeerToken(recordRoute)
	assert.Nil(t, err)

	peerToken2, err := FetchEitherProducePeerToken(recordRoute)
	assert.Nil(t, err)

	assert.Equal(t, peerToken, peerToken2)
}

func VerifyFetchPeerToken(t *testing.T) {
	recordRoute := filepath.Join(os.TempDir(), commitrand.Str(12)+"REDACTED")

	_, err := FetchPeerToken(recordRoute)
	assert.True(t, os.IsNotExist(err))

	_, err = FetchEitherProducePeerToken(recordRoute)
	require.NoError(t, err)

	peerToken, err := FetchPeerToken(recordRoute)
	assert.NoError(t, err)
	assert.NotNil(t, peerToken)
}

func VerifyPeerTokenPersistLike(t *testing.T) {
	recordRoute := filepath.Join(os.TempDir(), commitrand.Str(12)+"REDACTED")

	assert.NoFileExists(t, recordRoute)

	privateToken := edwards25519.ProducePrivateToken()
	peerToken := &PeerToken{
		PrivateToken: privateToken,
	}
	err := peerToken.PersistLike(recordRoute)
	assert.NoError(t, err)
	assert.FileExists(t, recordRoute)
}

//

func stuffOctets(bz []byte, objectiveOctets int) []byte {
	return append(bz, bytes.Repeat([]byte{0xFF}, objectiveOctets-len(bz))...)
}

func VerifyPositionWRObjective(t *testing.T) {
	objectiveOctets := 20
	scenarios := []struct {
		complexity uint
		objective     []byte
	}{
		{0, stuffOctets([]byte{}, objectiveOctets)},
		{1, stuffOctets([]byte{127}, objectiveOctets)},
		{8, stuffOctets([]byte{0}, objectiveOctets)},
		{9, stuffOctets([]byte{0, 127}, objectiveOctets)},
		{10, stuffOctets([]byte{0, 63}, objectiveOctets)},
		{16, stuffOctets([]byte{0, 0}, objectiveOctets)},
		{17, stuffOctets([]byte{0, 0, 127}, objectiveOctets)},
	}

	for _, c := range scenarios {
		assert.Equal(t, CreatePositionWRObjective(c.complexity, 20*8), c.objective)
	}
}
