package agreement

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"path/filepath"
	"testing"
	"time"

	db "github.com/valkyrieworks/-db"

	"github.com/valkyrieworks/iface/instance/objectdepot"
	cfg "github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/intrinsic/verify"
	"github.com/valkyrieworks/utils/log"
	engineseed "github.com/valkyrieworks/utils/random"
	"github.com/valkyrieworks/privatekey"
	"github.com/valkyrieworks/gateway"
	sm "github.com/valkyrieworks/status"
	"github.com/valkyrieworks/depot"
	"github.com/valkyrieworks/kinds"
)

//
//
//
//
//
func JournalComposeNLedgers(t *testing.T, wr io.Writer, countLedgers int, settings *cfg.Settings) (err error) {
	app := objectdepot.NewDurableSoftware(filepath.Join(settings.StoreFolder(), "REDACTED"))

	tracer := log.VerifyingTracer().With("REDACTED", "REDACTED")
	tracer.Details("REDACTED", "REDACTED", countLedgers)

	//
	//
	//
	privateRatifierKeyEntry := settings.PrivateRatifierKeyEntry()
	privateRatifierStatusEntry := settings.PrivateRatifierStatusEntry()
	privateRatifier := privatekey.ImportOrGenerateEntryPV(privateRatifierKeyEntry, privateRatifierStatusEntry)
	generatePaper, err := kinds.OriginPaperFromEntry(settings.OriginEntry())
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	ledgerDepotStore := db.NewMemoryStore()
	statusStore := ledgerDepotStore
	statusDepot := sm.NewDepot(statusStore, sm.DepotSettings{
		DropIfaceReplies: false,
	})
	status, err := sm.CreateOriginStatus(generatePaper)
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	status.Release.Agreement.App = objectdepot.ApplicationRelease
	if err = statusDepot.Persist(status); err != nil {
		t.Error(err)
	}

	ledgerDepot := depot.NewLedgerDepot(ledgerDepotStore)

	gatewayApplication := gateway.NewApplicationLinks(gateway.NewNativeCustomerOriginator(app), gateway.NoopStats())
	gatewayApplication.AssignTracer(tracer.With("REDACTED", "REDACTED"))
	if err := gatewayApplication.Begin(); err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	t.Cleanup(func() {
		if err := gatewayApplication.Halt(); err != nil {
			t.Error(err)
		}
	})

	eventBus := kinds.NewEventBus()
	eventBus.AssignTracer(tracer.With("REDACTED", "REDACTED"))
	if err := eventBus.Begin(); err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	t.Cleanup(func() {
		if err := eventBus.Halt(); err != nil {
			t.Error(err)
		}
	})
	txpool := emptyTxpool{}
	eventpool := sm.EmptyProofDepository{}
	ledgerExecute := sm.NewLedgerRunner(statusDepot, log.VerifyingTracer(), gatewayApplication.Agreement(), txpool, eventpool, ledgerDepot)
	agreementStatus := NewStatus(settings.Agreement, status.Clone(), ledgerExecute, ledgerDepot, txpool, eventpool)
	agreementStatus.AssignTracer(tracer)
	agreementStatus.AssignEventBus(eventBus)
	if privateRatifier != nil {
		agreementStatus.CollectionPrivateRatifier(privateRatifier)
	}
	//

	//
	countLedgersInscribed := make(chan struct{})
	wal := newOctetBufferJournal(tracer, NewJournalSerializer(wr), int64(countLedgers), countLedgersInscribed)
	//
	if err := wal.Record(TerminateLevelSignal{0}); err != nil {
		t.Error(err)
	}

	agreementStatus.wal = wal

	if err := agreementStatus.Begin(); err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	select {
	case <-countLedgersInscribed:
		if err := agreementStatus.Halt(); err != nil {
			t.Error(err)
		}
		return nil
	case <-time.After(1 * time.Minute):
		if err := agreementStatus.Halt(); err != nil {
			t.Error(err)
		}
		return fmt.Errorf("REDACTED", countLedgers)
	}
}

//
func JournalWithNLedgers(t *testing.T, countLedgers int, settings *cfg.Settings) (data []byte, err error) {
	var b bytes.Buffer
	wr := bufio.NewWriter(&b)

	if err := JournalComposeNLedgers(t, wr, countLedgers, settings); err != nil {
		return []byte{}, err
	}

	wr.Flush()
	return b.Bytes(), nil
}

func randomPort() int {
	//
	root, disperse := 20000, 20000
	return root + engineseed.Intn(disperse)
}

func createLocations() (string, string, string) {
	begin := randomPort()
	return fmt.Sprintf("REDACTED", begin),
		fmt.Sprintf("REDACTED", begin+1),
		fmt.Sprintf("REDACTED", begin+2)
}

//
func fetchSettings(t *testing.T) *cfg.Settings {
	c := verify.RestoreVerifyOrigin(t.Name())

	//
	cmt, rpc, grpc := createLocations()
	c.P2P.AcceptLocation = cmt
	c.RPC.AcceptLocation = rpc
	c.RPC.GRPCAcceptLocation = grpc
	return c
}

//
//
//
type octetBufferJournal struct {
	enc               *JournalSerializer
	ceased           bool
	levelToHalt      int64
	alertWhenHaltsTo chan<- struct{}

	tracer log.Tracer
}

//
var staticTime, _ = time.Parse(time.RFC3339, "REDACTED")

func newOctetBufferJournal(tracer log.Tracer, enc *JournalSerializer, nLedgers int64, alertHalt chan<- struct{}) *octetBufferJournal {
	return &octetBufferJournal{
		enc:               enc,
		levelToHalt:      nLedgers,
		alertWhenHaltsTo: alertHalt,
		tracer:            tracer,
	}
}

//
//
//
func (w *octetBufferJournal) Record(m JournalSignal) error {
	if w.ceased {
		w.tracer.Diagnose("REDACTED", "REDACTED", m)
		return nil
	}

	if terminateMessage, ok := m.(TerminateLevelSignal); ok {
		w.tracer.Diagnose("REDACTED", "REDACTED", terminateMessage.Level, "REDACTED", w.levelToHalt)
		if terminateMessage.Level == w.levelToHalt {
			w.tracer.Diagnose("REDACTED", "REDACTED", terminateMessage.Level)
			w.alertWhenHaltsTo <- struct{}{}
			w.ceased = true
			return nil
		}
	}

	w.tracer.Diagnose("REDACTED", "REDACTED", m)
	err := w.enc.Serialize(&ScheduledJournalSignal{staticTime, m})
	if err != nil {
		panic(fmt.Sprintf("REDACTED", m))
	}

	return nil
}

func (w *octetBufferJournal) RecordAlign(m JournalSignal) error {
	return w.Record(m)
}

func (w *octetBufferJournal) PurgeAndAlign() error { return nil }

func (w *octetBufferJournal) ScanForTerminateLevel(
	int64,
	*JournalScanSettings,
) (rd io.ReadCloser, located bool, err error) {
	return nil, false, nil
}

func (w *octetBufferJournal) Begin() error { return nil }
func (w *octetBufferJournal) Halt() error  { return nil }
func (w *octetBufferJournal) Wait()        {}
