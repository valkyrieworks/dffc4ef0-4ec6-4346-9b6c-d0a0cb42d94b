//

package simulations

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	kinds "github.com/valkyrieworks/iface/kinds"
)

//
type ApplicationLinkInquire struct {
	mock.Emulate
}

//
func (_m *ApplicationLinkInquire) Replicate(_a0 context.Context, _a1 string) (*kinds.ReplyReverberate, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyReverberate
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*kinds.ReplyReverberate, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *kinds.ReplyReverberate); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyReverberate)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *ApplicationLinkInquire) Fault() error {
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
func (_m *ApplicationLinkInquire) Details(_a0 context.Context, _a1 *kinds.QueryDetails) (*kinds.ReplyDetails, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyDetails
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryDetails) (*kinds.ReplyDetails, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryDetails) *kinds.ReplyDetails); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyDetails)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.QueryDetails) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *ApplicationLinkInquire) Inquire(_a0 context.Context, _a1 *kinds.QueryInquire) (*kinds.ReplyInquire, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyInquire
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryInquire) (*kinds.ReplyInquire, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryInquire) *kinds.ReplyInquire); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyInquire)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.QueryInquire) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
//
func NewApplicationLinkInquire(t interface {
	mock.TestingT
	Sanitize(func())
}) *ApplicationLinkInquire {
	emulate := &ApplicationLinkInquire{}
	mock.Emulate.Test(t)

	t.Sanitize(func() { mock.AssertExpectations(t) })

	return emulate
}
