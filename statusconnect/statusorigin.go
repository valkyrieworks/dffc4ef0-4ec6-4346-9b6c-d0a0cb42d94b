package statusconnect

import (
	"context"
	"fmt"
	"strings"
	"time"

	dbm "github.com/valkyrieworks/-db"

	"github.com/valkyrieworks/utils/log"
	engineconnect "github.com/valkyrieworks/utils/align"
	"github.com/valkyrieworks/rapid"
	rapidsource "github.com/valkyrieworks/rapid/source"
	rapidhttp "github.com/valkyrieworks/rapid/source/http"
	rapidrpc "github.com/valkyrieworks/rapid/rpc"
	rapidstore "github.com/valkyrieworks/rapid/depot/db"
	cometstatus "github.com/valkyrieworks/schema/consensuscore/status"
	rpchttp "github.com/valkyrieworks/rpc/customer/http"
	sm "github.com/valkyrieworks/status"
	"github.com/valkyrieworks/kinds"
	"github.com/valkyrieworks/release"
)

//

//
//
type StatusSource interface {
	//
	ApplicationDigest(ctx context.Context, level uint64) ([]byte, error)
	//
	Endorse(ctx context.Context, level uint64) (*kinds.Endorse, error)
	//
	Status(ctx context.Context, level uint64) (sm.Status, error)
}

//
type rapidCustomerStatusSource struct {
	engineconnect.Lock //
	lc            *rapid.Customer
	release       cometstatus.Release
	primaryLevel int64
	sources     map[rapidsource.Source]string
}

//
func NewRapidCustomerStatusSource(
	ctx context.Context,
	ledgerUID string,
	release cometstatus.Release,
	primaryLevel int64,
	hosts []string,
	validateOptions rapid.ValidateOptions,
	tracer log.Tracer,
) (StatusSource, error) {
	if len(hosts) < 2 {
		return nil, fmt.Errorf("REDACTED", len(hosts))
	}

	sources := make([]rapidsource.Source, 0, len(hosts))
	sourceDistant := make(map[rapidsource.Source]string)
	for _, host := range hosts {
		customer, err := rpcCustomer(host)
		if err != nil {
			return nil, fmt.Errorf("REDACTED", err)
		}
		source := rapidhttp.NewWithCustomer(ledgerUID, customer)
		sources = append(sources, source)
		//
		//
		sourceDistant[source] = host
	}

	lc, err := rapid.NewCustomer(ctx, ledgerUID, validateOptions, sources[0], sources[1:],
		rapidstore.New(dbm.NewMemoryStore(), "REDACTED"), rapid.Tracer(tracer), rapid.MaximumReprocessTries(5))
	if err != nil {
		return nil, err
	}
	return &rapidCustomerStatusSource{
		lc:            lc,
		release:       release,
		primaryLevel: primaryLevel,
		sources:     sourceDistant,
	}, nil
}

//
func (s *rapidCustomerStatusSource) ApplicationDigest(ctx context.Context, level uint64) ([]byte, error) {
	s.Lock()
	defer s.Unlock()

	//
	heading, err := s.lc.ValidateRapidLedgerAtLevel(ctx, int64(level+1), time.Now())
	if err != nil {
		return nil, err
	}
	//
	//
	//
	//
	//
	//
	//
	//
	_, err = s.lc.ValidateRapidLedgerAtLevel(ctx, int64(level+2), time.Now())
	if err != nil {
		return nil, err
	}
	return heading.ApplicationDigest, nil
}

//
func (s *rapidCustomerStatusSource) Endorse(ctx context.Context, level uint64) (*kinds.Endorse, error) {
	s.Lock()
	defer s.Unlock()
	heading, err := s.lc.ValidateRapidLedgerAtLevel(ctx, int64(level), time.Now())
	if err != nil {
		return nil, err
	}
	return heading.Endorse, nil
}

//
func (s *rapidCustomerStatusSource) Status(ctx context.Context, level uint64) (sm.Status, error) {
	s.Lock()
	defer s.Unlock()

	status := sm.Status{
		LedgerUID:       s.lc.LedgerUID(),
		Release:       s.release,
		PrimaryLevel: s.primaryLevel,
	}
	if status.PrimaryLevel == 0 {
		status.PrimaryLevel = 1
	}

	//
	//
	//
	//
	//
	//
	//
	//
	finalRapidLedger, err := s.lc.ValidateRapidLedgerAtLevel(ctx, int64(level), time.Now())
	if err != nil {
		return sm.Status{}, err
	}
	presentRapidLedger, err := s.lc.ValidateRapidLedgerAtLevel(ctx, int64(level+1), time.Now())
	if err != nil {
		return sm.Status{}, err
	}
	followingRapidLedger, err := s.lc.ValidateRapidLedgerAtLevel(ctx, int64(level+2), time.Now())
	if err != nil {
		return sm.Status{}, err
	}

	status.Release = cometstatus.Release{
		Agreement: presentRapidLedger.Release,
		Software:  release.TMCoreSemaphoreRev,
	}
	status.FinalLedgerLevel = finalRapidLedger.Level
	status.FinalLedgerTime = finalRapidLedger.Time
	status.FinalLedgerUID = finalRapidLedger.Endorse.LedgerUID
	status.ApplicationDigest = presentRapidLedger.ApplicationDigest
	status.FinalOutcomesDigest = presentRapidLedger.FinalOutcomesDigest
	status.FinalRatifiers = finalRapidLedger.RatifierAssign
	status.Ratifiers = presentRapidLedger.RatifierAssign
	status.FollowingRatifiers = followingRapidLedger.RatifierAssign
	status.FinalLevelRatifiersModified = followingRapidLedger.Level

	//
	leadingURL, ok := s.sources[s.lc.Leading()]
	if !ok || leadingURL == "REDACTED" {
		return sm.Status{}, fmt.Errorf("REDACTED")
	}
	leadingRPC, err := rpcCustomer(leadingURL)
	if err != nil {
		return sm.Status{}, fmt.Errorf("REDACTED", err)
	}
	rpccustomer := rapidrpc.NewCustomer(leadingRPC, s.lc)
	outcome, err := rpccustomer.AgreementOptions(ctx, &presentRapidLedger.Level)
	if err != nil {
		return sm.Status{}, fmt.Errorf("REDACTED",
			followingRapidLedger.Level, err)
	}
	status.AgreementOptions = outcome.AgreementOptions
	status.FinalLevelAgreementOptionsModified = presentRapidLedger.Level

	return status, nil
}

//
func rpcCustomer(host string) (*rpchttp.HTTP, error) {
	if !strings.Contains(host, "REDACTED") {
		host = "REDACTED" + host
	}
	c, err := rpchttp.New(host, "REDACTED")
	if err != nil {
		return nil, err
	}
	return c, nil
}
