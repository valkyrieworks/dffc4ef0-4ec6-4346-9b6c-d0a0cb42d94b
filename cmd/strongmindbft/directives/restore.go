package directives

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	strongos "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/os"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/privatevalue"
)

//
//
var RestoreEveryDirective = &cobra.Command{
	Use:     "REDACTED",
	Aliases: []string{"REDACTED"},
	Short:   "REDACTED",
	RunE:    restoreEveryDirective,
}

var retainLocationRegister bool

//
var RestoreStatusDirective = &cobra.Command{
	Use:     "REDACTED",
	Aliases: []string{"REDACTED"},
	Short:   "REDACTED",
	RunE: func(cmd *cobra.Command, arguments []string) (err error) {
		settings, err = AnalyzeSettings(cmd)
		if err != nil {
			return err
		}

		return restoreStatus(settings.DatastorePath(), tracer)
	},
}

func initialize() {
	RestoreEveryDirective.Flags().BoolVar(&retainLocationRegister, "REDACTED", false, "REDACTED")
}

//
var RestorePrivateAssessorDirective = &cobra.Command{
	Use:     "REDACTED",
	Aliases: []string{"REDACTED"},
	Short:   "REDACTED",
	RunE:    restorePrivateAssessor,
}

//
//
func restoreEveryDirective(cmd *cobra.Command, _ []string) (err error) {
	settings, err = AnalyzeSettings(cmd)
	if err != nil {
		return err
	}

	return restoreEvery(
		settings.DatastorePath(),
		settings.P2P.LocationRegisterRecord(),
		settings.PrivateAssessorTokenRecord(),
		settings.PrivateAssessorStatusRecord(),
		tracer,
	)
}

//
//
func restorePrivateAssessor(cmd *cobra.Command, _ []string) (err error) {
	settings, err = AnalyzeSettings(cmd)
	if err != nil {
		return err
	}

	restoreRecordPRV(settings.PrivateAssessorTokenRecord(), settings.PrivateAssessorStatusRecord(), tracer)
	return nil
}

//
func restoreEvery(datastorePath, locationRegisterRecord, privateItemTokenRecord, privateItemStatusRecord string, tracer log.Tracer) error {
	if retainLocationRegister {
		tracer.Details("REDACTED")
	} else {
		discardLocationRegister(locationRegisterRecord, tracer)
	}

	if err := os.RemoveAll(datastorePath); err == nil {
		tracer.Details("REDACTED", "REDACTED", datastorePath)
	} else {
		tracer.Failure("REDACTED", "REDACTED", datastorePath, "REDACTED", err)
	}

	if err := strongos.AssurePath(datastorePath, 0o700); err != nil {
		tracer.Failure("REDACTED", "REDACTED", err)
	}

	//
	restoreRecordPRV(privateItemTokenRecord, privateItemStatusRecord, tracer)
	return nil
}

//
func restoreStatus(datastorePath string, tracer log.Tracer) error {
	blockdatastore := filepath.Join(datastorePath, "REDACTED")
	status := filepath.Join(datastorePath, "REDACTED")
	wal := filepath.Join(datastorePath, "REDACTED")
	proof := filepath.Join(datastorePath, "REDACTED")
	transferOrdinal := filepath.Join(datastorePath, "REDACTED")

	if strongos.RecordPresent(blockdatastore) {
		if err := os.RemoveAll(blockdatastore); err == nil {
			tracer.Details("REDACTED", "REDACTED", blockdatastore)
		} else {
			tracer.Failure("REDACTED", "REDACTED", blockdatastore, "REDACTED", err)
		}
	}

	if strongos.RecordPresent(status) {
		if err := os.RemoveAll(status); err == nil {
			tracer.Details("REDACTED", "REDACTED", status)
		} else {
			tracer.Failure("REDACTED", "REDACTED", status, "REDACTED", err)
		}
	}

	if strongos.RecordPresent(wal) {
		if err := os.RemoveAll(wal); err == nil {
			tracer.Details("REDACTED", "REDACTED", wal)
		} else {
			tracer.Failure("REDACTED", "REDACTED", wal, "REDACTED", err)
		}
	}

	if strongos.RecordPresent(proof) {
		if err := os.RemoveAll(proof); err == nil {
			tracer.Details("REDACTED", "REDACTED", proof)
		} else {
			tracer.Failure("REDACTED", "REDACTED", proof, "REDACTED", err)
		}
	}

	if strongos.RecordPresent(transferOrdinal) {
		if err := os.RemoveAll(transferOrdinal); err == nil {
			tracer.Details("REDACTED", "REDACTED", transferOrdinal)
		} else {
			tracer.Failure("REDACTED", "REDACTED", transferOrdinal, "REDACTED", err)
		}
	}

	if err := strongos.AssurePath(datastorePath, 0o700); err != nil {
		tracer.Failure("REDACTED", "REDACTED", err)
	}
	return nil
}

func restoreRecordPRV(privateItemTokenRecord, privateItemStatusRecord string, tracer log.Tracer) {
	if _, err := os.Stat(privateItemTokenRecord); err == nil {
		pv := privatevalue.FetchRecordPRVBlankStatus(privateItemTokenRecord, privateItemStatusRecord)
		pv.Restore()
		tracer.Details(
			"REDACTED",
			"REDACTED", privateItemTokenRecord,
			"REDACTED", privateItemStatusRecord,
		)
	} else {
		pv := privatevalue.ProduceRecordPRV(privateItemTokenRecord, privateItemStatusRecord)
		pv.Persist()
		tracer.Details(
			"REDACTED",
			"REDACTED", privateItemTokenRecord,
			"REDACTED", privateItemStatusRecord,
		)
	}
}

func discardLocationRegister(locationRegisterRecord string, tracer log.Tracer) {
	if err := os.Remove(locationRegisterRecord); err == nil {
		tracer.Details("REDACTED", "REDACTED", locationRegisterRecord)
	} else if !os.IsNotExist(err) {
		tracer.Details("REDACTED", "REDACTED", locationRegisterRecord, "REDACTED", err)
	}
}
