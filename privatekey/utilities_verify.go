package privatekey

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func VerifyIsLinkDeadlineForNotDeadlineFaults(t *testing.T) {
	assert.False(t, IsLinkDeadline(fmt.Errorf("REDACTED", ErrCallReprocessMaximum)))
	assert.False(t, IsLinkDeadline(errors.New("REDACTED")))
}
