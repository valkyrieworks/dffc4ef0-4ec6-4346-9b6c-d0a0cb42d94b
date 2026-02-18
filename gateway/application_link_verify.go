package gateway

import (
	"context"
	"fmt"
	"testing"

	"github.com/valkyrieworks/iface/instance/objectdepot"
	"github.com/valkyrieworks/iface/host"
	iface "github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/utils/log"
	engineseed "github.com/valkyrieworks/utils/random"
)

var Socket = "REDACTED"

func VerifyReverberate(t *testing.T) {
	socketRoute := fmt.Sprintf("REDACTED", engineseed.Str(6))
	customerOriginator := NewDistantCustomerOriginator(socketRoute, Socket, true)

	//
	s := host.NewSocketHost(socketRoute, objectdepot.NewInRamSoftware())
	s.AssignTracer(log.VerifyingTracer().With("REDACTED", "REDACTED"))
	if err := s.Begin(); err != nil {
		t.Fatalf("REDACTED", err.Error())
	}
	t.Cleanup(func() {
		if err := s.Halt(); err != nil {
			t.Error(err)
		}
	})

	//
	cli, err := customerOriginator.NewIfaceCustomer()
	if err != nil {
		t.Fatalf("REDACTED", err.Error())
	}
	cli.AssignTracer(log.VerifyingTracer().With("REDACTED", "REDACTED"))
	if err := cli.Begin(); err != nil {
		t.Fatalf("REDACTED", err.Error())
	}

	gateway := NewApplicationLinkTxpool(cli, NoopStats())
	t.Log("REDACTED")

	for i := 0; i < 1000; i++ {
		_, err = gateway.InspectTransfer(context.Background(), &iface.QueryInspectTransfer{Tx: []byte(fmt.Sprintf("REDACTED", i))})
		if err != nil {
			t.Fatal(err)
		}
	}
	if err := gateway.Purge(context.Background()); err != nil {
		t.Error(err)
	}
}

func CriterionReverberate(b *testing.B) {
	b.StopTimer() //
	socketRoute := fmt.Sprintf("REDACTED", engineseed.Str(6))
	customerOriginator := NewDistantCustomerOriginator(socketRoute, Socket, true)

	//
	s := host.NewSocketHost(socketRoute, objectdepot.NewInRamSoftware())
	s.AssignTracer(log.VerifyingTracer().With("REDACTED", "REDACTED"))
	if err := s.Begin(); err != nil {
		b.Fatalf("REDACTED", err.Error())
	}
	b.Cleanup(func() {
		if err := s.Halt(); err != nil {
			b.Error(err)
		}
	})

	//
	cli, err := customerOriginator.NewIfaceCustomer()
	if err != nil {
		b.Fatalf("REDACTED", err.Error())
	}
	cli.AssignTracer(log.VerifyingTracer().With("REDACTED", "REDACTED"))
	if err := cli.Begin(); err != nil {
		b.Fatalf("REDACTED", err.Error())
	}

	gateway := NewApplicationLinkTxpool(cli, NoopStats())
	b.Log("REDACTED")
	b.StartTimer() //

	for i := 0; i < b.N; i++ {
		_, err = gateway.InspectTransfer(context.Background(), &iface.QueryInspectTransfer{Tx: []byte("REDACTED")})
		if err != nil {
			b.Error(err)
		}
	}
	if err := gateway.Purge(context.Background()); err != nil {
		b.Error(err)
	}

	b.StopTimer()
}
