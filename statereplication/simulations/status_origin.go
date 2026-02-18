//

package simulations

import (
	context "context"

	status "github.com/valkyrieworks/status"
	mock "github.com/stretchr/testify/mock"

	kinds "github.com/valkyrieworks/kinds"
)

//
type StateProvider struct {
	mock.Mock
}

//
func (_m *StateProvider) AppHash(ctx context.Context, height uint64) ([]byte, error) {
	ret := _m.Called(ctx, height)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) ([]byte, error)); ok {
		return rf(ctx, height)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) []byte); ok {
		r0 = rf(ctx, height)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, height)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *StateProvider) Commit(ctx context.Context, height uint64) (*kinds.Commit, error) {
	ret := _m.Called(ctx, height)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.Commit
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (*kinds.Commit, error)); ok {
		return rf(ctx, height)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *kinds.Commit); ok {
		r0 = rf(ctx, height)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.Commit)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, height)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *StateProvider) State(ctx context.Context, height uint64) (status.State, error) {
	ret := _m.Called(ctx, height)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 status.State
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (status.State, error)); ok {
		return rf(ctx, height)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) status.State); ok {
		r0 = rf(ctx, height)
	} else {
		r0 = ret.Get(0).(status.State)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, height)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
//
func NewStateProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *StateProvider {
	mock := &StateProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
