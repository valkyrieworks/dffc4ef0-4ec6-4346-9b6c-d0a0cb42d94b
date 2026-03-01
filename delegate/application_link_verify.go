package delegate

import (
	"context"
	"fmt"
	"testing"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/instance/statedepot"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/node"
	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
)

var PORT = "REDACTED"

func VerifyReverberate(t *testing.T) {
	terminalRoute := fmt.Sprintf("REDACTED", commitrand.Str(6))
	customerOriginator := FreshDistantCustomerOriginator(terminalRoute, PORT, true)

	//
	s := node.FreshPortDaemon(terminalRoute, statedepot.FreshInsideRamPlatform())
	s.AssignTracer(log.VerifyingTracer().Using("REDACTED", "REDACTED"))
	if err := s.Initiate(); err != nil {
		t.Fatalf("REDACTED", err.Error())
	}
	t.Cleanup(func() {
		if err := s.Halt(); err != nil {
			t.Error(err)
		}
	})

	//
	cli, err := customerOriginator.FreshIfaceCustomer()
	if err != nil {
		t.Fatalf("REDACTED", err.Error())
	}
	cli.AssignTracer(log.VerifyingTracer().Using("REDACTED", "REDACTED"))
	if err := cli.Initiate(); err != nil {
		t.Fatalf("REDACTED", err.Error())
	}

	delegate := FreshApplicationLinkTxpool(cli, NooperationTelemetry())
	t.Log("REDACTED")

	for i := 0; i < 1000; i++ {
		_, err = delegate.InspectTransfer(context.Background(), &iface.SolicitInspectTransfer{Tx: []byte(fmt.Sprintf("REDACTED", i))})
		if err != nil {
			t.Fatal(err)
		}
	}
	if err := delegate.Purge(context.Background()); err != nil {
		t.Error(err)
	}
}

func AssessmentReverberate(b *testing.B) {
	b.StopTimer() //
	terminalRoute := fmt.Sprintf("REDACTED", commitrand.Str(6))
	customerOriginator := FreshDistantCustomerOriginator(terminalRoute, PORT, true)

	//
	s := node.FreshPortDaemon(terminalRoute, statedepot.FreshInsideRamPlatform())
	s.AssignTracer(log.VerifyingTracer().Using("REDACTED", "REDACTED"))
	if err := s.Initiate(); err != nil {
		b.Fatalf("REDACTED", err.Error())
	}
	b.Cleanup(func() {
		if err := s.Halt(); err != nil {
			b.Error(err)
		}
	})

	//
	cli, err := customerOriginator.FreshIfaceCustomer()
	if err != nil {
		b.Fatalf("REDACTED", err.Error())
	}
	cli.AssignTracer(log.VerifyingTracer().Using("REDACTED", "REDACTED"))
	if err := cli.Initiate(); err != nil {
		b.Fatalf("REDACTED", err.Error())
	}

	delegate := FreshApplicationLinkTxpool(cli, NooperationTelemetry())
	b.Log("REDACTED")
	b.StartTimer() //

	for i := 0; i < b.N; i++ {
		_, err = delegate.InspectTransfer(context.Background(), &iface.SolicitInspectTransfer{Tx: []byte("REDACTED")})
		if err != nil {
			b.Error(err)
		}
	}
	if err := delegate.Purge(context.Background()); err != nil {
		b.Error(err)
	}

	b.StopTimer()
}
