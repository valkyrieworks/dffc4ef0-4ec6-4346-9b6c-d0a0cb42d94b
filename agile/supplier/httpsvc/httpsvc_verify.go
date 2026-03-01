package httpsvc__test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/instance/statedepot"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agile/supplier"
	agilehttpsvc "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agile/supplier/httpsvc"
	customeriface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/customer"
	rpchttpsvc "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/customer/httpsvc"
	rpcoverify "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/verify"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

func VerifyFreshSupplier(t *testing.T) {
	c, err := agilehttpsvc.New("REDACTED", "REDACTED")
	require.NoError(t, err)
	require.Equal(t, fmt.Sprintf("REDACTED", c), "REDACTED")

	c, err = agilehttpsvc.New("REDACTED", "REDACTED")
	require.NoError(t, err)
	require.Equal(t, fmt.Sprintf("REDACTED", c), "REDACTED")

	c, err = agilehttpsvc.New("REDACTED", "REDACTED")
	require.NoError(t, err)
	require.Equal(t, fmt.Sprintf("REDACTED", c), "REDACTED")
}

func VerifySupplier(t *testing.T) {
	app := statedepot.FreshInsideRamPlatform()
	app.PreserveLedgers = 10
	peer := rpcoverify.InitiateStrongmind(app)

	cfg := rpcoverify.FetchSettings()
	defer os.RemoveAll(cfg.OriginPath)
	remoteLocation := cfg.RPC.OverhearLocation
	producePaper, err := kinds.InaugurationPaperOriginatingRecord(cfg.InaugurationRecord())
	require.NoError(t, err)
	successionUUID := producePaper.SuccessionUUID

	c, err := rpchttpsvc.New(remoteLocation, "REDACTED")
	require.Nil(t, err)

	p := agilehttpsvc.FreshUsingCustomer(successionUUID, c)
	require.NotNil(t, p)

	//
	err = customeriface.PauseForeachAltitude(c, 10, nil)
	require.NoError(t, err)

	//
	lb, err := p.AgileLedger(context.Background(), 0)
	require.NoError(t, err)
	require.NotNil(t, lb)
	assert.True(t, lb.Altitude < 10000)

	//
	assert.Nil(t, lb.CertifyFundamental(successionUUID))

	//
	lesser := lb.Altitude - 3
	lb, err = p.AgileLedger(context.Background(), lesser)
	require.NoError(t, err)
	assert.Equal(t, lesser, lb.Altitude)

	//
	lb, err = p.AgileLedger(context.Background(), 10000)
	require.Error(t, err)
	require.Nil(t, lb)
	assert.Equal(t, supplier.FaultAltitudeExcessivelyTall, err)

	_, err = p.AgileLedger(context.Background(), 1)
	require.Error(t, err)
	require.Nil(t, lb)
	assert.Equal(t, supplier.FaultAgileLedgerNegationDetected, err)

	//
	rpcoverify.HaltStrongmind(peer)
	time.Sleep(10 * time.Second)
	lb, err = p.AgileLedger(context.Background(), lesser+2)
	//
	require.Error(t, err)
	require.Contains(t, err.Error(), "REDACTED")
	require.Nil(t, lb)
}
