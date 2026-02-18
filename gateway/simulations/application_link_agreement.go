//

package simulations

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	kinds "github.com/valkyrieworks/iface/kinds"
)

//
type ApplicationLinkAgreement struct {
	mock.Emulate
}

//
func (_m *ApplicationLinkAgreement) Endorse(_a0 context.Context) (*kinds.ReplyEndorse, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyEndorse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*kinds.ReplyEndorse, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *kinds.ReplyEndorse); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyEndorse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *ApplicationLinkAgreement) Fault() error {
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
func (_m *ApplicationLinkAgreement) ExpandBallot(_a0 context.Context, _a1 *kinds.QueryExpandBallot) (*kinds.ReplyExpandBallot, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyExpandBallot
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryExpandBallot) (*kinds.ReplyExpandBallot, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryExpandBallot) *kinds.ReplyExpandBallot); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyExpandBallot)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.QueryExpandBallot) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *ApplicationLinkAgreement) CompleteLedger(_a0 context.Context, _a1 *kinds.QueryCompleteLedger) (*kinds.ReplyCompleteLedger, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyCompleteLedger
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryCompleteLedger) (*kinds.ReplyCompleteLedger, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryCompleteLedger) *kinds.ReplyCompleteLedger); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyCompleteLedger)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.QueryCompleteLedger) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *ApplicationLinkAgreement) InitSeries(_a0 context.Context, _a1 *kinds.QueryInitSeries) (*kinds.ReplyInitSeries, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyInitSeries
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryInitSeries) (*kinds.ReplyInitSeries, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryInitSeries) *kinds.ReplyInitSeries); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyInitSeries)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.QueryInitSeries) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *ApplicationLinkAgreement) ArrangeNomination(_a0 context.Context, _a1 *kinds.QueryArrangeNomination) (*kinds.ReplyArrangeNomination, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyArrangeNomination
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryArrangeNomination) (*kinds.ReplyArrangeNomination, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryArrangeNomination) *kinds.ReplyArrangeNomination); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyArrangeNomination)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.QueryArrangeNomination) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *ApplicationLinkAgreement) HandleNomination(_a0 context.Context, _a1 *kinds.QueryHandleNomination) (*kinds.ReplyHandleNomination, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyHandleNomination
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryHandleNomination) (*kinds.ReplyHandleNomination, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryHandleNomination) *kinds.ReplyHandleNomination); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyHandleNomination)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.QueryHandleNomination) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *ApplicationLinkAgreement) ValidateBallotAddition(_a0 context.Context, _a1 *kinds.QueryValidateBallotAddition) (*kinds.ReplyValidateBallotAddition, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyValidateBallotAddition
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryValidateBallotAddition) (*kinds.ReplyValidateBallotAddition, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryValidateBallotAddition) *kinds.ReplyValidateBallotAddition); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyValidateBallotAddition)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.QueryValidateBallotAddition) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
//
func NewApplicationLinkAgreement(t interface {
	mock.TestingT
	Sanitize(func())
}) *ApplicationLinkAgreement {
	emulate := &ApplicationLinkAgreement{}
	mock.Emulate.Test(t)

	t.Sanitize(func() { mock.AssertExpectations(t) })

	return emulate
}
