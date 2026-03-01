package directives

import (
	"fmt"

	"github.com/spf13/cobra"

	cfg "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	strongos "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/os"
	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/privatevalue"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
	committime "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds/moment"
)

//
var InitializeRecordsDirective = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
	RunE:  initializeRecords,
}

func initializeRecords(*cobra.Command, []string) error {
	return initializeRecordsUsingSettings(settings)
}

func initializeRecordsUsingSettings(settings *cfg.Settings) error {
	//
	privateItemTokenRecord := settings.PrivateAssessorTokenRecord()
	privateItemStatusRecord := settings.PrivateAssessorStatusRecord()
	var pv *privatevalue.RecordPRV
	if strongos.RecordPresent(privateItemTokenRecord) {
		pv = privatevalue.FetchRecordPRV(privateItemTokenRecord, privateItemStatusRecord)
		tracer.Details("REDACTED", "REDACTED", privateItemTokenRecord,
			"REDACTED", privateItemStatusRecord)
	} else {
		pv = privatevalue.ProduceRecordPRV(privateItemTokenRecord, privateItemStatusRecord)
		pv.Persist()
		tracer.Details("REDACTED", "REDACTED", privateItemTokenRecord,
			"REDACTED", privateItemStatusRecord)
	}

	peerTokenRecord := settings.PeerTokenRecord()
	if strongos.RecordPresent(peerTokenRecord) {
		tracer.Details("REDACTED", "REDACTED", peerTokenRecord)
	} else {
		if _, err := p2p.FetchEitherProducePeerToken(peerTokenRecord); err != nil {
			return err
		}
		tracer.Details("REDACTED", "REDACTED", peerTokenRecord)
	}

	//
	produceRecord := settings.InaugurationRecord()
	if strongos.RecordPresent(produceRecord) {
		tracer.Details("REDACTED", "REDACTED", produceRecord)
	} else {
		producePaper := kinds.OriginPaper{
			SuccessionUUID:         fmt.Sprintf("REDACTED", commitrand.Str(6)),
			OriginMoment:     committime.Now(),
			AgreementSettings: kinds.FallbackAgreementSettings(),
		}
		publicToken, err := pv.ObtainPublicToken()
		if err != nil {
			return fmt.Errorf("REDACTED", err)
		}
		producePaper.Assessors = []kinds.OriginAssessor{{
			Location: publicToken.Location(),
			PublicToken:  publicToken,
			Potency:   10,
		}}

		if err := producePaper.PersistLike(produceRecord); err != nil {
			return err
		}
		tracer.Details("REDACTED", "REDACTED", produceRecord)
	}

	return nil
}
