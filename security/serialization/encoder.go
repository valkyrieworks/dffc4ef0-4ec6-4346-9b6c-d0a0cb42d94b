package serialization

import (
	"fmt"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/signature381"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/edwards25519"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/ellipticp256"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/jsn"
	pc "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/security"
)

//
//
type FaultUnservicedToken struct {
	Key any
}

func (e FaultUnservicedToken) Failure() string {
	return fmt.Sprintf("REDACTED", e.Key)
}

//
//
type FaultUnfitTokenSize struct {
	Key       any
	Got, Desire int
}

func (e FaultUnfitTokenSize) Failure() string {
	return fmt.Sprintf("REDACTED", e.Key, e.Got, e.Desire)
}

func initialize() {
	jsn.EnrollKind((*pc.CommonToken)(nil), "REDACTED")
	jsn.EnrollKind((*pc.Commonkey_Edwards25519)(nil), "REDACTED")
	jsn.EnrollKind((*pc.Commonkey_Ellipticp256)(nil), "REDACTED")
	if signature381.Activated {
		jsn.EnrollKind((*pc.Commonkey_Signature381)(nil), "REDACTED")
	}
}

//
func PublicTokenTowardSchema(k security.PublicToken) (pc.CommonToken, error) {
	var kp pc.CommonToken
	switch k := k.(type) {
	case edwards25519.PublicToken:
		kp = pc.CommonToken{
			Sum: &pc.Commonkey_Edwards25519{
				Edwards25519: k,
			},
		}
	case ellipticp256.PublicToken:
		kp = pc.CommonToken{
			Sum: &pc.Commonkey_Ellipticp256{
				Ellipticp256: k,
			},
		}
	case signature381.PublicToken:
		if !signature381.Activated {
			return kp, FaultUnservicedToken{Key: k}
		}

		kp = pc.CommonToken{
			Sum: &pc.Commonkey_Signature381{
				Signature381: k.Octets(),
			},
		}
	default:
		return kp, fmt.Errorf("REDACTED", k)
	}
	return kp, nil
}

//
func PublicTokenOriginatingSchema(k pc.CommonToken) (security.PublicToken, error) {
	switch k := k.Sum.(type) {
	case *pc.Commonkey_Edwards25519:
		if len(k.Edwards25519) != edwards25519.PublicTokenExtent {
			return nil, fmt.Errorf("REDACTED",
				len(k.Edwards25519), edwards25519.PublicTokenExtent)
		}
		pk := make(edwards25519.PublicToken, edwards25519.PublicTokenExtent)
		copy(pk, k.Edwards25519)
		return pk, nil
	case *pc.Commonkey_Ellipticp256:
		if len(k.Ellipticp256) != ellipticp256.PublicTokenExtent {
			return nil, fmt.Errorf("REDACTED",
				len(k.Ellipticp256), ellipticp256.PublicTokenExtent)
		}
		pk := make(ellipticp256.PublicToken, ellipticp256.PublicTokenExtent)
		copy(pk, k.Ellipticp256)
		return pk, nil
	case *pc.Commonkey_Signature381:
		if !signature381.Activated {
			return nil, FaultUnservicedToken{Key: k}
		}

		if len(k.Signature381) != signature381.PublicTokenExtent {
			return nil, FaultUnfitTokenSize{
				Key:  k,
				Got:  len(k.Signature381),
				Desire: signature381.PublicTokenExtent,
			}
		}
		return signature381.FreshCommonTokenOriginatingOctets(k.Signature381)
	default:
		return nil, fmt.Errorf("REDACTED", k)
	}
}

//
//
//
func PublicTokenOriginatingKindAlsoOctets(keyKind string, octets []byte) (security.PublicToken, error) {
	var publicToken security.PublicToken
	switch keyKind {
	case edwards25519.TokenKind:
		if len(octets) != edwards25519.PublicTokenExtent {
			return nil, FaultUnfitTokenSize{
				Key:  keyKind,
				Got:  len(octets),
				Desire: edwards25519.PublicTokenExtent,
			}
		}

		pk := make(edwards25519.PublicToken, edwards25519.PublicTokenExtent)
		copy(pk, octets)
		publicToken = pk
	case ellipticp256.TokenKind:
		if len(octets) != ellipticp256.PublicTokenExtent {
			return nil, FaultUnfitTokenSize{
				Key:  keyKind,
				Got:  len(octets),
				Desire: ellipticp256.PublicTokenExtent,
			}
		}

		pk := make(ellipticp256.PublicToken, ellipticp256.PublicTokenExtent)
		copy(pk, octets)
		publicToken = pk
	case signature381.TokenKind:
		if !signature381.Activated {
			return nil, FaultUnservicedToken{Key: keyKind}
		}

		if len(octets) != signature381.PublicTokenExtent {
			return nil, FaultUnfitTokenSize{
				Key:  keyKind,
				Got:  len(octets),
				Desire: signature381.PublicTokenExtent,
			}
		}

		return signature381.FreshCommonTokenOriginatingOctets(octets)
	default:
		return nil, FaultUnservicedToken{Key: keyKind}
	}
	return publicToken, nil
}
