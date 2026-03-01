package kinds

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/edwards25519"
	cryptocode "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/serialization"
)

func VerifyIfacePublicToken(t *testing.T) {
	keyEdwards := edwards25519.ProducePrivateToken().PublicToken()
	err := verifyIfacePublicToken(t, keyEdwards)
	assert.NoError(t, err)
}

func verifyIfacePublicToken(t *testing.T, pk security.PublicToken) error {
	ifacePublicToken, err := cryptocode.PublicTokenTowardSchema(pk)
	require.NoError(t, err)
	pk2, err := cryptocode.PublicTokenOriginatingSchema(ifacePublicToken)
	require.NoError(t, err)
	require.Equal(t, pk, pk2)
	return nil
}

func VerifyIfaceAssessors(t *testing.T) {
	keyEdwards := edwards25519.ProducePrivateToken().PublicToken()

	//
	cometItemAnticipated := FreshAssessor(keyEdwards, 10)

	cometItem := FreshAssessor(keyEdwards, 10)

	ifaceItem := Temp2buffer.AssessorRevise(cometItem)
	cometValues, err := Buffer2temp.AssessorRevisions([]iface.AssessorRevise{ifaceItem})
	assert.Nil(t, err)
	assert.Equal(t, cometItemAnticipated, cometValues[0])

	ifaceValues := Temp2buffer.AssessorRevisions(FreshAssessorAssign(cometValues))
	assert.Equal(t, []iface.AssessorRevise{ifaceItem}, ifaceValues)

	//
	cometItem.Location = keyEdwards.Location()

	ifaceItem = Temp2buffer.AssessorRevise(cometItem)
	cometValues, err = Buffer2temp.AssessorRevisions([]iface.AssessorRevise{ifaceItem})
	assert.Nil(t, err)
	assert.Equal(t, cometItemAnticipated, cometValues[0])
}

type publicTokenErnie struct{}

func (publicTokenErnie) Location() Location                    { return []byte{} }
func (publicTokenErnie) Octets() []byte                       { return []byte{} }
func (publicTokenErnie) ValidateNotation([]byte, []byte) bool { return false }
func (publicTokenErnie) Matches(security.PublicToken) bool           { return false }
func (publicTokenErnie) Text() string                      { return "REDACTED" }
func (publicTokenErnie) Kind() string                        { return "REDACTED" }

func VerifyIfaceAssessorOriginatingPublicTokenAlsoPotency(t *testing.T) {
	publickey := edwards25519.ProducePrivateToken().PublicToken()

	ifaceItem := Temp2buffer.FreshAssessorRevise(publickey, 10)
	assert.Equal(t, int64(10), ifaceItem.Potency)

	assert.Panics(t, func() { Temp2buffer.FreshAssessorRevise(nil, 10) })
	assert.Panics(t, func() { Temp2buffer.FreshAssessorRevise(publicTokenErnie{}, 10) })
}

func VerifyIfaceAssessorLackingPublicToken(t *testing.T) {
	keyEdwards := edwards25519.ProducePrivateToken().PublicToken()

	ifaceItem := Temp2buffer.Assessor(FreshAssessor(keyEdwards, 10))

	//
	cometItemAnticipated := iface.Assessor{
		Location: keyEdwards.Location(),
		Potency:   10,
	}

	assert.Equal(t, cometItemAnticipated, ifaceItem)
}
