//

package simulations

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	kinds "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
)

//
type ApplicationLinkAgreement struct {
	mock.Simulate
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
func (_m *ApplicationLinkAgreement) Failure() error {
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
func (_m *ApplicationLinkAgreement) BroadenBallot(_a0 context.Context, _a1 *kinds.SolicitBroadenBallot) (*kinds.ReplyBroadenBallot, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyBroadenBallot
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.SolicitBroadenBallot) (*kinds.ReplyBroadenBallot, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.SolicitBroadenBallot) *kinds.ReplyBroadenBallot); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyBroadenBallot)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.SolicitBroadenBallot) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *ApplicationLinkAgreement) CulminateLedger(_a0 context.Context, _a1 *kinds.SolicitCulminateLedger) (*kinds.ReplyCulminateLedger, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyCulminateLedger
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.SolicitCulminateLedger) (*kinds.ReplyCulminateLedger, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.SolicitCulminateLedger) *kinds.ReplyCulminateLedger); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyCulminateLedger)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.SolicitCulminateLedger) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *ApplicationLinkAgreement) InitializeSuccession(_a0 context.Context, _a1 *kinds.SolicitInitializeSuccession) (*kinds.ReplyInitializeSuccession, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyInitializeSuccession
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.SolicitInitializeSuccession) (*kinds.ReplyInitializeSuccession, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.SolicitInitializeSuccession) *kinds.ReplyInitializeSuccession); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyInitializeSuccession)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.SolicitInitializeSuccession) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *ApplicationLinkAgreement) ArrangeNomination(_a0 context.Context, _a1 *kinds.SolicitArrangeNomination) (*kinds.ReplyArrangeNomination, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyArrangeNomination
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.SolicitArrangeNomination) (*kinds.ReplyArrangeNomination, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.SolicitArrangeNomination) *kinds.ReplyArrangeNomination); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyArrangeNomination)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.SolicitArrangeNomination) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *ApplicationLinkAgreement) HandleNomination(_a0 context.Context, _a1 *kinds.SolicitHandleNomination) (*kinds.ReplyHandleNomination, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyHandleNomination
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.SolicitHandleNomination) (*kinds.ReplyHandleNomination, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.SolicitHandleNomination) *kinds.ReplyHandleNomination); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyHandleNomination)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.SolicitHandleNomination) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *ApplicationLinkAgreement) ValidateBallotAddition(_a0 context.Context, _a1 *kinds.SolicitValidateBallotAddition) (*kinds.ReplyValidateBallotAddition, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyValidateBallotAddition
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.SolicitValidateBallotAddition) (*kinds.ReplyValidateBallotAddition, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.SolicitValidateBallotAddition) *kinds.ReplyValidateBallotAddition); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyValidateBallotAddition)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.SolicitValidateBallotAddition) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
//
func FreshApplicationLinkAgreement(t interface {
	mock.TestingT
	Sanitize(func())
}) *ApplicationLinkAgreement {
	simulate := &ApplicationLinkAgreement{}
	mock.Simulate.Test(t)

	t.Sanitize(func() { mock.AssertExpectations(t) })

	return simulate
}
