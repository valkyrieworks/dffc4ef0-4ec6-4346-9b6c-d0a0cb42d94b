//

package simulations

import (
	octets "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/octets"
	customer "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/customer"

	context "context"

	basetypes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/base/kinds"

	log "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"

	mock "github.com/stretchr/testify/mock"

	kinds "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

//
type Customer struct {
	mock.Simulate
}

//
func (_m *Customer) IfaceDetails(_a0 context.Context) (*basetypes.OutcomeIfaceDetails, error) {
	ret := _m.Called(_a0)

	var r0 *basetypes.OutcomeIfaceDetails
	if rf, ok := ret.Get(0).(func(context.Context) *basetypes.OutcomeIfaceDetails); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeIfaceDetails)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) IfaceInquire(ctx context.Context, route string, data octets.HexadecimalOctets) (*basetypes.OutcomeIfaceInquire, error) {
	ret := _m.Called(ctx, route, data)

	var r0 *basetypes.OutcomeIfaceInquire
	if rf, ok := ret.Get(0).(func(context.Context, string, octets.HexadecimalOctets) *basetypes.OutcomeIfaceInquire); ok {
		r0 = rf(ctx, route, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeIfaceInquire)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, octets.HexadecimalOctets) error); ok {
		r1 = rf(ctx, route, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) IfaceInquireUsingChoices(ctx context.Context, route string, data octets.HexadecimalOctets, choices customer.IfaceInquireChoices) (*basetypes.OutcomeIfaceInquire, error) {
	ret := _m.Called(ctx, route, data, choices)

	var r0 *basetypes.OutcomeIfaceInquire
	if rf, ok := ret.Get(0).(func(context.Context, string, octets.HexadecimalOctets, customer.IfaceInquireChoices) *basetypes.OutcomeIfaceInquire); ok {
		r0 = rf(ctx, route, data, choices)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeIfaceInquire)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, octets.HexadecimalOctets, customer.IfaceInquireChoices) error); ok {
		r1 = rf(ctx, route, data, choices)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) Ledger(ctx context.Context, altitude *int64) (*basetypes.OutcomeLedger, error) {
	ret := _m.Called(ctx, altitude)

	var r0 *basetypes.OutcomeLedger
	if rf, ok := ret.Get(0).(func(context.Context, *int64) *basetypes.OutcomeLedger); ok {
		r0 = rf(ctx, altitude)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeLedger)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *int64) error); ok {
		r1 = rf(ctx, altitude)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) LedgerViaDigest(ctx context.Context, digest []byte) (*basetypes.OutcomeLedger, error) {
	ret := _m.Called(ctx, digest)

	var r0 *basetypes.OutcomeLedger
	if rf, ok := ret.Get(0).(func(context.Context, []byte) *basetypes.OutcomeLedger); ok {
		r0 = rf(ctx, digest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeLedger)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []byte) error); ok {
		r1 = rf(ctx, digest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) LedgerOutcomes(ctx context.Context, altitude *int64) (*basetypes.OutcomeLedgerOutcomes, error) {
	ret := _m.Called(ctx, altitude)

	var r0 *basetypes.OutcomeLedgerOutcomes
	if rf, ok := ret.Get(0).(func(context.Context, *int64) *basetypes.OutcomeLedgerOutcomes); ok {
		r0 = rf(ctx, altitude)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeLedgerOutcomes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *int64) error); ok {
		r1 = rf(ctx, altitude)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) LedgerLookup(ctx context.Context, inquire string, screen *int, everyScreen *int, sequenceVia string) (*basetypes.OutcomeLedgerLookup, error) {
	ret := _m.Called(ctx, inquire, screen, everyScreen, sequenceVia)

	var r0 *basetypes.OutcomeLedgerLookup
	if rf, ok := ret.Get(0).(func(context.Context, string, *int, *int, string) *basetypes.OutcomeLedgerLookup); ok {
		r0 = rf(ctx, inquire, screen, everyScreen, sequenceVia)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeLedgerLookup)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *int, *int, string) error); ok {
		r1 = rf(ctx, inquire, screen, everyScreen, sequenceVia)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) LedgerchainDetails(ctx context.Context, minimumAltitude int64, maximumAltitude int64) (*basetypes.OutcomeLedgerchainDetails, error) {
	ret := _m.Called(ctx, minimumAltitude, maximumAltitude)

	var r0 *basetypes.OutcomeLedgerchainDetails
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) *basetypes.OutcomeLedgerchainDetails); ok {
		r0 = rf(ctx, minimumAltitude, maximumAltitude)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeLedgerchainDetails)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, int64) error); ok {
		r1 = rf(ctx, minimumAltitude, maximumAltitude)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) MulticastProof(_a0 context.Context, _a1 kinds.Proof) (*basetypes.OutcomeMulticastProof, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *basetypes.OutcomeMulticastProof
	if rf, ok := ret.Get(0).(func(context.Context, kinds.Proof) *basetypes.OutcomeMulticastProof); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeMulticastProof)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, kinds.Proof) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) MulticastTransferAsyncronous(_a0 context.Context, _a1 kinds.Tx) (*basetypes.OutcomeMulticastTransfer, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *basetypes.OutcomeMulticastTransfer
	if rf, ok := ret.Get(0).(func(context.Context, kinds.Tx) *basetypes.OutcomeMulticastTransfer); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeMulticastTransfer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, kinds.Tx) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) MulticastTransferEndorse(_a0 context.Context, _a1 kinds.Tx) (*basetypes.OutcomeMulticastTransferEndorse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *basetypes.OutcomeMulticastTransferEndorse
	if rf, ok := ret.Get(0).(func(context.Context, kinds.Tx) *basetypes.OutcomeMulticastTransferEndorse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeMulticastTransferEndorse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, kinds.Tx) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) MulticastTransferChronize(_a0 context.Context, _a1 kinds.Tx) (*basetypes.OutcomeMulticastTransfer, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *basetypes.OutcomeMulticastTransfer
	if rf, ok := ret.Get(0).(func(context.Context, kinds.Tx) *basetypes.OutcomeMulticastTransfer); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeMulticastTransfer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, kinds.Tx) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) InspectTransfer(_a0 context.Context, _a1 kinds.Tx) (*basetypes.OutcomeInspectTransfer, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *basetypes.OutcomeInspectTransfer
	if rf, ok := ret.Get(0).(func(context.Context, kinds.Tx) *basetypes.OutcomeInspectTransfer); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeInspectTransfer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, kinds.Tx) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) Endorse(ctx context.Context, altitude *int64) (*basetypes.OutcomeEndorse, error) {
	ret := _m.Called(ctx, altitude)

	var r0 *basetypes.OutcomeEndorse
	if rf, ok := ret.Get(0).(func(context.Context, *int64) *basetypes.OutcomeEndorse); ok {
		r0 = rf(ctx, altitude)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeEndorse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *int64) error); ok {
		r1 = rf(ctx, altitude)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) AgreementSettings(ctx context.Context, altitude *int64) (*basetypes.OutcomeAgreementParameters, error) {
	ret := _m.Called(ctx, altitude)

	var r0 *basetypes.OutcomeAgreementParameters
	if rf, ok := ret.Get(0).(func(context.Context, *int64) *basetypes.OutcomeAgreementParameters); ok {
		r0 = rf(ctx, altitude)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeAgreementParameters)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *int64) error); ok {
		r1 = rf(ctx, altitude)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) AgreementStatus(_a0 context.Context) (*basetypes.OutcomeAgreementStatus, error) {
	ret := _m.Called(_a0)

	var r0 *basetypes.OutcomeAgreementStatus
	if rf, ok := ret.Get(0).(func(context.Context) *basetypes.OutcomeAgreementStatus); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeAgreementStatus)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) ExportAgreementStatus(_a0 context.Context) (*basetypes.OutcomeExportAgreementStatus, error) {
	ret := _m.Called(_a0)

	var r0 *basetypes.OutcomeExportAgreementStatus
	if rf, ok := ret.Get(0).(func(context.Context) *basetypes.OutcomeExportAgreementStatus); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeExportAgreementStatus)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) Inauguration(_a0 context.Context) (*basetypes.OutcomeInauguration, error) {
	ret := _m.Called(_a0)

	var r0 *basetypes.OutcomeInauguration
	if rf, ok := ret.Get(0).(func(context.Context) *basetypes.OutcomeInauguration); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeInauguration)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) InaugurationSegmented(_a0 context.Context, _a1 uint) (*basetypes.OutcomeInaugurationSegment, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *basetypes.OutcomeInaugurationSegment
	if rf, ok := ret.Get(0).(func(context.Context, uint) *basetypes.OutcomeInaugurationSegment); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeInaugurationSegment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) Heading(ctx context.Context, altitude *int64) (*basetypes.OutcomeHeadline, error) {
	ret := _m.Called(ctx, altitude)

	var r0 *basetypes.OutcomeHeadline
	if rf, ok := ret.Get(0).(func(context.Context, *int64) *basetypes.OutcomeHeadline); ok {
		r0 = rf(ctx, altitude)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeHeadline)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *int64) error); ok {
		r1 = rf(ctx, altitude)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) HeadingViaDigest(ctx context.Context, digest octets.HexadecimalOctets) (*basetypes.OutcomeHeadline, error) {
	ret := _m.Called(ctx, digest)

	var r0 *basetypes.OutcomeHeadline
	if rf, ok := ret.Get(0).(func(context.Context, octets.HexadecimalOctets) *basetypes.OutcomeHeadline); ok {
		r0 = rf(ctx, digest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeHeadline)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, octets.HexadecimalOctets) error); ok {
		r1 = rf(ctx, digest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) Vitality(_a0 context.Context) (*basetypes.OutcomeVitality, error) {
	ret := _m.Called(_a0)

	var r0 *basetypes.OutcomeVitality
	if rf, ok := ret.Get(0).(func(context.Context) *basetypes.OutcomeVitality); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeVitality)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) EqualsActive() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

//
func (_m *Customer) NetworkDetails(_a0 context.Context) (*basetypes.OutcomeNetworkDetails, error) {
	ret := _m.Called(_a0)

	var r0 *basetypes.OutcomeNetworkDetails
	if rf, ok := ret.Get(0).(func(context.Context) *basetypes.OutcomeNetworkDetails); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeNetworkDetails)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) CountPendingTrans(_a0 context.Context) (*basetypes.OutcomePendingTrans, error) {
	ret := _m.Called(_a0)

	var r0 *basetypes.OutcomePendingTrans
	if rf, ok := ret.Get(0).(func(context.Context) *basetypes.OutcomePendingTrans); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomePendingTrans)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) UponRestore() error {
	ret := _m.Called()

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
func (_m *Customer) Exit() <-chan struct{} {
	ret := _m.Called()

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
func (_m *Customer) Restore() error {
	ret := _m.Called()

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
func (_m *Customer) Initiate() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

//
func (_m *Customer) Condition(_a0 context.Context) (*basetypes.OutcomeCondition, error) {
	ret := _m.Called(_a0)

	var r0 *basetypes.OutcomeCondition
	if rf, ok := ret.Get(0).(func(context.Context) *basetypes.OutcomeCondition); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeCondition)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) Halt() error {
	ret := _m.Called()

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

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

//
func (_m *Customer) Listen(ctx context.Context, listener string, inquire string, outputVolume ...int) (<-chan basetypes.OutcomeIncident, error) {
	_va := make([]any, len(outputVolume))
	for _i := range outputVolume {
		_va[_i] = outputVolume[_i]
	}
	var _ca []any
	_ca = append(_ca, ctx, listener, inquire)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 <-chan basetypes.OutcomeIncident
	if rf, ok := ret.Get(0).(func(context.Context, string, string, ...int) <-chan basetypes.OutcomeIncident); ok {
		r0 = rf(ctx, listener, inquire, outputVolume...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan basetypes.OutcomeIncident)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, ...int) error); ok {
		r1 = rf(ctx, listener, inquire, outputVolume...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) Tx(ctx context.Context, digest []byte, ascertain bool) (*basetypes.OutcomeTransfer, error) {
	ret := _m.Called(ctx, digest, ascertain)

	var r0 *basetypes.OutcomeTransfer
	if rf, ok := ret.Get(0).(func(context.Context, []byte, bool) *basetypes.OutcomeTransfer); ok {
		r0 = rf(ctx, digest, ascertain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeTransfer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []byte, bool) error); ok {
		r1 = rf(ctx, digest, ascertain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) TransferLookup(ctx context.Context, inquire string, ascertain bool, screen *int, everyScreen *int, sequenceVia string) (*basetypes.OutcomeTransferLookup, error) {
	ret := _m.Called(ctx, inquire, ascertain, screen, everyScreen, sequenceVia)

	var r0 *basetypes.OutcomeTransferLookup
	if rf, ok := ret.Get(0).(func(context.Context, string, bool, *int, *int, string) *basetypes.OutcomeTransferLookup); ok {
		r0 = rf(ctx, inquire, ascertain, screen, everyScreen, sequenceVia)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeTransferLookup)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, bool, *int, *int, string) error); ok {
		r1 = rf(ctx, inquire, ascertain, screen, everyScreen, sequenceVia)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) PendingTrans(ctx context.Context, threshold *int) (*basetypes.OutcomePendingTrans, error) {
	ret := _m.Called(ctx, threshold)

	var r0 *basetypes.OutcomePendingTrans
	if rf, ok := ret.Get(0).(func(context.Context, *int) *basetypes.OutcomePendingTrans); ok {
		r0 = rf(ctx, threshold)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomePendingTrans)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *int) error); ok {
		r1 = rf(ctx, threshold)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) Unlisten(ctx context.Context, listener string, inquire string) error {
	ret := _m.Called(ctx, listener, inquire)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, listener, inquire)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

//
func (_m *Customer) UnlistenEvery(ctx context.Context, listener string) error {
	ret := _m.Called(ctx, listener)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, listener)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

//
func (_m *Customer) Assessors(ctx context.Context, altitude *int64, screen *int, everyScreen *int) (*basetypes.OutcomeAssessors, error) {
	ret := _m.Called(ctx, altitude, screen, everyScreen)

	var r0 *basetypes.OutcomeAssessors
	if rf, ok := ret.Get(0).(func(context.Context, *int64, *int, *int) *basetypes.OutcomeAssessors); ok {
		r0 = rf(ctx, altitude, screen, everyScreen)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeAssessors)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *int64, *int, *int) error); ok {
		r1 = rf(ctx, altitude, screen, everyScreen)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
