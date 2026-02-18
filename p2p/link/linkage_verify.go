package link

import (
	"encoding/hex"
	"net"
	"testing"
	"time"

	"github.com/cosmos/gogoproto/proto"
	"github.com/fortytw2/leaktest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/utils/log"
	"github.com/valkyrieworks/utils/protoio"
	tmp2p "github.com/valkyrieworks/schema/consensuscore/p2p"
	"github.com/valkyrieworks/schema/consensuscore/kinds"
)

const maximumPingPongPackageVolume = 1024 //

func instantiateVerifyMLinkage(link net.Conn) *MLinkage {
	onAccept := func(chanUID byte, messageOctets []byte) {
	}
	onFault := func(r any) {
	}
	c := instantiateMLinkageWithResponses(link, onAccept, onFault)
	c.AssignTracer(log.VerifyingTracer())
	return c
}

func instantiateMLinkageWithResponses(
	link net.Conn,
	onAccept func(chanUID byte, messageOctets []byte),
	onFault func(r any),
) *MLinkage {
	cfg := StandardMLinkSettings()
	cfg.PingCadence = 90 * time.Millisecond
	cfg.PongDeadline = 45 * time.Millisecond
	chanTraits := []*StreamDefinition{{ID: 0x01, Urgency: 1, TransmitBufferVolume: 1}}
	c := NewMLinkageWithSettings(link, chanTraits, onAccept, onFault, cfg)
	c.AssignTracer(log.VerifyingTracer())
	return c
}

func VerifyMLinkageTransmitPurgeHalt(t *testing.T) {
	host, customer := NetPipe()
	defer host.Close()
	defer customer.Close()

	customerLink := instantiateVerifyMLinkage(customer)
	err := customerLink.Begin()
	require.Nil(t, err)
	defer customerLink.Halt() //

	msg := []byte("REDACTED")
	assert.True(t, customerLink.Transmit(0x01, msg))

	messageExtent := 14

	//
	errChan := make(chan error)
	go func() {
		messageBYTE := make([]byte, messageExtent)
		_, err := host.Read(messageBYTE)
		if err != nil {
			t.Error(err)
			return
		}
		errChan <- err
	}()

	//
	customerLink.PurgeHalt()

	clock := time.NewTimer(3 * time.Second)
	select {
	case <-errChan:
	case <-clock.C:
		t.Error("REDACTED")
	}
}

func VerifyMLinkageTransmit(t *testing.T) {
	host, customer := NetPipe()
	defer host.Close()
	defer customer.Close()

	mconn := instantiateVerifyMLinkage(customer)
	err := mconn.Begin()
	require.Nil(t, err)
	defer mconn.Halt() //

	msg := []byte("REDACTED")
	assert.True(t, mconn.Transmit(0x01, msg))
	//
	//
	_, err = host.Read(make([]byte, len(msg)))
	if err != nil {
		t.Error(err)
	}
	assert.True(t, mconn.MayTransmit(0x01))

	msg = []byte("REDACTED")
	assert.True(t, mconn.AttemptTransmit(0x01, msg))
	_, err = host.Read(make([]byte, len(msg)))
	if err != nil {
		t.Error(err)
	}

	assert.False(t, mconn.MayTransmit(0x05), "REDACTED")
	assert.False(t, mconn.Transmit(0x05, []byte("REDACTED")), "REDACTED")
}

func VerifyMLinkageAccept(t *testing.T) {
	host, customer := NetPipe()
	defer host.Close()
	defer customer.Close()

	acceptedChan := make(chan []byte)
	faultsChan := make(chan any)
	onAccept := func(chanUID byte, messageOctets []byte) {
		acceptedChan <- messageOctets
	}
	onFault := func(r any) {
		faultsChan <- r
	}
	mconn1 := instantiateMLinkageWithResponses(customer, onAccept, onFault)
	err := mconn1.Begin()
	require.Nil(t, err)
	defer mconn1.Halt() //

	mconn2 := instantiateVerifyMLinkage(host)
	err = mconn2.Begin()
	require.Nil(t, err)
	defer mconn2.Halt() //

	msg := []byte("REDACTED")
	assert.True(t, mconn2.Transmit(0x01, msg))

	select {
	case acceptedOctets := <-acceptedChan:
		assert.Equal(t, msg, acceptedOctets)
	case err := <-faultsChan:
		t.Fatalf("REDACTED", msg, err)
	case <-time.After(500 * time.Millisecond):
		t.Fatalf("REDACTED", msg)
	}
}

func VerifyMLinkageState(t *testing.T) {
	host, customer := NetPipe()
	defer host.Close()
	defer customer.Close()

	mconn := instantiateVerifyMLinkage(customer)
	err := mconn.Begin()
	require.Nil(t, err)
	defer mconn.Halt() //

	state := mconn.Status()
	assert.NotNil(t, state)
	assert.Zero(t, state.Streams[0].TransmitBufferVolume)
}

func VerifyMLinkagePongDeadlineOutcomesInFault(t *testing.T) {
	host, customer := net.Pipe()
	defer host.Close()
	defer customer.Close()

	acceptedChan := make(chan []byte)
	faultsChan := make(chan any)
	onAccept := func(chanUID byte, messageOctets []byte) {
		acceptedChan <- messageOctets
	}
	onFault := func(r any) {
		faultsChan <- r
	}
	mconn := instantiateMLinkageWithResponses(customer, onAccept, onFault)
	err := mconn.Begin()
	require.Nil(t, err)
	defer mconn.Halt() //

	hostAcquiredPing := make(chan struct{})
	go func() {
		//
		var pkt tmp2p.Package
		_, err := protoio.NewSeparatedScanner(host, maximumPingPongPackageVolume).ScanMessage(&pkt)
		require.NoError(t, err)
		hostAcquiredPing <- struct{}{}
	}()
	<-hostAcquiredPing

	pongClockLapsed := mconn.settings.PongDeadline + 200*time.Millisecond
	select {
	case messageOctets := <-acceptedChan:
		t.Fatalf("REDACTED", messageOctets)
	case err := <-faultsChan:
		assert.NotNil(t, err)
	case <-time.After(pongClockLapsed):
		t.Fatalf("REDACTED", pongClockLapsed)
	}
}

func VerifyMLinkageVariedPongsInTheCommencement(t *testing.T) {
	host, customer := net.Pipe()
	defer host.Close()
	defer customer.Close()

	acceptedChan := make(chan []byte)
	faultsChan := make(chan any)
	onAccept := func(chanUID byte, messageOctets []byte) {
		acceptedChan <- messageOctets
	}
	onFault := func(r any) {
		faultsChan <- r
	}
	mconn := instantiateMLinkageWithResponses(customer, onAccept, onFault)
	err := mconn.Begin()
	require.Nil(t, err)
	defer mconn.Halt() //

	//
	schemaRecorder := protoio.NewSeparatedRecorder(host)

	_, err = schemaRecorder.RecordMessage(shouldEnclosePackage(&tmp2p.PackagePong{}))
	require.NoError(t, err)

	_, err = schemaRecorder.RecordMessage(shouldEnclosePackage(&tmp2p.PackagePong{}))
	require.NoError(t, err)

	_, err = schemaRecorder.RecordMessage(shouldEnclosePackage(&tmp2p.PackagePong{}))
	require.NoError(t, err)

	hostAcquiredPing := make(chan struct{})
	go func() {
		//
		var package tmp2p.Package
		_, err := protoio.NewSeparatedScanner(host, maximumPingPongPackageVolume).ScanMessage(&package)
		require.NoError(t, err)
		hostAcquiredPing <- struct{}{}

		//
		_, err = schemaRecorder.RecordMessage(shouldEnclosePackage(&tmp2p.PackagePong{}))
		require.NoError(t, err)
	}()
	<-hostAcquiredPing

	pongClockLapsed := mconn.settings.PongDeadline + 20*time.Millisecond
	select {
	case messageOctets := <-acceptedChan:
		t.Fatalf("REDACTED", messageOctets)
	case err := <-faultsChan:
		t.Fatalf("REDACTED", err)
	case <-time.After(pongClockLapsed):
		assert.True(t, mconn.IsActive())
	}
}

func VerifyMLinkageVariedPings(t *testing.T) {
	host, customer := net.Pipe()
	defer host.Close()
	defer customer.Close()

	acceptedChan := make(chan []byte)
	faultsChan := make(chan any)
	onAccept := func(chanUID byte, messageOctets []byte) {
		acceptedChan <- messageOctets
	}
	onFault := func(r any) {
		faultsChan <- r
	}
	mconn := instantiateMLinkageWithResponses(customer, onAccept, onFault)
	err := mconn.Begin()
	require.Nil(t, err)
	defer mconn.Halt() //

	//
	//
	schemaScanner := protoio.NewSeparatedScanner(host, maximumPingPongPackageVolume)
	schemaRecorder := protoio.NewSeparatedRecorder(host)
	var pkt tmp2p.Package

	_, err = schemaRecorder.RecordMessage(shouldEnclosePackage(&tmp2p.PackagePing{}))
	require.NoError(t, err)

	_, err = schemaScanner.ScanMessage(&pkt)
	require.NoError(t, err)

	_, err = schemaRecorder.RecordMessage(shouldEnclosePackage(&tmp2p.PackagePing{}))
	require.NoError(t, err)

	_, err = schemaScanner.ScanMessage(&pkt)
	require.NoError(t, err)

	_, err = schemaRecorder.RecordMessage(shouldEnclosePackage(&tmp2p.PackagePing{}))
	require.NoError(t, err)

	_, err = schemaScanner.ScanMessage(&pkt)
	require.NoError(t, err)

	assert.True(t, mconn.IsActive())
}

func VerifyMLinkagePingPongs(t *testing.T) {
	//
	defer leaktest.CheckTimeout(t, 10*time.Second)()

	host, customer := net.Pipe()

	defer host.Close()
	defer customer.Close()

	acceptedChan := make(chan []byte)
	faultsChan := make(chan any)
	onAccept := func(chanUID byte, messageOctets []byte) {
		acceptedChan <- messageOctets
	}
	onFault := func(r any) {
		faultsChan <- r
	}
	mconn := instantiateMLinkageWithResponses(customer, onAccept, onFault)
	err := mconn.Begin()
	require.Nil(t, err)
	defer mconn.Halt() //

	hostAcquiredPing := make(chan struct{})
	go func() {
		schemaScanner := protoio.NewSeparatedScanner(host, maximumPingPongPackageVolume)
		schemaRecorder := protoio.NewSeparatedRecorder(host)
		var pkt tmp2p.PackagePing

		//
		_, err = schemaScanner.ScanMessage(&pkt)
		require.NoError(t, err)
		hostAcquiredPing <- struct{}{}

		//
		_, err = schemaRecorder.RecordMessage(shouldEnclosePackage(&tmp2p.PackagePong{}))
		require.NoError(t, err)

		time.Sleep(mconn.settings.PingCadence)

		//
		_, err = schemaScanner.ScanMessage(&pkt)
		require.NoError(t, err)
		hostAcquiredPing <- struct{}{}

		//
		_, err = schemaRecorder.RecordMessage(shouldEnclosePackage(&tmp2p.PackagePong{}))
		require.NoError(t, err)
	}()
	<-hostAcquiredPing
	<-hostAcquiredPing

	pongClockLapsed := (mconn.settings.PongDeadline + 20*time.Millisecond) * 2
	select {
	case messageOctets := <-acceptedChan:
		t.Fatalf("REDACTED", messageOctets)
	case err := <-faultsChan:
		t.Fatalf("REDACTED", err)
	case <-time.After(2 * pongClockLapsed):
		assert.True(t, mconn.IsActive())
	}
}

func VerifyMLinkageHaltsAndYieldsFault(t *testing.T) {
	host, customer := NetPipe()
	defer host.Close()
	defer customer.Close()

	acceptedChan := make(chan []byte)
	faultsChan := make(chan any)
	onAccept := func(chanUID byte, messageOctets []byte) {
		acceptedChan <- messageOctets
	}
	onFault := func(r any) {
		faultsChan <- r
	}
	mconn := instantiateMLinkageWithResponses(customer, onAccept, onFault)
	err := mconn.Begin()
	require.Nil(t, err)
	defer mconn.Halt() //

	if err := customer.Close(); err != nil {
		t.Error(err)
	}

	select {
	case acceptedOctets := <-acceptedChan:
		t.Fatalf("REDACTED", acceptedOctets)
	case err := <-faultsChan:
		assert.NotNil(t, err)
		assert.False(t, mconn.IsActive())
	case <-time.After(500 * time.Millisecond):
		t.Fatal("REDACTED")
	}
}

func newCustomerAndHostLinksForReadFaults(t *testing.T, chanOnErr chan struct{}) (*MLinkage, *MLinkage) {
	host, customer := NetPipe()

	onAccept := func(chanUID byte, messageOctets []byte) {}
	onFault := func(r any) {}

	//
	chanTraits := []*StreamDefinition{
		{ID: 0x01, Urgency: 1, TransmitBufferVolume: 1},
		{ID: 0x02, Urgency: 1, TransmitBufferVolume: 1},
	}
	mconnCustomer := NewMLinkage(customer, chanTraits, onAccept, onFault)
	mconnCustomer.AssignTracer(log.VerifyingTracer().With("REDACTED", "REDACTED"))
	err := mconnCustomer.Begin()
	require.Nil(t, err)

	//
	//
	hostTracer := log.VerifyingTracer().With("REDACTED", "REDACTED")
	onFault = func(r any) {
		chanOnErr <- struct{}{}
	}
	mconnHost := instantiateMLinkageWithResponses(host, onAccept, onFault)
	mconnHost.AssignTracer(hostTracer)
	err = mconnHost.Begin()
	require.Nil(t, err)
	return mconnCustomer, mconnHost
}

func anticipateTransmit(ch chan struct{}) bool {
	after := time.After(time.Second * 5)
	select {
	case <-ch:
		return true
	case <-after:
		return false
	}
}

func VerifyMLinkageReadFaultFlawedCodec(t *testing.T) {
	chanOnErr := make(chan struct{})
	mconnCustomer, mconnHost := newCustomerAndHostLinksForReadFaults(t, chanOnErr)

	customer := mconnCustomer.link

	//
	_, err := customer.Write([]byte{1, 2, 3, 4, 5})
	require.NoError(t, err)
	assert.True(t, anticipateTransmit(chanOnErr), "REDACTED")

	t.Cleanup(func() {
		if err := mconnCustomer.Halt(); err != nil {
			t.Log(err)
		}
	})

	t.Cleanup(func() {
		if err := mconnHost.Halt(); err != nil {
			t.Log(err)
		}
	})
}

func VerifyMLinkageReadFaultUnclearConduit(t *testing.T) {
	chanOnErr := make(chan struct{})
	mconnCustomer, mconnHost := newCustomerAndHostLinksForReadFaults(t, chanOnErr)

	msg := []byte("REDACTED")

	//
	assert.False(t, mconnCustomer.Transmit(0x03, msg))

	//
	//
	assert.True(t, mconnCustomer.Transmit(0x02, msg))
	assert.True(t, anticipateTransmit(chanOnErr), "REDACTED")

	t.Cleanup(func() {
		if err := mconnCustomer.Halt(); err != nil {
			t.Log(err)
		}
	})

	t.Cleanup(func() {
		if err := mconnHost.Halt(); err != nil {
			t.Log(err)
		}
	})
}

func VerifyMLinkageReadFaultLengthySignal(t *testing.T) {
	chanOnErr := make(chan struct{})
	chanOnRcv := make(chan struct{})

	mconnCustomer, mconnHost := newCustomerAndHostLinksForReadFaults(t, chanOnErr)
	defer mconnCustomer.Halt() //
	defer mconnHost.Halt() //

	mconnHost.onAccept = func(chanUID byte, messageOctets []byte) {
		chanOnRcv <- struct{}{}
	}

	customer := mconnCustomer.link
	schemaRecorder := protoio.NewSeparatedRecorder(customer)

	//
	package := tmp2p.PackageMessage{
		StreamUID: 0x01,
		EOF:       true,
		Data:      make([]byte, mconnCustomer.settings.MaximumPackageMessageShipmentVolume),
	}

	_, err := schemaRecorder.RecordMessage(shouldEnclosePackage(&package))
	require.NoError(t, err)
	assert.True(t, anticipateTransmit(chanOnRcv), "REDACTED")

	//
	package = tmp2p.PackageMessage{
		StreamUID: 0x01,
		EOF:       true,
		Data:      make([]byte, mconnCustomer.settings.MaximumPackageMessageShipmentVolume+100),
	}

	_, err = schemaRecorder.RecordMessage(shouldEnclosePackage(&package))
	require.Error(t, err)
	assert.True(t, anticipateTransmit(chanOnErr), "REDACTED")
}

func VerifyMLinkageReadFaultUnclearMessageKind(t *testing.T) {
	chanOnErr := make(chan struct{})
	mconnCustomer, mconnHost := newCustomerAndHostLinksForReadFaults(t, chanOnErr)
	defer mconnCustomer.Halt() //
	defer mconnHost.Halt() //

	//
	_, err := protoio.NewSeparatedRecorder(mconnCustomer.link).RecordMessage(&kinds.Heading{LedgerUID: "REDACTED"})
	require.NoError(t, err)
	assert.True(t, anticipateTransmit(chanOnErr), "REDACTED")
}

func VerifyMLinkageAttemptTransmit(t *testing.T) {
	host, customer := NetPipe()
	defer host.Close()
	defer customer.Close()

	mconn := instantiateVerifyMLinkage(customer)
	err := mconn.Begin()
	require.Nil(t, err)
	defer mconn.Halt() //

	msg := []byte("REDACTED")
	outcomeChan := make(chan string, 2)
	assert.True(t, mconn.AttemptTransmit(0x01, msg))
	_, err = host.Read(make([]byte, len(msg)))
	require.NoError(t, err)
	assert.True(t, mconn.MayTransmit(0x01))
	assert.True(t, mconn.AttemptTransmit(0x01, msg))
	assert.False(t, mconn.MayTransmit(0x01))
	go func() {
		mconn.AttemptTransmit(0x01, msg)
		outcomeChan <- "REDACTED"
	}()
	assert.False(t, mconn.MayTransmit(0x01))
	assert.False(t, mconn.AttemptTransmit(0x01, msg))
	assert.Equal(t, "REDACTED", <-outcomeChan)
}

//
func VerifyLinkArrays(t *testing.T) {
	verifyScenarios := []struct {
		verifyLabel string
		msg      proto.Message
		expirationOctets string
	}{
		{"REDACTED", &tmp2p.PackagePing{}, "REDACTED"},
		{"REDACTED", &tmp2p.PackagePong{}, "REDACTED"},
		{"REDACTED", &tmp2p.PackageMessage{StreamUID: 1, EOF: false, Data: []byte("REDACTED")}, "REDACTED"},
	}

	for _, tc := range verifyScenarios {

		pm := shouldEnclosePackage(tc.msg)
		bz, err := pm.Serialize()
		require.NoError(t, err, tc.verifyLabel)

		require.Equal(t, tc.expirationOctets, hex.EncodeToString(bz), tc.verifyLabel)
	}
}

func VerifyMLinkageConduitOverload(t *testing.T) {
	chanOnErr := make(chan struct{})
	chanOnRcv := make(chan struct{})

	mconnCustomer, mconnHost := newCustomerAndHostLinksForReadFaults(t, chanOnErr)
	t.Cleanup(haltAll(t, mconnCustomer, mconnHost))

	mconnHost.onAccept = func(chanUID byte, messageOctets []byte) {
		chanOnRcv <- struct{}{}
	}

	customer := mconnCustomer.link
	schemaRecorder := protoio.NewSeparatedRecorder(customer)

	package := tmp2p.PackageMessage{
		StreamUID: 0x01,
		EOF:       true,
		Data:      []byte("REDACTED"),
	}
	_, err := schemaRecorder.RecordMessage(shouldEnclosePackage(&package))
	require.NoError(t, err)
	assert.True(t, anticipateTransmit(chanOnRcv))

	package.StreamUID = int32(1025)
	_, err = schemaRecorder.RecordMessage(shouldEnclosePackage(&package))
	require.NoError(t, err)
	assert.False(t, anticipateTransmit(chanOnRcv))
}

type terminator interface {
	Halt() error
}

func haltAll(t *testing.T, terminators ...terminator) func() {
	return func() {
		for _, s := range terminators {
			if err := s.Halt(); err != nil {
				t.Log(err)
			}
		}
	}
}
