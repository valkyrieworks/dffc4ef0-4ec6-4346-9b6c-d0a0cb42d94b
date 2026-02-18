//
package psql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/cosmos/gogoproto/proto"
	"github.com/lib/pq"

	iface "github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/utils/broadcast/inquire"
	"github.com/valkyrieworks/kinds"
)

const (
	sheetLedgers     = "REDACTED"
	sheetTransferOutcomes  = "REDACTED"
	sheetEvents     = "REDACTED"
	sheetProperties = "REDACTED"
	handlerLabel      = "REDACTED"
)

//
//
//
type EventDrain struct {
	depot   *sql.DB
	ledgerUID string
}

//
//
//
func NewEventDrain(linkStr, ledgerUID string) (*EventDrain, error) {
	db, err := sql.Open(handlerLabel, linkStr)
	if err != nil {
		return nil, err
	}
	return &EventDrain{
		depot:   db,
		ledgerUID: ledgerUID,
	}, nil
}

//
//
func (es *EventDrain) DB() *sql.DB { return es.depot }

//
//
//
//
func runInTransfer(db *sql.DB, inquire func(*sql.Tx) error) error {
	storedtransfer, err := db.Begin()
	if err != nil {
		return err
	}
	if err := inquire(storedtransfer); err != nil {
		_ = storedtransfer.Rollback() //
		return err
	}
	return storedtransfer.Commit()
}

func runBatchEmbed(db *sql.DB, sheetLabel string, fields []string, embeddings [][]any) error {
	return runInTransfer(db, func(tx *sql.Tx) error {
		command, err := tx.Prepare(pq.CopyIn(sheetLabel, fields...))
		if err != nil {
			return fmt.Errorf("REDACTED", err)
		}
		defer command.Close()
		for _, embed := range embeddings {
			if _, err := command.Exec(embed...); err != nil {
				return fmt.Errorf("REDACTED", err)
			}
		}
		if _, err := command.Exec(); err != nil {
			return fmt.Errorf("REDACTED", err)
		}
		return nil
	})
}

func arbitraryLargeseq() int64 {
	return rand.Int63()
}

var (
	transferrEmbedFields   = []string{"REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED"}
	eventEmbedFields = []string{"REDACTED", "REDACTED", "REDACTED", "REDACTED"}
	propertyEmbedFields  = []string{"REDACTED", "REDACTED", "REDACTED", "REDACTED"}
)

func batchEmbedEvents(ledgerUID, transferUID int64, events []iface.Event) (eventEmbeddings, propertyEmbeddings [][]any) {
	//
	var transferUIDArgument any
	if transferUID > 0 {
		transferUIDArgument = transferUID
	}
	for _, event := range events {
		//
		if event.Kind == "REDACTED" {
			continue
		}
		eventUID := arbitraryLargeseq()
		eventEmbeddings = append(eventEmbeddings, []any{eventUID, ledgerUID, transferUIDArgument, event.Kind})
		for _, property := range event.Properties {
			if !property.Ordinal {
				continue
			}
			compoundKey := event.Kind + "REDACTED" + property.Key
			propertyEmbeddings = append(propertyEmbeddings, []any{eventUID, property.Key, compoundKey, property.Item})
		}
	}
	return eventEmbeddings, propertyEmbeddings
}

//
//
//
//
func createCatalogedEvent(compoundKey, item string) iface.Event {
	i := strings.Index(compoundKey, "REDACTED")
	if i < 0 {
		return iface.Event{Kind: compoundKey}
	}
	return iface.Event{Kind: compoundKey[:i], Properties: []iface.EventProperty{
		{Key: compoundKey[i+1:], Item: item, Ordinal: true},
	}}
}

//
//
func (es *EventDrain) OrdinalLedgerEvents(h kinds.EventDataNewLedgerEvents) error {
	ts := time.Now().UTC()

	//
	//
	var ledgerUID int64
	//
	err := es.depot.QueryRow(`
REDACTED`+sheetLedgers+`REDACTED)
REDACTED)
REDACTEDG
REDACTED;
REDACTED`, h.Level, es.ledgerUID, ts).Scan(&ledgerUID)
	if err == sql.ErrNoRows {
		return nil //
	} else if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	//
	events := append([]iface.Event{createCatalogedEvent(kinds.LedgerLevelKey, strconv.FormatInt(h.Level, 10))}, h.Events...)
	//
	eventEmbeddings, propertyEmbeddings := batchEmbedEvents(ledgerUID, 0, events)
	if err := runBatchEmbed(es.depot, sheetEvents, eventEmbedFields, eventEmbeddings); err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	if err := runBatchEmbed(es.depot, sheetProperties, propertyEmbedFields, propertyEmbeddings); err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	return nil
}

//
func (es *EventDrain) fetchLedgerIDXDatastore(levels []int64) ([]int64, error) {
	var ledgerIDXDatastore pq.Int64Array
	if err := es.depot.QueryRow(`
REDACTED(
REDACTED`+sheetLedgers+`REDACTED1
REDACTED`,
		es.ledgerUID, pq.Array(levels)).Scan(&ledgerIDXDatastore); err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}
	return ledgerIDXDatastore, nil
}

func preloadTransferrPresence(db *sql.DB, ledgerIDXDatastore []int64, listings []uint32) ([]bool, error) {
	var presence []bool
	if err := db.QueryRow(`
REDACTED(
REDACTED`+sheetTransferOutcomes+`REDACTED)
REDACTED`,
		pq.Array(ledgerIDXDatastore), pq.Array(listings)).Scan((*pq.BoolArray)(&presence)); err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}
	return presence, nil
}

func (es *EventDrain) OrdinalTransferEvents(transfers []*iface.TransOutcome) error {
	ts := time.Now().UTC()
	levels := make([]int64, len(transfers))
	listings := make([]uint32, len(transfers))
	for i, txr := range transfers {
		levels[i] = txr.Level
		listings[i] = txr.Ordinal
	}
	//
	//
	ledgerIDXDatastore, err := es.fetchLedgerIDXDatastore(levels)
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	yetCataloged, err := preloadTransferrPresence(es.depot, ledgerIDXDatastore, listings)
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	transferrEmbeddings, propertyEmbeddings, eventEmbeddings := make([][]any, 0, len(transfers)), make([][]any, 0, len(transfers)), make([][]any, 0, len(transfers))
	for i, txr := range transfers {
		if yetCataloged[i] {
			continue
		}
		//
		outcomeData, err := proto.Marshal(txr)
		if err != nil {
			return fmt.Errorf("REDACTED", err)
		}
		//
		transferDigest := fmt.Sprintf("REDACTED", kinds.Tx(txr.Tx).Digest())
		//
		transferUID := arbitraryLargeseq()
		transferrEmbeddings = append(transferrEmbeddings, []any{transferUID, ledgerIDXDatastore[i], txr.Ordinal, ts, transferDigest, outcomeData})
		//
		events := append([]iface.Event{
			createCatalogedEvent(kinds.TransferDigestKey, transferDigest),
			createCatalogedEvent(kinds.TransferLevelKey, strconv.FormatInt(txr.Level, 10)),
		},
			txr.Outcome.Events...,
		)
		newEventEmbeddings, newPropertyEmbeddings := batchEmbedEvents(ledgerIDXDatastore[i], transferUID, events)
		eventEmbeddings = append(eventEmbeddings, newEventEmbeddings...)
		propertyEmbeddings = append(propertyEmbeddings, newPropertyEmbeddings...)
	}
	if err := runBatchEmbed(es.depot, sheetTransferOutcomes, transferrEmbedFields, transferrEmbeddings); err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	if err := runBatchEmbed(es.depot, sheetEvents, eventEmbedFields, eventEmbeddings); err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	if err := runBatchEmbed(es.depot, sheetProperties, propertyEmbedFields, propertyEmbeddings); err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	return nil
}

//
func (es *EventDrain) ScanLedgerEvents(_ context.Context, _ *inquire.Inquire) ([]int64, error) {
	return nil, errors.New("REDACTED")
}

//
func (es *EventDrain) ScanTransferEvents(_ context.Context, _ *inquire.Inquire) ([]*iface.TransOutcome, error) {
	return nil, errors.New("REDACTED")
}

//
func (es *EventDrain) FetchTransferByDigest(_ []byte) (*iface.TransOutcome, error) {
	return nil, errors.New("REDACTED")
}

//
func (es *EventDrain) HasLedger(_ int64) (bool, error) {
	return false, errors.New("REDACTED")
}

//
func (es *EventDrain) Halt() error { return es.depot.Close() }
