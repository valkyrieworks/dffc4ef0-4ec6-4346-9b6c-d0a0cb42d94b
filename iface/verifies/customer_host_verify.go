package verifies

import (
	"testing"

	"github.com/stretchr/testify/assert"

	ifacecustomer "github.com/valkyrieworks/iface/customer"
	"github.com/valkyrieworks/iface/instance/objectdepot"
	ifaceservice "github.com/valkyrieworks/iface/host"
)

func VerifyCustomerHostNoAddressPrefix(t *testing.T) {
	t.Helper()

	address := "REDACTED"
	carrier := "REDACTED"
	app := objectdepot.NewInRamSoftware()

	host, err := ifaceservice.NewHost(address, carrier, app)
	assert.NoError(t, err, "REDACTED")
	err = host.Begin()
	assert.NoError(t, err, "REDACTED")
	t.Cleanup(func() {
		if err := host.Halt(); err != nil {
			t.Error(err)
		}
	})

	customer, err := ifacecustomer.NewCustomer(address, carrier, true)
	assert.NoError(t, err, "REDACTED")
	err = customer.Begin()
	assert.NoError(t, err, "REDACTED")
	t.Cleanup(func() {
		if err := customer.Halt(); err != nil {
			t.Error(err)
		}
	})
}
