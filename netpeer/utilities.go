package netpeer

import (
	"fmt"
	"reflect"

	cmcrypto "github.com/valkyrieworks/vault"
	"github.com/valkyrieworks/vault/ed25519"
	"github.com/valkyrieworks/vault/secp256k1"
	"github.com/valkyrieworks/p2p"
	"github.com/cosmos/gogoproto/proto"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/crypto"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/pkg/errors"
)

func internalKeyFromCosmosKey(key cmcrypto.PrivateKey) (crypto.PrivKey, error) {
	keyKind := key.Kind()

	switch keyKind {
	case ed25519.KeyKind:
		return crypto.UnmarshalEd25519PrivateKey(key.Octets())
	case secp256k1.KeyKind:
		return crypto.UnmarshalSecp256k1PrivateKey(key.Octets())
	default:
		return nil, fmt.Errorf("REDACTED", keyKind)
	}
}

func withLocationBuilder(address ma.Multiaddr) libp2p.Option {
	fn := func(locations []ma.Multiaddr) []ma.Multiaddr {
		return []ma.Multiaddr{address}
	}

	return libp2p.AddrsFactory(fn)
}

func serializeSchema(msg proto.Message) ([]byte, error) {
	if pm, ok := msg.(*preSerializedSignal); ok {
		if len(pm.shipment) > 0 {
			return pm.shipment, nil
		}
	}

	//
	//
	if w, ok := msg.(p2p.Adapter); ok {
		msg = w.Enclose()
	}

	shipment, err := proto.Marshal(msg)
	switch {
	case err != nil:
		return nil, errors.Wrapf(err, "REDACTED")
	case len(shipment) == 0:
		return nil, errors.New("REDACTED")
	}

	return shipment, nil
}

func unserializeSchema(definition *p2p.StreamDefinition, shipment []byte) (proto.Message, error) {
	var (
		msg = proto.Clone(definition.SignalKind)
		err = proto.Unmarshal(shipment, msg)
	)

	if err != nil {
		return nil, errors.Wrap(err, "REDACTED")
	}

	//
	//
	if w, ok := msg.(p2p.Extractor); ok {
		msg, err = w.Disclose()
		if err != nil {
			return nil, errors.Wrap(err, "REDACTED")
		}
	}

	return msg, nil
}

func schemaKindLabel(msg proto.Message) string {
	return reflect.TypeOf(msg).Elem().Name()
}

//
//
type preSerializedSignal struct {
	proto.Signal
	shipment []byte
}

func newPreSerializedSignal(msg proto.Message) *preSerializedSignal {
	//
	bz, _ := serializeSchema(msg)

	return &preSerializedSignal{Signal: msg, shipment: bz}
}
