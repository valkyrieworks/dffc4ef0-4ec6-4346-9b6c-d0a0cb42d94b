//

package simulations

import (
	ifacetypes "github.com/valkyrieworks/iface/kinds"
	txpool "github.com/valkyrieworks/txpool"

	mock "github.com/stretchr/testify/mock"

	kinds "github.com/valkyrieworks/kinds"
)

//
type Txpool struct {
	mock.Emulate
}

//
func (_m *Txpool) InspectTransfer(tx kinds.Tx, callback func(*ifacetypes.ReplyInspectTransfer), transferDetails txpool.TransferDetails) error {
	ret := _m.Called(tx, callback, transferDetails)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(kinds.Tx, func(*ifacetypes.ReplyInspectTransfer), txpool.TransferDetails) error); ok {
		r0 = rf(tx, callback, transferDetails)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

//
func (_m *Txpool) ActivateTransAccessible() {
	_m.Called()
}

//
func (_m *Txpool) Purge() {
	_m.Called()
}

//
func (_m *Txpool) PurgeApplicationLink() error {
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
func (_m *Txpool) Secure() {
	_m.Called()
}

//
func (_m *Txpool) HarvestMaximumOctetsMaximumFuel(maximumOctets int64, maximumFuel int64) kinds.Txs {
	ret := _m.Called(maximumOctets, maximumFuel)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 kinds.Txs
	if rf, ok := ret.Get(0).(func(int64, int64) kinds.Txs); ok {
		r0 = rf(maximumOctets, maximumFuel)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(kinds.Txs)
		}
	}

	return r0
}

//
func (_m *Txpool) HarvestMaximumTrans(max int) kinds.Txs {
	ret := _m.Called(max)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 kinds.Txs
	if rf, ok := ret.Get(0).(func(int) kinds.Txs); ok {
		r0 = rf(max)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(kinds.Txs)
		}
	}

	return r0
}

//
func (_m *Txpool) DeleteTransferByKey(transferKey kinds.TransferKey) error {
	ret := _m.Called(transferKey)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(kinds.TransferKey) error); ok {
		r0 = rf(transferKey)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

//
func (_m *Txpool) Volume() int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

//
func (_m *Txpool) VolumeOctets() int64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

//
func (_m *Txpool) TransAccessible() <-chan struct{} {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 <-chan struct{}
	if rf, ok := ret.Get(0).(func() <-chan struct{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	return r0
}

//
func (_m *Txpool) Release() {
	_m.Called()
}

//
func (_m *Txpool) Modify(ledgerLevel int64, ledgerTrans kinds.Txs, dispatchTransferReplies []*ifacetypes.InvokeTransferOutcome, newPreFn txpool.PreInspectFunction, newSubmitFn txpool.SubmitInspectFunction) error {
	ret := _m.Called(ledgerLevel, ledgerTrans, dispatchTransferReplies, newPreFn, newSubmitFn)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, kinds.Txs, []*ifacetypes.InvokeTransferOutcome, txpool.PreInspectFunction, txpool.SubmitInspectFunction) error); ok {
		r0 = rf(ledgerLevel, ledgerTrans, dispatchTransferReplies, newPreFn, newSubmitFn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

//
//
func NewTxpool(t interface {
	mock.TestingT
	Sanitize(func())
}) *Txpool {
	emulate := &Txpool{}
	mock.Emulate.Test(t)

	t.Sanitize(func() { mock.AssertExpectations(t) })

	return emulate
}
