package assessment

import (
	"io"
	"testing"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
)

//
//
//
//
//

type nullFetcher struct{}

func (nullFetcher) Obtain(buf []byte) (int, error) {
	for i := range buf {
		buf[i] = 0
	}
	return len(buf), nil
}

//
//
func AssessmentTokenComposition(b *testing.B, composeToken func(fetcher io.Reader) security.PrivateToken) {
	var null nullFetcher
	for i := 0; i < b.N; i++ {
		composeToken(null)
	}
}

//
//
func AssessmentSignature(b *testing.B, private security.PrivateToken) {
	artifact := []byte("REDACTED")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := private.Attest(artifact)
		if err != nil {
			b.FailNow()
		}
	}
}

//
//
func AssessmentValidation(b *testing.B, private security.PrivateToken) {
	pub := private.PublicToken()
	//
	artifact := []byte("REDACTED")
	signing, err := private.Attest(artifact)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pub.ValidateSigning(artifact, signing)
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
