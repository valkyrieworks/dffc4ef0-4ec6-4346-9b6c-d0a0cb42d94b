package codec

import (
	"fmt"

	"github.com/valkyrieworks/vault"
	"github.com/valkyrieworks/vault/bls12381"
	"github.com/valkyrieworks/vault/ed25519"
	"github.com/valkyrieworks/vault/secp256k1"
	"github.com/valkyrieworks/utils/json"
	pc "github.com/valkyrieworks/schema/consensuscore/vault"
)

//
//
type ErrUnacceptedKey struct {
	Key any
}

func (e ErrUnacceptedKey) Fault() string {
	return fmt.Sprintf("REDACTED", e.Key)
}

//
//
type ErrCorruptKeySize struct {
	Key       any
	Got, Desire int
}

func (e ErrCorruptKeySize) Fault() string {
	return fmt.Sprintf("REDACTED", e.Key, e.Got, e.Desire)
}

func init() {
	json.EnrollKind((*pc.PublicKey)(nil), "REDACTED")
	json.EnrollKind((*pc.Publickey_Ed25519)(nil), "REDACTED")
	json.EnrollKind((*pc.Publickey_Secp256k1)(nil), "REDACTED")
	if bls12381.Activated {
		json.EnrollKind((*pc.Publickey_Bls12381)(nil), "REDACTED")
	}
}

//
func PublicKeyToSchema(k vault.PublicKey) (pc.PublicKey, error) {
	var kp pc.PublicKey
	switch k := k.(type) {
	case ed25519.PublicKey:
		kp = pc.PublicKey{
			Sum: &pc.Publickey_Ed25519{
				Ed25519: k,
			},
		}
	case secp256k1.PublicKey:
		kp = pc.PublicKey{
			Sum: &pc.Publickey_Secp256k1{
				Secp256k1: k,
			},
		}
	case bls12381.PublicKey:
		if !bls12381.Activated {
			return kp, ErrUnacceptedKey{Key: k}
		}

		kp = pc.PublicKey{
			Sum: &pc.Publickey_Bls12381{
				Bls12381: k.Octets(),
			},
		}
	default:
		return kp, fmt.Errorf("REDACTED", k)
	}
	return kp, nil
}

//
func PublicKeyFromSchema(k pc.PublicKey) (vault.PublicKey, error) {
	switch k := k.Sum.(type) {
	case *pc.Publickey_Ed25519:
		if len(k.Ed25519) != ed25519.PublicKeyVolume {
			return nil, fmt.Errorf("REDACTED",
				len(k.Ed25519), ed25519.PublicKeyVolume)
		}
		pk := make(ed25519.PublicKey, ed25519.PublicKeyVolume)
		copy(pk, k.Ed25519)
		return pk, nil
	case *pc.Publickey_Secp256k1:
		if len(k.Secp256k1) != secp256k1.PublicKeyVolume {
			return nil, fmt.Errorf("REDACTED",
				len(k.Secp256k1), secp256k1.PublicKeyVolume)
		}
		pk := make(secp256k1.PublicKey, secp256k1.PublicKeyVolume)
		copy(pk, k.Secp256k1)
		return pk, nil
	case *pc.Publickey_Bls12381:
		if !bls12381.Activated {
			return nil, ErrUnacceptedKey{Key: k}
		}

		if len(k.Bls12381) != bls12381.PublicKeyVolume {
			return nil, ErrCorruptKeySize{
				Key:  k,
				Got:  len(k.Bls12381),
				Desire: bls12381.PublicKeyVolume,
			}
		}
		return bls12381.NewPublicKeyFromOctets(k.Bls12381)
	default:
		return nil, fmt.Errorf("REDACTED", k)
	}
}

//
//
//
func PublicKeyFromKindAndOctets(publicidKind string, octets []byte) (vault.PublicKey, error) {
	var publicKey vault.PublicKey
	switch publicidKind {
	case ed25519.KeyKind:
		if len(octets) != ed25519.PublicKeyVolume {
			return nil, ErrCorruptKeySize{
				Key:  publicidKind,
				Got:  len(octets),
				Desire: ed25519.PublicKeyVolume,
			}
		}

		pk := make(ed25519.PublicKey, ed25519.PublicKeyVolume)
		copy(pk, octets)
		publicKey = pk
	case secp256k1.KeyKind:
		if len(octets) != secp256k1.PublicKeyVolume {
			return nil, ErrCorruptKeySize{
				Key:  publicidKind,
				Got:  len(octets),
				Desire: secp256k1.PublicKeyVolume,
			}
		}

		pk := make(secp256k1.PublicKey, secp256k1.PublicKeyVolume)
		copy(pk, octets)
		publicKey = pk
	case bls12381.KeyKind:
		if !bls12381.Activated {
			return nil, ErrUnacceptedKey{Key: publicidKind}
		}

		if len(octets) != bls12381.PublicKeyVolume {
			return nil, ErrCorruptKeySize{
				Key:  publicidKind,
				Got:  len(octets),
				Desire: bls12381.PublicKeyVolume,
			}
		}

		return bls12381.NewPublicKeyFromOctets(octets)
	default:
		return nil, ErrUnacceptedKey{Key: publicidKind}
	}
	return publicKey, nil
}
