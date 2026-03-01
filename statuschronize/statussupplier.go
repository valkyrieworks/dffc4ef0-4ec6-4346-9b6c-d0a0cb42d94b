package statuschronize

import (
	"context"
	"fmt"
	"strings"
	"time"

	dbm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/-db"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	commitchronize "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agile"
	agilesupplier "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agile/supplier"
	agilehttpsvc "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agile/supplier/httpsvc"
	agileiface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agile/rpc"
	agiledatastore "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agile/depot/db"
	strongstatus "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/status"
	rpchttpsvc "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/customer/httpsvc"
	sm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/edition"
)

//

//
//
type StatusSupplier interface {
	//
	PlatformDigest(ctx context.Context, altitude uint64) ([]byte, error)
	//
	Endorse(ctx context.Context, altitude uint64) (*kinds.Endorse, error)
	//
	Status(ctx context.Context, altitude uint64) (sm.Status, error)
}

//
type agileCustomerStatusSupplier struct {
	commitchronize.Exclusion //
	lc            *agile.Customer
	edition       strongstatus.Edition
	primaryAltitude int64
	suppliers     map[agilesupplier.Supplier]string
}

//
func FreshAgileCustomerStatusSupplier(
	ctx context.Context,
	successionUUID string,
	edition strongstatus.Edition,
	primaryAltitude int64,
	nodes []string,
	relianceChoices agile.RelianceChoices,
	tracer log.Tracer,
) (StatusSupplier, error) {
	if len(nodes) < 2 {
		return nil, fmt.Errorf("REDACTED", len(nodes))
	}

	suppliers := make([]agilesupplier.Supplier, 0, len(nodes))
	supplierDistant := make(map[agilesupplier.Supplier]string)
	for _, node := range nodes {
		customer, err := ifaceCustomer(node)
		if err != nil {
			return nil, fmt.Errorf("REDACTED", err)
		}
		supplier := agilehttpsvc.FreshUsingCustomer(successionUUID, customer)
		suppliers = append(suppliers, supplier)
		//
		//
		supplierDistant[supplier] = node
	}

	lc, err := agile.FreshCustomer(ctx, successionUUID, relianceChoices, suppliers[0], suppliers[1:],
		agiledatastore.New(dbm.FreshMemoryDatastore(), "REDACTED"), agile.Tracer(tracer), agile.MaximumReissueEndeavors(5))
	if err != nil {
		return nil, err
	}
	return &agileCustomerStatusSupplier{
		lc:            lc,
		edition:       edition,
		primaryAltitude: primaryAltitude,
		suppliers:     supplierDistant,
	}, nil
}

//
func (s *agileCustomerStatusSupplier) PlatformDigest(ctx context.Context, altitude uint64) ([]byte, error) {
	s.Lock()
	defer s.Unlock()

	//
	heading, err := s.lc.ValidateAgileLedgerLocatedAltitude(ctx, int64(altitude+1), time.Now())
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
	_, err = s.lc.ValidateAgileLedgerLocatedAltitude(ctx, int64(altitude+2), time.Now())
	if err != nil {
		return nil, err
	}
	return heading.PlatformDigest, nil
}

//
func (s *agileCustomerStatusSupplier) Endorse(ctx context.Context, altitude uint64) (*kinds.Endorse, error) {
	s.Lock()
	defer s.Unlock()
	heading, err := s.lc.ValidateAgileLedgerLocatedAltitude(ctx, int64(altitude), time.Now())
	if err != nil {
		return nil, err
	}
	return heading.Endorse, nil
}

//
func (s *agileCustomerStatusSupplier) Status(ctx context.Context, altitude uint64) (sm.Status, error) {
	s.Lock()
	defer s.Unlock()

	status := sm.Status{
		SuccessionUUID:       s.lc.SuccessionUUID(),
		Edition:       s.edition,
		PrimaryAltitude: s.primaryAltitude,
	}
	if status.PrimaryAltitude == 0 {
		status.PrimaryAltitude = 1
	}

	//
	//
	//
	//
	//
	//
	//
	//
	finalAgileLedger, err := s.lc.ValidateAgileLedgerLocatedAltitude(ctx, int64(altitude), time.Now())
	if err != nil {
		return sm.Status{}, err
	}
	prevailingAgileLedger, err := s.lc.ValidateAgileLedgerLocatedAltitude(ctx, int64(altitude+1), time.Now())
	if err != nil {
		return sm.Status{}, err
	}
	followingAgileLedger, err := s.lc.ValidateAgileLedgerLocatedAltitude(ctx, int64(altitude+2), time.Now())
	if err != nil {
		return sm.Status{}, err
	}

	status.Edition = strongstatus.Edition{
		Agreement: prevailingAgileLedger.Edition,
		Package:  edition.TEMPBaseSemaphoreEdtn,
	}
	status.FinalLedgerAltitude = finalAgileLedger.Altitude
	status.FinalLedgerMoment = finalAgileLedger.Moment
	status.FinalLedgerUUID = finalAgileLedger.Endorse.LedgerUUID
	status.PlatformDigest = prevailingAgileLedger.PlatformDigest
	status.FinalOutcomesDigest = prevailingAgileLedger.FinalOutcomesDigest
	status.FinalAssessors = finalAgileLedger.AssessorAssign
	status.Assessors = prevailingAgileLedger.AssessorAssign
	status.FollowingAssessors = followingAgileLedger.AssessorAssign
	status.FinalAltitudeAssessorsAltered = followingAgileLedger.Altitude

	//
	leadingWebroute, ok := s.suppliers[s.lc.Leading()]
	if !ok || leadingWebroute == "REDACTED" {
		return sm.Status{}, fmt.Errorf("REDACTED")
	}
	leadingIface, err := ifaceCustomer(leadingWebroute)
	if err != nil {
		return sm.Status{}, fmt.Errorf("REDACTED", err)
	}
	customeriface := agileiface.FreshCustomer(leadingIface, s.lc)
	outcome, err := customeriface.AgreementSettings(ctx, &prevailingAgileLedger.Altitude)
	if err != nil {
		return sm.Status{}, fmt.Errorf("REDACTED",
			followingAgileLedger.Altitude, err)
	}
	status.AgreementSettings = outcome.AgreementSettings
	status.FinalAltitudeAgreementParametersAltered = prevailingAgileLedger.Altitude

	return status, nil
}

//
func ifaceCustomer(node string) (*rpchttpsvc.Httpsvc, error) {
	if !strings.Contains(node, "REDACTED") {
		node = "REDACTED" + node
	}
	c, err := rpchttpsvc.New(node, "REDACTED")
	if err != nil {
		return nil, err
	}
	return c, nil
}
