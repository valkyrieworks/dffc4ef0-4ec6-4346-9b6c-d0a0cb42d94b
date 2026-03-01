//

package simulations

import (
	context "context"

	status "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status"
	mock "github.com/stretchr/testify/mock"

	kinds "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

//
type StatusSupplier struct {
	mock.Simulate
}

//
func (_m *StatusSupplier) PlatformDigest(ctx context.Context, altitude uint64) ([]byte, error) {
	ret := _m.Called(ctx, altitude)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) ([]byte, error)); ok {
		return rf(ctx, altitude)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) []byte); ok {
		r0 = rf(ctx, altitude)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, altitude)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *StatusSupplier) Endorse(ctx context.Context, altitude uint64) (*kinds.Endorse, error) {
	ret := _m.Called(ctx, altitude)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.Endorse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (*kinds.Endorse, error)); ok {
		return rf(ctx, altitude)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *kinds.Endorse); ok {
		r0 = rf(ctx, altitude)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.Endorse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, altitude)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *StatusSupplier) Status(ctx context.Context, altitude uint64) (status.Status, error) {
	ret := _m.Called(ctx, altitude)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 status.Status
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (status.Status, error)); ok {
		return rf(ctx, altitude)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) status.Status); ok {
		r0 = rf(ctx, altitude)
	} else {
		r0 = ret.Get(0).(status.Status)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, altitude)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
//
func FreshStatusSupplier(t interface {
	mock.TestingT
	Sanitize(func())
}) *StatusSupplier {
	simulate := &StatusSupplier{}
	mock.Simulate.Test(t)

	t.Sanitize(func() { mock.AssertExpectations(t) })

	return simulate
}
