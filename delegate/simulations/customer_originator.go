//

package simulations

import (
	abcicustomer "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/customer"
	mock "github.com/stretchr/testify/mock"
)

//
type CustomerOriginator struct {
	mock.Simulate
}

//
func (_m *CustomerOriginator) FreshIfaceCustomer() (abcicustomer.Customer, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 abcicustomer.Customer
	var r1 error
	if rf, ok := ret.Get(0).(func() (abcicustomer.Customer, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() abcicustomer.Customer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(abcicustomer.Customer)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
//
func FreshCustomerOriginator(t interface {
	mock.TestingT
	Sanitize(func())
}) *CustomerOriginator {
	simulate := &CustomerOriginator{}
	mock.Simulate.Test(t)

	t.Sanitize(func() { mock.AssertExpectations(t) })

	return simulate
}
