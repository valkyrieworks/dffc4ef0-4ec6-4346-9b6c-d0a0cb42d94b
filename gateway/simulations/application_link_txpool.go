//

package simulations

import (
	context "context"

	abciend "github.com/valkyrieworks/iface/customer"

	mock "github.com/stretchr/testify/mock"

	kinds "github.com/valkyrieworks/iface/kinds"
)

//
type ApplicationLinkTxpool struct {
	mock.Emulate
}

//
func (_m *ApplicationLinkTxpool) InspectTransfer(_a0 context.Context, _a1 *kinds.QueryInspectTransfer) (*kinds.ReplyInspectTransfer, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyInspectTransfer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryInspectTransfer) (*kinds.ReplyInspectTransfer, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryInspectTransfer) *kinds.ReplyInspectTransfer); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyInspectTransfer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.QueryInspectTransfer) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *ApplicationLinkTxpool) InspectTransferAsync(_a0 context.Context, _a1 *kinds.QueryInspectTransfer) (*abciend.RequestOutput, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *abciend.RequestOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryInspectTransfer) (*abciend.RequestOutput, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryInspectTransfer) *abciend.RequestOutput); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abciend.RequestOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.QueryInspectTransfer) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *ApplicationLinkTxpool) Fault() error {
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
func (_m *ApplicationLinkTxpool) Purge(_a0 context.Context) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

//
func (_m *ApplicationLinkTxpool) EmbedTransfer(_a0 context.Context, _a1 *kinds.QueryEmbedTransfer) (*kinds.ReplyEmbedTransfer, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyEmbedTransfer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryEmbedTransfer) (*kinds.ReplyEmbedTransfer, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryEmbedTransfer) *kinds.ReplyEmbedTransfer); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyEmbedTransfer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.QueryEmbedTransfer) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *ApplicationLinkTxpool) HarvestTrans(_a0 context.Context, _a1 *kinds.QueryHarvestTrans) (*kinds.ReplyHarvestTrans, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyHarvestTrans
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryHarvestTrans) (*kinds.ReplyHarvestTrans, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryHarvestTrans) *kinds.ReplyHarvestTrans); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyHarvestTrans)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.QueryHarvestTrans) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *ApplicationLinkTxpool) CollectionReplyCallback(_a0 abciend.Callback) {
	_m.Called(_a0)
}

//
//
func NewApplicationLinkTxpool(t interface {
	mock.TestingT
	Sanitize(func())
}) *ApplicationLinkTxpool {
	emulate := &ApplicationLinkTxpool{}
	mock.Emulate.Test(t)

	t.Sanitize(func() { mock.AssertExpectations(t) })

	return emulate
}
