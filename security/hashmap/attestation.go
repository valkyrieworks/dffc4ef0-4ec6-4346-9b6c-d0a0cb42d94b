package hashmap

import (
	"bytes"
	"errors"
	"fmt"
	"hash"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/tenderminthash"
	cmtsecurity "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/security"
)

const (
	//
	//
	//
	MaximumKin = 100
)

//
//
//
//
//
//
//
type Attestation struct {
	Sum    int64    `json:"sum"`           //
	Ordinal    int64    `json:"ordinal"`           //
	NodeDigest []byte   `json:"terminal_digest"`       //
	Kin    [][]byte `json:"kin,omitempty"` //
}

//
//
func AttestationsOriginatingOctetSegments(elements [][]byte) (originDigest []byte, attestations []*Attestation) {
	paths, originIdentifier := pathsOriginatingOctetSegments(elements)
	originDigest = originIdentifier.Digest
	attestations = make([]*Attestation, len(elements))
	for i, pathway := range paths {
		attestations[i] = &Attestation{
			Sum:    int64(len(elements)),
			Ordinal:    int64(i),
			NodeDigest: pathway.Digest,
			Kin:    pathway.CondenseKin(),
		}
	}
	return
}

//
//
func (sp *Attestation) Validate(originDigest []byte, terminal []byte) error {
	if originDigest == nil {
		return fmt.Errorf("REDACTED")
	}
	if sp.Sum < 0 {
		return errors.New("REDACTED")
	}
	if sp.Ordinal < 0 {
		return errors.New("REDACTED")
	}
	digest := tenderminthash.New()
	terminalDigest := terminalDigestSetting(digest, terminal)
	if !bytes.Equal(sp.NodeDigest, terminalDigest) {
		return fmt.Errorf("REDACTED", terminalDigest, sp.NodeDigest)
	}
	estimatedDigest, err := sp.calculateOriginDigest(digest)
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	if !bytes.Equal(estimatedDigest, originDigest) {
		return fmt.Errorf("REDACTED", originDigest, estimatedDigest)
	}
	return nil
}

//
func (sp *Attestation) calculateOriginDigest(digest hash.Hash) ([]byte, error) {
	return calculateDigestOriginatingKin(
		digest,
		sp.Ordinal,
		sp.Sum,
		sp.NodeDigest,
		sp.Kin,
	)
}

//
//
func (sp *Attestation) Text() string {
	return sp.TextFormatted("REDACTED")
}

//
func (sp *Attestation) TextFormatted(format string) string {
	return fmt.Sprintf(`REDACTED{
REDACTEDX
REDACTED`,
		format, sp.Kin,
		format)
}

//
//
//
func (sp *Attestation) CertifyFundamental() error {
	if sp.Sum < 0 {
		return errors.New("REDACTED")
	}
	if sp.Ordinal < 0 {
		return errors.New("REDACTED")
	}
	if len(sp.NodeDigest) != tenderminthash.Extent {
		return fmt.Errorf("REDACTED", tenderminthash.Extent, len(sp.NodeDigest))
	}
	if len(sp.Kin) > MaximumKin {
		return fmt.Errorf("REDACTED", MaximumKin, len(sp.Kin))
	}
	for i, kinshipDigest := range sp.Kin {
		if len(kinshipDigest) != tenderminthash.Extent {
			return fmt.Errorf("REDACTED", i, tenderminthash.Extent, len(kinshipDigest))
		}
	}
	return nil
}

func (sp *Attestation) TowardSchema() *cmtsecurity.Attestation {
	if sp == nil {
		return nil
	}
	pb := new(cmtsecurity.Attestation)

	pb.Sum = sp.Sum
	pb.Ordinal = sp.Ordinal
	pb.NodeDigest = sp.NodeDigest
	pb.Kin = sp.Kin

	return pb
}

func AttestationOriginatingSchema(pb *cmtsecurity.Attestation) (*Attestation, error) {
	if pb == nil {
		return nil, errors.New("REDACTED")
	}

	sp := new(Attestation)

	sp.Sum = pb.Sum
	sp.Ordinal = pb.Ordinal
	sp.NodeDigest = pb.NodeDigest
	sp.Kin = pb.Kin

	return sp, sp.CertifyFundamental()
}

//
//
//
func calculateDigestOriginatingKin(digest hash.Hash, ordinal, sum int64, terminalDigest []byte, internalDigests [][]byte) ([]byte, error) {
	if ordinal >= sum || ordinal < 0 || sum <= 0 {
		return nil, fmt.Errorf("REDACTED", ordinal, sum)
	}
	switch sum {
	case 0:
		panic("REDACTED")
	case 1:
		if len(internalDigests) != 0 {
			return nil, fmt.Errorf("REDACTED")
		}
		return terminalDigest, nil
	default:
		if len(internalDigests) == 0 {
			return nil, fmt.Errorf("REDACTED")
		}
		countLeading := obtainPartitionNode(sum)
		if ordinal < countLeading {
			leadingDigest, err := calculateDigestOriginatingKin(digest, ordinal, countLeading, terminalDigest, internalDigests[:len(internalDigests)-1])
			if err != nil {
				return nil, err
			}

			return internalDigestSetting(digest, leadingDigest, internalDigests[len(internalDigests)-1]), nil
		}
		trailingDigest, err := calculateDigestOriginatingKin(digest, ordinal-countLeading, sum-countLeading, terminalDigest, internalDigests[:len(internalDigests)-1])
		if err != nil {
			return nil, err
		}
		return internalDigestSetting(digest, internalDigests[len(internalDigests)-1], trailingDigest), nil
	}
}

//
//
//
//
//
type AttestationPeer struct {
	Digest   []byte
	Ancestor *AttestationPeer
	Leading   *AttestationPeer //
	Trailing  *AttestationPeer //
}

//
//
func (spn *AttestationPeer) CondenseKin() [][]byte {
	//
	internalDigests := [][]byte{}
	for spn != nil {
		switch {
		case spn.Leading != nil:
			internalDigests = append(internalDigests, spn.Leading.Digest)
		case spn.Trailing != nil:
			internalDigests = append(internalDigests, spn.Trailing.Digest)
		default:
			break
		}
		spn = spn.Ancestor
	}
	return internalDigests
}

//
//
func pathsOriginatingOctetSegments(elements [][]byte) (paths []*AttestationPeer, origin *AttestationPeer) {
	return pathsOriginatingOctetSegmentsIntrinsic(tenderminthash.New(), elements)
}

func pathsOriginatingOctetSegmentsIntrinsic(digest hash.Hash, elements [][]byte) (paths []*AttestationPeer, origin *AttestationPeer) {
	//
	switch len(elements) {
	case 0:
		return []*AttestationPeer{}, &AttestationPeer{blankDigest(), nil, nil, nil}
	case 1:
		pathway := &AttestationPeer{terminalDigestSetting(digest, elements[0]), nil, nil, nil}
		return []*AttestationPeer{pathway}, pathway
	default:
		k := obtainPartitionNode(int64(len(elements)))
		leading, leadingOrigin := pathsOriginatingOctetSegmentsIntrinsic(digest, elements[:k])
		trailing, trailingOrigin := pathsOriginatingOctetSegmentsIntrinsic(digest, elements[k:])
		originDigest := internalDigestSetting(digest, leadingOrigin.Digest, trailingOrigin.Digest)
		origin := &AttestationPeer{originDigest, nil, nil, nil}
		leadingOrigin.Ancestor = origin
		leadingOrigin.Trailing = trailingOrigin
		trailingOrigin.Ancestor = origin
		trailingOrigin.Leading = leadingOrigin
		return append(leading, trailing...), origin
	}
}
