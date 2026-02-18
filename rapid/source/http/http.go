package http

import (
	"context"
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/valkyrieworks/rapid/source"
	rpccustomer "github.com/valkyrieworks/rpc/customer"
	rpchttp "github.com/valkyrieworks/rpc/customer/http"
	"github.com/valkyrieworks/kinds"
)

var (
	//
	patternAbsentLevel = regexp.MustCompile("REDACTED")
	patternTooElevated       = regexp.MustCompile("REDACTED")
	patternScheduledOut      = regexp.MustCompile("REDACTED")

	maximumReprocessTries      = 5
	deadline          uint = 5 //
)

//
type http struct {
	ledgerUID string
	customer  rpccustomer.ExternalCustomer
}

//
//
//
func New(ledgerUID, external string) (source.Source, error) {
	//
	if !strings.Contains(external, "REDACTED") {
		external = "REDACTED" + external
	}

	httpCustomer, err := rpchttp.NewWithDeadline(external, "REDACTED", deadline)
	if err != nil {
		return nil, err
	}

	return NewWithCustomer(ledgerUID, httpCustomer), nil
}

//
func NewWithCustomer(ledgerUID string, customer rpccustomer.ExternalCustomer) source.Source {
	return &http{
		customer:  customer,
		ledgerUID: ledgerUID,
	}
}

//
func (p *http) LedgerUID() string {
	return p.ledgerUID
}

func (p *http) String() string {
	return fmt.Sprintf("REDACTED", p.customer.External())
}

//
//
func (p *http) RapidLedger(ctx context.Context, level int64) (*kinds.RapidLedger, error) {
	h, err := certifyLevel(level)
	if err != nil {
		return nil, source.ErrFlawedRapidLedger{Cause: err}
	}

	sh, err := p.attestedHeading(ctx, h)
	if err != nil {
		return nil, err
	}

	if level != 0 && sh.Level != level {
		return nil, source.ErrFlawedRapidLedger{
			Cause: fmt.Errorf("REDACTED", sh.Level, level),
		}
	}

	vs, err := p.ratifierCollection(ctx, &sh.Level)
	if err != nil {
		return nil, err
	}

	lb := &kinds.RapidLedger{
		AttestedHeading: sh,
		RatifierAssign: vs,
	}

	err = lb.CertifySimple(p.ledgerUID)
	if err != nil {
		return nil, source.ErrFlawedRapidLedger{Cause: err}
	}

	return lb, nil
}

//
func (p *http) NotifyProof(ctx context.Context, ev kinds.Proof) error {
	_, err := p.customer.MulticastProof(ctx, ev)
	return err
}

func (p *http) ratifierCollection(ctx context.Context, level *int64) (*kinds.RatifierAssign, error) {
	//
	//
	//
	const maximumSections = 100

	var (
		eachScreen = 100
		values    = []*kinds.Ratifier{}
		screen    = 1
		sum   = -1
	)

EXTERNAL_CYCLE:
	for len(values) != sum && screen <= maximumSections {
		for endeavor := 1; endeavor <= maximumReprocessTries; endeavor++ {
			res, err := p.customer.Ratifiers(ctx, level, &screen, &eachScreen)
			switch {
			case err == nil:
				//
				if len(res.Ratifiers) == 0 {
					return nil, source.ErrFlawedRapidLedger{
						Cause: fmt.Errorf("REDACTED",
							level, screen, eachScreen),
					}
				}
				if res.Sum <= 0 {
					return nil, source.ErrFlawedRapidLedger{
						Cause: fmt.Errorf("REDACTED",
							res.Sum, level, screen, eachScreen),
					}
				}

				sum = res.Sum
				values = append(values, res.Ratifiers...)
				screen++
				continue EXTERNAL_CYCLE

			case patternTooElevated.MatchString(err.Error()):
				return nil, source.ErrLevelTooElevated

			case patternAbsentLevel.MatchString(err.Error()):
				return nil, source.ErrRapidLedgerNegateLocated

			//
			case endeavor == maximumReprocessTries:
				return nil, source.ErrNoReply

			case patternScheduledOut.MatchString(err.Error()):
				//
				time.Sleep(retreatDeadline(uint16(endeavor)))
				continue

			//
			default:
				return nil, err
			}

		}
	}

	valueCollection, err := kinds.RatifierCollectionFromPresentRatifiers(values)
	if err != nil {
		return nil, source.ErrFlawedRapidLedger{Cause: err}
	}
	return valueCollection, nil
}

func (p *http) attestedHeading(ctx context.Context, level *int64) (*kinds.AttestedHeading, error) {
	for endeavor := 1; endeavor <= maximumReprocessTries; endeavor++ {
		endorse, err := p.customer.Endorse(ctx, level)
		switch {
		case err == nil:
			//
			//
			//
			//
			if endorse.IsEmpty() {
				//
				//
				return nil, source.ErrLevelTooElevated
			}
			return &endorse.AttestedHeading, nil

		case patternTooElevated.MatchString(err.Error()):
			return nil, source.ErrLevelTooElevated

		case patternAbsentLevel.MatchString(err.Error()):
			return nil, source.ErrRapidLedgerNegateLocated

		case patternScheduledOut.MatchString(err.Error()):
			//
			time.Sleep(retreatDeadline(uint16(endeavor)))
			continue

		//
		default:
			return nil, err
		}
	}
	return nil, source.ErrNoReply
}

func certifyLevel(level int64) (*int64, error) {
	if level < 0 {
		return nil, fmt.Errorf("REDACTED", level)
	}

	h := &level
	if level == 0 {
		h = nil
	}
	return h, nil
}

//
//
func retreatDeadline(endeavor uint16) time.Duration {
	//
	return time.Duration(500*endeavor*endeavor)*time.Millisecond + time.Duration(rand.Intn(1000))*time.Millisecond
}
