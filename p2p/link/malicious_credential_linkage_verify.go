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

	"github.com/valkyrieworks/vault"
	"github.com/valkyrieworks/vault/ed25519"
	cryptocode "github.com/valkyrieworks/vault/codec"
	"github.com/valkyrieworks/utils/protoio"
	tmp2p "github.com/valkyrieworks/schema/consensuscore/p2p"
)

type buffer struct {
	following bytes.Buffer
}

func (b *buffer) Scan(data []byte) (n int, err error) {
	return b.following.Read(data)
}

func (b *buffer) Record(data []byte) (n int, err error) {
	return b.following.Write(data)
}

func (b *buffer) Octets() []byte {
	return b.following.Bytes()
}

func (b *buffer) End() error {
	return nil
}

type maliciousLink struct {
	tokenLink *TokenLinkage
	buffer     *buffer

	locationEphPublic  *[32]byte
	locationEphPrivate *[32]byte
	modEphPublic  *[32]byte
	privateKey    vault.PrivateKey

	readPhase   int
	recordPhase  int
	readDisplacement int

	allocateEphKey        bool
	flawedEphKey          bool
	allocateAuthAutograph bool
	flawedAuthAutograph   bool
}

func newMaliciousLink(allocateEphKey, flawedEphKey, allocateAuthAutograph, flawedAuthAutograph bool) *maliciousLink {
	privateKey := ed25519.GeneratePrivateKey()
	locationEphPublic, locationEphPrivate := generateEphKeys()
	var rep [32]byte
	c := &maliciousLink{
		locationEphPublic:  locationEphPublic,
		locationEphPrivate: locationEphPrivate,
		modEphPublic:  &rep,
		privateKey:    privateKey,

		allocateEphKey:        allocateEphKey,
		flawedEphKey:          flawedEphKey,
		allocateAuthAutograph: allocateAuthAutograph,
		flawedAuthAutograph:   flawedAuthAutograph,
	}

	return c
}

func (c *maliciousLink) Scan(data []byte) (n int, err error) {
	if !c.allocateEphKey {
		return 0, io.EOF
	}

	switch c.readPhase {
	case 0:
		if !c.flawedEphKey {
			lc := *c.locationEphPublic
			bz, err := protoio.SerializeSeparated(&gogotypes.BytesValue{Value: lc[:]})
			if err != nil {
				panic(err)
			}
			copy(data, bz[c.readDisplacement:])
			n = len(data)
		} else {
			bz, err := protoio.SerializeSeparated(&gogotypes.BytesValue{Value: []byte("REDACTED")})
			if err != nil {
				panic(err)
			}
			copy(data, bz)
			n = len(data)
		}
		c.readDisplacement += n

		if n >= 32 {
			c.readDisplacement = 0
			c.readPhase = 1
			if !c.allocateAuthAutograph {
				c.readPhase = 2
			}
		}

		return n, nil
	case 1:
		autograph := c.attestDispute()
		if !c.flawedAuthAutograph {
			pkpb, err := cryptocode.PublicKeyToSchema(c.privateKey.PublicKey())
			if err != nil {
				panic(err)
			}
			bz, err := protoio.SerializeSeparated(&tmp2p.AuthSignatureSignal{PublicKey: pkpb, Sig: autograph})
			if err != nil {
				panic(err)
			}
			n, err = c.tokenLink.Record(bz)
			if err != nil {
				panic(err)
			}
			if c.readDisplacement > len(c.buffer.Octets()) {
				return len(data), nil
			}
			copy(data, c.buffer.Octets()[c.readDisplacement:])
		} else {
			bz, err := protoio.SerializeSeparated(&gogotypes.BytesValue{Value: []byte("REDACTED")})
			if err != nil {
				panic(err)
			}
			n, err = c.tokenLink.Record(bz)
			if err != nil {
				panic(err)
			}
			if c.readDisplacement > len(c.buffer.Octets()) {
				return len(data), nil
			}
			copy(data, c.buffer.Octets())
		}
		c.readDisplacement += len(data)
		return n, nil
	default:
		return 0, io.EOF
	}
}

func (c *maliciousLink) Record(data []byte) (n int, err error) {
	switch c.recordPhase {
	case 0:
		var (
			octets     gogotypes.BytesValue
			modEphPublic [32]byte
		)
		err := protoio.UnserializeSeparated(data, &octets)
		if err != nil {
			panic(err)
		}
		copy(modEphPublic[:], bytes.Value)
		c.modEphPublic = &modEphPublic
		c.recordPhase = 1
		if !c.allocateAuthAutograph {
			c.recordPhase = 2
		}
		return len(data), nil
	case 1:
		//
		return len(data), nil
	default:
		return 0, io.EOF
	}
}

func (c *maliciousLink) End() error {
	return nil
}

func (c *maliciousLink) attestDispute() []byte {
	//
	loEphPublic, greetingEphPublic := sort32(c.locationEphPublic, c.modEphPublic)

	log := merlin.NewTranscript("REDACTED")

	log.AppendMessage(tagTemporaryLesserExternalKey, loEphPublic[:])
	log.AppendMessage(tagTemporaryUpperExternalKey, greetingEphPublic[:])

	//
	//
	locationIsMinimum := bytes.Equal(c.locationEphPublic[:], loEphPublic[:])

	//
	dhCredential, err := calculateDHCredential(c.modEphPublic, c.locationEphPrivate)
	if err != nil {
		panic(err)
	}

	log.AppendMessage(tagDHCredential, dhCredential[:])

	//
	//
	//
	receiveCredential, transmitCredential := deduceCredentials(dhCredential, locationIsMinimum)

	const disputeVolume = 32
	var dispute [disputeVolume]byte
	log.ExtractBytes(dispute[:], tagCredentialLinkageMac)

	transmitAead, err := chacha20poly1305.New(transmitCredential[:])
	if err != nil {
		panic(errors.New("REDACTED"))
	}
	receiveAead, err := chacha20poly1305.New(receiveCredential[:])
	if err != nil {
		panic(errors.New("REDACTED"))
	}

	b := &buffer{}
	c.tokenLink = &TokenLinkage{
		link:            b,
		linkRecorder:      bufio.NewWriterSize(b, standardRecordFrameVolume),
		linkScanner:      b,
		receiveFrame:      nil,
		receiveNonce:       new([aeadNonceVolume]byte),
		transmitNonce:       new([aeadNonceVolume]byte),
		receiveAead:        receiveAead,
		transmitAead:        transmitAead,
		receiveBorder:       make([]byte, sumBorderVolume),
		receiveSecuredBorder: make([]byte, sumBorderVolume+aeadVolumeBurden),
		transmitBorder:       make([]byte, sumBorderVolume),
		transmitSecuredBorder: make([]byte, sumBorderVolume+aeadVolumeBurden),
	}
	c.buffer = b

	//
	locationAutograph, err := attestDispute(&dispute, c.privateKey)
	if err != nil {
		panic(err)
	}

	return locationAutograph
}

//
//
func VerifyCreateCredentialLinkage(t *testing.T) {
	verifyScenarios := []struct {
		label   string
		link   *maliciousLink
		errMessage string
	}{
		{"REDACTED", newMaliciousLink(false, false, false, false), "REDACTED"},
		{"REDACTED", newMaliciousLink(true, true, false, false), "REDACTED"},
		{"REDACTED", newMaliciousLink(true, false, false, false), "REDACTED"},
		{"REDACTED", newMaliciousLink(true, false, true, true), "REDACTED"},
	}

	for _, tc := range verifyScenarios {
		t.Run(tc.label, func(t *testing.T) {
			privateKey := ed25519.GeneratePrivateKey()
			_, err := CreateTokenLinkage(tc.link, privateKey)
			if tc.errMessage != "REDACTED" {
				if assert.Error(t, err) {
					assert.Contains(t, err.Error(), tc.errMessage)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
