package kinds

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	iface "github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/vault"
	"github.com/valkyrieworks/vault/ed25519"
	cryptocode "github.com/valkyrieworks/vault/codec"
)

func VerifyIfacePublicKey(t *testing.T) {
	publicidEd := ed25519.GeneratePrivateKey().PublicKey()
	err := verifyIfacePublicKey(t, publicidEd)
	assert.NoError(t, err)
}

func verifyIfacePublicKey(t *testing.T, pk vault.PublicKey) error {
	ifacePublicKey, err := cryptocode.PublicKeyToSchema(pk)
	require.NoError(t, err)
	pk2, err := cryptocode.PublicKeyFromSchema(ifacePublicKey)
	require.NoError(t, err)
	require.Equal(t, pk, pk2)
	return nil
}

func VerifyIfaceRatifiers(t *testing.T) {
	publicidEd := ed25519.GeneratePrivateKey().PublicKey()

	//
	cometValueAnticipated := NewRatifier(publicidEd, 10)

	cometValue := NewRatifier(publicidEd, 10)

	ifaceValue := Tm2schema.RatifierModify(cometValue)
	cometValues, err := Schema2tm.RatifierRefreshes([]iface.RatifierModify{ifaceValue})
	assert.Nil(t, err)
	assert.Equal(t, cometValueAnticipated, cometValues[0])

	ifaceValues := Tm2schema.RatifierRefreshes(NewRatifierCollection(cometValues))
	assert.Equal(t, []iface.RatifierModify{ifaceValue}, ifaceValues)

	//
	cometValue.Location = publicidEd.Location()

	ifaceValue = Tm2schema.RatifierModify(cometValue)
	cometValues, err = Schema2tm.RatifierRefreshes([]iface.RatifierModify{ifaceValue})
	assert.Nil(t, err)
	assert.Equal(t, cometValueAnticipated, cometValues[0])
}

type publicKeyEddie struct{}

func (publicKeyEddie) Location() Location                    { return []byte{} }
func (publicKeyEddie) Octets() []byte                       { return []byte{} }
func (publicKeyEddie) ValidateAutograph([]byte, []byte) bool { return false }
func (publicKeyEddie) Matches(vault.PublicKey) bool           { return false }
func (publicKeyEddie) String() string                      { return "REDACTED" }
func (publicKeyEddie) Kind() string                        { return "REDACTED" }

func VerifyIfaceRatifierFromPublicKeyAndEnergy(t *testing.T) {
	publickey := ed25519.GeneratePrivateKey().PublicKey()

	ifaceValue := Tm2schema.NewRatifierModify(publickey, 10)
	assert.Equal(t, int64(10), ifaceValue.Energy)

	assert.Panics(t, func() { Tm2schema.NewRatifierModify(nil, 10) })
	assert.Panics(t, func() { Tm2schema.NewRatifierModify(publicKeyEddie{}, 10) })
}

func VerifyIfaceRatifierLackingPublicKey(t *testing.T) {
	publicidEd := ed25519.GeneratePrivateKey().PublicKey()

	ifaceValue := Tm2schema.Ratifier(NewRatifier(publicidEd, 10))

	//
	cometValueAnticipated := iface.Ratifier{
		Location: publicidEd.Location(),
		Energy:   10,
	}

	assert.Equal(t, cometValueAnticipated, ifaceValue)
}
