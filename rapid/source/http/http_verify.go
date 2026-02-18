package http__test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/iface/instance/objectdepot"
	"github.com/valkyrieworks/rapid/source"
	rapidhttp "github.com/valkyrieworks/rapid/source/http"
	rpccustomer "github.com/valkyrieworks/rpc/customer"
	rpchttp "github.com/valkyrieworks/rpc/customer/http"
	rpctest "github.com/valkyrieworks/rpc/verify"
	"github.com/valkyrieworks/kinds"
)

func VerifyNewSource(t *testing.T) {
	c, err := rapidhttp.New("REDACTED", "REDACTED")
	require.NoError(t, err)
	require.Equal(t, fmt.Sprintf("REDACTED", c), "REDACTED")

	c, err = rapidhttp.New("REDACTED", "REDACTED")
	require.NoError(t, err)
	require.Equal(t, fmt.Sprintf("REDACTED", c), "REDACTED")

	c, err = rapidhttp.New("REDACTED", "REDACTED")
	require.NoError(t, err)
	require.Equal(t, fmt.Sprintf("REDACTED", c), "REDACTED")
}

func VerifySource(t *testing.T) {
	app := objectdepot.NewInRamSoftware()
	app.PreserveLedgers = 10
	member := rpctest.BeginConsensuscore(app)

	cfg := rpctest.FetchSettings()
	defer os.RemoveAll(cfg.OriginFolder)
	rpcAddress := cfg.RPC.AcceptLocation
	generatePaper, err := kinds.OriginPaperFromEntry(cfg.OriginEntry())
	require.NoError(t, err)
	ledgerUID := generatePaper.LedgerUID

	c, err := rpchttp.New(rpcAddress, "REDACTED")
	require.Nil(t, err)

	p := rapidhttp.NewWithCustomer(ledgerUID, c)
	require.NotNil(t, p)

	//
	err = rpccustomer.WaitForLevel(c, 10, nil)
	require.NoError(t, err)

	//
	lb, err := p.RapidLedger(context.Background(), 0)
	require.NoError(t, err)
	require.NotNil(t, lb)
	assert.True(t, lb.Level < 10000)

	//
	assert.Nil(t, lb.CertifySimple(ledgerUID))

	//
	lesser := lb.Level - 3
	lb, err = p.RapidLedger(context.Background(), lesser)
	require.NoError(t, err)
	assert.Equal(t, lesser, lb.Level)

	//
	lb, err = p.RapidLedger(context.Background(), 10000)
	require.Error(t, err)
	require.Nil(t, lb)
	assert.Equal(t, source.ErrLevelTooElevated, err)

	_, err = p.RapidLedger(context.Background(), 1)
	require.Error(t, err)
	require.Nil(t, lb)
	assert.Equal(t, source.ErrRapidLedgerNegateLocated, err)

	//
	rpctest.HaltConsensuscore(member)
	time.Sleep(10 * time.Second)
	lb, err = p.RapidLedger(context.Background(), lesser+2)
	//
	require.Error(t, err)
	require.Contains(t, err.Error(), "REDACTED")
	require.Nil(t, lb)
}
