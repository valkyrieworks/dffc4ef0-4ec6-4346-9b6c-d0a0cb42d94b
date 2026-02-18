package privatekey

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/vault/ed25519"
)

func fetchCallerVerifyScenarios(t *testing.T) []callerVerifyScenario {
	tcpAddress := FetchReleaseLocalhostAddressPort()
	unixEntryRoute, err := verifyUnixAddress()
	require.NoError(t, err)
	unixAddress := fmt.Sprintf("REDACTED", unixEntryRoute)

	return []callerVerifyScenario{
		{
			address:   tcpAddress,
			caller: CallTCPFn(tcpAddress, verifyDeadlineScanRecord, ed25519.GeneratePrivateKey()),
		},
		{
			address:   unixAddress,
			caller: CallUnixFn(unixEntryRoute),
		},
	}
}

func VerifyIsLinkDeadlineForBasicDelays(t *testing.T) {
	//
	tcpAddress := FetchReleaseLocalhostAddressPort()
	caller := CallTCPFn(tcpAddress, time.Millisecond, ed25519.GeneratePrivateKey())
	_, err := caller()
	assert.Error(t, err)
	assert.True(t, IsLinkDeadline(err))
}

func VerifyIsLinkDeadlineForEncapsulatedLinkDelays(t *testing.T) {
	tcpAddress := FetchReleaseLocalhostAddressPort()
	caller := CallTCPFn(tcpAddress, time.Millisecond, ed25519.GeneratePrivateKey())
	_, err := caller()
	assert.Error(t, err)
	err = fmt.Errorf("REDACTED", err, ErrLinkageDeadline)
	assert.True(t, IsLinkDeadline(err))
}
