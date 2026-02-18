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

	"github.com/valkyrieworks/vault/ed25519"
	"github.com/valkyrieworks/utils/protoio"
	"github.com/valkyrieworks/p2p/link"
	tmp2p "github.com/valkyrieworks/schema/consensuscore/p2p"
)

var standardMemberLabel = "REDACTED"

func emptyMemberDetails() MemberDetails {
	return StandardMemberDetails{}
}

//
//
//
func newMultiplexCarrier(
	memberDetails MemberDetails,
	memberKey MemberKey,
) *MulticastCarrier {
	return NewMulticastCarrier(
		memberDetails, memberKey, link.StandardMLinkSettings(),
	)
}

func VerifyCarrierMultiplexLinkRefine(t *testing.T) {
	mt := newMultiplexCarrier(
		emptyMemberDetails(),
		MemberKey{
			PrivateKey: ed25519.GeneratePrivateKey(),
		},
	)
	id := mt.memberKey.ID()

	MulticastCarrierLinkScreens(
		func(_ LinkCollection, _ net.Conn, _ []net.IP) error { return nil },
		func(_ LinkCollection, _ net.Conn, _ []net.IP) error { return nil },
		func(_ LinkCollection, _ net.Conn, _ []net.IP) error {
			return fmt.Errorf("REDACTED")
		},
	)(mt)

	address, err := NewNetLocationString(UIDLocationString(id, "REDACTED"))
	if err != nil {
		t.Fatal(err)
	}

	if err := mt.Observe(*address); err != nil {
		t.Fatal(err)
	}

	faultc := make(chan error)

	go func() {
		address := NewNetLocation(id, mt.observer.Addr())

		_, err := address.Call()
		if err != nil {
			faultc <- err
			return
		}

		close(faultc)
	}()

	if err := <-faultc; err != nil {
		t.Errorf("REDACTED", err)
	}

	_, err = mt.Allow(nodeSettings{})
	if e, ok := err.(ErrDeclined); ok {
		if !e.IsScreened() {
			t.Errorf("REDACTED", err)
		}
	} else {
		t.Errorf("REDACTED", err)
	}
}

func VerifyCarrierMultiplexLinkRefineDeadline(t *testing.T) {
	mt := newMultiplexCarrier(
		emptyMemberDetails(),
		MemberKey{
			PrivateKey: ed25519.GeneratePrivateKey(),
		},
	)
	id := mt.memberKey.ID()

	MulticastCarrierRefineDeadline(5 * time.Millisecond)(mt)
	MulticastCarrierLinkScreens(
		func(_ LinkCollection, _ net.Conn, _ []net.IP) error {
			time.Sleep(1 * time.Second)
			return nil
		},
	)(mt)

	address, err := NewNetLocationString(UIDLocationString(id, "REDACTED"))
	if err != nil {
		t.Fatal(err)
	}

	if err := mt.Observe(*address); err != nil {
		t.Fatal(err)
	}

	faultc := make(chan error)
	go func() {
		address := NewNetLocation(id, mt.observer.Addr())

		_, err := address.Call()
		if err != nil {
			faultc <- err
			return
		}

		close(faultc)
	}()

	if err := <-faultc; err != nil {
		t.Errorf("REDACTED", err)
	}

	_, err = mt.Allow(nodeSettings{})
	if _, ok := err.(ErrRefineDeadline); !ok {
		t.Errorf("REDACTED", err)
	}
}

func VerifyCarrierMultiplexMaximumArrivingLinkages(t *testing.T) {
	pv := ed25519.GeneratePrivateKey()
	id := PublicKeyToUID(pv.PublicKey())
	mt := newMultiplexCarrier(
		verifyMemberDetails(
			id, "REDACTED",
		),
		MemberKey{
			PrivateKey: pv,
		},
	)

	MulticastCarrierMaximumIncomingLinkages(0)(mt)

	address, err := NewNetLocationString(UIDLocationString(id, "REDACTED"))
	if err != nil {
		t.Fatal(err)
	}
	const maximumArrivingLinks = 2
	MulticastCarrierMaximumIncomingLinkages(maximumArrivingLinks)(mt)
	if err := mt.Observe(*address); err != nil {
		t.Fatal(err)
	}

	laddress := NewNetLocation(mt.memberKey.ID(), mt.observer.Addr())

	//
	for i := 0; i <= maximumArrivingLinks; i++ {
		faultc := make(chan error)
		go verifyCaller(*laddress, faultc)

		err = <-faultc
		if i < maximumArrivingLinks {
			if err != nil {
				t.Errorf("REDACTED", err)
			}
			_, err = mt.Allow(nodeSettings{})
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

func VerifyCarrierMultiplexAllowVaried(t *testing.T) {
	mt := verifyConfigureMultiplexCarrier(t)
	laddress := NewNetLocation(mt.memberKey.ID(), mt.observer.Addr())

	var (
		origin     = rand.New(rand.NewSource(time.Now().UnixNano()))
		nCallers = origin.Intn(64) + 64
		faultc     = make(chan error, nCallers)
	)

	//
	for i := 0; i < nCallers; i++ {
		go verifyCaller(*laddress, faultc)
	}

	//
	for i := 0; i < nCallers; i++ {
		if err := <-faultc; err != nil {
			t.Fatal(err)
		}
	}

	ps := []Node{}

	//
	for i := 0; i < cap(faultc); i++ {
		p, err := mt.Allow(nodeSettings{})
		if err != nil {
			t.Fatal(err)
		}

		if err := p.Begin(); err != nil {
			t.Fatal(err)
		}

		ps = append(ps, p)
	}

	if possess, desire := len(ps), cap(faultc); possess != desire {
		t.Errorf("REDACTED", possess, desire)
	}

	//
	for _, p := range ps {
		if err := p.Halt(); err != nil {
			t.Fatal(err)
		}
	}

	if err := mt.End(); err != nil {
		t.Errorf("REDACTED", err)
	}
}

func verifyCaller(callAddress NetLocation, faultc chan error) {
	var (
		pv     = ed25519.GeneratePrivateKey()
		caller = newMultiplexCarrier(
			verifyMemberDetails(PublicKeyToUID(pv.PublicKey()), standardMemberLabel),
			MemberKey{
				PrivateKey: pv,
			},
		)
	)

	_, err := caller.Call(callAddress, nodeSettings{})
	if err != nil {
		faultc <- err
		return
	}

	//
	faultc <- nil
}

func VerifyCarrierMultiplexAllowNotHalting(t *testing.T) {
	mt := verifyConfigureMultiplexCarrier(t)

	var (
		quickMemberPV   = ed25519.GeneratePrivateKey()
		quickMemberDetails = verifyMemberDetails(PublicKeyToUID(quickMemberPV.PublicKey()), "REDACTED")
		faultc         = make(chan error)
		fastc        = make(chan struct{})
		slowc        = make(chan struct{})
		slowdonec    = make(chan struct{})
	)

	//
	go func() {
		address := NewNetLocation(mt.memberKey.ID(), mt.observer.Addr())

		c, err := address.Call()
		if err != nil {
			faultc <- err
			return
		}

		close(slowc)
		defer func() {
			close(slowdonec)
		}()

		//
		runtime.Gosched()

		select {
		case <-fastc:
			//
		case <-time.After(200 * time.Millisecond):
			//
			faultc <- fmt.Errorf("REDACTED")
		}

		sc, err := enhanceTokenLink(c, 200*time.Millisecond, ed25519.GeneratePrivateKey())
		if err != nil {
			faultc <- err
			return
		}

		_, err = greeting(sc, 200*time.Millisecond,
			verifyMemberDetails(
				PublicKeyToUID(ed25519.GeneratePrivateKey().PublicKey()),
				"REDACTED",
			))
		if err != nil {
			faultc <- err
		}
	}()

	//
	go func() {
		<-slowc

		caller := newMultiplexCarrier(
			quickMemberDetails,
			MemberKey{
				PrivateKey: quickMemberPV,
			},
		)
		address := NewNetLocation(mt.memberKey.ID(), mt.observer.Addr())

		_, err := caller.Call(*address, nodeSettings{})
		if err != nil {
			faultc <- err
			return
		}

		close(fastc)
		<-slowdonec
		close(faultc)
	}()

	if err := <-faultc; err != nil {
		t.Logf("REDACTED", err)
	}

	p, err := mt.Allow(nodeSettings{})
	if err != nil {
		t.Fatal(err)
	}

	if possess, desire := p.MemberDetails(), quickMemberDetails; !reflect.DeepEqual(possess, desire) {
		t.Errorf("REDACTED", possess, desire)
	}
}

func VerifyCarrierMultiplexCertifyMemberDetails(t *testing.T) {
	mt := verifyConfigureMultiplexCarrier(t)

	faultc := make(chan error)

	go func() {
		var (
			pv     = ed25519.GeneratePrivateKey()
			caller = newMultiplexCarrier(
				verifyMemberDetails(PublicKeyToUID(pv.PublicKey()), "REDACTED"), //
				MemberKey{
					PrivateKey: pv,
				},
			)
		)

		address := NewNetLocation(mt.memberKey.ID(), mt.observer.Addr())

		_, err := caller.Call(*address, nodeSettings{})
		if err != nil {
			faultc <- err
			return
		}

		close(faultc)
	}()

	if err := <-faultc; err != nil {
		t.Errorf("REDACTED", err)
	}

	_, err := mt.Allow(nodeSettings{})
	if e, ok := err.(ErrDeclined); ok {
		if !e.IsMemberDetailsCorrupt() {
			t.Errorf("REDACTED", err)
		}
	} else {
		t.Errorf("REDACTED", err)
	}
}

func VerifyCarrierMultiplexDeclineDiscrepancyUID(t *testing.T) {
	mt := verifyConfigureMultiplexCarrier(t)

	faultc := make(chan error)

	go func() {
		caller := newMultiplexCarrier(
			verifyMemberDetails(
				PublicKeyToUID(ed25519.GeneratePrivateKey().PublicKey()), "REDACTED",
			),
			MemberKey{
				PrivateKey: ed25519.GeneratePrivateKey(),
			},
		)
		address := NewNetLocation(mt.memberKey.ID(), mt.observer.Addr())

		_, err := caller.Call(*address, nodeSettings{})
		if err != nil {
			faultc <- err
			return
		}

		close(faultc)
	}()

	if err := <-faultc; err != nil {
		t.Errorf("REDACTED", err)
	}

	_, err := mt.Allow(nodeSettings{})
	if e, ok := err.(ErrDeclined); ok {
		if !e.IsAuthBreakdown() {
			t.Errorf("REDACTED", e)
		}
	} else {
		t.Errorf("REDACTED", err)
	}
}

func VerifyCarrierMultiplexCallDeclineIncorrectUID(t *testing.T) {
	mt := verifyConfigureMultiplexCarrier(t)

	var (
		pv     = ed25519.GeneratePrivateKey()
		caller = newMultiplexCarrier(
			verifyMemberDetails(PublicKeyToUID(pv.PublicKey()), "REDACTED"), //
			MemberKey{
				PrivateKey: pv,
			},
		)
	)

	incorrectUID := PublicKeyToUID(ed25519.GeneratePrivateKey().PublicKey())
	address := NewNetLocation(incorrectUID, mt.observer.Addr())

	_, err := caller.Call(*address, nodeSettings{})
	if err != nil {
		t.Logf("REDACTED", err)
		if e, ok := err.(ErrDeclined); ok {
			if !e.IsAuthBreakdown() {
				t.Errorf("REDACTED", e)
			}
		} else {
			t.Errorf("REDACTED", err)
		}
	}
}

func VerifyCarrierMultiplexDeclineDiscordant(t *testing.T) {
	mt := verifyConfigureMultiplexCarrier(t)

	faultc := make(chan error)

	go func() {
		var (
			pv     = ed25519.GeneratePrivateKey()
			caller = newMultiplexCarrier(
				verifyMemberDetailsWithFabric(PublicKeyToUID(pv.PublicKey()), "REDACTED", "REDACTED"),
				MemberKey{
					PrivateKey: pv,
				},
			)
		)
		address := NewNetLocation(mt.memberKey.ID(), mt.observer.Addr())

		_, err := caller.Call(*address, nodeSettings{})
		if err != nil {
			faultc <- err
			return
		}

		close(faultc)
	}()

	_, err := mt.Allow(nodeSettings{})
	if e, ok := err.(ErrDeclined); ok {
		if !e.IsDiscordant() {
			t.Errorf("REDACTED", e)
		}
	} else {
		t.Errorf("REDACTED", err)
	}
}

func VerifyCarrierMultiplexDeclineEgo(t *testing.T) {
	mt := verifyConfigureMultiplexCarrier(t)

	faultc := make(chan error)

	go func() {
		address := NewNetLocation(mt.memberKey.ID(), mt.observer.Addr())

		_, err := mt.Call(*address, nodeSettings{})
		if err != nil {
			faultc <- err
			return
		}

		close(faultc)
	}()

	if err := <-faultc; err != nil {
		if e, ok := err.(ErrDeclined); ok {
			if !e.IsEgo() {
				t.Errorf("REDACTED", e)
			}
		} else {
			t.Errorf("REDACTED", err)
		}
	} else {
		t.Errorf("REDACTED")
	}

	_, err := mt.Allow(nodeSettings{})
	if err, ok := err.(ErrDeclined); ok {
		if !err.IsEgo() {
			t.Errorf("REDACTED", err)
		}
	} else {
		t.Errorf("REDACTED", nil)
	}
}

func VerifyCarrierLinkReplicatedIPRefine(t *testing.T) {
	refine := LinkReplicatedIPRefine()

	if err := refine(nil, &verifyCarrierLink{}, nil); err != nil {
		t.Fatal(err)
	}

	var (
		c  = &verifyCarrierLink{}
		cs = NewLinkCollection()
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

func VerifyCarrierGreeting(t *testing.T) {
	ln, err := net.Listen("REDACTED", "REDACTED")
	if err != nil {
		t.Fatal(err)
	}

	var (
		nodePV       = ed25519.GeneratePrivateKey()
		nodeMemberDetails = verifyMemberDetails(PublicKeyToUID(nodePV.PublicKey()), standardMemberLabel)
	)

	go func() {
		c, err := net.Dial(ln.Addr().Network(), ln.Addr().String())
		if err != nil {
			t.Error(err)
			return
		}

		go func(c net.Conn) {
			_, err := protoio.NewSeparatedRecorder(c).RecordMessage(nodeMemberDetails.(StandardMemberDetails).ToSchema())
			if err != nil {
				t.Error(err)
			}
		}(c)
		go func(c net.Conn) {

			//
			var pbni tmp2p.StandardMemberDetails

			schemaScanner := protoio.NewSeparatedScanner(c, MaximumMemberDetailsVolume())
			_, err := schemaScanner.ScanMessage(&pbni)
			if err != nil {
				t.Error(err)
			}

			_, err = StandardMemberDetailsFromToSchema(&pbni)
			if err != nil {
				t.Error(err)
			}
		}(c)
	}()

	c, err := ln.Accept()
	if err != nil {
		t.Fatal(err)
	}

	ni, err := greeting(c, 20*time.Millisecond, emptyMemberDetails())
	if err != nil {
		t.Fatal(err)
	}

	if possess, desire := ni, nodeMemberDetails; !reflect.DeepEqual(possess, desire) {
		t.Errorf("REDACTED", possess, desire)
	}
}

func VerifyCarrierAppendConduit(t *testing.T) {
	mt := newMultiplexCarrier(
		emptyMemberDetails(),
		MemberKey{
			PrivateKey: ed25519.GeneratePrivateKey(),
		},
	)
	verifyConduit := byte(0x01)

	mt.AppendConduit(verifyConduit)
	if !mt.memberDetails.(StandardMemberDetails).HasConduit(verifyConduit) {
		t.Errorf("REDACTED", verifyConduit, mt.memberDetails.(StandardMemberDetails).Streams)
	}
}

//
func verifyConfigureMultiplexCarrier(t *testing.T) *MulticastCarrier {
	var (
		pv = ed25519.GeneratePrivateKey()
		id = PublicKeyToUID(pv.PublicKey())
		mt = newMultiplexCarrier(
			verifyMemberDetails(
				id, "REDACTED",
			),
			MemberKey{
				PrivateKey: pv,
			},
		)
	)

	address, err := NewNetLocationString(UIDLocationString(id, "REDACTED"))
	if err != nil {
		t.Fatal(err)
	}

	if err := mt.Observe(*address); err != nil {
		t.Fatal(err)
	}

	//
	time.Sleep(20 * time.Millisecond)

	return mt
}

type verifyCarrierAddress struct{}

func (a *verifyCarrierAddress) Fabric() string { return "REDACTED" }
func (a *verifyCarrierAddress) String() string  { return "REDACTED" }

type verifyCarrierLink struct{}

func (c *verifyCarrierLink) End() error {
	return fmt.Errorf("REDACTED")
}

func (c *verifyCarrierLink) NativeAddress() net.Addr {
	return &verifyCarrierAddress{}
}

func (c *verifyCarrierLink) DistantAddress() net.Addr {
	return &verifyCarrierAddress{}
}

func (c *verifyCarrierLink) Scan(_ []byte) (int, error) {
	return -1, fmt.Errorf("REDACTED")
}

func (c *verifyCarrierLink) CollectionLimit(_ time.Time) error {
	return fmt.Errorf("REDACTED")
}

func (c *verifyCarrierLink) CollectionReadLimit(_ time.Time) error {
	return fmt.Errorf("REDACTED")
}

func (c *verifyCarrierLink) CollectionRecordLimit(_ time.Time) error {
	return fmt.Errorf("REDACTED")
}

func (c *verifyCarrierLink) Record(_ []byte) (int, error) {
	return -1, fmt.Errorf("REDACTED")
}
