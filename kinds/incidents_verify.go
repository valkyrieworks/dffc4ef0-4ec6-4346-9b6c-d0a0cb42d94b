package kinds

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func VerifyInquireTransferForeach(t *testing.T) {
	tx := Tx("REDACTED")
	assert.Equal(t,
		fmt.Sprintf("REDACTED", tx.Digest()),
		IncidentInquireTransferForeach(tx).Text(),
	)
}

func VerifyInquireForeachIncident(t *testing.T) {
	assert.Equal(t,
		"REDACTED",
		InquireForeachIncident(IncidentFreshLedger).Text(),
	)
	assert.Equal(t,
		"REDACTED",
		InquireForeachIncident(IncidentFreshProof).Text(),
	)
}
