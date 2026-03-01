package netp2p

import (
	"fmt"
	"reflect"

	cryptography "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/edwards25519"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/ellipticp256"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	"github.com/cosmos/gogoproto/proto"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/crypto"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/pkg/errors"
)

func secludedTokenOriginatingUniverseToken(key cryptography.PrivateToken) (crypto.PrivKey, error) {
	tokenKind := key.Kind()

	switch tokenKind {
	case edwards25519.TokenKind:
		return crypto.UnmarshalEd25519PrivateKey(key.Octets())
	case ellipticp256.TokenKind:
		return crypto.UnmarshalSecp256k1PrivateKey(key.Octets())
	default:
		return nil, fmt.Errorf("REDACTED", tokenKind)
	}
}

func usingLocatorBuilder(location ma.Multiaddr) libp2p.Option {
	fn := func(locations []ma.Multiaddr) []ma.Multiaddr {
		return []ma.Multiaddr{location}
	}

	return libp2p.AddrsFactory(fn)
}

func serializeSchema(msg proto.Message) ([]byte, error) {
	if pm, ok := msg.(*priorSerializedArtifact); ok {
		if len(pm.content) > 0 {
			return pm.content, nil
		}
	}

	//
	//
	if w, ok := msg.(p2p.Encapsulator); ok {
		msg = w.Enclose()
	}

	content, err := proto.Marshal(msg)
	switch {
	case err != nil:
		return nil, errors.Wrapf(err, "REDACTED")
	case len(content) == 0:
		return nil, errors.New("REDACTED")
	}

	return content, nil
}

func decodeSchema(definition *p2p.ConduitDefinition, content []byte) (proto.Message, error) {
	var (
		msg = proto.Clone(definition.SignalKind)
		err = proto.Unmarshal(content, msg)
	)

	if err != nil {
		return nil, errors.Wrap(err, "REDACTED")
	}

	//
	//
	if w, ok := msg.(p2p.Unwrapper); ok {
		msg, err = w.Disclose()
		if err != nil {
			return nil, errors.Wrap(err, "REDACTED")
		}
	}

	return msg, nil
}

//
//
type priorSerializedArtifact struct {
	proto.Signal
	content []byte
}

func freshPriorSerializedArtifact(msg proto.Message) *priorSerializedArtifact {
	//
	bz, _ := serializeSchema(msg)

	return &priorSerializedArtifact{Signal: msg, content: bz}
}

func schemaKindAlias(msg proto.Message) string {
	if pm, ok := msg.(*priorSerializedArtifact); ok {
		return reflect.TypeOf(pm.Signal).Elem().Name()
	}

	return reflect.TypeOf(msg).Elem().Name()
}
