//

package simulations

import (
	context "context"

	log "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	mock "github.com/stretchr/testify/mock"

	inquire "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/broadcastlisten/inquire"

	transferordinal "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/transferordinal"

	kinds "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
)

//
type TransferOrdinalizer struct {
	mock.Simulate
}

//
func (_m *TransferOrdinalizer) AppendCluster(b *transferordinal.Cluster) error {
	ret := _m.Called(b)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*transferordinal.Cluster) error); ok {
		r0 = rf(b)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

//
func (_m *TransferOrdinalizer) Get(digest []byte) (*kinds.TransferOutcome, error) {
	ret := _m.Called(digest)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.TransferOutcome
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte) (*kinds.TransferOutcome, error)); ok {
		return rf(digest)
	}
	if rf, ok := ret.Get(0).(func([]byte) *kinds.TransferOutcome); ok {
		r0 = rf(digest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.TransferOutcome)
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
func (_m *TransferOrdinalizer) Ordinal(outcome *kinds.TransferOutcome) error {
	ret := _m.Called(outcome)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*kinds.TransferOutcome) error); ok {
		r0 = rf(outcome)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

//
func (_m *TransferOrdinalizer) Lookup(ctx context.Context, q *inquire.Inquire) ([]*kinds.TransferOutcome, error) {
	ret := _m.Called(ctx, q)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 []*kinds.TransferOutcome
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *inquire.Inquire) ([]*kinds.TransferOutcome, error)); ok {
		return rf(ctx, q)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *inquire.Inquire) []*kinds.TransferOutcome); ok {
		r0 = rf(ctx, q)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*kinds.TransferOutcome)
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
func (_m *TransferOrdinalizer) AssignTracer(l log.Tracer) {
	_m.Called(l)
}

//
//
func FreshTransferOrdinalizer(t interface {
	mock.TestingT
	Sanitize(func())
}) *TransferOrdinalizer {
	simulate := &TransferOrdinalizer{}
	mock.Simulate.Test(t)

	t.Sanitize(func() { mock.AssertExpectations(t) })

	return simulate
}
