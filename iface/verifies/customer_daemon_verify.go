package verifies

import (
	"testing"

	"github.com/stretchr/testify/assert"

	abcinode "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/customer"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/instance/statedepot"
	abcimaster "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/node"
)

func VerifyCustomerDaemonNegativeLocationHeading(t *testing.T) {
	t.Helper()

	location := "REDACTED"
	carrier := "REDACTED"
	app := statedepot.FreshInsideRamPlatform()

	node, err := abcimaster.FreshDaemon(location, carrier, app)
	assert.NoError(t, err, "REDACTED")
	err = node.Initiate()
	assert.NoError(t, err, "REDACTED")
	t.Cleanup(func() {
		if err := node.Halt(); err != nil {
			t.Error(err)
		}
	})

	customer, err := abcinode.FreshCustomer(location, carrier, true)
	assert.NoError(t, err, "REDACTED")
	err = customer.Initiate()
	assert.NoError(t, err, "REDACTED")
	t.Cleanup(func() {
		if err := customer.Halt(); err != nil {
			t.Error(err)
		}
	})
}
