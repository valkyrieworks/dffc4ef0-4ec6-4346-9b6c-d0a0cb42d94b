//

package simulations

import (
	context "context"

	log "github.com/valkyrieworks/utils/log"

	mock "github.com/stretchr/testify/mock"

	inquire "github.com/valkyrieworks/utils/broadcast/inquire"

	kinds "github.com/valkyrieworks/kinds"
)

//
type LedgerOrdinaler struct {
	mock.Emulate
}

//
func (_m *LedgerOrdinaler) Has(level int64) (bool, error) {
	ret := _m.Called(level)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (bool, error)); ok {
		return rf(level)
	}
	if rf, ok := ret.Get(0).(func(int64) bool); ok {
		r0 = rf(level)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(level)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *LedgerOrdinaler) Ordinal(_a0 kinds.EventDataNewLedgerEvents) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(kinds.EventDataNewLedgerEvents) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

//
func (_m *LedgerOrdinaler) Scan(ctx context.Context, q *inquire.Inquire) ([]int64, error) {
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
func (_m *LedgerOrdinaler) AssignTracer(l log.Tracer) {
	_m.Called(l)
}

//
//
func NewLedgerOrdinaler(t interface {
	mock.TestingT
	Sanitize(func())
}) *LedgerOrdinaler {
	emulate := &LedgerOrdinaler{}
	mock.Emulate.Test(t)

	t.Sanitize(func() { mock.AssertExpectations(t) })

	return emulate
}
