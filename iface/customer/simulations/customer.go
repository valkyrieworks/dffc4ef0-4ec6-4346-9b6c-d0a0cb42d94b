//

package simulations

import (
	context "context"

	abcicustomer "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/customer"

	log "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"

	mock "github.com/stretchr/testify/mock"

	kinds "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
)

//
type Customer struct {
	mock.Simulate
}

//
func (_m *Customer) ExecuteImageSegment(_a0 context.Context, _a1 *kinds.SolicitExecuteImageSegment) (*kinds.ReplyExecuteImageSegment, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyExecuteImageSegment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.SolicitExecuteImageSegment) (*kinds.ReplyExecuteImageSegment, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.SolicitExecuteImageSegment) *kinds.ReplyExecuteImageSegment); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyExecuteImageSegment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.SolicitExecuteImageSegment) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) InspectTransfer(_a0 context.Context, _a1 *kinds.SolicitInspectTransfer) (*kinds.ReplyInspectTransfer, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyInspectTransfer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.SolicitInspectTransfer) (*kinds.ReplyInspectTransfer, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.SolicitInspectTransfer) *kinds.ReplyInspectTransfer); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyInspectTransfer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.SolicitInspectTransfer) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) InspectTransferAsyncronous(_a0 context.Context, _a1 *kinds.SolicitInspectTransfer) (*abcicustomer.RequestResult, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *abcicustomer.RequestResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.SolicitInspectTransfer) (*abcicustomer.RequestResult, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.SolicitInspectTransfer) *abcicustomer.RequestResult); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abcicustomer.RequestResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.SolicitInspectTransfer) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) Endorse(_a0 context.Context, _a1 *kinds.SolicitEndorse) (*kinds.ReplyEndorse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyEndorse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.SolicitEndorse) (*kinds.ReplyEndorse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.SolicitEndorse) *kinds.ReplyEndorse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyEndorse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.SolicitEndorse) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) Reverberate(_a0 context.Context, _a1 string) (*kinds.ReplyReverberate, error) {
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
func (_m *Customer) Failure() error {
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
func (_m *Customer) BroadenBallot(_a0 context.Context, _a1 *kinds.SolicitBroadenBallot) (*kinds.ReplyBroadenBallot, error) {
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
func (_m *Customer) CulminateLedger(_a0 context.Context, _a1 *kinds.SolicitCulminateLedger) (*kinds.ReplyCulminateLedger, error) {
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
func (_m *Customer) Purge(_a0 context.Context) error {
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
func (_m *Customer) Details(_a0 context.Context, _a1 *kinds.SolicitDetails) (*kinds.ReplyDetails, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyDetails
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.SolicitDetails) (*kinds.ReplyDetails, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.SolicitDetails) *kinds.ReplyDetails); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyDetails)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.SolicitDetails) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) InitializeSuccession(_a0 context.Context, _a1 *kinds.SolicitInitializeSuccession) (*kinds.ReplyInitializeSuccession, error) {
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
func (_m *Customer) AppendTransfer(_a0 context.Context, _a1 *kinds.SolicitAppendTransfer) (*kinds.ReplyAppendTransfer, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyAppendTransfer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.SolicitAppendTransfer) (*kinds.ReplyAppendTransfer, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.SolicitAppendTransfer) *kinds.ReplyAppendTransfer); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyAppendTransfer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.SolicitAppendTransfer) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) EqualsActive() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

//
func (_m *Customer) CollectionImages(_a0 context.Context, _a1 *kinds.SolicitCollectionImages) (*kinds.ReplyCatalogImages, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyCatalogImages
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.SolicitCollectionImages) (*kinds.ReplyCatalogImages, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.SolicitCollectionImages) *kinds.ReplyCatalogImages); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyCatalogImages)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.SolicitCollectionImages) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) FetchImageSegment(_a0 context.Context, _a1 *kinds.SolicitFetchImageSegment) (*kinds.ReplyFetchImageSegment, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyFetchImageSegment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.SolicitFetchImageSegment) (*kinds.ReplyFetchImageSegment, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.SolicitFetchImageSegment) *kinds.ReplyFetchImageSegment); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyFetchImageSegment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.SolicitFetchImageSegment) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) ExtendImage(_a0 context.Context, _a1 *kinds.SolicitExtendImage) (*kinds.ReplyExtendImage, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyExtendImage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.SolicitExtendImage) (*kinds.ReplyExtendImage, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.SolicitExtendImage) *kinds.ReplyExtendImage); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyExtendImage)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.SolicitExtendImage) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) UponRestore() error {
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
func (_m *Customer) UponInitiate() error {
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
func (_m *Customer) UponHalt() {
	_m.Called()
}

//
func (_m *Customer) ArrangeNomination(_a0 context.Context, _a1 *kinds.SolicitArrangeNomination) (*kinds.ReplyArrangeNomination, error) {
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
func (_m *Customer) HandleNomination(_a0 context.Context, _a1 *kinds.SolicitHandleNomination) (*kinds.ReplyHandleNomination, error) {
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
func (_m *Customer) Inquire(_a0 context.Context, _a1 *kinds.SolicitInquire) (*kinds.ReplyInquire, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyInquire
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.SolicitInquire) (*kinds.ReplyInquire, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.SolicitInquire) *kinds.ReplyInquire); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyInquire)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.SolicitInquire) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) Exit() <-chan struct{} {
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
func (_m *Customer) HarvestTrans(_a0 context.Context, _a1 *kinds.SolicitHarvestTrans) (*kinds.ReplyHarvestTrans, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyHarvestTrans
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.SolicitHarvestTrans) (*kinds.ReplyHarvestTrans, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.SolicitHarvestTrans) *kinds.ReplyHarvestTrans); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyHarvestTrans)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.SolicitHarvestTrans) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) Restore() error {
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
func (_m *Customer) AssignTracer(_a0 log.Tracer) {
	_m.Called(_a0)
}

//
func (_m *Customer) AssignReplyClbk(_a0 abcicustomer.Clbk) {
	_m.Called(_a0)
}

//
func (_m *Customer) Initiate() error {
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
func (_m *Customer) Halt() error {
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
func (_m *Customer) Text() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

//
func (_m *Customer) ValidateBallotAddition(_a0 context.Context, _a1 *kinds.SolicitValidateBallotAddition) (*kinds.ReplyValidateBallotAddition, error) {
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
func FreshCustomer(t interface {
	mock.TestingT
	Sanitize(func())
}) *Customer {
	simulate := &Customer{}
	mock.Simulate.Test(t)

	t.Sanitize(func() { mock.AssertExpectations(t) })

	return simulate
}
