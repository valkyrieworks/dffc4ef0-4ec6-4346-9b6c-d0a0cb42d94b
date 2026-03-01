package link

import (
	"bufio"
	"bytes"
	"crypto/cipher"
	crand "crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math"
	"net"
	"time"

	gogotypes "github.com/cosmos/gogoproto/types"
	"github.com/oasisprotocol/curve25519-voi/primitives/merlin"
	"golang.org/x/crypto/chacha20poly1305"
	"golang.org/x/crypto/curve25519"
	"golang.org/x/crypto/hkdf"
	"golang.org/x/crypto/nacl/box"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/edwards25519"
	cryptocode "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/serialization"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/asyncronous"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/protocolio"
	commitchronize "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
	tmpfabric "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/p2p"
)

//
const (
	dataLengthExtent      = 4
	dataMaximumExtent      = 1024
	sumStructureExtent   = dataMaximumExtent + dataLengthExtent
	aeadExtentMargin = 16 //
	aeadTokenExtent      = chacha20poly1305.KeySize
	aeadNumberExtent    = chacha20poly1305.NonceSize

	tagTemporaryLesserCommonToken = "REDACTED"
	tagTemporaryHigherCommonToken = "REDACTED"
	tagDHCredential                = "REDACTED"
	tagCredentialLinkageMac     = "REDACTED"

	fallbackPersistReserveExtent = 128 * 1024
	//
	//
	fallbackFetchReserveExtent = 65 * 1024
)

var (
	FaultMinorSequenceDistantPublicToken = errors.New("REDACTED")

	credentialLinkTokenAlsoQueryProduce = []byte("REDACTED")
)

//
//
//
//
//
//
//
//
//
type CredentialLinkage struct {
	//
	obtainAead cipher.AEAD
	transmitAead cipher.AEAD

	modPublicToken security.PublicToken

	link       io.ReadWriteCloser
	linkPersistor *bufio.Writer
	linkFetcher io.Reader

	//
	//
	//
	//
	//
	//
	//
	obtainMutex         commitchronize.Exclusion
	obtainReserve      []byte
	obtainNumber       *[aeadNumberExtent]byte
	obtainStructure       []byte
	obtainSecuredStructure []byte

	transmitMutex         commitchronize.Exclusion
	transmitNumber       *[aeadNumberExtent]byte
	transmitStructure       []byte
	transmitSecuredStructure []byte
}

//
//
//
//
//
func CreateCredentialLinkage(link io.ReadWriteCloser, positionPrivateToken security.PrivateToken) (*CredentialLinkage, error) {
	positionPublicToken := positionPrivateToken.PublicToken()

	//
	positionTempPublic, positionTempPrivate := produceTempTokens()

	//
	//
	//
	modTempPublic, err := allocateTempPublicToken(link, positionTempPublic)
	if err != nil {
		return nil, err
	}

	//
	minimumTempPublic, highTempPublic := sort32(positionTempPublic, modTempPublic)

	record := merlin.NewTranscript("REDACTED")

	record.AppendMessage(tagTemporaryLesserCommonToken, minimumTempPublic[:])
	record.AppendMessage(tagTemporaryHigherCommonToken, highTempPublic[:])

	//
	//
	positionEqualsMinimal := bytes.Equal(positionTempPublic[:], minimumTempPublic[:])

	//
	dhCredential, err := calculateDHCredential(modTempPublic, positionTempPrivate)
	if err != nil {
		return nil, err
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
		return nil, errors.New("REDACTED")
	}
	obtainAead, err := chacha20poly1305.New(obtainCredential[:])
	if err != nil {
		return nil, errors.New("REDACTED")
	}

	sc := &CredentialLinkage{
		link:            link,
		linkPersistor:      bufio.NewWriterSize(link, fallbackPersistReserveExtent),
		linkFetcher:      bufio.NewReaderSize(link, fallbackFetchReserveExtent),
		obtainReserve:      nil,
		obtainNumber:       new([aeadNumberExtent]byte),
		transmitNumber:       new([aeadNumberExtent]byte),
		obtainAead:        obtainAead,
		transmitAead:        transmitAead,
		obtainStructure:       make([]byte, sumStructureExtent),
		obtainSecuredStructure: make([]byte, aeadExtentMargin+sumStructureExtent),
		transmitStructure:       make([]byte, sumStructureExtent),
		transmitSecuredStructure: make([]byte, aeadExtentMargin+sumStructureExtent),
	}

	//
	positionNotation, err := attestQuery(&query, positionPrivateToken)
	if err != nil {
		return nil, err
	}

	//
	authSignatureSignal, err := allocateAuthNotation(sc, positionPublicToken, positionNotation)
	if err != nil {
		return nil, err
	}

	modPublicToken, modNotation := authSignatureSignal.Key, authSignatureSignal.Sig
	if _, ok := modPublicToken.(edwards25519.PublicToken); !ok {
		return nil, fmt.Errorf("REDACTED", modPublicToken)
	}
	if !modPublicToken.ValidateNotation(query[:], modNotation) {
		return nil, errors.New("REDACTED")
	}

	//
	sc.modPublicToken = modPublicToken
	return sc, nil
}

//
func (sc *CredentialLinkage) DistantPublicToken() security.PublicToken {
	return sc.modPublicToken
}

//
//
func (sc *CredentialLinkage) Record(data []byte) (n int, err error) {
	sc.transmitMutex.Lock()
	defer sc.transmitMutex.Unlock()
	securedStructure, structure := sc.transmitSecuredStructure, sc.transmitStructure

	for 0 < len(data) {
		if err := func() error {
			var segment []byte
			if dataMaximumExtent < len(data) {
				segment = data[:dataMaximumExtent]
				data = data[dataMaximumExtent:]
			} else {
				segment = data
				data = nil
			}
			segmentMagnitude := len(segment)
			binary.LittleEndian.PutUint32(structure, uint32(segmentMagnitude))
			copy(structure[dataLengthExtent:], segment)

			//
			sc.transmitAead.Seal(securedStructure[:0], sc.transmitNumber[:], structure, nil)
			increaseNumber(sc.transmitNumber)
			//

			_, err = sc.linkPersistor.Write(securedStructure)
			if err != nil {
				return err
			}
			n += len(segment)
			return nil
		}(); err != nil {
			return n, err
		}
	}
	sc.linkPersistor.Flush()
	return n, err
}

//
func (sc *CredentialLinkage) Obtain(data []byte) (n int, err error) {
	sc.obtainMutex.Lock()
	defer sc.obtainMutex.Unlock()

	//
	if 0 < len(sc.obtainReserve) {
		n = copy(data, sc.obtainReserve)
		sc.obtainReserve = sc.obtainReserve[n:]
		return n, err
	}

	//
	securedStructure := sc.obtainSecuredStructure
	_, err = io.ReadFull(sc.linkFetcher, securedStructure)
	if err != nil {
		return n, err
	}

	//
	//
	structure := sc.obtainStructure
	_, err = sc.obtainAead.Open(structure[:0], sc.obtainNumber[:], securedStructure, nil)
	if err != nil {
		return n, fmt.Errorf("REDACTED", err)
	}
	increaseNumber(sc.obtainNumber)
	//

	//
	//
	segmentMagnitude := binary.LittleEndian.Uint32(structure) //
	if segmentMagnitude > dataMaximumExtent {
		return 0, errors.New("REDACTED")
	}
	segment := structure[dataLengthExtent : dataLengthExtent+segmentMagnitude]
	n = copy(data, segment)
	if n < len(segment) {
		sc.obtainReserve = make([]byte, len(segment)-n)
		copy(sc.obtainReserve, segment[n:])
	}
	return n, err
}

//
func (sc *CredentialLinkage) Shutdown() error                  { return sc.link.Close() }
func (sc *CredentialLinkage) RegionalLocation() net.Addr           { return sc.link.(net.Conn).LocalAddr() }
func (sc *CredentialLinkage) DistantLocation() net.Addr          { return sc.link.(net.Conn).RemoteAddr() }
func (sc *CredentialLinkage) AssignExpiration(t time.Time) error { return sc.link.(net.Conn).SetDeadline(t) }
func (sc *CredentialLinkage) AssignFetchLimit(t time.Time) error {
	return sc.link.(net.Conn).SetReadDeadline(t)
}

func (sc *CredentialLinkage) AssignPersistLimit(t time.Time) error {
	return sc.link.(net.Conn).SetWriteDeadline(t)
}

func produceTempTokens() (tempPublic, tempPrivate *[32]byte) {
	var err error
	//
	//
	//
	tempPublic, tempPrivate, err = box.GenerateKey(crand.Reader)
	if err != nil {
		panic("REDACTED")
	}
	return
}

func allocateTempPublicToken(link io.ReadWriter, positionTempPublic *[32]byte) (modTempPublic *[32]byte, err error) {
	//
	trs, _ := asyncronous.Concurrent(
		func(_ int) (val any, cancel bool, err error) {
			lc := *positionTempPublic
			_, err = protocolio.FreshSeparatedPersistor(link).PersistSignal(&gogotypes.BytesValue{Value: lc[:]})
			if err != nil {
				return nil, true, err //
			}
			return nil, false, nil
		},
		func(_ int) (val any, cancel bool, err error) {
			var octets gogotypes.BytesValue
			_, err = protocolio.FreshSeparatedFetcher(link, 1024*1024).FetchSignal(&octets)
			if err != nil {
				return nil, true, err //
			}

			var _remephpub [32]byte
			copy(_remephpub[:], bytes.Value)
			return _remephpub, false, nil
		},
	)

	//
	if trs.InitialFailure() != nil {
		err = trs.InitialFailure()
		return modTempPublic, err
	}

	//
	_remephpub := trs.InitialDatum().([32]byte)
	return &_remephpub, nil
}

func deduceCredentials(
	dhCredential *[32]byte,
	positionEqualsMinimal bool,
) (obtainCredential, transmitCredential *[aeadTokenExtent]byte) {
	digest := sha256.New
	hkdf := hkdf.New(digest, dhCredential[:], nil, credentialLinkTokenAlsoQueryProduce)
	//
	res := new([2*aeadTokenExtent + 32]byte)
	_, err := io.ReadFull(hkdf, res[:])
	if err != nil {
		panic(err)
	}

	obtainCredential = new([aeadTokenExtent]byte)
	transmitCredential = new([aeadTokenExtent]byte)

	//
	//
	//
	//
	if positionEqualsMinimal {
		copy(obtainCredential[:], res[0:aeadTokenExtent])
		copy(transmitCredential[:], res[aeadTokenExtent:aeadTokenExtent*2])
	} else {
		copy(transmitCredential[:], res[0:aeadTokenExtent])
		copy(obtainCredential[:], res[aeadTokenExtent:aeadTokenExtent*2])
	}

	return
}

//
//
func calculateDHCredential(modPublicToken, positionPrivateToken *[32]byte) (*[32]byte, error) {
	rightshiftToken, err := curve25519.X25519(positionPrivateToken[:], modPublicToken[:])
	if err != nil {
		return nil, err
	}
	var rightshiftTokenSeries [32]byte
	copy(rightshiftTokenSeries[:], rightshiftToken)
	return &rightshiftTokenSeries, nil
}

func sort32(foo, bar *[32]byte) (lo, hi *[32]byte) {
	if bytes.Compare(foo[:], bar[:]) < 0 {
		lo = foo
		hi = bar
	} else {
		lo = bar
		hi = foo
	}
	return
}

func attestQuery(query *[32]byte, positionPrivateToken security.PrivateToken) ([]byte, error) {
	signing, err := positionPrivateToken.Attest(query[:])
	if err != nil {
		return nil, err
	}
	return signing, nil
}

type authSignatureArtifact struct {
	Key security.PublicToken
	Sig []byte
}

func allocateAuthNotation(sc io.ReadWriter, publicToken security.PublicToken, signing []byte) (obtainSignal authSignatureArtifact, err error) {
	//
	trs, _ := asyncronous.Concurrent(
		func(_ int) (val any, cancel bool, err error) {
			pbpk, err := cryptocode.PublicTokenTowardSchema(publicToken)
			if err != nil {
				return nil, true, err
			}
			_, err = protocolio.FreshSeparatedPersistor(sc).PersistSignal(&tmpfabric.AuthSignatureArtifact{PublicToken: pbpk, Sig: signing})
			if err != nil {
				return nil, true, err //
			}
			return nil, false, nil
		},
		func(_ int) (val any, cancel bool, err error) {
			var pba tmpfabric.AuthSignatureArtifact
			_, err = protocolio.FreshSeparatedFetcher(sc, 1024*1024).FetchSignal(&pba)
			if err != nil {
				return nil, true, err //
			}

			pk, err := cryptocode.PublicTokenOriginatingSchema(pba.PublicToken)
			if err != nil {
				return nil, true, err //
			}

			_acceptartifact := authSignatureArtifact{
				Key: pk,
				Sig: pba.Sig,
			}
			return _acceptartifact, false, nil
		},
	)

	//
	if trs.InitialFailure() != nil {
		err = trs.InitialFailure()
		return obtainSignal, err
	}

	_acceptartifact := trs.InitialDatum().(authSignatureArtifact)
	return _acceptartifact, nil
}

//

//
//
//
//
func increaseNumber(number *[aeadNumberExtent]byte) {
	tally := binary.LittleEndian.Uint64(number[4:])
	if tally == math.MaxUint64 {
		//
		//
		panic("REDACTED")
	}
	tally++
	binary.LittleEndian.PutUint64(number[4:], tally)
}
