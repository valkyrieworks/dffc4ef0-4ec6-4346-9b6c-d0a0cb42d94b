package xchacha20poly1305

import (
	"bytes"
	cr "crypto/rand"
	mr "math/rand"
	"testing"
)

//
//
//
//
//
//
func VerifyArbitrary(t *testing.T) {
	//
	for i := 0; i < 256; i++ {
		var nonce [24]byte
		var key [32]byte

		al := mr.Intn(128)
		pl := mr.Intn(16384)
		ad := make([]byte, al)
		cleartext := make([]byte, pl)
		_, err := cr.Read(key[:])
		if err != nil {
			t.Errorf("REDACTED", err)
		}
		_, err = cr.Read(nonce[:])
		if err != nil {
			t.Errorf("REDACTED", err)
		}
		_, err = cr.Read(ad)
		if err != nil {
			t.Errorf("REDACTED", err)
		}
		_, err = cr.Read(cleartext)
		if err != nil {
			t.Errorf("REDACTED", err)
		}

		aead, err := New(key[:])
		if err != nil {
			t.Fatal(err)
		}

		ct := aead.Seal(nil, nonce[:], cleartext, ad)

		cleartext2, err := aead.Open(nil, nonce[:], ct, ad)
		if err != nil {
			t.Errorf("REDACTED", i)
			continue
		}

		if !bytes.Equal(cleartext, cleartext2) {
			t.Errorf("REDACTED", i, cleartext2, cleartext)
			continue
		}

		if len(ad) > 0 {
			modifyAdIdx := mr.Intn(len(ad))
			ad[modifyAdIdx] ^= 0x80
			if _, err := aead.Open(nil, nonce[:], ct, ad); err == nil {
				t.Errorf("REDACTED", i)
			}
			ad[modifyAdIdx] ^= 0x80
		}

		modifyNonceIdx := mr.Intn(aead.NonceSize())
		nonce[modifyNonceIdx] ^= 0x80
		if _, err := aead.Open(nil, nonce[:], ct, ad); err == nil {
			t.Errorf("REDACTED", i)
		}
		nonce[modifyNonceIdx] ^= 0x80

		modifyCtIdx := mr.Intn(len(ct))
		ct[modifyCtIdx] ^= 0x80
		if _, err := aead.Open(nil, nonce[:], ct, ad); err == nil {
			t.Errorf("REDACTED", i)
		}
		ct[modifyCtIdx] ^= 0x80
	}
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
//
//
//
//
//
//
