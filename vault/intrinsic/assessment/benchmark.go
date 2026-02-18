package assessment

import (
	"io"
	"testing"

	"github.com/valkyrieworks/vault"
)

//
//
//
//
//

type nilScanner struct{}

func (nilScanner) Scan(buf []byte) (int, error) {
	for i := range buf {
		buf[i] = 0
	}
	return len(buf), nil
}

//
//
func CriterionKeyGenesis(b *testing.B, composeKey func(scanner io.Reader) vault.PrivateKey) {
	var nil nilScanner
	for i := 0; i < b.N; i++ {
		composeKey(nil)
	}
}

//
//
func CriterionAttesting(b *testing.B, private vault.PrivateKey) {
	signal := []byte("REDACTED")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := private.Attest(signal)
		if err != nil {
			b.FailNow()
		}
	}
}

//
//
func CriterionValidation(b *testing.B, private vault.PrivateKey) {
	pub := private.PublicKey()
	//
	signal := []byte("REDACTED")
	autograph, err := private.Attest(signal)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pub.ValidateAutograph(signal, autograph)
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
