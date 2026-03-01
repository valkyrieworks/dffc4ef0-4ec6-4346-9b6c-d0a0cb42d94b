//

package simulations

import (
	ifacetypes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	mock "github.com/stretchr/testify/mock"

	status "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status"

	kinds "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

//
type Depot struct {
	mock.Simulate
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
func (_m *Depot) Shutdown() error {
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
func (_m *Depot) ObtainInactiveStatusChronizeAltitude() (int64, error) {
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
func (_m *Depot) Fetch() (status.Status, error) {
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
func (_m *Depot) FetchAgreementParameters(_a0 int64) (kinds.AgreementSettings, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 kinds.AgreementSettings
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (kinds.AgreementSettings, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(int64) kinds.AgreementSettings); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(kinds.AgreementSettings)
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Depot) FetchCulminateLedgerReply(_a0 int64) (*ifacetypes.ReplyCulminateLedger, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *ifacetypes.ReplyCulminateLedger
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (*ifacetypes.ReplyCulminateLedger, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(int64) *ifacetypes.ReplyCulminateLedger); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ifacetypes.ReplyCulminateLedger)
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
func (_m *Depot) FetchOriginatingDatastoreEitherOriginPaper(_a0 *kinds.OriginPaper) (status.Status, error) {
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
func (_m *Depot) FetchOriginatingDatastoreEitherInaugurationRecord(_a0 string) (status.Status, error) {
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
func (_m *Depot) FetchFinalCulminateLedgerReply(_a0 int64) (*ifacetypes.ReplyCulminateLedger, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *ifacetypes.ReplyCulminateLedger
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (*ifacetypes.ReplyCulminateLedger, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(int64) *ifacetypes.ReplyCulminateLedger); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ifacetypes.ReplyCulminateLedger)
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
func (_m *Depot) FetchAssessors(_a0 int64) (*kinds.AssessorAssign, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.AssessorAssign
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (*kinds.AssessorAssign, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(int64) *kinds.AssessorAssign); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.AssessorAssign)
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
func (_m *Depot) TrimStatuses(_a0 int64, _a1 int64, _a2 int64) error {
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
func (_m *Depot) PersistCulminateLedgerReply(_a0 int64, _a1 *ifacetypes.ReplyCulminateLedger) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, *ifacetypes.ReplyCulminateLedger) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

//
func (_m *Depot) AssignInactiveStatusChronizeAltitude(altitude int64) error {
	ret := _m.Called(altitude)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(altitude)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

//
//
func FreshDepot(t interface {
	mock.TestingT
	Sanitize(func())
}) *Depot {
	simulate := &Depot{}
	mock.Simulate.Test(t)

	t.Sanitize(func() { mock.AssertExpectations(t) })

	return simulate
}
