package p2p

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/vault/ed25519"
	engineseed "github.com/valkyrieworks/utils/random"
)

func VerifyImportOrGenerateMemberKey(t *testing.T) {
	entryRoute := filepath.Join(os.TempDir(), engineseed.Str(12)+"REDACTED")

	memberKey, err := ImportOrGenerateMemberKey(entryRoute)
	assert.Nil(t, err)

	memberKey2, err := ImportOrGenerateMemberKey(entryRoute)
	assert.Nil(t, err)

	assert.Equal(t, memberKey, memberKey2)
}

func VerifyImportMemberKey(t *testing.T) {
	entryRoute := filepath.Join(os.TempDir(), engineseed.Str(12)+"REDACTED")

	_, err := ImportMemberKey(entryRoute)
	assert.True(t, os.IsNotExist(err))

	_, err = ImportOrGenerateMemberKey(entryRoute)
	require.NoError(t, err)

	memberKey, err := ImportMemberKey(entryRoute)
	assert.NoError(t, err)
	assert.NotNil(t, memberKey)
}

func VerifyMemberKeyPersistAs(t *testing.T) {
	entryRoute := filepath.Join(os.TempDir(), engineseed.Str(12)+"REDACTED")

	assert.NoFileExists(t, entryRoute)

	privateKey := ed25519.GeneratePrivateKey()
	memberKey := &MemberKey{
		PrivateKey: privateKey,
	}
	err := memberKey.PersistAs(entryRoute)
	assert.NoError(t, err)
	assert.FileExists(t, entryRoute)
}

//

func paddingOctets(bz []byte, objectiveOctets int) []byte {
	return append(bz, bytes.Repeat([]byte{0xFF}, objectiveOctets-len(bz))...)
}

func VerifyPoWriterObjective(t *testing.T) {
	objectiveOctets := 20
	scenarios := []struct {
		complexity uint
		objective     []byte
	}{
		{0, paddingOctets([]byte{}, objectiveOctets)},
		{1, paddingOctets([]byte{127}, objectiveOctets)},
		{8, paddingOctets([]byte{0}, objectiveOctets)},
		{9, paddingOctets([]byte{0, 127}, objectiveOctets)},
		{10, paddingOctets([]byte{0, 63}, objectiveOctets)},
		{16, paddingOctets([]byte{0, 0}, objectiveOctets)},
		{17, paddingOctets([]byte{0, 0, 127}, objectiveOctets)},
	}

	for _, c := range scenarios {
		assert.Equal(t, CreatePoWriterObjective(c.complexity, 20*8), c.objective)
	}
}
