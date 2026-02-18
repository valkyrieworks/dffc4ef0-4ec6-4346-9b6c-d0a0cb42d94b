package rapid

import (
	"context"
	"time"

	"github.com/valkyrieworks/rapid/source"
	"github.com/valkyrieworks/rapid/source/http"
	"github.com/valkyrieworks/rapid/depot"
)

//
//
//
//
//
//
func NewHTTPCustomer(
	ctx context.Context,
	ledgerUID string,
	validateOptions ValidateOptions,
	leadingLocation string,
	attestorsAddresses []string,
	validatedDepot depot.Depot,
	options ...Setting,
) (*Customer, error) {
	sources, err := sourcesFromAddresses(append(attestorsAddresses, leadingLocation), ledgerUID)
	if err != nil {
		return nil, err
	}

	return NewCustomer(
		ctx,
		ledgerUID,
		validateOptions,
		sources[len(sources)-1],
		sources[:len(sources)-1],
		validatedDepot,
		options...)
}

//
//
//
//
//
//
func NewHTTPCustomerFromValidatedDepot(
	ledgerUID string,
	validatingDuration time.Duration,
	leadingLocation string,
	attestorsAddresses []string,
	validatedDepot depot.Depot,
	options ...Setting,
) (*Customer, error) {
	sources, err := sourcesFromAddresses(append(attestorsAddresses, leadingLocation), ledgerUID)
	if err != nil {
		return nil, err
	}

	return NewCustomerFromValidatedDepot(
		ledgerUID,
		validatingDuration,
		sources[len(sources)-1],
		sources[:len(sources)-1],
		validatedDepot,
		options...)
}

func sourcesFromAddresses(locations []string, ledgerUID string) ([]source.Source, error) {
	sources := make([]source.Source, len(locations))
	for idx, location := range locations {
		p, err := http.New(ledgerUID, location)
		if err != nil {
			return nil, err
		}
		sources[idx] = p
	}
	return sources, nil
}
