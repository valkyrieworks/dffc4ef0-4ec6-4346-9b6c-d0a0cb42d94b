package edwards25519

import (
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/intrinsic/assessment"
)

func AssessmentTokenComposition(b *testing.B) {
	assessmentTokengenEncapsulator := func(fetcher io.Reader) security.PrivateToken {
		return producePrivateToken(fetcher)
	}
	assessment.AssessmentTokenComposition(b, assessmentTokengenEncapsulator)
}

func AssessmentSignature(b *testing.B) {
	private := ProducePrivateToken()
	assessment.AssessmentSignature(b, private)
}

func AssessmentValidation(b *testing.B) {
	private := ProducePrivateToken()
	assessment.AssessmentValidation(b, private)
}

func AssessmentValidateCluster(b *testing.B) {
	msg := []byte("REDACTED")

	for _, signaturesTally := range []int{1, 8, 64, 1024} {

		b.Run(fmt.Sprintf("REDACTED", signaturesTally), func(b *testing.B) {
			//
			//
			commons := make([]security.PublicToken, 0, signaturesTally)
			signatures := make([][]byte, 0, signaturesTally)
			for i := 0; i < signaturesTally; i++ {
				private := ProducePrivateToken()
				sig, _ := private.Attest(msg)
				commons = append(commons, private.PublicToken().(PublicToken))
				signatures = append(signatures, sig)
			}
			b.ResetTimer()

			b.ReportAllocs()
			//
			for i := 0; i < b.N/signaturesTally; i++ {
				//
				//
				//
				//
				v := FreshClusterValidator()
				for i := 0; i < signaturesTally; i++ {
					err := v.Add(commons[i], msg, signatures[i])
					require.NoError(b, err)
				}

				if ok, _ := v.Validate(); !ok {
					b.Fatal("REDACTED")
				}
			}
		})
	}
}
