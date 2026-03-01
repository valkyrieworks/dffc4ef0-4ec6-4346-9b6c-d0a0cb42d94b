package httpsvc

import (
	"context"
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agile/supplier"
	customeriface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/customer"
	rpchttpsvc "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/customer/httpsvc"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

var (
	//
	patternAbsentAltitude = regexp.MustCompile("REDACTED")
	patternExcessivelyTall       = regexp.MustCompile("REDACTED")
	patternScheduledOutput      = regexp.MustCompile("REDACTED")

	maximumReissueEndeavors      = 5
	deadline          uint = 5 //
)

//
type httpsvc struct {
	successionUUID string
	customer  customeriface.DistantCustomer
}

//
//
//
func New(successionUUID, distant string) (supplier.Supplier, error) {
	//
	if !strings.Contains(distant, "REDACTED") {
		distant = "REDACTED" + distant
	}

	httpsvcCustomer, err := rpchttpsvc.FreshUsingDeadline(distant, "REDACTED", deadline)
	if err != nil {
		return nil, err
	}

	return FreshUsingCustomer(successionUUID, httpsvcCustomer), nil
}

//
func FreshUsingCustomer(successionUUID string, customer customeriface.DistantCustomer) supplier.Supplier {
	return &httpsvc{
		customer:  customer,
		successionUUID: successionUUID,
	}
}

//
func (p *httpsvc) SuccessionUUID() string {
	return p.successionUUID
}

func (p *httpsvc) Text() string {
	return fmt.Sprintf("REDACTED", p.customer.Distant())
}

//
//
func (p *httpsvc) AgileLedger(ctx context.Context, altitude int64) (*kinds.AgileLedger, error) {
	h, err := certifyAltitude(altitude)
	if err != nil {
		return nil, supplier.FaultFlawedAgileLedger{Rationale: err}
	}

	sh, err := p.notatedHeadline(ctx, h)
	if err != nil {
		return nil, err
	}

	if altitude != 0 && sh.Altitude != altitude {
		return nil, supplier.FaultFlawedAgileLedger{
			Rationale: fmt.Errorf("REDACTED", sh.Altitude, altitude),
		}
	}

	vs, err := p.assessorAssign(ctx, &sh.Altitude)
	if err != nil {
		return nil, err
	}

	lb := &kinds.AgileLedger{
		NotatedHeading: sh,
		AssessorAssign: vs,
	}

	err = lb.CertifyFundamental(p.successionUUID)
	if err != nil {
		return nil, supplier.FaultFlawedAgileLedger{Rationale: err}
	}

	return lb, nil
}

//
func (p *httpsvc) NotifyProof(ctx context.Context, ev kinds.Proof) error {
	_, err := p.customer.MulticastProof(ctx, ev)
	return err
}

func (p *httpsvc) assessorAssign(ctx context.Context, altitude *int64) (*kinds.AssessorAssign, error) {
	//
	//
	//
	const maximumDisplays = 100

	var (
		everyScreen = 100
		values    = []*kinds.Assessor{}
		screen    = 1
		sum   = -1
	)

EXTERNAL_CYCLE:
	for len(values) != sum && screen <= maximumDisplays {
		for effort := 1; effort <= maximumReissueEndeavors; effort++ {
			res, err := p.customer.Assessors(ctx, altitude, &screen, &everyScreen)
			switch {
			case err == nil:
				//
				if len(res.Assessors) == 0 {
					return nil, supplier.FaultFlawedAgileLedger{
						Rationale: fmt.Errorf("REDACTED",
							altitude, screen, everyScreen),
					}
				}
				if res.Sum <= 0 {
					return nil, supplier.FaultFlawedAgileLedger{
						Rationale: fmt.Errorf("REDACTED",
							res.Sum, altitude, screen, everyScreen),
					}
				}

				sum = res.Sum
				values = append(values, res.Assessors...)
				screen++
				continue EXTERNAL_CYCLE

			case patternExcessivelyTall.MatchString(err.Error()):
				return nil, supplier.FaultAltitudeExcessivelyTall

			case patternAbsentAltitude.MatchString(err.Error()):
				return nil, supplier.FaultAgileLedgerNegationDetected

			//
			case effort == maximumReissueEndeavors:
				return nil, supplier.FaultNegativeReply

			case patternScheduledOutput.MatchString(err.Error()):
				//
				time.Sleep(retreatDeadline(uint16(effort)))
				continue

			//
			default:
				return nil, err
			}

		}
	}

	itemAssign, err := kinds.AssessorAssignOriginatingCurrentAssessors(values)
	if err != nil {
		return nil, supplier.FaultFlawedAgileLedger{Rationale: err}
	}
	return itemAssign, nil
}

func (p *httpsvc) notatedHeadline(ctx context.Context, altitude *int64) (*kinds.NotatedHeading, error) {
	for effort := 1; effort <= maximumReissueEndeavors; effort++ {
		endorse, err := p.customer.Endorse(ctx, altitude)
		switch {
		case err == nil:
			//
			//
			//
			//
			if endorse.EqualsBlank() {
				//
				//
				return nil, supplier.FaultAltitudeExcessivelyTall
			}
			return &endorse.NotatedHeading, nil

		case patternExcessivelyTall.MatchString(err.Error()):
			return nil, supplier.FaultAltitudeExcessivelyTall

		case patternAbsentAltitude.MatchString(err.Error()):
			return nil, supplier.FaultAgileLedgerNegationDetected

		case patternScheduledOutput.MatchString(err.Error()):
			//
			time.Sleep(retreatDeadline(uint16(effort)))
			continue

		//
		default:
			return nil, err
		}
	}
	return nil, supplier.FaultNegativeReply
}

func certifyAltitude(altitude int64) (*int64, error) {
	if altitude < 0 {
		return nil, fmt.Errorf("REDACTED", altitude)
	}

	h := &altitude
	if altitude == 0 {
		h = nil
	}
	return h, nil
}

//
//
func retreatDeadline(effort uint16) time.Duration {
	//
	return time.Duration(500*effort*effort)*time.Millisecond + time.Duration(rand.Intn(1000))*time.Millisecond
}
