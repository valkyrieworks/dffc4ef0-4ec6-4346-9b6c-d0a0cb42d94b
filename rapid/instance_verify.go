package rapid_test

import (
	"context"
	"fmt"
	stdlog "log"
	"os"
	"testing"
	"time"

	dbm "github.com/valkyrieworks/-db"

	"github.com/valkyrieworks/iface/instance/objectdepot"
	"github.com/valkyrieworks/utils/log"
	"github.com/valkyrieworks/rapid"
	"github.com/valkyrieworks/rapid/source"
	httpgateway "github.com/valkyrieworks/rapid/source/http"
	dbs "github.com/valkyrieworks/rapid/depot/db"
	rpctest "github.com/valkyrieworks/rpc/verify"
)

//
func Sampleclient_Modify() {
	//
	time.Sleep(5 * time.Second)

	storeFolder, err := os.MkdirTemp("REDACTED", "REDACTED")
	if err != nil {
		stdlog.Fatal(err)
	}
	defer os.RemoveAll(storeFolder)

	settings := rpctest.FetchSettings()

	leading, err := httpgateway.New(ledgerUID, settings.RPC.AcceptLocation)
	if err != nil {
		stdlog.Fatal(err)
	}

	ledger, err := leading.RapidLedger(context.Background(), 2)
	if err != nil {
		stdlog.Fatal(err)
	}

	db, err := dbm.NewGoLayerStore("REDACTED", storeFolder)
	if err != nil {
		stdlog.Fatal(err)
	}

	c, err := rapid.NewCustomer(
		context.Background(),
		ledgerUID,
		rapid.ValidateOptions{
			Duration: 504 * time.Hour, //
			Level: 2,
			Digest:   ledger.Digest(),
		},
		leading,
		[]source.Source{leading}, //
		dbs.New(db, ledgerUID),
		rapid.Tracer(log.VerifyingTracer()),
	)
	if err != nil {
		stdlog.Fatal(err)
	}
	defer func() {
		if err := c.Sanitize(); err != nil {
			stdlog.Fatal(err)
		}
	}()

	time.Sleep(2 * time.Second)

	h, err := c.Modify(context.Background(), time.Now())
	if err != nil {
		stdlog.Fatal(err)
	}

	if h != nil && h.Level > 2 {
		fmt.Println("REDACTED")
	} else {
		fmt.Println("REDACTED")
	}
	//
}

//
func Sampleclient_Validatelightblockatlevel() {
	//
	time.Sleep(5 * time.Second)

	storeFolder, err := os.MkdirTemp("REDACTED", "REDACTED")
	if err != nil {
		stdlog.Fatal(err)
	}
	defer os.RemoveAll(storeFolder)

	settings := rpctest.FetchSettings()

	leading, err := httpgateway.New(ledgerUID, settings.RPC.AcceptLocation)
	if err != nil {
		stdlog.Fatal(err)
	}

	ledger, err := leading.RapidLedger(context.Background(), 2)
	if err != nil {
		stdlog.Fatal(err)
	}

	db, err := dbm.NewGoLayerStore("REDACTED", storeFolder)
	if err != nil {
		stdlog.Fatal(err)
	}

	c, err := rapid.NewCustomer(
		context.Background(),
		ledgerUID,
		rapid.ValidateOptions{
			Duration: 504 * time.Hour, //
			Level: 2,
			Digest:   ledger.Digest(),
		},
		leading,
		[]source.Source{leading}, //
		dbs.New(db, ledgerUID),
		rapid.Tracer(log.VerifyingTracer()),
	)
	if err != nil {
		stdlog.Fatal(err)
	}
	defer func() {
		if err := c.Sanitize(); err != nil {
			stdlog.Fatal(err)
		}
	}()

	_, err = c.ValidateRapidLedgerAtLevel(context.Background(), 3, time.Now())
	if err != nil {
		stdlog.Fatal(err)
	}

	h, err := c.ValidatedRapidLedger(3)
	if err != nil {
		stdlog.Fatal(err)
	}

	fmt.Println("REDACTED", h.Level)
	//
}

func VerifyMain(m *testing.M) {
	//
	app := objectdepot.NewInRamSoftware()
	member := rpctest.BeginConsensuscore(app, rpctest.InhibitStdout)

	code := m.Run()

	//
	rpctest.HaltConsensuscore(member)
	os.Exit(code)
}
