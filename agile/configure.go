package agile

import (
	"context"
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agile/supplier"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agile/supplier/httpsvc"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agile/depot"
)

//
//
//
//
//
//
func FreshHttpsvcCustomer(
	ctx context.Context,
	successionUUID string,
	relianceChoices RelianceChoices,
	leadingLocator string,
	attestorsLocators []string,
	reliableDepot depot.Depot,
	choices ...Selection,
) (*Customer, error) {
	suppliers, err := suppliersOriginatingLocators(append(attestorsLocators, leadingLocator), successionUUID)
	if err != nil {
		return nil, err
	}

	return FreshCustomer(
		ctx,
		successionUUID,
		relianceChoices,
		suppliers[len(suppliers)-1],
		suppliers[:len(suppliers)-1],
		reliableDepot,
		choices...)
}

//
//
//
//
//
//
func FreshHttpsvcCustomerOriginatingReliableDepot(
	successionUUID string,
	relyingCycle time.Duration,
	leadingLocator string,
	attestorsLocators []string,
	reliableDepot depot.Depot,
	choices ...Selection,
) (*Customer, error) {
	suppliers, err := suppliersOriginatingLocators(append(attestorsLocators, leadingLocator), successionUUID)
	if err != nil {
		return nil, err
	}

	return FreshCustomerOriginatingReliableDepot(
		successionUUID,
		relyingCycle,
		suppliers[len(suppliers)-1],
		suppliers[:len(suppliers)-1],
		reliableDepot,
		choices...)
}

func suppliersOriginatingLocators(locations []string, successionUUID string) ([]supplier.Supplier, error) {
	suppliers := make([]supplier.Supplier, len(locations))
	for idx, location := range locations {
		p, err := httpsvc.New(successionUUID, location)
		if err != nil {
			return nil, err
		}
		suppliers[idx] = p
	}
	return suppliers, nil
}
