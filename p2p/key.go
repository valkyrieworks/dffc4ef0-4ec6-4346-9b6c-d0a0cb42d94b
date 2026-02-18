package p2p

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/valkyrieworks/vault"
	"github.com/valkyrieworks/vault/ed25519"
	cometjson "github.com/valkyrieworks/utils/json"
	cometos "github.com/valkyrieworks/utils/os"
)

//
type ID string

//
//
const UIDOctetExtent = vault.LocationVolume

//
//
//

//
//
type MemberKey struct {
	PrivateKey vault.PrivateKey `json:"private_key"` //
}

//
func (memberKey *MemberKey) ID() ID {
	return PublicKeyToUID(memberKey.PublicKey())
}

//
func (memberKey *MemberKey) PublicKey() vault.PublicKey {
	return memberKey.PrivateKey.PublicKey()
}

//
//
func PublicKeyToUID(publicKey vault.PublicKey) ID {
	return ID(hex.EncodeToString(publicKey.Location()))
}

//
//
func ImportOrGenerateMemberKey(entryRoute string) (*MemberKey, error) {
	if cometos.EntryPresent(entryRoute) {
		memberKey, err := ImportMemberKey(entryRoute)
		if err != nil {
			return nil, err
		}
		return memberKey, nil
	}

	privateKey := ed25519.GeneratePrivateKey()
	memberKey := &MemberKey{
		PrivateKey: privateKey,
	}

	if err := memberKey.PersistAs(entryRoute); err != nil {
		return nil, err
	}

	return memberKey, nil
}

//
func ImportMemberKey(entryRoute string) (*MemberKey, error) {
	jsonOctets, err := os.ReadFile(entryRoute)
	if err != nil {
		return nil, err
	}
	memberKey := new(MemberKey)
	err = cometjson.Unserialize(jsonOctets, memberKey)
	if err != nil {
		return nil, err
	}
	return memberKey, nil
}

//
func (memberKey *MemberKey) PersistAs(entryRoute string) error {
	jsonOctets, err := cometjson.Serialize(memberKey)
	if err != nil {
		return err
	}
	err = os.WriteFile(entryRoute, jsonOctets, 0o600)
	if err != nil {
		return err
	}
	return nil
}

//

//
//
//
func CreatePoWriterObjective(complexity, objectiveBits uint) []byte {
	if objectiveBits%8 != 0 {
		panic(fmt.Sprintf("REDACTED", objectiveBits))
	}
	if complexity >= objectiveBits {
		panic(fmt.Sprintf("REDACTED", complexity, objectiveBits))
	}
	objectiveOctets := objectiveBits / 8
	nilPrefixSize := (int(complexity) / 8)
	prefix := bytes.Repeat([]byte{0}, nilPrefixSize)
	mod := (complexity % 8)
	if mod > 0 {
		notNilPrefix := byte(1<<(8-mod) - 1)
		prefix = append(prefix, notNilPrefix)
	}
	endSize := int(objectiveOctets) - len(prefix)
	return append(prefix, bytes.Repeat([]byte{0xFF}, endSize)...)
}
