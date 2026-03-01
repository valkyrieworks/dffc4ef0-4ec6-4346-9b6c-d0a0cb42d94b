package xsalsa20poly1305

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
func VerifyUnpredictable(t *testing.T) {
	//
	for i := 0; i < 256; i++ {
		var number [24]byte
		var key [32]byte

		al := mr.Intn(128)
		pl := mr.Intn(16384)
		ad := make([]byte, al)
		cleartext := make([]byte, pl)
		_, err := cr.Read(key[:])
		if err != nil {
			t.Errorf("REDACTED", err)
		}
		_, err = cr.Read(number[:])
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

		cipher, err := New(key[:])
		if err != nil {
			t.Fatal(err)
		}

		ct := cipher.Seal(nil, number[:], cleartext, ad)

		cleartext2, err := cipher.Open(nil, number[:], ct, ad)
		if err != nil {
			t.Errorf("REDACTED", i)
			continue
		}

		if !bytes.Equal(cleartext, cleartext2) {
			t.Errorf("REDACTED", i, cleartext2, cleartext)
			continue
		}

		if len(ad) > 0 {
			modifyAssociateddataOffset := mr.Intn(len(ad))
			ad[modifyAssociateddataOffset] ^= 0x80
			if _, err := cipher.Open(nil, number[:], ct, ad); err == nil {
				t.Errorf("REDACTED", i)
			}
			ad[modifyAssociateddataOffset] ^= 0x80
		}

		modifyNumberOffset := mr.Intn(cipher.NonceSize())
		number[modifyNumberOffset] ^= 0x80
		if _, err := cipher.Open(nil, number[:], ct, ad); err == nil {
			t.Errorf("REDACTED", i)
		}
		number[modifyNumberOffset] ^= 0x80

		modifyCiphertextOffset := mr.Intn(len(ct))
		ct[modifyCiphertextOffset] ^= 0x80
		if _, err := cipher.Open(nil, number[:], ct, ad); err == nil {
			t.Errorf("REDACTED", i)
		}
		ct[modifyCiphertextOffset] ^= 0x80
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
