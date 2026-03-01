package privatevalue

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func VerifyEqualsLinkDeadlineForeachUnDeadlineFaults(t *testing.T) {
	assert.False(t, EqualsLinkDeadline(fmt.Errorf("REDACTED", FaultCallReissueMaximum)))
	assert.False(t, EqualsLinkDeadline(errors.New("REDACTED")))
}
