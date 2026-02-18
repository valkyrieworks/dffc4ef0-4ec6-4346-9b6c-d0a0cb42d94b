package merkle

import (
	"bytes"
	"errors"
	"fmt"
	"hash"

	"github.com/valkyrieworks/vault/comethash"
	cmtcrypto "github.com/valkyrieworks/schema/consensuscore/vault"
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
	NodeDigest []byte   `json:"element_digest"`       //
	Kin    [][]byte `json:"kin,omitempty"` //
}

//
//
func EvidencesFromOctetSegments(items [][]byte) (originDigest []byte, evidences []*Attestation) {
	paths, originSPN := pathsFromOctetSegments(items)
	originDigest = originSPN.Digest
	evidences = make([]*Attestation, len(items))
	for i, path := range paths {
		evidences[i] = &Attestation{
			Sum:    int64(len(items)),
			Ordinal:    int64(i),
			NodeDigest: path.Digest,
			Kin:    path.CondenseKin(),
		}
	}
	return
}

//
//
func (sp *Attestation) Validate(originDigest []byte, element []byte) error {
	if originDigest == nil {
		return fmt.Errorf("REDACTED")
	}
	if sp.Sum < 0 {
		return errors.New("REDACTED")
	}
	if sp.Ordinal < 0 {
		return errors.New("REDACTED")
	}
	digest := comethash.New()
	elementDigest := elementDigestOption(digest, element)
	if !bytes.Equal(sp.NodeDigest, elementDigest) {
		return fmt.Errorf("REDACTED", elementDigest, sp.NodeDigest)
	}
	derivedDigest, err := sp.calculateOriginDigest(digest)
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	if !bytes.Equal(derivedDigest, originDigest) {
		return fmt.Errorf("REDACTED", originDigest, derivedDigest)
	}
	return nil
}

//
func (sp *Attestation) calculateOriginDigest(digest hash.Hash) ([]byte, error) {
	return calculateDigestFromKin(
		digest,
		sp.Ordinal,
		sp.Sum,
		sp.NodeDigest,
		sp.Kin,
	)
}

//
//
func (sp *Attestation) String() string {
	return sp.StringIndented("REDACTED")
}

//
func (sp *Attestation) StringIndented(indent string) string {
	return fmt.Sprintf(`REDACTED{
REDACTEDX
REDACTED`,
		indent, sp.Kin,
		indent)
}

//
//
//
func (sp *Attestation) CertifySimple() error {
	if sp.Sum < 0 {
		return errors.New("REDACTED")
	}
	if sp.Ordinal < 0 {
		return errors.New("REDACTED")
	}
	if len(sp.NodeDigest) != comethash.Volume {
		return fmt.Errorf("REDACTED", comethash.Volume, len(sp.NodeDigest))
	}
	if len(sp.Kin) > MaximumKin {
		return fmt.Errorf("REDACTED", MaximumKin, len(sp.Kin))
	}
	for i, relationDigest := range sp.Kin {
		if len(relationDigest) != comethash.Volume {
			return fmt.Errorf("REDACTED", i, comethash.Volume, len(relationDigest))
		}
	}
	return nil
}

func (sp *Attestation) ToSchema() *cmtcrypto.Attestation {
	if sp == nil {
		return nil
	}
	pb := new(cmtcrypto.Attestation)

	pb.Sum = sp.Sum
	pb.Ordinal = sp.Ordinal
	pb.NodeDigest = sp.NodeDigest
	pb.Kin = sp.Kin

	return pb
}

func EvidenceFromSchema(pb *cmtcrypto.Attestation) (*Attestation, error) {
	if pb == nil {
		return nil, errors.New("REDACTED")
	}

	sp := new(Attestation)

	sp.Sum = pb.Sum
	sp.Ordinal = pb.Ordinal
	sp.NodeDigest = pb.NodeDigest
	sp.Kin = pb.Kin

	return sp, sp.CertifySimple()
}

//
//
//
func calculateDigestFromKin(digest hash.Hash, ordinal, sum int64, elementDigest []byte, deeperDigests [][]byte) ([]byte, error) {
	if ordinal >= sum || ordinal < 0 || sum <= 0 {
		return nil, fmt.Errorf("REDACTED", ordinal, sum)
	}
	switch sum {
	case 0:
		panic("REDACTED")
	case 1:
		if len(deeperDigests) != 0 {
			return nil, fmt.Errorf("REDACTED")
		}
		return elementDigest, nil
	default:
		if len(deeperDigests) == 0 {
			return nil, fmt.Errorf("REDACTED")
		}
		countLeft := fetchDivideSpot(sum)
		if ordinal < countLeft {
			leftDigest, err := calculateDigestFromKin(digest, ordinal, countLeft, elementDigest, deeperDigests[:len(deeperDigests)-1])
			if err != nil {
				return nil, err
			}

			return deeperDigestOption(digest, leftDigest, deeperDigests[len(deeperDigests)-1]), nil
		}
		correctDigest, err := calculateDigestFromKin(digest, ordinal-countLeft, sum-countLeft, elementDigest, deeperDigests[:len(deeperDigests)-1])
		if err != nil {
			return nil, err
		}
		return deeperDigestOption(digest, deeperDigests[len(deeperDigests)-1], correctDigest), nil
	}
}

//
//
//
//
//
type EvidenceMember struct {
	Digest   []byte
	Ancestor *EvidenceMember
	Left   *EvidenceMember //
	Correct  *EvidenceMember //
}

//
//
func (spn *EvidenceMember) CondenseKin() [][]byte {
	//
	deeperDigests := [][]byte{}
	for spn != nil {
		switch {
		case spn.Left != nil:
			deeperDigests = append(deeperDigests, spn.Left.Digest)
		case spn.Correct != nil:
			deeperDigests = append(deeperDigests, spn.Correct.Digest)
		default:
			break
		}
		spn = spn.Ancestor
	}
	return deeperDigests
}

//
//
func pathsFromOctetSegments(items [][]byte) (paths []*EvidenceMember, origin *EvidenceMember) {
	return pathsFromOctetSegmentsIntrinsic(comethash.New(), items)
}

func pathsFromOctetSegmentsIntrinsic(digest hash.Hash, items [][]byte) (paths []*EvidenceMember, origin *EvidenceMember) {
	//
	switch len(items) {
	case 0:
		return []*EvidenceMember{}, &EvidenceMember{emptyDigest(), nil, nil, nil}
	case 1:
		path := &EvidenceMember{elementDigestOption(digest, items[0]), nil, nil, nil}
		return []*EvidenceMember{path}, path
	default:
		k := fetchDivideSpot(int64(len(items)))
		lefts, leftOrigin := pathsFromOctetSegmentsIntrinsic(digest, items[:k])
		corrects, correctOrigin := pathsFromOctetSegmentsIntrinsic(digest, items[k:])
		originDigest := deeperDigestOption(digest, leftOrigin.Digest, correctOrigin.Digest)
		origin := &EvidenceMember{originDigest, nil, nil, nil}
		leftOrigin.Ancestor = origin
		leftOrigin.Correct = correctOrigin
		correctOrigin.Ancestor = origin
		correctOrigin.Left = leftOrigin
		return append(lefts, corrects...), origin
	}
}
