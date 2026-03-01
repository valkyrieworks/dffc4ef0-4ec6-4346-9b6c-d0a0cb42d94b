//

package simulations

import (
	status "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status"
	mock "github.com/stretchr/testify/mock"

	kinds "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

//
type LedgerDepot struct {
	mock.Simulate
}

//
func (_m *LedgerDepot) Foundation() int64 {
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
func (_m *LedgerDepot) Shutdown() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

//
func (_m *LedgerDepot) EraseNewestLedger() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
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
func (_m *LedgerDepot) FetchFoundationSummary() *kinds.LedgerSummary {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.LedgerSummary
	if rf, ok := ret.Get(0).(func() *kinds.LedgerSummary); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.LedgerSummary)
		}
	}

	return r0
}

//
func (_m *LedgerDepot) FetchLedger(altitude int64) *kinds.Ledger {
	ret := _m.Called(altitude)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.Ledger
	if rf, ok := ret.Get(0).(func(int64) *kinds.Ledger); ok {
		r0 = rf(altitude)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.Ledger)
		}
	}

	return r0
}

//
func (_m *LedgerDepot) FetchLedgerViaDigest(digest []byte) *kinds.Ledger {
	ret := _m.Called(digest)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.Ledger
	if rf, ok := ret.Get(0).(func([]byte) *kinds.Ledger); ok {
		r0 = rf(digest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.Ledger)
		}
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
func (_m *LedgerDepot) FetchLedgerExpandedEndorse(altitude int64) *kinds.ExpandedEndorse {
	ret := _m.Called(altitude)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ExpandedEndorse
	if rf, ok := ret.Get(0).(func(int64) *kinds.ExpandedEndorse); ok {
		r0 = rf(altitude)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ExpandedEndorse)
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
func (_m *LedgerDepot) FetchLedgerSummaryViaDigest(digest []byte) *kinds.LedgerSummary {
	ret := _m.Called(digest)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.LedgerSummary
	if rf, ok := ret.Get(0).(func([]byte) *kinds.LedgerSummary); ok {
		r0 = rf(digest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.LedgerSummary)
		}
	}

	return r0
}

//
func (_m *LedgerDepot) FetchLedgerFragment(altitude int64, ordinal int) *kinds.Fragment {
	ret := _m.Called(altitude, ordinal)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.Fragment
	if rf, ok := ret.Get(0).(func(int64, int) *kinds.Fragment); ok {
		r0 = rf(altitude, ordinal)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.Fragment)
		}
	}

	return r0
}

//
func (_m *LedgerDepot) FetchObservedEndorse(altitude int64) *kinds.Endorse {
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
func (_m *LedgerDepot) TrimLedgers(altitude int64, _a1 status.Status) (uint64, int64, error) {
	ret := _m.Called(altitude, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 uint64
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(int64, status.Status) (uint64, int64, error)); ok {
		return rf(altitude, _a1)
	}
	if rf, ok := ret.Get(0).(func(int64, status.Status) uint64); ok {
		r0 = rf(altitude, _a1)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(int64, status.Status) int64); ok {
		r1 = rf(altitude, _a1)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(int64, status.Status) error); ok {
		r2 = rf(altitude, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

//
func (_m *LedgerDepot) PersistLedger(ledger *kinds.Ledger, ledgerFragments *kinds.FragmentAssign, observedEndorse *kinds.Endorse) {
	_m.Called(ledger, ledgerFragments, observedEndorse)
}

//
func (_m *LedgerDepot) PersistLedgerUsingExpandedEndorse(ledger *kinds.Ledger, ledgerFragments *kinds.FragmentAssign, observedEndorse *kinds.ExpandedEndorse) {
	_m.Called(ledger, ledgerFragments, observedEndorse)
}

//
func (_m *LedgerDepot) Extent() int64 {
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
