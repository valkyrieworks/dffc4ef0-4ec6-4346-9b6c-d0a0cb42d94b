package transferordinal

import (
	"context"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/facility"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/ordinalizer"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

//

const (
	listener = "REDACTED"
)

//
//
type OrdinalizerFacility struct {
	facility.FoundationFacility

	transferIndexer           TransferOrdinalizer
	ledgerIndexer        ordinalizer.LedgerOrdinalizer
	incidentPipeline         *kinds.IncidentChannel
	concludeUponFailure bool
}

//
func FreshOrdinalizerFacility(
	transferIndexer TransferOrdinalizer,
	ledgerIndexer ordinalizer.LedgerOrdinalizer,
	incidentPipeline *kinds.IncidentChannel,
	concludeUponFailure bool,
) *OrdinalizerFacility {

	is := &OrdinalizerFacility{transferIndexer: transferIndexer, ledgerIndexer: ledgerIndexer, incidentPipeline: incidentPipeline, concludeUponFailure: concludeUponFailure}
	is.FoundationFacility = *facility.FreshFoundationFacility(nil, "REDACTED", is)
	return is
}

//
//
func (is *OrdinalizerFacility) UponInitiate() error {
	//
	//
	//
	ledgerUnder, err := is.incidentPipeline.ListenUncached(
		context.Background(),
		listener,
		kinds.IncidentInquireFreshLedgerIncidents)
	if err != nil {
		return err
	}

	transUnder, err := is.incidentPipeline.ListenUncached(context.Background(), listener, kinds.IncidentInquireTransfer)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case <-ledgerUnder.Aborted():
				return
			case msg := <-ledgerUnder.Out():
				incidentFreshLedgerIncidents := msg.Data().(kinds.IncidentDataFreshLedgerIncidents)
				altitude := incidentFreshLedgerIncidents.Altitude
				countTrans := incidentFreshLedgerIncidents.CountTrans

				cluster := FreshCluster(countTrans)

				for i := int64(0); i < countTrans; i++ {
					message2 := <-transUnder.Out()
					transferOutcome := message2.Data().(kinds.IncidentDataTransfer).TransferOutcome

					if err = cluster.Add(&transferOutcome); err != nil {
						is.Tracer.Failure(
							"REDACTED",
							"REDACTED", altitude,
							"REDACTED", transferOutcome.Ordinal,
							"REDACTED", err,
						)

						if is.concludeUponFailure {
							if err := is.Halt(); err != nil {
								is.Tracer.Failure("REDACTED", "REDACTED", err)
							}
							return
						}
					}
				}

				if err := is.ledgerIndexer.Ordinal(incidentFreshLedgerIncidents); err != nil {
					is.Tracer.Failure("REDACTED", "REDACTED", altitude, "REDACTED", err)
					if is.concludeUponFailure {
						if err := is.Halt(); err != nil {
							is.Tracer.Failure("REDACTED", "REDACTED", err)
						}
						return
					}
				} else {
					is.Tracer.Details("REDACTED", "REDACTED", altitude)
				}

				if err = is.transferIndexer.AppendCluster(cluster); err != nil {
					is.Tracer.Failure("REDACTED", "REDACTED", altitude, "REDACTED", err)
					if is.concludeUponFailure {
						if err := is.Halt(); err != nil {
							is.Tracer.Failure("REDACTED", "REDACTED", err)
						}
						return
					}
				} else {
					is.Tracer.Diagnose("REDACTED", "REDACTED", altitude, "REDACTED", countTrans)
				}
			}
		}
	}()
	return nil
}

//
func (is *OrdinalizerFacility) UponHalt() {
	if is.incidentPipeline.EqualsActive() {
		_ = is.incidentPipeline.UnlistenEvery(context.Background(), listener)
	}
}
