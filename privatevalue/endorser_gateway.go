package privatevalue

import (
	"fmt"
	"net"
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/protocolio"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/facility"
	commitchronize "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
	privatevalueschema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/privatevalue"
)

const (
	fallbackDeadlineRetrievePersistMoments = 5
)

type endorserGateway struct {
	facility.FoundationFacility

	linkMutex commitchronize.Exclusion
	link    net.Conn

	deadlineRetrievePersist time.Duration
}

//
func (se *endorserGateway) Shutdown() error {
	se.DiscardLinkage()
	return nil
}

//
func (se *endorserGateway) EqualsAssociated() bool {
	se.linkMutex.Lock()
	defer se.linkMutex.Unlock()
	return se.equalsAssociated()
}

//
func (se *endorserGateway) ObtainAccessibleLinkage(linkageAccessibleChnl chan net.Conn) bool {
	se.linkMutex.Lock()
	defer se.linkMutex.Unlock()

	//
	select {
	case se.link = <-linkageAccessibleChnl:
		return true
	default:
	}
	return false
}

//
func (se *endorserGateway) PauseLinkage(linkageAccessibleChnl chan net.Conn, maximumPause time.Duration) error {
	select {
	case link := <-linkageAccessibleChnl:
		se.AssignLinkage(link)
	case <-time.After(maximumPause):
		return FaultLinkageDeadline
	}

	return nil
}

//
func (se *endorserGateway) AssignLinkage(freshLinkage net.Conn) {
	se.linkMutex.Lock()
	defer se.linkMutex.Unlock()
	se.link = freshLinkage
}

//
func (se *endorserGateway) DiscardLinkage() {
	se.linkMutex.Lock()
	defer se.linkMutex.Unlock()
	se.discardLinkage()
}

//
func (se *endorserGateway) FetchArtifact() (msg privatevalueschema.Signal, err error) {
	se.linkMutex.Lock()
	defer se.linkMutex.Unlock()

	if !se.equalsAssociated() {
		return msg, fmt.Errorf("REDACTED", FaultNegativeLinkage)
	}
	//
	limit := time.Now().Add(se.deadlineRetrievePersist)

	err = se.link.SetReadDeadline(limit)
	if err != nil {
		return
	}
	const maximumDistantEndorserSignalExtent = 1024 * 10
	schemaFetcher := protocolio.FreshSeparatedFetcher(se.link, maximumDistantEndorserSignalExtent)
	_, err = schemaFetcher.FetchSignal(&msg)
	if _, ok := err.(deadlineFailure); ok {
		if err != nil {
			err = fmt.Errorf("REDACTED", err, FaultRetrieveDeadline)
		} else {
			err = fmt.Errorf("REDACTED", FaultRetrieveDeadline)
		}

		se.Tracer.Diagnose("REDACTED", "REDACTED", se)
		se.discardLinkage()
	}

	return
}

//
func (se *endorserGateway) PersistArtifact(msg privatevalueschema.Signal) (err error) {
	se.linkMutex.Lock()
	defer se.linkMutex.Unlock()

	if !se.equalsAssociated() {
		return fmt.Errorf("REDACTED", FaultNegativeLinkage)
	}

	schemaPersistor := protocolio.FreshSeparatedPersistor(se.link)

	//
	limit := time.Now().Add(se.deadlineRetrievePersist)
	err = se.link.SetWriteDeadline(limit)
	if err != nil {
		return
	}

	_, err = schemaPersistor.PersistSignal(&msg)
	if _, ok := err.(deadlineFailure); ok {
		if err != nil {
			err = fmt.Errorf("REDACTED", err, FaultPersistDeadline)
		} else {
			err = fmt.Errorf("REDACTED", FaultPersistDeadline)
		}
		se.discardLinkage()
	}

	return
}

func (se *endorserGateway) equalsAssociated() bool {
	return se.link != nil
}

func (se *endorserGateway) discardLinkage() {
	if se.link != nil {
		if err := se.link.Close(); err != nil {
			se.Tracer.Failure("REDACTED", "REDACTED", err)
		}
		se.link = nil
	}
}
