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

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/protocolio"
	tmpfabric "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/p2p"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
)

const maximumPingPongPacketExtent = 1024 //

func generateVerifyModuleLinkage(link net.Conn) *ModuleLinkage {
	uponAccept := func(chnlUUID byte, signalOctets []byte) {
	}
	uponFailure := func(r any) {
	}
	c := generateModuleLinkageUsingReacts(link, uponAccept, uponFailure)
	c.AssignTracer(log.VerifyingTracer())
	return c
}

func generateModuleLinkageUsingReacts(
	link net.Conn,
	uponAccept func(chnlUUID byte, signalOctets []byte),
	uponFailure func(r any),
) *ModuleLinkage {
	cfg := FallbackModuleLinkSettings()
	cfg.PingDuration = 90 * time.Millisecond
	cfg.PongDeadline = 45 * time.Millisecond
	chnlDescriptions := []*ConduitDefinition{{ID: 0x01, Urgency: 1, TransmitStagingVolume: 1}}
	c := FreshModuleLinkageUsingSettings(link, chnlDescriptions, uponAccept, uponFailure, cfg)
	c.AssignTracer(log.VerifyingTracer())
	return c
}

func VerifyModuleLinkageTransmitPurgeHalt(t *testing.T) {
	node, customer := NetworkTube()
	defer node.Close()
	defer customer.Close()

	customerLink := generateVerifyModuleLinkage(customer)
	err := customerLink.Initiate()
	require.Nil(t, err)
	defer customerLink.Halt() //

	msg := []byte("REDACTED")
	assert.True(t, customerLink.Transmit(0x01, msg))

	signalMagnitude := 14

	//
	faultChnl := make(chan error)
	go func() {
		signalBYTE := make([]byte, signalMagnitude)
		_, err := node.Read(signalBYTE)
		if err != nil {
			t.Error(err)
			return
		}
		faultChnl <- err
	}()

	//
	customerLink.PurgeHalt()

	clock := time.NewTimer(3 * time.Second)
	select {
	case <-faultChnl:
	case <-clock.C:
		t.Error("REDACTED")
	}
}

func VerifyModuleLinkageTransmit(t *testing.T) {
	node, customer := NetworkTube()
	defer node.Close()
	defer customer.Close()

	multilink := generateVerifyModuleLinkage(customer)
	err := multilink.Initiate()
	require.Nil(t, err)
	defer multilink.Halt() //

	msg := []byte("REDACTED")
	assert.True(t, multilink.Transmit(0x01, msg))
	//
	//
	_, err = node.Read(make([]byte, len(msg)))
	if err != nil {
		t.Error(err)
	}
	assert.True(t, multilink.AbleTransmit(0x01))

	msg = []byte("REDACTED")
	assert.True(t, multilink.AttemptTransmit(0x01, msg))
	_, err = node.Read(make([]byte, len(msg)))
	if err != nil {
		t.Error(err)
	}

	assert.False(t, multilink.AbleTransmit(0x05), "REDACTED")
	assert.False(t, multilink.Transmit(0x05, []byte("REDACTED")), "REDACTED")
}

func VerifyModuleLinkageAccept(t *testing.T) {
	node, customer := NetworkTube()
	defer node.Close()
	defer customer.Close()

	acceptedChnl := make(chan []byte)
	faultsStream := make(chan any)
	uponAccept := func(chnlUUID byte, signalOctets []byte) {
		acceptedChnl <- signalOctets
	}
	uponFailure := func(r any) {
		faultsStream <- r
	}
	mconn1 := generateModuleLinkageUsingReacts(customer, uponAccept, uponFailure)
	err := mconn1.Initiate()
	require.Nil(t, err)
	defer mconn1.Halt() //

	mconn2 := generateVerifyModuleLinkage(node)
	err = mconn2.Initiate()
	require.Nil(t, err)
	defer mconn2.Halt() //

	msg := []byte("REDACTED")
	assert.True(t, mconn2.Transmit(0x01, msg))

	select {
	case acceptedOctets := <-acceptedChnl:
		assert.Equal(t, msg, acceptedOctets)
	case err := <-faultsStream:
		t.Fatalf("REDACTED", msg, err)
	case <-time.After(500 * time.Millisecond):
		t.Fatalf("REDACTED", msg)
	}
}

func VerifyModuleLinkageCondition(t *testing.T) {
	node, customer := NetworkTube()
	defer node.Close()
	defer customer.Close()

	multilink := generateVerifyModuleLinkage(customer)
	err := multilink.Initiate()
	require.Nil(t, err)
	defer multilink.Halt() //

	condition := multilink.Condition()
	assert.NotNil(t, condition)
	assert.Zero(t, condition.Conduits[0].TransmitStagingExtent)
}

func VerifyModuleLinkagePongDeadlineOutcomesInsideFailure(t *testing.T) {
	node, customer := net.Pipe()
	defer node.Close()
	defer customer.Close()

	acceptedChnl := make(chan []byte)
	faultsStream := make(chan any)
	uponAccept := func(chnlUUID byte, signalOctets []byte) {
		acceptedChnl <- signalOctets
	}
	uponFailure := func(r any) {
		faultsStream <- r
	}
	multilink := generateModuleLinkageUsingReacts(customer, uponAccept, uponFailure)
	err := multilink.Initiate()
	require.Nil(t, err)
	defer multilink.Halt() //

	daemonAttainedPing := make(chan struct{})
	go func() {
		//
		var pkt tmpfabric.Packet
		_, err := protocolio.FreshSeparatedFetcher(node, maximumPingPongPacketExtent).FetchSignal(&pkt)
		require.NoError(t, err)
		daemonAttainedPing <- struct{}{}
	}()
	<-daemonAttainedPing

	pongClockLapsed := multilink.settings.PongDeadline + 200*time.Millisecond
	select {
	case signalOctets := <-acceptedChnl:
		t.Fatalf("REDACTED", signalOctets)
	case err := <-faultsStream:
		assert.NotNil(t, err)
	case <-time.After(pongClockLapsed):
		t.Fatalf("REDACTED", pongClockLapsed)
	}
}

func VerifyModuleLinkageVariousPongsInsideTheCommencement(t *testing.T) {
	node, customer := net.Pipe()
	defer node.Close()
	defer customer.Close()

	acceptedChnl := make(chan []byte)
	faultsStream := make(chan any)
	uponAccept := func(chnlUUID byte, signalOctets []byte) {
		acceptedChnl <- signalOctets
	}
	uponFailure := func(r any) {
		faultsStream <- r
	}
	multilink := generateModuleLinkageUsingReacts(customer, uponAccept, uponFailure)
	err := multilink.Initiate()
	require.Nil(t, err)
	defer multilink.Halt() //

	//
	schemaPersistor := protocolio.FreshSeparatedPersistor(node)

	_, err = schemaPersistor.PersistSignal(shouldEnclosePacket(&tmpfabric.PacketPong{}))
	require.NoError(t, err)

	_, err = schemaPersistor.PersistSignal(shouldEnclosePacket(&tmpfabric.PacketPong{}))
	require.NoError(t, err)

	_, err = schemaPersistor.PersistSignal(shouldEnclosePacket(&tmpfabric.PacketPong{}))
	require.NoError(t, err)

	daemonAttainedPing := make(chan struct{})
	go func() {
		//
		var packet tmpfabric.Packet
		_, err := protocolio.FreshSeparatedFetcher(node, maximumPingPongPacketExtent).FetchSignal(&packet)
		require.NoError(t, err)
		daemonAttainedPing <- struct{}{}

		//
		_, err = schemaPersistor.PersistSignal(shouldEnclosePacket(&tmpfabric.PacketPong{}))
		require.NoError(t, err)
	}()
	<-daemonAttainedPing

	pongClockLapsed := multilink.settings.PongDeadline + 20*time.Millisecond
	select {
	case signalOctets := <-acceptedChnl:
		t.Fatalf("REDACTED", signalOctets)
	case err := <-faultsStream:
		t.Fatalf("REDACTED", err)
	case <-time.After(pongClockLapsed):
		assert.True(t, multilink.EqualsActive())
	}
}

func VerifyModuleLinkageVariousPings(t *testing.T) {
	node, customer := net.Pipe()
	defer node.Close()
	defer customer.Close()

	acceptedChnl := make(chan []byte)
	faultsStream := make(chan any)
	uponAccept := func(chnlUUID byte, signalOctets []byte) {
		acceptedChnl <- signalOctets
	}
	uponFailure := func(r any) {
		faultsStream <- r
	}
	multilink := generateModuleLinkageUsingReacts(customer, uponAccept, uponFailure)
	err := multilink.Initiate()
	require.Nil(t, err)
	defer multilink.Halt() //

	//
	//
	schemaFetcher := protocolio.FreshSeparatedFetcher(node, maximumPingPongPacketExtent)
	schemaPersistor := protocolio.FreshSeparatedPersistor(node)
	var pkt tmpfabric.Packet

	_, err = schemaPersistor.PersistSignal(shouldEnclosePacket(&tmpfabric.PacketPing{}))
	require.NoError(t, err)

	_, err = schemaFetcher.FetchSignal(&pkt)
	require.NoError(t, err)

	_, err = schemaPersistor.PersistSignal(shouldEnclosePacket(&tmpfabric.PacketPing{}))
	require.NoError(t, err)

	_, err = schemaFetcher.FetchSignal(&pkt)
	require.NoError(t, err)

	_, err = schemaPersistor.PersistSignal(shouldEnclosePacket(&tmpfabric.PacketPing{}))
	require.NoError(t, err)

	_, err = schemaFetcher.FetchSignal(&pkt)
	require.NoError(t, err)

	assert.True(t, multilink.EqualsActive())
}

func VerifyModuleLinkagePingPongs(t *testing.T) {
	//
	defer leaktest.CheckTimeout(t, 10*time.Second)()

	node, customer := net.Pipe()

	defer node.Close()
	defer customer.Close()

	acceptedChnl := make(chan []byte)
	faultsStream := make(chan any)
	uponAccept := func(chnlUUID byte, signalOctets []byte) {
		acceptedChnl <- signalOctets
	}
	uponFailure := func(r any) {
		faultsStream <- r
	}
	multilink := generateModuleLinkageUsingReacts(customer, uponAccept, uponFailure)
	err := multilink.Initiate()
	require.Nil(t, err)
	defer multilink.Halt() //

	daemonAttainedPing := make(chan struct{})
	go func() {
		schemaFetcher := protocolio.FreshSeparatedFetcher(node, maximumPingPongPacketExtent)
		schemaPersistor := protocolio.FreshSeparatedPersistor(node)
		var pkt tmpfabric.PacketPing

		//
		_, err = schemaFetcher.FetchSignal(&pkt)
		require.NoError(t, err)
		daemonAttainedPing <- struct{}{}

		//
		_, err = schemaPersistor.PersistSignal(shouldEnclosePacket(&tmpfabric.PacketPong{}))
		require.NoError(t, err)

		time.Sleep(multilink.settings.PingDuration)

		//
		_, err = schemaFetcher.FetchSignal(&pkt)
		require.NoError(t, err)
		daemonAttainedPing <- struct{}{}

		//
		_, err = schemaPersistor.PersistSignal(shouldEnclosePacket(&tmpfabric.PacketPong{}))
		require.NoError(t, err)
	}()
	<-daemonAttainedPing
	<-daemonAttainedPing

	pongClockLapsed := (multilink.settings.PongDeadline + 20*time.Millisecond) * 2
	select {
	case signalOctets := <-acceptedChnl:
		t.Fatalf("REDACTED", signalOctets)
	case err := <-faultsStream:
		t.Fatalf("REDACTED", err)
	case <-time.After(2 * pongClockLapsed):
		assert.True(t, multilink.EqualsActive())
	}
}

func VerifyModuleLinkageHaltsAlsoYieldsFailure(t *testing.T) {
	node, customer := NetworkTube()
	defer node.Close()
	defer customer.Close()

	acceptedChnl := make(chan []byte)
	faultsStream := make(chan any)
	uponAccept := func(chnlUUID byte, signalOctets []byte) {
		acceptedChnl <- signalOctets
	}
	uponFailure := func(r any) {
		faultsStream <- r
	}
	multilink := generateModuleLinkageUsingReacts(customer, uponAccept, uponFailure)
	err := multilink.Initiate()
	require.Nil(t, err)
	defer multilink.Halt() //

	if err := customer.Close(); err != nil {
		t.Error(err)
	}

	select {
	case acceptedOctets := <-acceptedChnl:
		t.Fatalf("REDACTED", acceptedOctets)
	case err := <-faultsStream:
		assert.NotNil(t, err)
		assert.False(t, multilink.EqualsActive())
	case <-time.After(500 * time.Millisecond):
		t.Fatal("REDACTED")
	}
}

func freshCustomerAlsoDaemonLinksForeachFetchFaults(t *testing.T, chnlUponFault chan struct{}) (*ModuleLinkage, *ModuleLinkage) {
	node, customer := NetworkTube()

	uponAccept := func(chnlUUID byte, signalOctets []byte) {}
	uponFailure := func(r any) {}

	//
	chnlDescriptions := []*ConduitDefinition{
		{ID: 0x01, Urgency: 1, TransmitStagingVolume: 1},
		{ID: 0x02, Urgency: 1, TransmitStagingVolume: 1},
	}
	multilinkCustomer := FreshModuleLinkage(customer, chnlDescriptions, uponAccept, uponFailure)
	multilinkCustomer.AssignTracer(log.VerifyingTracer().Using("REDACTED", "REDACTED"))
	err := multilinkCustomer.Initiate()
	require.Nil(t, err)

	//
	//
	daemonTracer := log.VerifyingTracer().Using("REDACTED", "REDACTED")
	uponFailure = func(r any) {
		chnlUponFault <- struct{}{}
	}
	multilinkDaemon := generateModuleLinkageUsingReacts(node, uponAccept, uponFailure)
	multilinkDaemon.AssignTracer(daemonTracer)
	err = multilinkDaemon.Initiate()
	require.Nil(t, err)
	return multilinkCustomer, multilinkDaemon
}

func anticipateTransmit(ch chan struct{}) bool {
	subsequent := time.After(time.Second * 5)
	select {
	case <-ch:
		return true
	case <-subsequent:
		return false
	}
}

func VerifyModuleLinkageFetchFailureFlawedSerialization(t *testing.T) {
	chnlUponFault := make(chan struct{})
	multilinkCustomer, multilinkDaemon := freshCustomerAlsoDaemonLinksForeachFetchFaults(t, chnlUponFault)

	customer := multilinkCustomer.link

	//
	_, err := customer.Write([]byte{1, 2, 3, 4, 5})
	require.NoError(t, err)
	assert.True(t, anticipateTransmit(chnlUponFault), "REDACTED")

	t.Cleanup(func() {
		if err := multilinkCustomer.Halt(); err != nil {
			t.Log(err)
		}
	})

	t.Cleanup(func() {
		if err := multilinkDaemon.Halt(); err != nil {
			t.Log(err)
		}
	})
}

func VerifyModuleLinkageFetchFailureUnfamiliarConduit(t *testing.T) {
	chnlUponFault := make(chan struct{})
	multilinkCustomer, multilinkDaemon := freshCustomerAlsoDaemonLinksForeachFetchFaults(t, chnlUponFault)

	msg := []byte("REDACTED")

	//
	assert.False(t, multilinkCustomer.Transmit(0x03, msg))

	//
	//
	assert.True(t, multilinkCustomer.Transmit(0x02, msg))
	assert.True(t, anticipateTransmit(chnlUponFault), "REDACTED")

	t.Cleanup(func() {
		if err := multilinkCustomer.Halt(); err != nil {
			t.Log(err)
		}
	})

	t.Cleanup(func() {
		if err := multilinkDaemon.Halt(); err != nil {
			t.Log(err)
		}
	})
}

func VerifyModuleLinkageFetchFailureExtendedArtifact(t *testing.T) {
	chnlUponFault := make(chan struct{})
	chnlUponAcceptmsg := make(chan struct{})

	multilinkCustomer, multilinkDaemon := freshCustomerAlsoDaemonLinksForeachFetchFaults(t, chnlUponFault)
	defer multilinkCustomer.Halt() //
	defer multilinkDaemon.Halt() //

	multilinkDaemon.uponAccept = func(chnlUUID byte, signalOctets []byte) {
		chnlUponAcceptmsg <- struct{}{}
	}

	customer := multilinkCustomer.link
	schemaPersistor := protocolio.FreshSeparatedPersistor(customer)

	//
	packet := tmpfabric.PacketSignal{
		ConduitUUID: 0x01,
		EOF:       true,
		Data:      make([]byte, multilinkCustomer.settings.MaximumPacketSignalWorkloadExtent),
	}

	_, err := schemaPersistor.PersistSignal(shouldEnclosePacket(&packet))
	require.NoError(t, err)
	assert.True(t, anticipateTransmit(chnlUponAcceptmsg), "REDACTED")

	//
	packet = tmpfabric.PacketSignal{
		ConduitUUID: 0x01,
		EOF:       true,
		Data:      make([]byte, multilinkCustomer.settings.MaximumPacketSignalWorkloadExtent+100),
	}

	_, err = schemaPersistor.PersistSignal(shouldEnclosePacket(&packet))
	require.Error(t, err)
	assert.True(t, anticipateTransmit(chnlUponFault), "REDACTED")
}

func VerifyModuleLinkageFetchFailureUnfamiliarSignalKind(t *testing.T) {
	chnlUponFault := make(chan struct{})
	multilinkCustomer, multilinkDaemon := freshCustomerAlsoDaemonLinksForeachFetchFaults(t, chnlUponFault)
	defer multilinkCustomer.Halt() //
	defer multilinkDaemon.Halt() //

	//
	_, err := protocolio.FreshSeparatedPersistor(multilinkCustomer.link).PersistSignal(&kinds.Heading{SuccessionUUID: "REDACTED"})
	require.NoError(t, err)
	assert.True(t, anticipateTransmit(chnlUponFault), "REDACTED")
}

func VerifyModuleLinkageAttemptTransmit(t *testing.T) {
	node, customer := NetworkTube()
	defer node.Close()
	defer customer.Close()

	multilink := generateVerifyModuleLinkage(customer)
	err := multilink.Initiate()
	require.Nil(t, err)
	defer multilink.Halt() //

	msg := []byte("REDACTED")
	outcomeChnl := make(chan string, 2)
	assert.True(t, multilink.AttemptTransmit(0x01, msg))
	_, err = node.Read(make([]byte, len(msg)))
	require.NoError(t, err)
	assert.True(t, multilink.AbleTransmit(0x01))
	assert.True(t, multilink.AttemptTransmit(0x01, msg))
	assert.False(t, multilink.AbleTransmit(0x01))
	go func() {
		multilink.AttemptTransmit(0x01, msg)
		outcomeChnl <- "REDACTED"
	}()
	assert.False(t, multilink.AbleTransmit(0x01))
	assert.False(t, multilink.AttemptTransmit(0x01, msg))
	assert.Equal(t, "REDACTED", <-outcomeChnl)
}

//
func VerifyLinkArrays(t *testing.T) {
	verifyScenarios := []struct {
		verifyAlias string
		msg      proto.Message
		expirationOctets string
	}{
		{"REDACTED", &tmpfabric.PacketPing{}, "REDACTED"},
		{"REDACTED", &tmpfabric.PacketPong{}, "REDACTED"},
		{"REDACTED", &tmpfabric.PacketSignal{ConduitUUID: 1, EOF: false, Data: []byte("REDACTED")}, "REDACTED"},
	}

	for _, tc := range verifyScenarios {

		pm := shouldEnclosePacket(tc.msg)
		bz, err := pm.Serialize()
		require.NoError(t, err, tc.verifyAlias)

		require.Equal(t, tc.expirationOctets, hex.EncodeToString(bz), tc.verifyAlias)
	}
}

func VerifyModuleLinkageConduitOverrun(t *testing.T) {
	chnlUponFault := make(chan struct{})
	chnlUponAcceptmsg := make(chan struct{})

	multilinkCustomer, multilinkDaemon := freshCustomerAlsoDaemonLinksForeachFetchFaults(t, chnlUponFault)
	t.Cleanup(haltEvery(t, multilinkCustomer, multilinkDaemon))

	multilinkDaemon.uponAccept = func(chnlUUID byte, signalOctets []byte) {
		chnlUponAcceptmsg <- struct{}{}
	}

	customer := multilinkCustomer.link
	schemaPersistor := protocolio.FreshSeparatedPersistor(customer)

	packet := tmpfabric.PacketSignal{
		ConduitUUID: 0x01,
		EOF:       true,
		Data:      []byte("REDACTED"),
	}
	_, err := schemaPersistor.PersistSignal(shouldEnclosePacket(&packet))
	require.NoError(t, err)
	assert.True(t, anticipateTransmit(chnlUponAcceptmsg))

	packet.ConduitUUID = int32(1025)
	_, err = schemaPersistor.PersistSignal(shouldEnclosePacket(&packet))
	require.NoError(t, err)
	assert.False(t, anticipateTransmit(chnlUponAcceptmsg))
}

type terminator interface {
	Halt() error
}

func haltEvery(t *testing.T, terminators ...terminator) func() {
	return func() {
		for _, s := range terminators {
			if err := s.Halt(); err != nil {
				t.Log(err)
			}
		}
	}
}
