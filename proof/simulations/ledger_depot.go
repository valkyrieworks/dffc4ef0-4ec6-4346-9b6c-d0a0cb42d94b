//

package simulations

import (
	kinds "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
	mock "github.com/stretchr/testify/mock"
)

//
type LedgerDepot struct {
	mock.Simulate
}

//
func (_m *LedgerDepot) Altitude() int64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

//
func (_m *LedgerDepot) FetchLedgerEndorse(altitude int64) *kinds.Endorse {
	ret := _m.Called(altitude)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.Endorse
	if rf, ok := ret.Get(0).(func(int64) *kinds.Endorse); ok {
		r0 = rf(altitude)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.Endorse)
		}
	}

	return r0
}

//
func (_m *LedgerDepot) FetchLedgerSummary(altitude int64) *kinds.LedgerSummary {
	ret := _m.Called(altitude)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.LedgerSummary
	if rf, ok := ret.Get(0).(func(int64) *kinds.LedgerSummary); ok {
		r0 = rf(altitude)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.LedgerSummary)
		}
	}

	return r0
}

//
//
func FreshLedgerDepot(t interface {
	mock.TestingT
	Sanitize(func())
}) *LedgerDepot {
	simulate := &LedgerDepot{}
	mock.Simulate.Test(t)

	t.Sanitize(func() { mock.AssertExpectations(t) })

	return simulate
}
