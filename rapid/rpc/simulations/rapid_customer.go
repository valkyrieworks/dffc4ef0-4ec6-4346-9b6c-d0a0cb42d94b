//

package simulations

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	time "time"

	kinds "github.com/valkyrieworks/kinds"
)

//
type RapidCustomer struct {
	mock.Emulate
}

//
func (_m *RapidCustomer) LedgerUID() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

//
func (_m *RapidCustomer) ValidatedRapidLedger(level int64) (*kinds.RapidLedger, error) {
	ret := _m.Called(level)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.RapidLedger
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (*kinds.RapidLedger, error)); ok {
		return rf(level)
	}
	if rf, ok := ret.Get(0).(func(int64) *kinds.RapidLedger); ok {
		r0 = rf(level)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.RapidLedger)
		}
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(level)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *RapidCustomer) Modify(ctx context.Context, now time.Time) (*kinds.RapidLedger, error) {
	ret := _m.Called(ctx, now)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.RapidLedger
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, time.Time) (*kinds.RapidLedger, error)); ok {
		return rf(ctx, now)
	}
	if rf, ok := ret.Get(0).(func(context.Context, time.Time) *kinds.RapidLedger); ok {
		r0 = rf(ctx, now)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.RapidLedger)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, time.Time) error); ok {
		r1 = rf(ctx, now)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *RapidCustomer) ValidateRapidLedgerAtLevel(ctx context.Context, level int64, now time.Time) (*kinds.RapidLedger, error) {
	ret := _m.Called(ctx, level, now)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.RapidLedger
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, time.Time) (*kinds.RapidLedger, error)); ok {
		return rf(ctx, level, now)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, time.Time) *kinds.RapidLedger); ok {
		r0 = rf(ctx, level, now)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.RapidLedger)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, time.Time) error); ok {
		r1 = rf(ctx, level, now)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
//
func NewRapidCustomer(t interface {
	mock.TestingT
	Sanitize(func())
}) *RapidCustomer {
	emulate := &RapidCustomer{}
	mock.Emulate.Test(t)

	t.Sanitize(func() { mock.AssertExpectations(t) })

	return emulate
}
