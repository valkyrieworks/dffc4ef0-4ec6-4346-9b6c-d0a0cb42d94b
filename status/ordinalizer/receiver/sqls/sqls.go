//
package sqls

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

	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/broadcastlisten/inquire"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

const (
	registryLedgers     = "REDACTED"
	registryTransferOutcomes  = "REDACTED"
	registryIncidents     = "REDACTED"
	registryProperties = "REDACTED"
	handlerAlias      = "REDACTED"
)

//
//
//
type IncidentReceiver struct {
	depot   *sql.DB
	successionUUID string
}

//
//
//
func FreshIncidentReceiver(linkTxt, successionUUID string) (*IncidentReceiver, error) {
	db, err := sql.Open(handlerAlias, linkTxt)
	if err != nil {
		return nil, err
	}
	return &IncidentReceiver{
		depot:   db,
		successionUUID: successionUUID,
	}, nil
}

//
//
func (es *IncidentReceiver) DB() *sql.DB { return es.depot }

//
//
//
//
func executeInsideTransfer(db *sql.DB, inquire func(*sql.Tx) error) error {
	dstransaction, err := db.Begin()
	if err != nil {
		return err
	}
	if err := inquire(dstransaction); err != nil {
		_ = dstransaction.Rollback() //
		return err
	}
	return dstransaction.Commit()
}

func executeLumpAppend(db *sql.DB, registryAlias string, fields []string, appends [][]any) error {
	return executeInsideTransfer(db, func(tx *sql.Tx) error {
		command, err := tx.Prepare(pq.CopyIn(registryAlias, fields...))
		if err != nil {
			return fmt.Errorf("REDACTED", err)
		}
		defer command.Close()
		for _, append := range appends {
			if _, err := command.Exec(append...); err != nil {
				return fmt.Errorf("REDACTED", err)
			}
		}
		if _, err := command.Exec(); err != nil {
			return fmt.Errorf("REDACTED", err)
		}
		return nil
	})
}

func unpredictableLargesequence() int64 {
	return rand.Int63()
}

var (
	transferrAppendFields   = []string{"REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED"}
	incidentAppendFields = []string{"REDACTED", "REDACTED", "REDACTED", "REDACTED"}
	propertyAppendFields  = []string{"REDACTED", "REDACTED", "REDACTED", "REDACTED"}
)

func lumpAppendIncidents(ledgerUUID, transferUUID int64, incidents []iface.Incident) (incidentAppends, propertyAppends [][]any) {
	//
	var transferUUIDParam any
	if transferUUID > 0 {
		transferUUIDParam = transferUUID
	}
	for _, incident := range incidents {
		//
		if incident.Kind == "REDACTED" {
			continue
		}
		incidentUUID := unpredictableLargesequence()
		incidentAppends = append(incidentAppends, []any{incidentUUID, ledgerUUID, transferUUIDParam, incident.Kind})
		for _, property := range incident.Properties {
			if !property.Ordinal {
				continue
			}
			complexToken := incident.Kind + "REDACTED" + property.Key
			propertyAppends = append(propertyAppends, []any{incidentUUID, property.Key, complexToken, property.Datum})
		}
	}
	return incidentAppends, propertyAppends
}

//
//
//
//
func createPositionedIncident(complexToken, datum string) iface.Incident {
	i := strings.Index(complexToken, "REDACTED")
	if i < 0 {
		return iface.Incident{Kind: complexToken}
	}
	return iface.Incident{Kind: complexToken[:i], Properties: []iface.IncidentProperty{
		{Key: complexToken[i+1:], Datum: datum, Ordinal: true},
	}}
}

//
//
func (es *IncidentReceiver) PositionLedgerIncidents(h kinds.IncidentDataFreshLedgerIncidents) error {
	ts := time.Now().UTC()

	//
	//
	var ledgerUUID int64
	//
	err := es.depot.QueryRow(`
REDACTED`+registryLedgers+`REDACTED)
REDACTED)
REDACTEDG
REDACTED;
REDACTED`, h.Altitude, es.successionUUID, ts).Scan(&ledgerUUID)
	if err == sql.ErrNoRows {
		return nil //
	} else if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	//
	incidents := append([]iface.Incident{createPositionedIncident(kinds.LedgerAltitudeToken, strconv.FormatInt(h.Altitude, 10))}, h.Incidents...)
	//
	incidentAppends, propertyAppends := lumpAppendIncidents(ledgerUUID, 0, incidents)
	if err := executeLumpAppend(es.depot, registryIncidents, incidentAppendFields, incidentAppends); err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	if err := executeLumpAppend(es.depot, registryProperties, propertyAppendFields, propertyAppends); err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	return nil
}

//
func (es *IncidentReceiver) obtainLedgerIDXDstore(elevations []int64) ([]int64, error) {
	var ledgerIDXDstore pq.Int64Array
	if err := es.depot.QueryRow(`
REDACTED(
REDACTED`+registryLedgers+`REDACTED1
REDACTED`,
		es.successionUUID, pq.Array(elevations)).Scan(&ledgerIDXDstore); err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}
	return ledgerIDXDstore, nil
}

func precacheTransferrPresence(db *sql.DB, ledgerIDXDstore []int64, indices []uint32) ([]bool, error) {
	var presence []bool
	if err := db.QueryRow(`
REDACTED(
REDACTED`+registryTransferOutcomes+`REDACTED)
REDACTED`,
		pq.Array(ledgerIDXDstore), pq.Array(indices)).Scan((*pq.BoolArray)(&presence)); err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}
	return presence, nil
}

func (es *IncidentReceiver) PositionTransferIncidents(transfers []*iface.TransferOutcome) error {
	ts := time.Now().UTC()
	elevations := make([]int64, len(transfers))
	indices := make([]uint32, len(transfers))
	for i, txr := range transfers {
		elevations[i] = txr.Altitude
		indices[i] = txr.Ordinal
	}
	//
	//
	ledgerIDXDstore, err := es.obtainLedgerIDXDstore(elevations)
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	earlierPositioned, err := precacheTransferrPresence(es.depot, ledgerIDXDstore, indices)
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	transferrAppends, propertyAppends, incidentAppends := make([][]any, 0, len(transfers)), make([][]any, 0, len(transfers)), make([][]any, 0, len(transfers))
	for i, txr := range transfers {
		if earlierPositioned[i] {
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
		transferUUID := unpredictableLargesequence()
		transferrAppends = append(transferrAppends, []any{transferUUID, ledgerIDXDstore[i], txr.Ordinal, ts, transferDigest, outcomeData})
		//
		incidents := append([]iface.Incident{
			createPositionedIncident(kinds.TransferDigestToken, transferDigest),
			createPositionedIncident(kinds.TransferAltitudeToken, strconv.FormatInt(txr.Altitude, 10)),
		},
			txr.Outcome.Incidents...,
		)
		freshIncidentAppends, freshPropertyAppends := lumpAppendIncidents(ledgerIDXDstore[i], transferUUID, incidents)
		incidentAppends = append(incidentAppends, freshIncidentAppends...)
		propertyAppends = append(propertyAppends, freshPropertyAppends...)
	}
	if err := executeLumpAppend(es.depot, registryTransferOutcomes, transferrAppendFields, transferrAppends); err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	if err := executeLumpAppend(es.depot, registryIncidents, incidentAppendFields, incidentAppends); err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	if err := executeLumpAppend(es.depot, registryProperties, propertyAppendFields, propertyAppends); err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	return nil
}

//
func (es *IncidentReceiver) LookupLedgerIncidents(_ context.Context, _ *inquire.Inquire) ([]int64, error) {
	return nil, errors.New("REDACTED")
}

//
func (es *IncidentReceiver) LookupTransferIncidents(_ context.Context, _ *inquire.Inquire) ([]*iface.TransferOutcome, error) {
	return nil, errors.New("REDACTED")
}

//
func (es *IncidentReceiver) ObtainTransferViaDigest(_ []byte) (*iface.TransferOutcome, error) {
	return nil, errors.New("REDACTED")
}

//
func (es *IncidentReceiver) OwnsLedger(_ int64) (bool, error) {
	return false, errors.New("REDACTED")
}

//
func (es *IncidentReceiver) Halt() error { return es.depot.Close() }
