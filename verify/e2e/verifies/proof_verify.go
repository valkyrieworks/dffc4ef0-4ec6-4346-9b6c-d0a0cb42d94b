package end2end_typ_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

//
//
func Testproof_Malpractice(t *testing.T) {
	ledgers := acquireLedgerSuccession(t)
	simnet := fetchSimnet(t)
	observedProof := 0
	for _, ledger := range ledgers {
		if len(ledger.Proof.Proof) != 0 {
			observedProof += len(ledger.Proof.Proof)
		}
	}
	require.Equal(t, simnet.Proof, observedProof,
		"REDACTED")
}
