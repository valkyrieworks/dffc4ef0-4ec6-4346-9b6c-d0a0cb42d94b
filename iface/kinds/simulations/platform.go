//

package simulations

import (
	context "context"

	kinds "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	mock "github.com/stretchr/testify/mock"
)

//
type Platform struct {
	mock.Simulate
}

//
func (_m *Platform) ExecuteImageSegment(_a0 context.Context, _a1 *kinds.SolicitExecuteImageSegment) (*kinds.ReplyExecuteImageSegment, error) {
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
func (_m *Platform) InspectTransfer(_a0 context.Context, _a1 *kinds.SolicitInspectTransfer) (*kinds.ReplyInspectTransfer, error) {
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
func (_m *Platform) Endorse(_a0 context.Context, _a1 *kinds.SolicitEndorse) (*kinds.ReplyEndorse, error) {
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
func (_m *Platform) BroadenBallot(_a0 context.Context, _a1 *kinds.SolicitBroadenBallot) (*kinds.ReplyBroadenBallot, error) {
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
func (_m *Platform) CulminateLedger(_a0 context.Context, _a1 *kinds.SolicitCulminateLedger) (*kinds.ReplyCulminateLedger, error) {
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
func (_m *Platform) Details(_a0 context.Context, _a1 *kinds.SolicitDetails) (*kinds.ReplyDetails, error) {
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
func (_m *Platform) InitializeSuccession(_a0 context.Context, _a1 *kinds.SolicitInitializeSuccession) (*kinds.ReplyInitializeSuccession, error) {
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
func (_m *Platform) AppendTransfer(_a0 context.Context, _a1 *kinds.SolicitAppendTransfer) (*kinds.ReplyAppendTransfer, error) {
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
func (_m *Platform) CollectionImages(_a0 context.Context, _a1 *kinds.SolicitCollectionImages) (*kinds.ReplyCatalogImages, error) {
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
func (_m *Platform) FetchImageSegment(_a0 context.Context, _a1 *kinds.SolicitFetchImageSegment) (*kinds.ReplyFetchImageSegment, error) {
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
func (_m *Platform) ExtendImage(_a0 context.Context, _a1 *kinds.SolicitExtendImage) (*kinds.ReplyExtendImage, error) {
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
func (_m *Platform) ArrangeNomination(_a0 context.Context, _a1 *kinds.SolicitArrangeNomination) (*kinds.ReplyArrangeNomination, error) {
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
func (_m *Platform) HandleNomination(_a0 context.Context, _a1 *kinds.SolicitHandleNomination) (*kinds.ReplyHandleNomination, error) {
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
func (_m *Platform) Inquire(_a0 context.Context, _a1 *kinds.SolicitInquire) (*kinds.ReplyInquire, error) {
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
func (_m *Platform) HarvestTrans(_a0 context.Context, _a1 *kinds.SolicitHarvestTrans) (*kinds.ReplyHarvestTrans, error) {
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
func (_m *Platform) ValidateBallotAddition(_a0 context.Context, _a1 *kinds.SolicitValidateBallotAddition) (*kinds.ReplyValidateBallotAddition, error) {
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
func FreshPlatform(t interface {
	mock.TestingT
	Sanitize(func())
}) *Platform {
	simulate := &Platform{}
	mock.Simulate.Test(t)

	t.Sanitize(func() { mock.AssertExpectations(t) })

	return simulate
}
