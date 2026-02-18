package ed25519

import (
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/vault"
	"github.com/valkyrieworks/vault/intrinsic/assessment"
)

func CriterionKeyGenesis(b *testing.B) {
	criterionKeygenAdapter := func(scanner io.Reader) vault.PrivateKey {
		return generatePrivateKey(scanner)
	}
	assessment.CriterionKeyGenesis(b, criterionKeygenAdapter)
}

func CriterionAttesting(b *testing.B) {
	private := GeneratePrivateKey()
	assessment.CriterionAttesting(b, private)
}

func CriterionValidation(b *testing.B) {
	private := GeneratePrivateKey()
	assessment.CriterionValidation(b, private)
}

func CriterionValidateGroup(b *testing.B) {
	msg := []byte("REDACTED")

	for _, autographsTally := range []int{1, 8, 64, 1024} {

		b.Run(fmt.Sprintf("REDACTED", autographsTally), func(b *testing.B) {
			//
			//
			publics := make([]vault.PublicKey, 0, autographsTally)
			autographs := make([][]byte, 0, autographsTally)
			for i := 0; i < autographsTally; i++ {
				private := GeneratePrivateKey()
				sig, _ := private.Attest(msg)
				publics = append(publics, private.PublicKey().(PublicKey))
				autographs = append(autographs, sig)
			}
			b.ResetTimer()

			b.ReportAllocs()
			//
			for i := 0; i < b.N/autographsTally; i++ {
				//
				//
				//
				//
				v := NewGroupValidator()
				for i := 0; i < autographsTally; i++ {
					err := v.Add(publics[i], msg, autographs[i])
					require.NoError(b, err)
				}

				if ok, _ := v.Validate(); !ok {
					b.Fatal("REDACTED")
				}
			}
		})
	}
}
