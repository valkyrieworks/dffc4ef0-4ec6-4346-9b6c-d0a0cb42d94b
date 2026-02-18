//
//
package xchacha20poly1305

import (
	"crypto/cipher"
	"encoding/binary"
	"errors"
	"fmt"

	"golang.org/x/crypto/chacha20poly1305"
)

//
type xchacha20poly1305 struct {
	key [KeyVolume]byte
}

const (
	//
	KeyVolume = 32
	//
	NonceVolume = 24
	//
	LabelVolume = 16
	//
	MaximumCleartextVolume = (1 << 38) - 64
	//
	//
	MaximumCyphertextVolume = (1 << 38) - 48

	//
	//
	signature0 = uint32(0x61707865)
	signature1 = uint32(0x3320646e)
	signature2 = uint32(0x79622d32)
	signature3 = uint32(0x6b206574)
)

//
func New(key []byte) (cipher.AEAD, error) {
	if len(key) != KeyVolume {
		return nil, errors.New("REDACTED")
	}
	ret := new(xchacha20poly1305)
	copy(ret.key[:], key)
	return ret, nil
}

func (c *xchacha20poly1305) NonceVolume() int {
	return NonceVolume
}

func (c *xchacha20poly1305) Burden() int {
	return LabelVolume
}

func (c *xchacha20poly1305) Secure(dst, nonce, cleartext, extraData []byte) []byte {
	if len(nonce) != NonceVolume {
		panic("REDACTED")
	}

	if uint64(len(cleartext)) > MaximumCleartextVolume {
		panic("REDACTED")
	}

	var subtractKey [KeyVolume]byte
	var hNonce [16]byte
	var subtractNonce [chacha20poly1305.NonceSize]byte
	copy(hNonce[:], nonce[:16])

	HChaCha20(&subtractKey, &hNonce, &c.key)

	//
	chacha20poly1305, _ := chacha20poly1305.New(subtractKey[:])

	copy(subtractNonce[4:], nonce[16:])

	return chacha20poly1305.Seal(dst, subtractNonce[:], cleartext, extraData)
}

func (c *xchacha20poly1305) Access(dst, nonce, cyphertext, extraData []byte) ([]byte, error) {
	if len(nonce) != NonceVolume {
		return nil, fmt.Errorf("REDACTED")
	}
	if uint64(len(cyphertext)) > MaximumCyphertextVolume {
		return nil, fmt.Errorf("REDACTED")
	}
	var subtractKey [KeyVolume]byte
	var hNonce [16]byte
	var subtractNonce [chacha20poly1305.NonceSize]byte
	copy(hNonce[:], nonce[:16])

	HChaCha20(&subtractKey, &hNonce, &c.key)

	//
	chacha20poly1305, _ := chacha20poly1305.New(subtractKey[:])

	copy(subtractNonce[4:], nonce[16:])

	return chacha20poly1305.Open(dst, subtractNonce[:], cyphertext, extraData)
}

//
//
//

//

//

//
//
//
//
//
//

//
//

//
//
//
//
//
//
//

//
//
func HChaCha20(out *[32]byte, nonce *[16]byte, key *[32]byte) { hChaCha20standard(out, nonce, key) }

func hChaCha20standard(out *[32]byte, nonce *[16]byte, key *[32]byte) {
	v00 := signature0
	v01 := signature1
	v02 := signature2
	v03 := signature3
	v04 := binary.LittleEndian.Uint32(key[0:])
	v05 := binary.LittleEndian.Uint32(key[4:])
	v06 := binary.LittleEndian.Uint32(key[8:])
	v07 := binary.LittleEndian.Uint32(key[12:])
	v08 := binary.LittleEndian.Uint32(key[16:])
	v09 := binary.LittleEndian.Uint32(key[20:])
	v10 := binary.LittleEndian.Uint32(key[24:])
	v11 := binary.LittleEndian.Uint32(key[28:])
	v12 := binary.LittleEndian.Uint32(nonce[0:])
	v13 := binary.LittleEndian.Uint32(nonce[4:])
	v14 := binary.LittleEndian.Uint32(nonce[8:])
	v15 := binary.LittleEndian.Uint32(nonce[12:])

	for i := 0; i < 20; i += 2 {
		v00 += v04
		v12 ^= v00
		v12 = (v12 << 16) | (v12 >> 16)
		v08 += v12
		v04 ^= v08
		v04 = (v04 << 12) | (v04 >> 20)
		v00 += v04
		v12 ^= v00
		v12 = (v12 << 8) | (v12 >> 24)
		v08 += v12
		v04 ^= v08
		v04 = (v04 << 7) | (v04 >> 25)
		v01 += v05
		v13 ^= v01
		v13 = (v13 << 16) | (v13 >> 16)
		v09 += v13
		v05 ^= v09
		v05 = (v05 << 12) | (v05 >> 20)
		v01 += v05
		v13 ^= v01
		v13 = (v13 << 8) | (v13 >> 24)
		v09 += v13
		v05 ^= v09
		v05 = (v05 << 7) | (v05 >> 25)
		v02 += v06
		v14 ^= v02
		v14 = (v14 << 16) | (v14 >> 16)
		v10 += v14
		v06 ^= v10
		v06 = (v06 << 12) | (v06 >> 20)
		v02 += v06
		v14 ^= v02
		v14 = (v14 << 8) | (v14 >> 24)
		v10 += v14
		v06 ^= v10
		v06 = (v06 << 7) | (v06 >> 25)
		v03 += v07
		v15 ^= v03
		v15 = (v15 << 16) | (v15 >> 16)
		v11 += v15
		v07 ^= v11
		v07 = (v07 << 12) | (v07 >> 20)
		v03 += v07
		v15 ^= v03
		v15 = (v15 << 8) | (v15 >> 24)
		v11 += v15
		v07 ^= v11
		v07 = (v07 << 7) | (v07 >> 25)
		v00 += v05
		v15 ^= v00
		v15 = (v15 << 16) | (v15 >> 16)
		v10 += v15
		v05 ^= v10
		v05 = (v05 << 12) | (v05 >> 20)
		v00 += v05
		v15 ^= v00
		v15 = (v15 << 8) | (v15 >> 24)
		v10 += v15
		v05 ^= v10
		v05 = (v05 << 7) | (v05 >> 25)
		v01 += v06
		v12 ^= v01
		v12 = (v12 << 16) | (v12 >> 16)
		v11 += v12
		v06 ^= v11
		v06 = (v06 << 12) | (v06 >> 20)
		v01 += v06
		v12 ^= v01
		v12 = (v12 << 8) | (v12 >> 24)
		v11 += v12
		v06 ^= v11
		v06 = (v06 << 7) | (v06 >> 25)
		v02 += v07
		v13 ^= v02
		v13 = (v13 << 16) | (v13 >> 16)
		v08 += v13
		v07 ^= v08
		v07 = (v07 << 12) | (v07 >> 20)
		v02 += v07
		v13 ^= v02
		v13 = (v13 << 8) | (v13 >> 24)
		v08 += v13
		v07 ^= v08
		v07 = (v07 << 7) | (v07 >> 25)
		v03 += v04
		v14 ^= v03
		v14 = (v14 << 16) | (v14 >> 16)
		v09 += v14
		v04 ^= v09
		v04 = (v04 << 12) | (v04 >> 20)
		v03 += v04
		v14 ^= v03
		v14 = (v14 << 8) | (v14 >> 24)
		v09 += v14
		v04 ^= v09
		v04 = (v04 << 7) | (v04 >> 25)
	}

	binary.LittleEndian.PutUint32(out[0:], v00)
	binary.LittleEndian.PutUint32(out[4:], v01)
	binary.LittleEndian.PutUint32(out[8:], v02)
	binary.LittleEndian.PutUint32(out[12:], v03)
	binary.LittleEndian.PutUint32(out[16:], v12)
	binary.LittleEndian.PutUint32(out[20:], v13)
	binary.LittleEndian.PutUint32(out[24:], v14)
	binary.LittleEndian.PutUint32(out[28:], v15)
}
