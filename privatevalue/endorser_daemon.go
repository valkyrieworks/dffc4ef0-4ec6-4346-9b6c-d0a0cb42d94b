package privatevalue

import (
	"io"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/facility"
	commitchronize "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
	privatevalueschema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/privatevalue"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

//
type CertificationSolicitProcessorMethod func(
	privateItem kinds.PrivateAssessor,
	solicitArtifact privatevalueschema.Signal,
	successionUUID string) (privatevalueschema.Signal, error)

type EndorserDaemon struct {
	facility.FoundationFacility

	gateway *EndorserCallerGateway
	successionUUID  string
	privateItem  kinds.PrivateAssessor

	processorMutex               commitchronize.Exclusion
	certificationSolicitProcessor CertificationSolicitProcessorMethod
}

func FreshEndorserDaemon(gateway *EndorserCallerGateway, successionUUID string, privateItem kinds.PrivateAssessor) *EndorserDaemon {
	ss := &EndorserDaemon{
		gateway:                 gateway,
		successionUUID:                  successionUUID,
		privateItem:                  privateItem,
		certificationSolicitProcessor: FallbackCertificationSolicitProcessor,
	}

	ss.FoundationFacility = *facility.FreshFoundationFacility(gateway.Tracer, "REDACTED", ss)

	return ss
}

//
func (ss *EndorserDaemon) UponInitiate() error {
	go ss.facilityCycle()
	return nil
}

//
func (ss *EndorserDaemon) UponHalt() {
	ss.gateway.Tracer.Diagnose("REDACTED")
	_ = ss.gateway.Shutdown()
}

//
func (ss *EndorserDaemon) AssignSolicitProcessor(certificationSolicitProcessor CertificationSolicitProcessorMethod) {
	ss.processorMutex.Lock()
	defer ss.processorMutex.Unlock()
	ss.certificationSolicitProcessor = certificationSolicitProcessor
}

func (ss *EndorserDaemon) facilityAwaitingSolicit() {
	if !ss.EqualsActive() {
		return //
	}

	req, err := ss.gateway.FetchArtifact()
	if err != nil {
		if err != io.EOF {
			ss.Tracer.Failure("REDACTED", "REDACTED", err)
		}
		return
	}

	var res privatevalueschema.Signal
	{
		//
		ss.processorMutex.Lock()
		defer ss.processorMutex.Unlock()
		res, err = ss.certificationSolicitProcessor(ss.privateItem, req, ss.successionUUID)
		if err != nil {
			//
			ss.Tracer.Failure("REDACTED", "REDACTED", err)
		}
	}

	err = ss.gateway.PersistArtifact(res)
	if err != nil {
		ss.Tracer.Failure("REDACTED", "REDACTED", err)
	}
}

func (ss *EndorserDaemon) facilityCycle() {
	for {
		select {
		default:
			err := ss.gateway.assureLinkage()
			if err != nil {
				return
			}
			ss.facilityAwaitingSolicit()

		case <-ss.Exit():
			return
		}
	}
}
