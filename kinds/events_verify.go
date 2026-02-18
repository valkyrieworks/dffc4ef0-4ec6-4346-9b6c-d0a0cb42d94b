package kinds

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func VerifyInquireTransferFor(t *testing.T) {
	tx := Tx("REDACTED")
	assert.Equal(t,
		fmt.Sprintf("REDACTED", tx.Digest()),
		EventInquireTransferFor(tx).String(),
	)
}

func VerifyInquireForEvent(t *testing.T) {
	assert.Equal(t,
		"REDACTED",
		InquireForEvent(EventNewLedger).String(),
	)
	assert.Equal(t,
		"REDACTED",
		InquireForEvent(EventNewProof).String(),
	)
}
