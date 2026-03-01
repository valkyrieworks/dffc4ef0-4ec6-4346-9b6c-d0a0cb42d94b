package link

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"testing"

	gogotypes "github.com/cosmos/gogoproto/types"
	"github.com/oasisprotocol/curve25519-voi/primitives/merlin"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/chacha20poly1305"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/edwards25519"
	cryptocode "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/serialization"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/protocolio"
	tmpfabric "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/p2p"
)

type reserve struct {
	following bytes.Buffer
}

func (b *reserve) Obtain(data []byte) (n int, err error) {
	return b.following.Read(data)
}

func (b *reserve) Record(data []byte) (n int, err error) {
	return b.following.Write(data)
}

func (b *reserve) Octets() []byte {
	return b.following.Bytes()
}

func (b *reserve) Shutdown() error {
	return nil
}

type maliciousLink struct {
	credentialLink *CredentialLinkage
	reserve     *reserve

	positionTempPublic  *[32]byte
	positionTempPrivate *[32]byte
	modTempPublic  *[32]byte
	privateToken    security.PrivateToken

	fetchPhase   int
	persistPhase  int
	fetchDisplacement int

	allocateTempToken        bool
	flawedTempToken          bool
	allocateAuthNotation bool
	flawedAuthNotation   bool
}

func freshMaliciousLink(allocateTempToken, flawedTempToken, allocateAuthNotation, flawedAuthNotation bool) *maliciousLink {
	privateToken := edwards25519.ProducePrivateToken()
	positionTempPublic, positionTempPrivate := produceTempTokens()
	var rep [32]byte
	c := &maliciousLink{
		positionTempPublic:  positionTempPublic,
		positionTempPrivate: positionTempPrivate,
		modTempPublic:  &rep,
		privateToken:    privateToken,

		allocateTempToken:        allocateTempToken,
		flawedTempToken:          flawedTempToken,
		allocateAuthNotation: allocateAuthNotation,
		flawedAuthNotation:   flawedAuthNotation,
	}

	return c
}

func (c *maliciousLink) Obtain(data []byte) (n int, err error) {
	if !c.allocateTempToken {
		return 0, io.EOF
	}

	switch c.fetchPhase {
	case 0:
		if !c.flawedTempToken {
			lc := *c.positionTempPublic
			bz, err := protocolio.SerializeSeparated(&gogotypes.BytesValue{Value: lc[:]})
			if err != nil {
				panic(err)
			}
			copy(data, bz[c.fetchDisplacement:])
			n = len(data)
		} else {
			bz, err := protocolio.SerializeSeparated(&gogotypes.BytesValue{Value: []byte("REDACTED")})
			if err != nil {
				panic(err)
			}
			copy(data, bz)
			n = len(data)
		}
		c.fetchDisplacement += n

		if n >= 32 {
			c.fetchDisplacement = 0
			c.fetchPhase = 1
			if !c.allocateAuthNotation {
				c.fetchPhase = 2
			}
		}

		return n, nil
	case 1:
		signing := c.attestQuery()
		if !c.flawedAuthNotation {
			pkpb, err := cryptocode.PublicTokenTowardSchema(c.privateToken.PublicToken())
			if err != nil {
				panic(err)
			}
			bz, err := protocolio.SerializeSeparated(&tmpfabric.AuthSignatureArtifact{PublicToken: pkpb, Sig: signing})
			if err != nil {
				panic(err)
			}
			n, err = c.credentialLink.Record(bz)
			if err != nil {
				panic(err)
			}
			if c.fetchDisplacement > len(c.reserve.Octets()) {
				return len(data), nil
			}
			copy(data, c.reserve.Octets()[c.fetchDisplacement:])
		} else {
			bz, err := protocolio.SerializeSeparated(&gogotypes.BytesValue{Value: []byte("REDACTED")})
			if err != nil {
				panic(err)
			}
			n, err = c.credentialLink.Record(bz)
			if err != nil {
				panic(err)
			}
			if c.fetchDisplacement > len(c.reserve.Octets()) {
				return len(data), nil
			}
			copy(data, c.reserve.Octets())
		}
		c.fetchDisplacement += len(data)
		return n, nil
	default:
		return 0, io.EOF
	}
}

func (c *maliciousLink) Record(data []byte) (n int, err error) {
	switch c.persistPhase {
	case 0:
		var (
			octets     gogotypes.BytesValue
			modTempPublic [32]byte
		)
		err := protocolio.DecodeSeparated(data, &octets)
		if err != nil {
			panic(err)
		}
		copy(modTempPublic[:], bytes.Value)
		c.modTempPublic = &modTempPublic
		c.persistPhase = 1
		if !c.allocateAuthNotation {
			c.persistPhase = 2
		}
		return len(data), nil
	case 1:
		//
		return len(data), nil
	default:
		return 0, io.EOF
	}
}

func (c *maliciousLink) Shutdown() error {
	return nil
}

func (c *maliciousLink) attestQuery() []byte {
	//
	minimumTempPublic, highTempPublic := sort32(c.positionTempPublic, c.modTempPublic)

	record := merlin.NewTranscript("REDACTED")

	record.AppendMessage(tagTemporaryLesserCommonToken, minimumTempPublic[:])
	record.AppendMessage(tagTemporaryHigherCommonToken, highTempPublic[:])

	//
	//
	positionEqualsMinimal := bytes.Equal(c.positionTempPublic[:], minimumTempPublic[:])

	//
	dhCredential, err := calculateDHCredential(c.modTempPublic, c.positionTempPrivate)
	if err != nil {
		panic(err)
	}

	record.AppendMessage(tagDHCredential, dhCredential[:])

	//
	//
	//
	obtainCredential, transmitCredential := deduceCredentials(dhCredential, positionEqualsMinimal)

	const queryExtent = 32
	var query [queryExtent]byte
	record.ExtractBytes(query[:], tagCredentialLinkageMac)

	transmitAead, err := chacha20poly1305.New(transmitCredential[:])
	if err != nil {
		panic(errors.New("REDACTED"))
	}
	obtainAead, err := chacha20poly1305.New(obtainCredential[:])
	if err != nil {
		panic(errors.New("REDACTED"))
	}

	b := &reserve{}
	c.credentialLink = &CredentialLinkage{
		link:            b,
		linkPersistor:      bufio.NewWriterSize(b, fallbackPersistReserveExtent),
		linkFetcher:      b,
		obtainReserve:      nil,
		obtainNumber:       new([aeadNumberExtent]byte),
		transmitNumber:       new([aeadNumberExtent]byte),
		obtainAead:        obtainAead,
		transmitAead:        transmitAead,
		obtainStructure:       make([]byte, sumStructureExtent),
		obtainSecuredStructure: make([]byte, sumStructureExtent+aeadExtentMargin),
		transmitStructure:       make([]byte, sumStructureExtent),
		transmitSecuredStructure: make([]byte, sumStructureExtent+aeadExtentMargin),
	}
	c.reserve = b

	//
	positionNotation, err := attestQuery(&query, c.privateToken)
	if err != nil {
		panic(err)
	}

	return positionNotation
}

//
//
func VerifyCreateCredentialLinkage(t *testing.T) {
	verifyScenarios := []struct {
		alias   string
		link   *maliciousLink
		faultSignal string
	}{
		{"REDACTED", freshMaliciousLink(false, false, false, false), "REDACTED"},
		{"REDACTED", freshMaliciousLink(true, true, false, false), "REDACTED"},
		{"REDACTED", freshMaliciousLink(true, false, false, false), "REDACTED"},
		{"REDACTED", freshMaliciousLink(true, false, true, true), "REDACTED"},
	}

	for _, tc := range verifyScenarios {
		t.Run(tc.alias, func(t *testing.T) {
			privateToken := edwards25519.ProducePrivateToken()
			_, err := CreateCredentialLinkage(tc.link, privateToken)
			if tc.faultSignal != "REDACTED" {
				if assert.Error(t, err) {
					assert.Contains(t, err.Error(), tc.faultSignal)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
