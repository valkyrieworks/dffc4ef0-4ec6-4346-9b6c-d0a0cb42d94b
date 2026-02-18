//

package simulations

import (
	context "context"

	kinds "github.com/valkyrieworks/iface/kinds"
	mock "github.com/stretchr/testify/mock"
)

//
type Software struct {
	mock.Emulate
}

//
func (_m *Software) ExecuteMirrorSegment(_a0 context.Context, _a1 *kinds.QueryExecuteMirrorSegment) (*kinds.ReplyExecuteMirrorSegment, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyExecuteMirrorSegment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryExecuteMirrorSegment) (*kinds.ReplyExecuteMirrorSegment, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryExecuteMirrorSegment) *kinds.ReplyExecuteMirrorSegment); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyExecuteMirrorSegment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.QueryExecuteMirrorSegment) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Software) InspectTransfer(_a0 context.Context, _a1 *kinds.QueryInspectTransfer) (*kinds.ReplyInspectTransfer, error) {
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
func (_m *Software) Endorse(_a0 context.Context, _a1 *kinds.QueryEndorse) (*kinds.ReplyEndorse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyEndorse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryEndorse) (*kinds.ReplyEndorse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryEndorse) *kinds.ReplyEndorse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyEndorse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.QueryEndorse) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Software) ExpandBallot(_a0 context.Context, _a1 *kinds.QueryExpandBallot) (*kinds.ReplyExpandBallot, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyExpandBallot
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryExpandBallot) (*kinds.ReplyExpandBallot, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryExpandBallot) *kinds.ReplyExpandBallot); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyExpandBallot)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.QueryExpandBallot) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Software) CompleteLedger(_a0 context.Context, _a1 *kinds.QueryCompleteLedger) (*kinds.ReplyCompleteLedger, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyCompleteLedger
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryCompleteLedger) (*kinds.ReplyCompleteLedger, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryCompleteLedger) *kinds.ReplyCompleteLedger); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyCompleteLedger)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.QueryCompleteLedger) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Software) Details(_a0 context.Context, _a1 *kinds.QueryDetails) (*kinds.ReplyDetails, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyDetails
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryDetails) (*kinds.ReplyDetails, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryDetails) *kinds.ReplyDetails); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyDetails)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.QueryDetails) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Software) InitSeries(_a0 context.Context, _a1 *kinds.QueryInitSeries) (*kinds.ReplyInitSeries, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyInitSeries
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryInitSeries) (*kinds.ReplyInitSeries, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryInitSeries) *kinds.ReplyInitSeries); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyInitSeries)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.QueryInitSeries) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Software) EmbedTransfer(_a0 context.Context, _a1 *kinds.QueryEmbedTransfer) (*kinds.ReplyEmbedTransfer, error) {
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
func (_m *Software) CatalogMirrors(_a0 context.Context, _a1 *kinds.QueryCatalogMirrors) (*kinds.ReplyCatalogMirrors, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyCatalogMirrors
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryCatalogMirrors) (*kinds.ReplyCatalogMirrors, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryCatalogMirrors) *kinds.ReplyCatalogMirrors); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyCatalogMirrors)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.QueryCatalogMirrors) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Software) ImportMirrorSegment(_a0 context.Context, _a1 *kinds.QueryImportMirrorSegment) (*kinds.ReplyImportMirrorSegment, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyImportMirrorSegment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryImportMirrorSegment) (*kinds.ReplyImportMirrorSegment, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryImportMirrorSegment) *kinds.ReplyImportMirrorSegment); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyImportMirrorSegment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.QueryImportMirrorSegment) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Software) ProposalMirror(_a0 context.Context, _a1 *kinds.QueryProposalMirror) (*kinds.ReplyProposalMirror, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyProposalMirror
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryProposalMirror) (*kinds.ReplyProposalMirror, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryProposalMirror) *kinds.ReplyProposalMirror); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyProposalMirror)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.QueryProposalMirror) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Software) ArrangeNomination(_a0 context.Context, _a1 *kinds.QueryArrangeNomination) (*kinds.ReplyArrangeNomination, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyArrangeNomination
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryArrangeNomination) (*kinds.ReplyArrangeNomination, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryArrangeNomination) *kinds.ReplyArrangeNomination); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyArrangeNomination)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.QueryArrangeNomination) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Software) HandleNomination(_a0 context.Context, _a1 *kinds.QueryHandleNomination) (*kinds.ReplyHandleNomination, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyHandleNomination
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryHandleNomination) (*kinds.ReplyHandleNomination, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryHandleNomination) *kinds.ReplyHandleNomination); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyHandleNomination)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.QueryHandleNomination) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Software) Inquire(_a0 context.Context, _a1 *kinds.QueryInquire) (*kinds.ReplyInquire, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyInquire
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryInquire) (*kinds.ReplyInquire, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryInquire) *kinds.ReplyInquire); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyInquire)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.QueryInquire) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Software) HarvestTrans(_a0 context.Context, _a1 *kinds.QueryHarvestTrans) (*kinds.ReplyHarvestTrans, error) {
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
func (_m *Software) ValidateBallotAddition(_a0 context.Context, _a1 *kinds.QueryValidateBallotAddition) (*kinds.ReplyValidateBallotAddition, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyValidateBallotAddition
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryValidateBallotAddition) (*kinds.ReplyValidateBallotAddition, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryValidateBallotAddition) *kinds.ReplyValidateBallotAddition); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyValidateBallotAddition)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.QueryValidateBallotAddition) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
//
func NewSoftware(t interface {
	mock.TestingT
	Sanitize(func())
}) *Software {
	emulate := &Software{}
	mock.Emulate.Test(t)

	t.Sanitize(func() { mock.AssertExpectations(t) })

	return emulate
}
