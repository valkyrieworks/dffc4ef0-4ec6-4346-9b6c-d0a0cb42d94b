//

package simulations

import (
	ifacetypes "github.com/valkyrieworks/iface/kinds"
	mock "github.com/stretchr/testify/mock"

	status "github.com/valkyrieworks/status"

	kinds "github.com/valkyrieworks/kinds"
)

//
type Depot struct {
	mock.Emulate
}

//
func (_m *Depot) Onboard(_a0 status.Status) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(status.Status) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

//
func (_m *Depot) End() error {
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
func (_m *Depot) FetchInactiveStatusAlignLevel() (int64, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func() (int64, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Depot) Import() (status.Status, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 status.Status
	var r1 error
	if rf, ok := ret.Get(0).(func() (status.Status, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() status.Status); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(status.Status)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Depot) ImportAgreementOptions(_a0 int64) (kinds.AgreementOptions, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 kinds.AgreementOptions
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (kinds.AgreementOptions, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(int64) kinds.AgreementOptions); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(kinds.AgreementOptions)
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Depot) ImportCompleteLedgerReply(_a0 int64) (*ifacetypes.ReplyCompleteLedger, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *ifacetypes.ReplyCompleteLedger
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (*ifacetypes.ReplyCompleteLedger, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(int64) *ifacetypes.ReplyCompleteLedger); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ifacetypes.ReplyCompleteLedger)
		}
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Depot) ImportFromStoreOrOriginPaper(_a0 *kinds.OriginPaper) (status.Status, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 status.Status
	var r1 error
	if rf, ok := ret.Get(0).(func(*kinds.OriginPaper) (status.Status, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*kinds.OriginPaper) status.Status); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(status.Status)
	}

	if rf, ok := ret.Get(1).(func(*kinds.OriginPaper) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Depot) ImportFromStoreOrOriginEntry(_a0 string) (status.Status, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 status.Status
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (status.Status, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(string) status.Status); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(status.Status)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Depot) ImportFinalCompleteLedgerReply(_a0 int64) (*ifacetypes.ReplyCompleteLedger, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *ifacetypes.ReplyCompleteLedger
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (*ifacetypes.ReplyCompleteLedger, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(int64) *ifacetypes.ReplyCompleteLedger); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ifacetypes.ReplyCompleteLedger)
		}
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Depot) ImportRatifiers(_a0 int64) (*kinds.RatifierAssign, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.RatifierAssign
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (*kinds.RatifierAssign, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(int64) *kinds.RatifierAssign); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.RatifierAssign)
		}
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Depot) TrimConditions(_a0 int64, _a1 int64, _a2 int64) error {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, int64, int64) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

//
func (_m *Depot) Persist(_a0 status.Status) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(status.Status) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

//
func (_m *Depot) PersistCompleteLedgerReply(_a0 int64, _a1 *ifacetypes.ReplyCompleteLedger) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, *ifacetypes.ReplyCompleteLedger) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

//
func (_m *Depot) CollectionInactiveStatusAlignLevel(level int64) error {
	ret := _m.Called(level)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(level)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

//
//
func NewDepot(t interface {
	mock.TestingT
	Sanitize(func())
}) *Depot {
	emulate := &Depot{}
	mock.Emulate.Test(t)

	t.Sanitize(func() { mock.AssertExpectations(t) })

	return emulate
}
