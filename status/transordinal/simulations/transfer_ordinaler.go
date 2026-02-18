//

package simulations

import (
	context "context"

	log "github.com/valkyrieworks/utils/log"
	mock "github.com/stretchr/testify/mock"

	inquire "github.com/valkyrieworks/utils/broadcast/inquire"

	transordinal "github.com/valkyrieworks/status/transordinal"

	kinds "github.com/valkyrieworks/iface/kinds"
)

//
type TransOrdinaler struct {
	mock.Emulate
}

//
func (_m *TransOrdinaler) AppendGroup(b *transordinal.Group) error {
	ret := _m.Called(b)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*transordinal.Group) error); ok {
		r0 = rf(b)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

//
func (_m *TransOrdinaler) Get(digest []byte) (*kinds.TransOutcome, error) {
	ret := _m.Called(digest)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.TransOutcome
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte) (*kinds.TransOutcome, error)); ok {
		return rf(digest)
	}
	if rf, ok := ret.Get(0).(func([]byte) *kinds.TransOutcome); ok {
		r0 = rf(digest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.TransOutcome)
		}
	}

	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(digest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *TransOrdinaler) Ordinal(outcome *kinds.TransOutcome) error {
	ret := _m.Called(outcome)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*kinds.TransOutcome) error); ok {
		r0 = rf(outcome)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

//
func (_m *TransOrdinaler) Scan(ctx context.Context, q *inquire.Inquire) ([]*kinds.TransOutcome, error) {
	ret := _m.Called(ctx, q)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 []*kinds.TransOutcome
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *inquire.Inquire) ([]*kinds.TransOutcome, error)); ok {
		return rf(ctx, q)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *inquire.Inquire) []*kinds.TransOutcome); ok {
		r0 = rf(ctx, q)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*kinds.TransOutcome)
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
func (_m *TransOrdinaler) AssignTracer(l log.Tracer) {
	_m.Called(l)
}

//
//
func NewTransferOrdinaler(t interface {
	mock.TestingT
	Sanitize(func())
}) *TransOrdinaler {
	emulate := &TransOrdinaler{}
	mock.Emulate.Test(t)

	t.Sanitize(func() { mock.AssertExpectations(t) })

	return emulate
}
