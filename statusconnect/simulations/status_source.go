//

package simulations

import (
	context "context"

	status "github.com/valkyrieworks/status"
	mock "github.com/stretchr/testify/mock"

	kinds "github.com/valkyrieworks/kinds"
)

//
type StatusSource struct {
	mock.Emulate
}

//
func (_m *StatusSource) ApplicationDigest(ctx context.Context, level uint64) ([]byte, error) {
	ret := _m.Called(ctx, level)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) ([]byte, error)); ok {
		return rf(ctx, level)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) []byte); ok {
		r0 = rf(ctx, level)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, level)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *StatusSource) Endorse(ctx context.Context, level uint64) (*kinds.Endorse, error) {
	ret := _m.Called(ctx, level)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.Endorse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (*kinds.Endorse, error)); ok {
		return rf(ctx, level)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *kinds.Endorse); ok {
		r0 = rf(ctx, level)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.Endorse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, level)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *StatusSource) Status(ctx context.Context, level uint64) (status.Status, error) {
	ret := _m.Called(ctx, level)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 status.Status
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (status.Status, error)); ok {
		return rf(ctx, level)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) status.Status); ok {
		r0 = rf(ctx, level)
	} else {
		r0 = ret.Get(0).(status.Status)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, level)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
//
func NewStatusSource(t interface {
	mock.TestingT
	Sanitize(func())
}) *StatusSource {
	emulate := &StatusSource{}
	mock.Emulate.Test(t)

	t.Sanitize(func() { mock.AssertExpectations(t) })

	return emulate
}
