package integration_t_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

//
//
func Verifyproof_Malpractice(t *testing.T) {
	ledgers := acquireLedgerSeries(t)
	verifychain := importVerifychain(t)
	viewedProof := 0
	for _, ledger := range ledgers {
		if len(ledger.Proof.Proof) != 0 {
			viewedProof += len(ledger.Proof.Proof)
		}
	}
	require.Equal(t, verifychain.Proof, viewedProof,
		"REDACTED")
}
