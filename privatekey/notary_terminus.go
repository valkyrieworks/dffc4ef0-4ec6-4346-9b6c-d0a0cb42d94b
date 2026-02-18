package privatekey

import (
	"fmt"
	"net"
	"time"

	"github.com/valkyrieworks/utils/protoio"
	"github.com/valkyrieworks/utils/daemon"
	engineconnect "github.com/valkyrieworks/utils/align"
	privatekeyproto "github.com/valkyrieworks/schema/consensuscore/privatekey"
)

const (
	standardDeadlineScanRecordMoments = 5
)

type notaryTerminus struct {
	daemon.RootDaemon

	linkMutex engineconnect.Lock
	link    net.Conn

	deadlineScanRecord time.Duration
}

//
func (se *notaryTerminus) End() error {
	se.DiscardLinkage()
	return nil
}

//
func (se *notaryTerminus) IsLinked() bool {
	se.linkMutex.Lock()
	defer se.linkMutex.Unlock()
	return se.isLinked()
}

//
func (se *notaryTerminus) FetchAccessibleLinkage(linkageAccessibleChan chan net.Conn) bool {
	se.linkMutex.Lock()
	defer se.linkMutex.Unlock()

	//
	select {
	case se.link = <-linkageAccessibleChan:
		return true
	default:
	}
	return false
}

//
func (se *notaryTerminus) WaitLinkage(linkageAccessibleChan chan net.Conn, maximumWait time.Duration) error {
	select {
	case link := <-linkageAccessibleChan:
		se.CollectionLinkage(link)
	case <-time.After(maximumWait):
		return ErrLinkageDeadline
	}

	return nil
}

//
func (se *notaryTerminus) CollectionLinkage(newLinkage net.Conn) {
	se.linkMutex.Lock()
	defer se.linkMutex.Unlock()
	se.link = newLinkage
}

//
func (se *notaryTerminus) DiscardLinkage() {
	se.linkMutex.Lock()
	defer se.linkMutex.Unlock()
	se.discardLinkage()
}

//
func (se *notaryTerminus) ScanSignal() (msg privatekeyproto.Signal, err error) {
	se.linkMutex.Lock()
	defer se.linkMutex.Unlock()

	if !se.isLinked() {
		return msg, fmt.Errorf("REDACTED", ErrNoLinkage)
	}
	//
	limit := time.Now().Add(se.deadlineScanRecord)

	err = se.link.SetReadDeadline(limit)
	if err != nil {
		return
	}
	const maximumDistantNotaryMessageVolume = 1024 * 10
	schemaScanner := protoio.NewSeparatedScanner(se.link, maximumDistantNotaryMessageVolume)
	_, err = schemaScanner.ScanMessage(&msg)
	if _, ok := err.(deadlineFault); ok {
		if err != nil {
			err = fmt.Errorf("REDACTED", err, ErrScanDeadline)
		} else {
			err = fmt.Errorf("REDACTED", ErrScanDeadline)
		}

		se.Tracer.Diagnose("REDACTED", "REDACTED", se)
		se.discardLinkage()
	}

	return
}

//
func (se *notaryTerminus) RecordSignal(msg privatekeyproto.Signal) (err error) {
	se.linkMutex.Lock()
	defer se.linkMutex.Unlock()

	if !se.isLinked() {
		return fmt.Errorf("REDACTED", ErrNoLinkage)
	}

	schemaRecorder := protoio.NewSeparatedRecorder(se.link)

	//
	limit := time.Now().Add(se.deadlineScanRecord)
	err = se.link.SetWriteDeadline(limit)
	if err != nil {
		return
	}

	_, err = schemaRecorder.RecordMessage(&msg)
	if _, ok := err.(deadlineFault); ok {
		if err != nil {
			err = fmt.Errorf("REDACTED", err, ErrRecordDeadline)
		} else {
			err = fmt.Errorf("REDACTED", ErrRecordDeadline)
		}
		se.discardLinkage()
	}

	return
}

func (se *notaryTerminus) isLinked() bool {
	return se.link != nil
}

func (se *notaryTerminus) discardLinkage() {
	if se.link != nil {
		if err := se.link.Close(); err != nil {
			se.Tracer.Fault("REDACTED", "REDACTED", err)
		}
		se.link = nil
	}
}
