//

package simulations

import (
	context "context"

	log "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"

	mock "github.com/stretchr/testify/mock"

	inquire "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/broadcastlisten/inquire"

	kinds "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

//
type LedgerOrdinalizer struct {
	mock.Simulate
}

//
func (_m *LedgerOrdinalizer) Has(altitude int64) (bool, error) {
	ret := _m.Called(altitude)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (bool, error)); ok {
		return rf(altitude)
	}
	if rf, ok := ret.Get(0).(func(int64) bool); ok {
		r0 = rf(altitude)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(altitude)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *LedgerOrdinalizer) Ordinal(_a0 kinds.IncidentDataFreshLedgerIncidents) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(kinds.IncidentDataFreshLedgerIncidents) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

//
func (_m *LedgerOrdinalizer) Lookup(ctx context.Context, q *inquire.Inquire) ([]int64, error) {
	ret := _m.Called(ctx, q)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 []int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *inquire.Inquire) ([]int64, error)); ok {
		return rf(ctx, q)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *inquire.Inquire) []int64); ok {
		r0 = rf(ctx, q)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int64)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *inquire.Inquire) error); ok {
		r1 = rf(ctx, q)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *LedgerOrdinalizer) AssignTracer(l log.Tracer) {
	_m.Called(l)
}

//
//
func FreshLedgerOrdinalizer(t interface {
	mock.TestingT
	Sanitize(func())
}) *LedgerOrdinalizer {
	simulate := &LedgerOrdinalizer{}
	mock.Simulate.Test(t)

	t.Sanitize(func() { mock.AssertExpectations(t) })

	return simulate
}
