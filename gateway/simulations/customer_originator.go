//

package simulations

import (
	abciend "github.com/valkyrieworks/iface/customer"
	mock "github.com/stretchr/testify/mock"
)

//
type CustomerOriginator struct {
	mock.Emulate
}

//
func (_m *CustomerOriginator) NewIfaceCustomer() (abciend.Customer, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 abciend.Customer
	var r1 error
	if rf, ok := ret.Get(0).(func() (abciend.Customer, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() abciend.Customer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(abciend.Customer)
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
func NewCustomerOriginator(t interface {
	mock.TestingT
	Sanitize(func())
}) *CustomerOriginator {
	emulate := &CustomerOriginator{}
	mock.Emulate.Test(t)

	t.Sanitize(func() { mock.AssertExpectations(t) })

	return emulate
}
