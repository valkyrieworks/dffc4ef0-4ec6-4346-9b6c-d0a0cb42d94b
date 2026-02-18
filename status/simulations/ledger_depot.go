//

package simulations

import (
	status "github.com/valkyrieworks/status"
	mock "github.com/stretchr/testify/mock"

	kinds "github.com/valkyrieworks/kinds"
)

//
type LedgerDepot struct {
	mock.Emulate
}

//
func (_m *LedgerDepot) Root() int64 {
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
func (_m *LedgerDepot) End() error {
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
func (_m *LedgerDepot) RemoveNewestLedger() error {
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
func (_m *LedgerDepot) Level() int64 {
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
func (_m *LedgerDepot) ImportRootMeta() *kinds.LedgerMeta {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.LedgerMeta
	if rf, ok := ret.Get(0).(func() *kinds.LedgerMeta); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.LedgerMeta)
		}
	}

	return r0
}

//
func (_m *LedgerDepot) ImportLedger(level int64) *kinds.Ledger {
	ret := _m.Called(level)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.Ledger
	if rf, ok := ret.Get(0).(func(int64) *kinds.Ledger); ok {
		r0 = rf(level)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.Ledger)
		}
	}

	return r0
}

//
func (_m *LedgerDepot) ImportLedgerByDigest(digest []byte) *kinds.Ledger {
	ret := _m.Called(digest)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.Ledger
	if rf, ok := ret.Get(0).(func([]byte) *kinds.Ledger); ok {
		r0 = rf(digest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.Ledger)
		}
	}

	return r0
}

//
func (_m *LedgerDepot) ImportLedgerEndorse(level int64) *kinds.Endorse {
	ret := _m.Called(level)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.Endorse
	if rf, ok := ret.Get(0).(func(int64) *kinds.Endorse); ok {
		r0 = rf(level)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.Endorse)
		}
	}

	return r0
}

//
func (_m *LedgerDepot) ImportLedgerExpandedEndorse(level int64) *kinds.ExpandedEndorse {
	ret := _m.Called(level)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ExpandedEndorse
	if rf, ok := ret.Get(0).(func(int64) *kinds.ExpandedEndorse); ok {
		r0 = rf(level)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ExpandedEndorse)
		}
	}

	return r0
}

//
func (_m *LedgerDepot) ImportLedgerMeta(level int64) *kinds.LedgerMeta {
	ret := _m.Called(level)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.LedgerMeta
	if rf, ok := ret.Get(0).(func(int64) *kinds.LedgerMeta); ok {
		r0 = rf(level)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.LedgerMeta)
		}
	}

	return r0
}

//
func (_m *LedgerDepot) ImportLedgerMetaByDigest(digest []byte) *kinds.LedgerMeta {
	ret := _m.Called(digest)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.LedgerMeta
	if rf, ok := ret.Get(0).(func([]byte) *kinds.LedgerMeta); ok {
		r0 = rf(digest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.LedgerMeta)
		}
	}

	return r0
}

//
func (_m *LedgerDepot) ImportLedgerSegment(level int64, ordinal int) *kinds.Segment {
	ret := _m.Called(level, ordinal)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.Segment
	if rf, ok := ret.Get(0).(func(int64, int) *kinds.Segment); ok {
		r0 = rf(level, ordinal)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.Segment)
		}
	}

	return r0
}

//
func (_m *LedgerDepot) ImportViewedEndorse(level int64) *kinds.Endorse {
	ret := _m.Called(level)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.Endorse
	if rf, ok := ret.Get(0).(func(int64) *kinds.Endorse); ok {
		r0 = rf(level)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.Endorse)
		}
	}

	return r0
}

//
func (_m *LedgerDepot) TrimLedgers(level int64, _a1 status.Status) (uint64, int64, error) {
	ret := _m.Called(level, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 uint64
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(int64, status.Status) (uint64, int64, error)); ok {
		return rf(level, _a1)
	}
	if rf, ok := ret.Get(0).(func(int64, status.Status) uint64); ok {
		r0 = rf(level, _a1)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(int64, status.Status) int64); ok {
		r1 = rf(level, _a1)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(int64, status.Status) error); ok {
		r2 = rf(level, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

//
func (_m *LedgerDepot) PersistLedger(ledger *kinds.Ledger, ledgerSegments *kinds.SegmentCollection, viewedEndorse *kinds.Endorse) {
	_m.Called(ledger, ledgerSegments, viewedEndorse)
}

//
func (_m *LedgerDepot) PersistLedgerWithExpandedEndorse(ledger *kinds.Ledger, ledgerSegments *kinds.SegmentCollection, viewedEndorse *kinds.ExpandedEndorse) {
	_m.Called(ledger, ledgerSegments, viewedEndorse)
}

//
func (_m *LedgerDepot) Volume() int64 {
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
//
func NewLedgerDepot(t interface {
	mock.TestingT
	Sanitize(func())
}) *LedgerDepot {
	emulate := &LedgerDepot{}
	mock.Emulate.Test(t)

	t.Sanitize(func() { mock.AssertExpectations(t) })

	return emulate
}
