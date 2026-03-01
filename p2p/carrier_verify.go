package p2p

import (
	"fmt"
	"math/rand"
	"net"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/edwards25519"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/protocolio"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p/link"
	tmpfabric "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/p2p"
)

var fallbackPeerAlias = "REDACTED"

func blankPeerDetails() PeerDetails {
	return FallbackPeerDetails{}
}

//
//
//
func freshMultiplexCarrier(
	peerDetails PeerDetails,
	peerToken PeerToken,
) *MultiplexCarrier {
	return FreshMultiplexCarrier(
		peerDetails, peerToken, link.FallbackModuleLinkSettings(),
	)
}

func VerifyCarrierMultiplexLinkRefine(t *testing.T) {
	mt := freshMultiplexCarrier(
		blankPeerDetails(),
		PeerToken{
			PrivateToken: edwards25519.ProducePrivateToken(),
		},
	)
	id := mt.peerToken.ID()

	MultiplexCarrierLinkCriteria(
		func(_ LinkAssign, _ net.Conn, _ []net.IP) error { return nil },
		func(_ LinkAssign, _ net.Conn, _ []net.IP) error { return nil },
		func(_ LinkAssign, _ net.Conn, _ []net.IP) error {
			return fmt.Errorf("REDACTED")
		},
	)(mt)

	location, err := FreshNetworkLocatorText(UUIDLocationText(id, "REDACTED"))
	if err != nil {
		t.Fatal(err)
	}

	if err := mt.Overhear(*location); err != nil {
		t.Fatal(err)
	}

	faultchnl := make(chan error)

	go func() {
		location := FreshNetworkLocator(id, mt.observer.Addr())

		_, err := location.Call()
		if err != nil {
			faultchnl <- err
			return
		}

		close(faultchnl)
	}()

	if err := <-faultchnl; err != nil {
		t.Errorf("REDACTED", err)
	}

	_, err = mt.Embrace(nodeSettings{})
	if e, ok := err.(FaultDeclined); ok {
		if !e.EqualsScreened() {
			t.Errorf("REDACTED", err)
		}
	} else {
		t.Errorf("REDACTED", err)
	}
}

func VerifyCarrierMultiplexLinkRefineDeadline(t *testing.T) {
	mt := freshMultiplexCarrier(
		blankPeerDetails(),
		PeerToken{
			PrivateToken: edwards25519.ProducePrivateToken(),
		},
	)
	id := mt.peerToken.ID()

	MultiplexCarrierRefineDeadline(5 * time.Millisecond)(mt)
	MultiplexCarrierLinkCriteria(
		func(_ LinkAssign, _ net.Conn, _ []net.IP) error {
			time.Sleep(1 * time.Second)
			return nil
		},
	)(mt)

	location, err := FreshNetworkLocatorText(UUIDLocationText(id, "REDACTED"))
	if err != nil {
		t.Fatal(err)
	}

	if err := mt.Overhear(*location); err != nil {
		t.Fatal(err)
	}

	faultchnl := make(chan error)
	go func() {
		location := FreshNetworkLocator(id, mt.observer.Addr())

		_, err := location.Call()
		if err != nil {
			faultchnl <- err
			return
		}

		close(faultchnl)
	}()

	if err := <-faultchnl; err != nil {
		t.Errorf("REDACTED", err)
	}

	_, err = mt.Embrace(nodeSettings{})
	if _, ok := err.(FaultRefineDeadline); !ok {
		t.Errorf("REDACTED", err)
	}
}

func VerifyCarrierMultiplexMaximumArrivingLinkages(t *testing.T) {
	pv := edwards25519.ProducePrivateToken()
	id := PublicTokenTowardUUID(pv.PublicToken())
	mt := freshMultiplexCarrier(
		verifyPeerDetails(
			id, "REDACTED",
		),
		PeerToken{
			PrivateToken: pv,
		},
	)

	MultiplexCarrierMaximumArrivingLinkages(0)(mt)

	location, err := FreshNetworkLocatorText(UUIDLocationText(id, "REDACTED"))
	if err != nil {
		t.Fatal(err)
	}
	const maximumArrivingLinks = 2
	MultiplexCarrierMaximumArrivingLinkages(maximumArrivingLinks)(mt)
	if err := mt.Overhear(*location); err != nil {
		t.Fatal(err)
	}

	localaddr := FreshNetworkLocator(mt.peerToken.ID(), mt.observer.Addr())

	//
	for i := 0; i <= maximumArrivingLinks; i++ {
		faultchnl := make(chan error)
		go verifyCaller(*localaddr, faultchnl)

		err = <-faultchnl
		if i < maximumArrivingLinks {
			if err != nil {
				t.Errorf("REDACTED", err)
			}
			_, err = mt.Embrace(nodeSettings{})
			if err != nil {
				t.Errorf("REDACTED", err)
			}
		} else if err == nil || !strings.Contains(err.Error(), "REDACTED") {
			//
			//
			//
			t.Errorf("REDACTED", err)
		}
	}
}

func VerifyCarrierMultiplexEmbraceVarious(t *testing.T) {
	mt := verifyConfigureMultiplexCarrier(t)
	localaddr := FreshNetworkLocator(mt.peerToken.ID(), mt.observer.Addr())

	var (
		germ     = rand.New(rand.NewSource(time.Now().UnixNano()))
		nthCallers = germ.Intn(64) + 64
		faultchnl     = make(chan error, nthCallers)
	)

	//
	for i := 0; i < nthCallers; i++ {
		go verifyCaller(*localaddr, faultchnl)
	}

	//
	for i := 0; i < nthCallers; i++ {
		if err := <-faultchnl; err != nil {
			t.Fatal(err)
		}
	}

	ps := []Node{}

	//
	for i := 0; i < cap(faultchnl); i++ {
		p, err := mt.Embrace(nodeSettings{})
		if err != nil {
			t.Fatal(err)
		}

		if err := p.Initiate(); err != nil {
			t.Fatal(err)
		}

		ps = append(ps, p)
	}

	if possess, desire := len(ps), cap(faultchnl); possess != desire {
		t.Errorf("REDACTED", possess, desire)
	}

	//
	for _, p := range ps {
		if err := p.Halt(); err != nil {
			t.Fatal(err)
		}
	}

	if err := mt.Shutdown(); err != nil {
		t.Errorf("REDACTED", err)
	}
}

func verifyCaller(callLocation NetworkLocator, faultchnl chan error) {
	var (
		pv     = edwards25519.ProducePrivateToken()
		caller = freshMultiplexCarrier(
			verifyPeerDetails(PublicTokenTowardUUID(pv.PublicToken()), fallbackPeerAlias),
			PeerToken{
				PrivateToken: pv,
			},
		)
	)

	_, err := caller.Call(callLocation, nodeSettings{})
	if err != nil {
		faultchnl <- err
		return
	}

	//
	faultchnl <- nil
}

func VerifyCarrierMultiplexEmbraceUnHalting(t *testing.T) {
	mt := verifyConfigureMultiplexCarrier(t)

	var (
		swiftPeerPRV   = edwards25519.ProducePrivateToken()
		swiftPeerDetails = verifyPeerDetails(PublicTokenTowardUUID(swiftPeerPRV.PublicToken()), "REDACTED")
		faultchnl         = make(chan error)
		quickchnl        = make(chan struct{})
		sluggishchnl        = make(chan struct{})
		laggarddonechnl    = make(chan struct{})
	)

	//
	go func() {
		location := FreshNetworkLocator(mt.peerToken.ID(), mt.observer.Addr())

		c, err := location.Call()
		if err != nil {
			faultchnl <- err
			return
		}

		close(sluggishchnl)
		defer func() {
			close(laggarddonechnl)
		}()

		//
		runtime.Gosched()

		select {
		case <-quickchnl:
			//
		case <-time.After(200 * time.Millisecond):
			//
			faultchnl <- fmt.Errorf("REDACTED")
		}

		sc, err := modernizeCredentialLink(c, 200*time.Millisecond, edwards25519.ProducePrivateToken())
		if err != nil {
			faultchnl <- err
			return
		}

		_, err = negotiation(sc, 200*time.Millisecond,
			verifyPeerDetails(
				PublicTokenTowardUUID(edwards25519.ProducePrivateToken().PublicToken()),
				"REDACTED",
			))
		if err != nil {
			faultchnl <- err
		}
	}()

	//
	go func() {
		<-sluggishchnl

		caller := freshMultiplexCarrier(
			swiftPeerDetails,
			PeerToken{
				PrivateToken: swiftPeerPRV,
			},
		)
		location := FreshNetworkLocator(mt.peerToken.ID(), mt.observer.Addr())

		_, err := caller.Call(*location, nodeSettings{})
		if err != nil {
			faultchnl <- err
			return
		}

		close(quickchnl)
		<-laggarddonechnl
		close(faultchnl)
	}()

	if err := <-faultchnl; err != nil {
		t.Logf("REDACTED", err)
	}

	p, err := mt.Embrace(nodeSettings{})
	if err != nil {
		t.Fatal(err)
	}

	if possess, desire := p.PeerDetails(), swiftPeerDetails; !reflect.DeepEqual(possess, desire) {
		t.Errorf("REDACTED", possess, desire)
	}
}

func VerifyCarrierMultiplexCertifyPeerDetails(t *testing.T) {
	mt := verifyConfigureMultiplexCarrier(t)

	faultchnl := make(chan error)

	go func() {
		var (
			pv     = edwards25519.ProducePrivateToken()
			caller = freshMultiplexCarrier(
				verifyPeerDetails(PublicTokenTowardUUID(pv.PublicToken()), "REDACTED"), //
				PeerToken{
					PrivateToken: pv,
				},
			)
		)

		location := FreshNetworkLocator(mt.peerToken.ID(), mt.observer.Addr())

		_, err := caller.Call(*location, nodeSettings{})
		if err != nil {
			faultchnl <- err
			return
		}

		close(faultchnl)
	}()

	if err := <-faultchnl; err != nil {
		t.Errorf("REDACTED", err)
	}

	_, err := mt.Embrace(nodeSettings{})
	if e, ok := err.(FaultDeclined); ok {
		if !e.EqualsPeerDetailsUnfit() {
			t.Errorf("REDACTED", err)
		}
	} else {
		t.Errorf("REDACTED", err)
	}
}

func VerifyCarrierMultiplexDeclineDiscrepancyUUID(t *testing.T) {
	mt := verifyConfigureMultiplexCarrier(t)

	faultchnl := make(chan error)

	go func() {
		caller := freshMultiplexCarrier(
			verifyPeerDetails(
				PublicTokenTowardUUID(edwards25519.ProducePrivateToken().PublicToken()), "REDACTED",
			),
			PeerToken{
				PrivateToken: edwards25519.ProducePrivateToken(),
			},
		)
		location := FreshNetworkLocator(mt.peerToken.ID(), mt.observer.Addr())

		_, err := caller.Call(*location, nodeSettings{})
		if err != nil {
			faultchnl <- err
			return
		}

		close(faultchnl)
	}()

	if err := <-faultchnl; err != nil {
		t.Errorf("REDACTED", err)
	}

	_, err := mt.Embrace(nodeSettings{})
	if e, ok := err.(FaultDeclined); ok {
		if !e.EqualsAuthBreakdown() {
			t.Errorf("REDACTED", e)
		}
	} else {
		t.Errorf("REDACTED", err)
	}
}

func VerifyCarrierMultiplexCallDeclineIncorrectUUID(t *testing.T) {
	mt := verifyConfigureMultiplexCarrier(t)

	var (
		pv     = edwards25519.ProducePrivateToken()
		caller = freshMultiplexCarrier(
			verifyPeerDetails(PublicTokenTowardUUID(pv.PublicToken()), "REDACTED"), //
			PeerToken{
				PrivateToken: pv,
			},
		)
	)

	incorrectUUID := PublicTokenTowardUUID(edwards25519.ProducePrivateToken().PublicToken())
	location := FreshNetworkLocator(incorrectUUID, mt.observer.Addr())

	_, err := caller.Call(*location, nodeSettings{})
	if err != nil {
		t.Logf("REDACTED", err)
		if e, ok := err.(FaultDeclined); ok {
			if !e.EqualsAuthBreakdown() {
				t.Errorf("REDACTED", e)
			}
		} else {
			t.Errorf("REDACTED", err)
		}
	}
}

func VerifyCarrierMultiplexDeclineUnmatched(t *testing.T) {
	mt := verifyConfigureMultiplexCarrier(t)

	faultchnl := make(chan error)

	go func() {
		var (
			pv     = edwards25519.ProducePrivateToken()
			caller = freshMultiplexCarrier(
				verifyPeerDetailsUsingFabric(PublicTokenTowardUUID(pv.PublicToken()), "REDACTED", "REDACTED"),
				PeerToken{
					PrivateToken: pv,
				},
			)
		)
		location := FreshNetworkLocator(mt.peerToken.ID(), mt.observer.Addr())

		_, err := caller.Call(*location, nodeSettings{})
		if err != nil {
			faultchnl <- err
			return
		}

		close(faultchnl)
	}()

	_, err := mt.Embrace(nodeSettings{})
	if e, ok := err.(FaultDeclined); ok {
		if !e.EqualsUnmatched() {
			t.Errorf("REDACTED", e)
		}
	} else {
		t.Errorf("REDACTED", err)
	}
}

func VerifyCarrierMultiplexDeclineEgo(t *testing.T) {
	mt := verifyConfigureMultiplexCarrier(t)

	faultchnl := make(chan error)

	go func() {
		location := FreshNetworkLocator(mt.peerToken.ID(), mt.observer.Addr())

		_, err := mt.Call(*location, nodeSettings{})
		if err != nil {
			faultchnl <- err
			return
		}

		close(faultchnl)
	}()

	if err := <-faultchnl; err != nil {
		if e, ok := err.(FaultDeclined); ok {
			if !e.EqualsEgo() {
				t.Errorf("REDACTED", e)
			}
		} else {
			t.Errorf("REDACTED", err)
		}
	} else {
		t.Errorf("REDACTED")
	}

	_, err := mt.Embrace(nodeSettings{})
	if err, ok := err.(FaultDeclined); ok {
		if !err.EqualsEgo() {
			t.Errorf("REDACTED", err)
		}
	} else {
		t.Errorf("REDACTED", nil)
	}
}

func VerifyCarrierLinkReplicatedINETRefine(t *testing.T) {
	refine := LinkReplicatedINETRefine()

	if err := refine(nil, &verifyCarrierLink{}, nil); err != nil {
		t.Fatal(err)
	}

	var (
		c  = &verifyCarrierLink{}
		cs = FreshLinkAssign()
	)

	cs.Set(c, []net.IP{
		{10, 0, 10, 1},
		{10, 0, 10, 2},
		{10, 0, 10, 3},
	})

	if err := refine(cs, c, []net.IP{
		{10, 0, 10, 2},
	}); err == nil {
		t.Errorf("REDACTED")
	}
}

func VerifyCarrierNegotiation(t *testing.T) {
	ln, err := net.Listen("REDACTED", "REDACTED")
	if err != nil {
		t.Fatal(err)
	}

	var (
		nodePRV       = edwards25519.ProducePrivateToken()
		nodePeerDetails = verifyPeerDetails(PublicTokenTowardUUID(nodePRV.PublicToken()), fallbackPeerAlias)
	)

	go func() {
		c, err := net.Dial(ln.Addr().Network(), ln.Addr().String())
		if err != nil {
			t.Error(err)
			return
		}

		go func(c net.Conn) {
			_, err := protocolio.FreshSeparatedPersistor(c).PersistSignal(nodePeerDetails.(FallbackPeerDetails).TowardSchema())
			if err != nil {
				t.Error(err)
			}
		}(c)
		go func(c net.Conn) {

			//
			var protobufferindex tmpfabric.FallbackPeerDetails

			schemaFetcher := protocolio.FreshSeparatedFetcher(c, MaximumPeerDetailsExtent())
			_, err := schemaFetcher.FetchSignal(&protobufferindex)
			if err != nil {
				t.Error(err)
			}

			_, err = FallbackPeerDetailsOriginatingTowardSchema(&protobufferindex)
			if err != nil {
				t.Error(err)
			}
		}(c)
	}()

	c, err := ln.Accept()
	if err != nil {
		t.Fatal(err)
	}

	ni, err := negotiation(c, 20*time.Millisecond, blankPeerDetails())
	if err != nil {
		t.Fatal(err)
	}

	if possess, desire := ni, nodePeerDetails; !reflect.DeepEqual(possess, desire) {
		t.Errorf("REDACTED", possess, desire)
	}
}

func VerifyCarrierAppendConduit(t *testing.T) {
	mt := freshMultiplexCarrier(
		blankPeerDetails(),
		PeerToken{
			PrivateToken: edwards25519.ProducePrivateToken(),
		},
	)
	verifyConduit := byte(0x01)

	mt.AppendConduit(verifyConduit)
	if !mt.peerDetails.(FallbackPeerDetails).OwnsConduit(verifyConduit) {
		t.Errorf("REDACTED", verifyConduit, mt.peerDetails.(FallbackPeerDetails).Conduits)
	}
}

//
func verifyConfigureMultiplexCarrier(t *testing.T) *MultiplexCarrier {
	var (
		pv = edwards25519.ProducePrivateToken()
		id = PublicTokenTowardUUID(pv.PublicToken())
		mt = freshMultiplexCarrier(
			verifyPeerDetails(
				id, "REDACTED",
			),
			PeerToken{
				PrivateToken: pv,
			},
		)
	)

	location, err := FreshNetworkLocatorText(UUIDLocationText(id, "REDACTED"))
	if err != nil {
		t.Fatal(err)
	}

	if err := mt.Overhear(*location); err != nil {
		t.Fatal(err)
	}

	//
	time.Sleep(20 * time.Millisecond)

	return mt
}

type verifyCarrierLocation struct{}

func (a *verifyCarrierLocation) Fabric() string { return "REDACTED" }
func (a *verifyCarrierLocation) Text() string  { return "REDACTED" }

type verifyCarrierLink struct{}

func (c *verifyCarrierLink) Shutdown() error {
	return fmt.Errorf("REDACTED")
}

func (c *verifyCarrierLink) RegionalLocation() net.Addr {
	return &verifyCarrierLocation{}
}

func (c *verifyCarrierLink) DistantLocation() net.Addr {
	return &verifyCarrierLocation{}
}

func (c *verifyCarrierLink) Obtain(_ []byte) (int, error) {
	return -1, fmt.Errorf("REDACTED")
}

func (c *verifyCarrierLink) AssignExpiration(_ time.Time) error {
	return fmt.Errorf("REDACTED")
}

func (c *verifyCarrierLink) AssignFetchExpiration(_ time.Time) error {
	return fmt.Errorf("REDACTED")
}

func (c *verifyCarrierLink) AssignPersistExpiration(_ time.Time) error {
	return fmt.Errorf("REDACTED")
}

func (c *verifyCarrierLink) Record(_ []byte) (int, error) {
	return -1, fmt.Errorf("REDACTED")
}
