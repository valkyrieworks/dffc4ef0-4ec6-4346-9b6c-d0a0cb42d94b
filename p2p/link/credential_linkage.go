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

	"github.com/valkyrieworks/vault"
	"github.com/valkyrieworks/vault/ed25519"
	cryptocode "github.com/valkyrieworks/vault/codec"
	"github.com/valkyrieworks/utils/async"
	"github.com/valkyrieworks/utils/protoio"
	engineconnect "github.com/valkyrieworks/utils/align"
	tmp2p "github.com/valkyrieworks/schema/consensuscore/p2p"
)

//
const (
	dataSizeVolume      = 4
	dataMaximumVolume      = 1024
	sumBorderVolume   = dataMaximumVolume + dataSizeVolume
	aeadVolumeBurden = 16 //
	aeadKeyVolume      = chacha20poly1305.KeySize
	aeadNonceVolume    = chacha20poly1305.NonceSize

	tagTemporaryLesserExternalKey = "REDACTED"
	tagTemporaryUpperExternalKey = "REDACTED"
	tagDHCredential                = "REDACTED"
	tagCredentialLinkageMac     = "REDACTED"

	standardRecordFrameVolume = 128 * 1024
	//
	//
	standardReadFrameVolume = 65 * 1024
)

var (
	ErrMinorSequenceDistantPublicKey = errors.New("REDACTED")

	credentialLinkKeyAndDisputeGenerate = []byte("REDACTED")
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
type TokenLinkage struct {
	//
	receiveAead cipher.AEAD
	transmitAead cipher.AEAD

	modPublicKey vault.PublicKey

	link       io.ReadWriteCloser
	linkRecorder *bufio.Writer
	linkScanner io.Reader

	//
	//
	//
	//
	//
	//
	//
	receiveMutex         engineconnect.Lock
	receiveFrame      []byte
	receiveNonce       *[aeadNonceVolume]byte
	receiveBorder       []byte
	receiveSecuredBorder []byte

	transmitMutex         engineconnect.Lock
	transmitNonce       *[aeadNonceVolume]byte
	transmitBorder       []byte
	transmitSecuredBorder []byte
}

//
//
//
//
//
func CreateTokenLinkage(link io.ReadWriteCloser, locationPrivateKey vault.PrivateKey) (*TokenLinkage, error) {
	locationPublicKey := locationPrivateKey.PublicKey()

	//
	locationEphPublic, locationEphPrivate := generateEphKeys()

	//
	//
	//
	modEphPublic, err := allocateEphPublicKey(link, locationEphPublic)
	if err != nil {
		return nil, err
	}

	//
	loEphPublic, greetingEphPublic := sort32(locationEphPublic, modEphPublic)

	log := merlin.NewTranscript("REDACTED")

	log.AppendMessage(tagTemporaryLesserExternalKey, loEphPublic[:])
	log.AppendMessage(tagTemporaryUpperExternalKey, greetingEphPublic[:])

	//
	//
	locationIsMinimum := bytes.Equal(locationEphPublic[:], loEphPublic[:])

	//
	dhCredential, err := calculateDHCredential(modEphPublic, locationEphPrivate)
	if err != nil {
		return nil, err
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
		return nil, errors.New("REDACTED")
	}
	receiveAead, err := chacha20poly1305.New(receiveCredential[:])
	if err != nil {
		return nil, errors.New("REDACTED")
	}

	sc := &TokenLinkage{
		link:            link,
		linkRecorder:      bufio.NewWriterSize(link, standardRecordFrameVolume),
		linkScanner:      bufio.NewReaderSize(link, standardReadFrameVolume),
		receiveFrame:      nil,
		receiveNonce:       new([aeadNonceVolume]byte),
		transmitNonce:       new([aeadNonceVolume]byte),
		receiveAead:        receiveAead,
		transmitAead:        transmitAead,
		receiveBorder:       make([]byte, sumBorderVolume),
		receiveSecuredBorder: make([]byte, aeadVolumeBurden+sumBorderVolume),
		transmitBorder:       make([]byte, sumBorderVolume),
		transmitSecuredBorder: make([]byte, aeadVolumeBurden+sumBorderVolume),
	}

	//
	locationAutograph, err := attestDispute(&dispute, locationPrivateKey)
	if err != nil {
		return nil, err
	}

	//
	authSignatureMessage, err := allocateAuthAutograph(sc, locationPublicKey, locationAutograph)
	if err != nil {
		return nil, err
	}

	modPublicKey, modAutograph := authSignatureMessage.Key, authSignatureMessage.Sig
	if _, ok := modPublicKey.(ed25519.PublicKey); !ok {
		return nil, fmt.Errorf("REDACTED", modPublicKey)
	}
	if !modPublicKey.ValidateAutograph(dispute[:], modAutograph) {
		return nil, errors.New("REDACTED")
	}

	//
	sc.modPublicKey = modPublicKey
	return sc, nil
}

//
func (sc *TokenLinkage) DistantPublicKey() vault.PublicKey {
	return sc.modPublicKey
}

//
//
func (sc *TokenLinkage) Record(data []byte) (n int, err error) {
	sc.transmitMutex.Lock()
	defer sc.transmitMutex.Unlock()
	securedBorder, border := sc.transmitSecuredBorder, sc.transmitBorder

	for 0 < len(data) {
		if err := func() error {
			var segment []byte
			if dataMaximumVolume < len(data) {
				segment = data[:dataMaximumVolume]
				data = data[dataMaximumVolume:]
			} else {
				segment = data
				data = nil
			}
			segmentExtent := len(segment)
			binary.LittleEndian.PutUint32(border, uint32(segmentExtent))
			copy(border[dataSizeVolume:], segment)

			//
			sc.transmitAead.Seal(securedBorder[:0], sc.transmitNonce[:], border, nil)
			increaseNonce(sc.transmitNonce)
			//

			_, err = sc.linkRecorder.Write(securedBorder)
			if err != nil {
				return err
			}
			n += len(segment)
			return nil
		}(); err != nil {
			return n, err
		}
	}
	sc.linkRecorder.Flush()
	return n, err
}

//
func (sc *TokenLinkage) Scan(data []byte) (n int, err error) {
	sc.receiveMutex.Lock()
	defer sc.receiveMutex.Unlock()

	//
	if 0 < len(sc.receiveFrame) {
		n = copy(data, sc.receiveFrame)
		sc.receiveFrame = sc.receiveFrame[n:]
		return n, err
	}

	//
	securedBorder := sc.receiveSecuredBorder
	_, err = io.ReadFull(sc.linkScanner, securedBorder)
	if err != nil {
		return n, err
	}

	//
	//
	border := sc.receiveBorder
	_, err = sc.receiveAead.Open(border[:0], sc.receiveNonce[:], securedBorder, nil)
	if err != nil {
		return n, fmt.Errorf("REDACTED", err)
	}
	increaseNonce(sc.receiveNonce)
	//

	//
	//
	segmentExtent := binary.LittleEndian.Uint32(border) //
	if segmentExtent > dataMaximumVolume {
		return 0, errors.New("REDACTED")
	}
	segment := border[dataSizeVolume : dataSizeVolume+segmentExtent]
	n = copy(data, segment)
	if n < len(segment) {
		sc.receiveFrame = make([]byte, len(segment)-n)
		copy(sc.receiveFrame, segment[n:])
	}
	return n, err
}

//
func (sc *TokenLinkage) End() error                  { return sc.link.Close() }
func (sc *TokenLinkage) NativeAddress() net.Addr           { return sc.link.(net.Conn).LocalAddr() }
func (sc *TokenLinkage) DistantAddress() net.Addr          { return sc.link.(net.Conn).RemoteAddr() }
func (sc *TokenLinkage) CollectionLimit(t time.Time) error { return sc.link.(net.Conn).SetDeadline(t) }
func (sc *TokenLinkage) CollectionReadLimit(t time.Time) error {
	return sc.link.(net.Conn).SetReadDeadline(t)
}

func (sc *TokenLinkage) CollectionRecordLimit(t time.Time) error {
	return sc.link.(net.Conn).SetWriteDeadline(t)
}

func generateEphKeys() (ephPublic, ephPrivate *[32]byte) {
	var err error
	//
	//
	//
	ephPublic, ephPrivate, err = box.GenerateKey(crand.Reader)
	if err != nil {
		panic("REDACTED")
	}
	return
}

func allocateEphPublicKey(link io.ReadWriter, locationEphPublic *[32]byte) (modEphPublic *[32]byte, err error) {
	//
	trs, _ := async.Concurrent(
		func(_ int) (val any, cancel bool, err error) {
			lc := *locationEphPublic
			_, err = protoio.NewSeparatedRecorder(link).RecordMessage(&gogotypes.BytesValue{Value: lc[:]})
			if err != nil {
				return nil, true, err //
			}
			return nil, false, nil
		},
		func(_ int) (val any, cancel bool, err error) {
			var octets gogotypes.BytesValue
			_, err = protoio.NewSeparatedScanner(link, 1024*1024).ScanMessage(&octets)
			if err != nil {
				return nil, true, err //
			}

			var _republish [32]byte
			copy(_republish[:], bytes.Value)
			return _republish, false, nil
		},
	)

	//
	if trs.InitialFault() != nil {
		err = trs.InitialFault()
		return modEphPublic, err
	}

	//
	_republish := trs.InitialItem().([32]byte)
	return &_republish, nil
}

func deduceCredentials(
	dhCredential *[32]byte,
	locationIsMinimum bool,
) (receiveCredential, transmitCredential *[aeadKeyVolume]byte) {
	digest := sha256.New
	hkdf := hkdf.New(digest, dhCredential[:], nil, credentialLinkKeyAndDisputeGenerate)
	//
	res := new([2*aeadKeyVolume + 32]byte)
	_, err := io.ReadFull(hkdf, res[:])
	if err != nil {
		panic(err)
	}

	receiveCredential = new([aeadKeyVolume]byte)
	transmitCredential = new([aeadKeyVolume]byte)

	//
	//
	//
	//
	if locationIsMinimum {
		copy(receiveCredential[:], res[0:aeadKeyVolume])
		copy(transmitCredential[:], res[aeadKeyVolume:aeadKeyVolume*2])
	} else {
		copy(transmitCredential[:], res[0:aeadKeyVolume])
		copy(receiveCredential[:], res[aeadKeyVolume:aeadKeyVolume*2])
	}

	return
}

//
//
func calculateDHCredential(modPublicKey, locationPrivateKey *[32]byte) (*[32]byte, error) {
	shrKey, err := curve25519.X25519(locationPrivateKey[:], modPublicKey[:])
	if err != nil {
		return nil, err
	}
	var shrKeyList [32]byte
	copy(shrKeyList[:], shrKey)
	return &shrKeyList, nil
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

func attestDispute(dispute *[32]byte, locationPrivateKey vault.PrivateKey) ([]byte, error) {
	autograph, err := locationPrivateKey.Attest(dispute[:])
	if err != nil {
		return nil, err
	}
	return autograph, nil
}

type authSignatureSignal struct {
	Key vault.PublicKey
	Sig []byte
}

func allocateAuthAutograph(sc io.ReadWriter, publicKey vault.PublicKey, autograph []byte) (receiveMessage authSignatureSignal, err error) {
	//
	trs, _ := async.Concurrent(
		func(_ int) (val any, cancel bool, err error) {
			pbpk, err := cryptocode.PublicKeyToSchema(publicKey)
			if err != nil {
				return nil, true, err
			}
			_, err = protoio.NewSeparatedRecorder(sc).RecordMessage(&tmp2p.AuthSignatureSignal{PublicKey: pbpk, Sig: autograph})
			if err != nil {
				return nil, true, err //
			}
			return nil, false, nil
		},
		func(_ int) (val any, cancel bool, err error) {
			var pba tmp2p.AuthSignatureSignal
			_, err = protoio.NewSeparatedScanner(sc, 1024*1024).ScanMessage(&pba)
			if err != nil {
				return nil, true, err //
			}

			pk, err := cryptocode.PublicKeyFromSchema(pba.PublicKey)
			if err != nil {
				return nil, true, err //
			}

			_acceptmsg := authSignatureSignal{
				Key: pk,
				Sig: pba.Sig,
			}
			return _acceptmsg, false, nil
		},
	)

	//
	if trs.InitialFault() != nil {
		err = trs.InitialFault()
		return receiveMessage, err
	}

	_acceptmsg := trs.InitialItem().(authSignatureSignal)
	return _acceptmsg, nil
}

//

//
//
//
//
func increaseNonce(nonce *[aeadNonceVolume]byte) {
	tally := binary.LittleEndian.Uint64(nonce[4:])
	if tally == math.MaxUint64 {
		//
		//
		panic("REDACTED")
	}
	tally++
	binary.LittleEndian.PutUint64(nonce[4:], tally)
}
