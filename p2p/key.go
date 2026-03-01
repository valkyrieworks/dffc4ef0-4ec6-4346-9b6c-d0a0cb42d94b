package p2p

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/edwards25519"
	strongmindjson "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/jsn"
	strongos "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/os"
)

//
type ID string

//
//
const UUIDOctetMagnitude = security.LocatorExtent

//
//
//

//
//
type PeerToken struct {
	PrivateToken security.PrivateToken `json:"private_token"` //
}

//
func (peerToken *PeerToken) ID() ID {
	return PublicTokenTowardUUID(peerToken.PublicToken())
}

//
func (peerToken *PeerToken) PublicToken() security.PublicToken {
	return peerToken.PrivateToken.PublicToken()
}

//
//
func PublicTokenTowardUUID(publicToken security.PublicToken) ID {
	return ID(hex.EncodeToString(publicToken.Location()))
}

//
//
func FetchEitherProducePeerToken(recordRoute string) (*PeerToken, error) {
	if strongos.RecordPresent(recordRoute) {
		peerToken, err := FetchPeerToken(recordRoute)
		if err != nil {
			return nil, err
		}
		return peerToken, nil
	}

	privateToken := edwards25519.ProducePrivateToken()
	peerToken := &PeerToken{
		PrivateToken: privateToken,
	}

	if err := peerToken.PersistLike(recordRoute); err != nil {
		return nil, err
	}

	return peerToken, nil
}

//
func FetchPeerToken(recordRoute string) (*PeerToken, error) {
	jsnOctets, err := os.ReadFile(recordRoute)
	if err != nil {
		return nil, err
	}
	peerToken := new(PeerToken)
	err = strongmindjson.Decode(jsnOctets, peerToken)
	if err != nil {
		return nil, err
	}
	return peerToken, nil
}

//
func (peerToken *PeerToken) PersistLike(recordRoute string) error {
	jsnOctets, err := strongmindjson.Serialize(peerToken)
	if err != nil {
		return err
	}
	err = os.WriteFile(recordRoute, jsnOctets, 0o600)
	if err != nil {
		return err
	}
	return nil
}

//

//
//
//
func CreatePositionWRObjective(complexity, objectiveDigits uint) []byte {
	if objectiveDigits%8 != 0 {
		panic(fmt.Sprintf("REDACTED", objectiveDigits))
	}
	if complexity >= objectiveDigits {
		panic(fmt.Sprintf("REDACTED", complexity, objectiveDigits))
	}
	objectiveOctets := objectiveDigits / 8
	nullHeadingLength := (int(complexity) / 8)
	heading := bytes.Repeat([]byte{0}, nullHeadingLength)
	mod := (complexity % 8)
	if mod > 0 {
		unNullHeading := byte(1<<(8-mod) - 1)
		heading = append(heading, unNullHeading)
	}
	endLength := int(objectiveOctets) - len(heading)
	return append(heading, bytes.Repeat([]byte{0xFF}, endLength)...)
}
