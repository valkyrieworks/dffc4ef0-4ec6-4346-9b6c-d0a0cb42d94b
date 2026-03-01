package privatevalue

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/edwards25519"
)

func obtainCallerVerifyScenarios(t *testing.T) []callerVerifyScenario {
	tcpsocketLocation := ObtainReleaseLocalmachineLocationChannel()
	posixRecordRoute, err := verifyPosixLocation()
	require.NoError(t, err)
	posixLocation := fmt.Sprintf("REDACTED", posixRecordRoute)

	return []callerVerifyScenario{
		{
			location:   tcpsocketLocation,
			caller: CallStreamProc(tcpsocketLocation, verifyDeadlineRetrievePersist, edwards25519.ProducePrivateToken()),
		},
		{
			location:   posixLocation,
			caller: CallPosixProc(posixRecordRoute),
		},
	}
}

func VerifyEqualsLinkDeadlineForeachEssentialExpirations(t *testing.T) {
	//
	tcpsocketLocation := ObtainReleaseLocalmachineLocationChannel()
	caller := CallStreamProc(tcpsocketLocation, time.Millisecond, edwards25519.ProducePrivateToken())
	_, err := caller()
	assert.Error(t, err)
	assert.True(t, EqualsLinkDeadline(err))
}

func VerifyEqualsLinkDeadlineForeachEncapsulatedLinkExpirations(t *testing.T) {
	tcpsocketLocation := ObtainReleaseLocalmachineLocationChannel()
	caller := CallStreamProc(tcpsocketLocation, time.Millisecond, edwards25519.ProducePrivateToken())
	_, err := caller()
	assert.Error(t, err)
	err = fmt.Errorf("REDACTED", err, FaultLinkageDeadline)
	assert.True(t, EqualsLinkDeadline(err))
}
