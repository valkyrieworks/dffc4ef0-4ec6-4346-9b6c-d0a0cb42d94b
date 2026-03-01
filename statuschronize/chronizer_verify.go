package statuschronize

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	commitchronize "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	netmocks "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p/simulations"
	strongstatus "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/status"
	sschema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/statuschronize"
	strongmindedition "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/edition"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/delegate"
	delegatesimulate "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/delegate/simulations"
	sm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/statuschronize/simulations"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/edition"
)

const verifyApplicationEdition = 9

//
func configureExtendChronizer() (*chronizer, *delegatesimulate.PlatformLinkImage) {
	linkInquire := &delegatesimulate.PlatformLinkInquire{}
	linkImage := &delegatesimulate.PlatformLinkImage{}
	statusSupplier := &simulations.StatusSupplier{}
	statusSupplier.On("REDACTED", mock.Anything, mock.Anything).Return([]byte("REDACTED"), nil)
	cfg := settings.FallbackStatusChronizeSettings()
	chronizer := freshChronizer(*cfg, log.FreshNooperationTracer(), linkImage, linkInquire, statusSupplier, "REDACTED")

	return chronizer, linkImage
}

//
func plainNode(id string) *netmocks.Node {
	node := &netmocks.Node{}
	node.On("REDACTED").Return(p2p.ID(id))
	return node
}

func Chronizertest_Chronizeany(t *testing.T) {
	status := sm.Status{
		SuccessionUUID: "REDACTED",
		Edition: strongstatus.Edition{
			Agreement: strongmindedition.Agreement{
				Ledger: edition.LedgerScheme,
				App:   verifyApplicationEdition,
			},
			Package: edition.TEMPBaseSemaphoreEdtn,
		},

		FinalLedgerAltitude: 1,
		FinalLedgerUUID:     kinds.LedgerUUID{Digest: []byte("REDACTED")},
		FinalLedgerMoment:   time.Now(),
		FinalOutcomesDigest: []byte("REDACTED"),
		PlatformDigest:         []byte("REDACTED"),

		FinalAssessors: &kinds.AssessorAssign{Nominator: &kinds.Assessor{Location: []byte("REDACTED")}},
		Assessors:     &kinds.AssessorAssign{Nominator: &kinds.Assessor{Location: []byte("REDACTED")}},
		FollowingAssessors: &kinds.AssessorAssign{Nominator: &kinds.Assessor{Location: []byte("REDACTED")}},

		AgreementSettings:                  *kinds.FallbackAgreementSettings(),
		FinalAltitudeAgreementParametersAltered: 1,
	}
	endorse := &kinds.Endorse{LedgerUUID: kinds.LedgerUUID{Digest: []byte("REDACTED")}}

	segments := []*segment{
		{Altitude: 1, Layout: 1, Ordinal: 0, Segment: []byte{1, 1, 0}},
		{Altitude: 1, Layout: 1, Ordinal: 1, Segment: []byte{1, 1, 1}},
		{Altitude: 1, Layout: 1, Ordinal: 2, Segment: []byte{1, 1, 2}},
	}
	s := &image{Altitude: 1, Layout: 1, Segments: 3, Digest: []byte{1, 2, 3}}

	statusSupplier := &simulations.StatusSupplier{}
	statusSupplier.On("REDACTED", mock.Anything, uint64(1)).Return(status.PlatformDigest, nil)
	statusSupplier.On("REDACTED", mock.Anything, uint64(2)).Return([]byte("REDACTED"), nil)
	statusSupplier.On("REDACTED", mock.Anything, uint64(1)).Return(endorse, nil)
	statusSupplier.On("REDACTED", mock.Anything, uint64(1)).Return(status, nil)
	linkImage := &delegatesimulate.PlatformLinkImage{}
	linkInquire := &delegatesimulate.PlatformLinkInquire{}

	cfg := settings.FallbackStatusChronizeSettings()
	chronizer := freshChronizer(*cfg, log.FreshNooperationTracer(), linkImage, linkInquire, statusSupplier, "REDACTED")

	//
	_, err := chronizer.AppendSegment(&segment{Altitude: 1, Layout: 1, Ordinal: 0, Segment: []byte{1}})
	require.Error(t, err)

	//
	nodeAN := &netmocks.Node{}
	nodeAN.On("REDACTED").Return(p2p.ID("REDACTED"))
	nodeAN.On("REDACTED", mock.MatchedBy(func(i any) bool {
		e, ok := i.(p2p.Wrapper)
		if !ok {
			return false
		}
		req, ok := e.Signal.(*sschema.ImagesSolicit)
		return ok && e.ConduitUUID == ImageConduit && req != nil
	})).Return(true)
	chronizer.AppendNode(nodeAN)
	nodeAN.AssertExpectations(t)

	nodeBYTE := &netmocks.Node{}
	nodeBYTE.On("REDACTED").Return(p2p.ID("REDACTED"))
	nodeBYTE.On("REDACTED", mock.MatchedBy(func(i any) bool {
		e, ok := i.(p2p.Wrapper)
		if !ok {
			return false
		}
		req, ok := e.Signal.(*sschema.ImagesSolicit)
		return ok && e.ConduitUUID == ImageConduit && req != nil
	})).Return(true)
	chronizer.AppendNode(nodeBYTE)
	nodeBYTE.AssertExpectations(t)

	//
	//
	equalsFresh, err := chronizer.AppendImage(nodeAN, s)
	require.NoError(t, err)
	assert.True(t, equalsFresh)

	equalsFresh, err = chronizer.AppendImage(nodeBYTE, s)
	require.NoError(t, err)
	assert.False(t, equalsFresh)

	equalsFresh, err = chronizer.AppendImage(nodeBYTE, &image{Altitude: 2, Layout: 2, Segments: 3, Digest: []byte{1}})
	require.NoError(t, err)
	assert.True(t, equalsFresh)

	//
	//
	linkImage.On("REDACTED", mock.Anything, &iface.SolicitExtendImage{
		Image: &iface.Image{
			Altitude: 2,
			Layout: 2,
			Segments: 3,
			Digest:   []byte{1},
		},
		PlatformDigest: []byte("REDACTED"),
	}).Return(&iface.ReplyExtendImage{Outcome: iface.Replyextendimage_DECLINE_LAYOUT}, nil)
	linkImage.On("REDACTED", mock.Anything, &iface.SolicitExtendImage{
		Image: &iface.Image{
			Altitude:   s.Altitude,
			Layout:   s.Layout,
			Segments:   s.Segments,
			Digest:     s.Digest,
			Attributes: s.Attributes,
		},
		PlatformDigest: []byte("REDACTED"),
	}).Times(2).Return(&iface.ReplyExtendImage{Outcome: iface.Replyextendimage_EMBRACE}, nil)

	segmentSolicits := make(map[uint32]int)
	segmentSolicitsMutex := commitchronize.Exclusion{}
	uponSegmentSolicit := func(arguments mock.Arguments) {
		e, ok := arguments[0].(p2p.Wrapper)
		require.True(t, ok)
		msg := e.Signal.(*sschema.SegmentSolicit)
		require.EqualValues(t, 1, msg.Altitude)
		require.EqualValues(t, 1, msg.Layout)
		require.LessOrEqual(t, msg.Ordinal, uint32(len(segments)))

		appended, err := chronizer.AppendSegment(segments[msg.Ordinal])
		require.NoError(t, err)
		assert.True(t, appended)

		segmentSolicitsMutex.Lock()
		segmentSolicits[msg.Ordinal]++
		segmentSolicitsMutex.Unlock()
	}
	nodeAN.On("REDACTED", mock.MatchedBy(func(i any) bool {
		e, ok := i.(p2p.Wrapper)
		return ok && e.ConduitUUID == SegmentConduit
	})).Maybe().Run(uponSegmentSolicit).Return(true)
	nodeBYTE.On("REDACTED", mock.MatchedBy(func(i any) bool {
		e, ok := i.(p2p.Wrapper)
		return ok && e.ConduitUUID == SegmentConduit
	})).Maybe().Run(uponSegmentSolicit).Return(true)

	//
	//
	//
	linkImage.On("REDACTED", mock.Anything, &iface.SolicitExecuteImageSegment{
		Ordinal: 2, Segment: []byte{1, 1, 2},
	}).Once().Run(func(arguments mock.Arguments) { time.Sleep(2 * time.Second) }).Return(
		&iface.ReplyExecuteImageSegment{
			Outcome:        iface.Replyapplyimagefragment_REISSUE_IMAGE,
			RetrieveSegments: []uint32{1},
		}, nil)

	linkImage.On("REDACTED", mock.Anything, &iface.SolicitExecuteImageSegment{
		Ordinal: 0, Segment: []byte{1, 1, 0},
	}).Times(2).Return(&iface.ReplyExecuteImageSegment{Outcome: iface.Replyapplyimagefragment_EMBRACE}, nil)
	linkImage.On("REDACTED", mock.Anything, &iface.SolicitExecuteImageSegment{
		Ordinal: 1, Segment: []byte{1, 1, 1},
	}).Times(2).Return(&iface.ReplyExecuteImageSegment{Outcome: iface.Replyapplyimagefragment_EMBRACE}, nil)
	linkImage.On("REDACTED", mock.Anything, &iface.SolicitExecuteImageSegment{
		Ordinal: 2, Segment: []byte{1, 1, 2},
	}).Once().Return(&iface.ReplyExecuteImageSegment{Outcome: iface.Replyapplyimagefragment_EMBRACE}, nil)
	linkInquire.On("REDACTED", mock.Anything, delegate.SolicitDetails).Return(&iface.ReplyDetails{
		PlatformEdition:       verifyApplicationEdition,
		FinalLedgerAltitude:  1,
		FinalLedgerPlatformDigest: []byte("REDACTED"),
	}, nil)

	freshStatus, finalEndorse, err := chronizer.ChronizeSome(0, func() {})
	require.NoError(t, err)

	time.Sleep(50 * time.Millisecond) //

	segmentSolicitsMutex.Lock()
	assert.Equal(t, map[uint32]int{0: 1, 1: 2, 2: 1}, segmentSolicits)
	segmentSolicitsMutex.Unlock()

	anticipateStatus := status

	assert.Equal(t, anticipateStatus, freshStatus)
	assert.Equal(t, endorse, finalEndorse)

	linkImage.AssertExpectations(t)
	linkInquire.AssertExpectations(t)
	nodeAN.AssertExpectations(t)
	nodeBYTE.AssertExpectations(t)
}

func Chronizertest_Chronizeany_unimages(t *testing.T) {
	chronizer, _ := configureExtendChronizer()
	_, _, err := chronizer.ChronizeSome(0, func() {})
	assert.Equal(t, faultNegativeImages, err)
}

func Chronizertest_Chronizeany_cancel(t *testing.T) {
	chronizer, linkImage := configureExtendChronizer()

	s := &image{Altitude: 1, Layout: 1, Segments: 3, Digest: []byte{1, 2, 3}}
	_, err := chronizer.AppendImage(plainNode("REDACTED"), s)
	require.NoError(t, err)
	linkImage.On("REDACTED", mock.Anything, &iface.SolicitExtendImage{
		Image: towardIface(s), PlatformDigest: []byte("REDACTED"),
	}).Once().Return(&iface.ReplyExtendImage{Outcome: iface.Replyextendimage_CANCEL}, nil)

	_, _, err = chronizer.ChronizeSome(0, func() {})
	assert.Equal(t, faultCancel, err)
	linkImage.AssertExpectations(t)
}

func Chronizertest_Chronizeany_decline(t *testing.T) {
	chronizer, linkImage := configureExtendChronizer()

	//
	s22 := &image{Altitude: 2, Layout: 2, Segments: 3, Digest: []byte{1, 2, 3}}
	s12 := &image{Altitude: 1, Layout: 2, Segments: 3, Digest: []byte{1, 2, 3}}
	s11 := &image{Altitude: 1, Layout: 1, Segments: 3, Digest: []byte{1, 2, 3}}
	_, err := chronizer.AppendImage(plainNode("REDACTED"), s22)
	require.NoError(t, err)
	_, err = chronizer.AppendImage(plainNode("REDACTED"), s12)
	require.NoError(t, err)
	_, err = chronizer.AppendImage(plainNode("REDACTED"), s11)
	require.NoError(t, err)

	linkImage.On("REDACTED", mock.Anything, &iface.SolicitExtendImage{
		Image: towardIface(s22), PlatformDigest: []byte("REDACTED"),
	}).Once().Return(&iface.ReplyExtendImage{Outcome: iface.Replyextendimage_DECLINE}, nil)

	linkImage.On("REDACTED", mock.Anything, &iface.SolicitExtendImage{
		Image: towardIface(s12), PlatformDigest: []byte("REDACTED"),
	}).Once().Return(&iface.ReplyExtendImage{Outcome: iface.Replyextendimage_DECLINE}, nil)

	linkImage.On("REDACTED", mock.Anything, &iface.SolicitExtendImage{
		Image: towardIface(s11), PlatformDigest: []byte("REDACTED"),
	}).Once().Return(&iface.ReplyExtendImage{Outcome: iface.Replyextendimage_DECLINE}, nil)

	_, _, err = chronizer.ChronizeSome(0, func() {})
	assert.Equal(t, faultNegativeImages, err)
	linkImage.AssertExpectations(t)
}

func Chronizertest_Chronizeany_decline_layout(t *testing.T) {
	chronizer, linkImage := configureExtendChronizer()

	//
	s22 := &image{Altitude: 2, Layout: 2, Segments: 3, Digest: []byte{1, 2, 3}}
	s12 := &image{Altitude: 1, Layout: 2, Segments: 3, Digest: []byte{1, 2, 3}}
	s11 := &image{Altitude: 1, Layout: 1, Segments: 3, Digest: []byte{1, 2, 3}}
	_, err := chronizer.AppendImage(plainNode("REDACTED"), s22)
	require.NoError(t, err)
	_, err = chronizer.AppendImage(plainNode("REDACTED"), s12)
	require.NoError(t, err)
	_, err = chronizer.AppendImage(plainNode("REDACTED"), s11)
	require.NoError(t, err)

	linkImage.On("REDACTED", mock.Anything, &iface.SolicitExtendImage{
		Image: towardIface(s22), PlatformDigest: []byte("REDACTED"),
	}).Once().Return(&iface.ReplyExtendImage{Outcome: iface.Replyextendimage_DECLINE_LAYOUT}, nil)

	linkImage.On("REDACTED", mock.Anything, &iface.SolicitExtendImage{
		Image: towardIface(s11), PlatformDigest: []byte("REDACTED"),
	}).Once().Return(&iface.ReplyExtendImage{Outcome: iface.Replyextendimage_CANCEL}, nil)

	_, _, err = chronizer.ChronizeSome(0, func() {})
	assert.Equal(t, faultCancel, err)
	linkImage.AssertExpectations(t)
}

func Chronizertest_Chronizeany_decline_originator(t *testing.T) {
	chronizer, linkImage := configureExtendChronizer()

	nodeAN := plainNode("REDACTED")
	nodeBYTE := plainNode("REDACTED")
	nodeCount := plainNode("REDACTED")

	//
	//
	//
	sa := &image{Altitude: 1, Layout: 1, Segments: 3, Digest: []byte{1, 2, 3}}
	sb := &image{Altitude: 2, Layout: 1, Segments: 3, Digest: []byte{1, 2, 3}}
	sc := &image{Altitude: 3, Layout: 1, Segments: 3, Digest: []byte{1, 2, 3}}
	sbc := &image{Altitude: 4, Layout: 1, Segments: 3, Digest: []byte{1, 2, 3}}
	_, err := chronizer.AppendImage(nodeAN, sa)
	require.NoError(t, err)
	_, err = chronizer.AppendImage(nodeBYTE, sb)
	require.NoError(t, err)
	_, err = chronizer.AppendImage(nodeCount, sc)
	require.NoError(t, err)
	_, err = chronizer.AppendImage(nodeBYTE, sbc)
	require.NoError(t, err)
	_, err = chronizer.AppendImage(nodeCount, sbc)
	require.NoError(t, err)

	linkImage.On("REDACTED", mock.Anything, &iface.SolicitExtendImage{
		Image: towardIface(sbc), PlatformDigest: []byte("REDACTED"),
	}).Once().Return(&iface.ReplyExtendImage{Outcome: iface.Replyextendimage_DECLINE_ORIGINATOR}, nil)

	linkImage.On("REDACTED", mock.Anything, &iface.SolicitExtendImage{
		Image: towardIface(sa), PlatformDigest: []byte("REDACTED"),
	}).Once().Return(&iface.ReplyExtendImage{Outcome: iface.Replyextendimage_DECLINE}, nil)

	_, _, err = chronizer.ChronizeSome(0, func() {})
	assert.Equal(t, faultNegativeImages, err)
	linkImage.AssertExpectations(t)
}

func Chronizertest_Chronizeany_ifacefault(t *testing.T) {
	chronizer, linkImage := configureExtendChronizer()

	faultDetonate := errors.New("REDACTED")
	s := &image{Altitude: 1, Layout: 1, Segments: 3, Digest: []byte{1, 2, 3}}
	_, err := chronizer.AppendImage(plainNode("REDACTED"), s)
	require.NoError(t, err)
	linkImage.On("REDACTED", mock.Anything, &iface.SolicitExtendImage{
		Image: towardIface(s), PlatformDigest: []byte("REDACTED"),
	}).Once().Return(nil, faultDetonate)

	_, _, err = chronizer.ChronizeSome(0, func() {})
	assert.True(t, errors.Is(err, faultDetonate))
	linkImage.AssertExpectations(t)
}

func Chronizertest_extendimage(t *testing.T) {
	unfamiliarFault := errors.New("REDACTED")
	detonate := errors.New("REDACTED")

	verifycases := map[string]struct {
		outcome    iface.Replyextendimage_Outcome
		err       error
		anticipateFault error
	}{
		"REDACTED":           {iface.Replyextendimage_EMBRACE, nil, nil},
		"REDACTED":            {iface.Replyextendimage_CANCEL, nil, faultCancel},
		"REDACTED":           {iface.Replyextendimage_DECLINE, nil, faultDeclineImage},
		"REDACTED":    {iface.Replyextendimage_DECLINE_LAYOUT, nil, faultDeclineLayout},
		"REDACTED":    {iface.Replyextendimage_DECLINE_ORIGINATOR, nil, faultDeclineOriginator},
		"REDACTED":          {iface.Replyextendimage_UNFAMILIAR, nil, unfamiliarFault},
		"REDACTED":            {0, detonate, detonate},
		"REDACTED": {9, nil, unfamiliarFault},
	}
	for alias, tc := range verifycases {

		t.Run(alias, func(t *testing.T) {
			chronizer, linkImage := configureExtendChronizer()
			s := &image{Altitude: 1, Layout: 1, Segments: 3, Digest: []byte{1, 2, 3}, reliablePlatformDigest: []byte("REDACTED")}
			linkImage.On("REDACTED", mock.Anything, &iface.SolicitExtendImage{
				Image: towardIface(s),
				PlatformDigest:  []byte("REDACTED"),
			}).Return(&iface.ReplyExtendImage{Outcome: tc.outcome}, tc.err)
			err := chronizer.extendImage(s)
			if tc.anticipateFault == unfamiliarFault {
				require.Error(t, err)
			} else {
				revealed := errors.Unwrap(err)
				if revealed != nil {
					err = revealed
				}
				assert.Equal(t, tc.anticipateFault, err)
			}
		})
	}
}

func Chronizertest_executefragments_Outcomes(t *testing.T) {
	unfamiliarFault := errors.New("REDACTED")
	detonate := errors.New("REDACTED")

	verifycases := map[string]struct {
		outcome    iface.Replyapplyimagefragment_Outcome
		err       error
		anticipateFault error
	}{
		"REDACTED":           {iface.Replyapplyimagefragment_EMBRACE, nil, nil},
		"REDACTED":            {iface.Replyapplyimagefragment_CANCEL, nil, faultCancel},
		"REDACTED":            {iface.Replyapplyimagefragment_REISSUE, nil, nil},
		"REDACTED":   {iface.Replyapplyimagefragment_REISSUE_IMAGE, nil, faultReissueImage},
		"REDACTED":  {iface.Replyapplyimagefragment_DECLINE_IMAGE, nil, faultDeclineImage},
		"REDACTED":          {iface.Replyapplyimagefragment_UNFAMILIAR, nil, unfamiliarFault},
		"REDACTED":            {0, detonate, detonate},
		"REDACTED": {9, nil, unfamiliarFault},
	}
	for alias, tc := range verifycases {

		t.Run(alias, func(t *testing.T) {
			linkInquire := &delegatesimulate.PlatformLinkInquire{}
			linkImage := &delegatesimulate.PlatformLinkImage{}
			statusSupplier := &simulations.StatusSupplier{}
			statusSupplier.On("REDACTED", mock.Anything, mock.Anything).Return([]byte("REDACTED"), nil)

			cfg := settings.FallbackStatusChronizeSettings()
			chronizer := freshChronizer(*cfg, log.FreshNooperationTracer(), linkImage, linkInquire, statusSupplier, "REDACTED")

			content := []byte{1, 2, 3}
			segments, err := freshSegmentStaging(&image{Altitude: 1, Layout: 1, Segments: 1}, "REDACTED")
			require.NoError(t, err)
			_, err = segments.Add(&segment{Altitude: 1, Layout: 1, Ordinal: 0, Segment: content})
			require.NoError(t, err)

			linkImage.On("REDACTED", mock.Anything, &iface.SolicitExecuteImageSegment{
				Ordinal: 0, Segment: content,
			}).Once().Return(&iface.ReplyExecuteImageSegment{Outcome: tc.outcome}, tc.err)
			if tc.outcome == iface.Replyapplyimagefragment_REISSUE {
				linkImage.On("REDACTED", mock.Anything, &iface.SolicitExecuteImageSegment{
					Ordinal: 0, Segment: content,
				}).Once().Return(&iface.ReplyExecuteImageSegment{
					Outcome: iface.Replyapplyimagefragment_EMBRACE,
				}, nil)
			}

			err = chronizer.executeSegments(segments)
			if tc.anticipateFault == unfamiliarFault {
				require.Error(t, err)
			} else {
				revealed := errors.Unwrap(err)
				if revealed != nil {
					err = revealed
				}
				assert.Equal(t, tc.anticipateFault, err)
			}
			linkImage.AssertExpectations(t)
		})
	}
}

func Chronizertest_executefragments_Reclaimfragments(t *testing.T) {
	//
	verifycases := map[string]struct {
		outcome iface.Replyapplyimagefragment_Outcome
	}{
		"REDACTED":          {iface.Replyapplyimagefragment_EMBRACE},
		"REDACTED":           {iface.Replyapplyimagefragment_CANCEL},
		"REDACTED":           {iface.Replyapplyimagefragment_REISSUE},
		"REDACTED":  {iface.Replyapplyimagefragment_REISSUE_IMAGE},
		"REDACTED": {iface.Replyapplyimagefragment_DECLINE_IMAGE},
	}
	for alias, tc := range verifycases {

		t.Run(alias, func(t *testing.T) {
			linkInquire := &delegatesimulate.PlatformLinkInquire{}
			linkImage := &delegatesimulate.PlatformLinkImage{}
			statusSupplier := &simulations.StatusSupplier{}
			statusSupplier.On("REDACTED", mock.Anything, mock.Anything).Return([]byte("REDACTED"), nil)

			cfg := settings.FallbackStatusChronizeSettings()
			chronizer := freshChronizer(*cfg, log.FreshNooperationTracer(), linkImage, linkInquire, statusSupplier, "REDACTED")

			segments, err := freshSegmentStaging(&image{Altitude: 1, Layout: 1, Segments: 3}, "REDACTED")
			require.NoError(t, err)
			appended, err := segments.Add(&segment{Altitude: 1, Layout: 1, Ordinal: 0, Segment: []byte{0}})
			require.True(t, appended)
			require.NoError(t, err)
			appended, err = segments.Add(&segment{Altitude: 1, Layout: 1, Ordinal: 1, Segment: []byte{1}})
			require.True(t, appended)
			require.NoError(t, err)
			appended, err = segments.Add(&segment{Altitude: 1, Layout: 1, Ordinal: 2, Segment: []byte{2}})
			require.True(t, appended)
			require.NoError(t, err)

			//
			linkImage.On("REDACTED", mock.Anything, &iface.SolicitExecuteImageSegment{
				Ordinal: 0, Segment: []byte{0},
			}).Once().Return(&iface.ReplyExecuteImageSegment{Outcome: iface.Replyapplyimagefragment_EMBRACE}, nil)
			linkImage.On("REDACTED", mock.Anything, &iface.SolicitExecuteImageSegment{
				Ordinal: 1, Segment: []byte{1},
			}).Once().Return(&iface.ReplyExecuteImageSegment{Outcome: iface.Replyapplyimagefragment_EMBRACE}, nil)
			linkImage.On("REDACTED", mock.Anything, &iface.SolicitExecuteImageSegment{
				Ordinal: 2, Segment: []byte{2},
			}).Once().Return(&iface.ReplyExecuteImageSegment{
				Outcome:        tc.outcome,
				RetrieveSegments: []uint32{1},
			}, nil)

			//
			//
			//
			go func() {
				chronizer.executeSegments(segments) //
			}()

			time.Sleep(50 * time.Millisecond)
			assert.True(t, segments.Has(0))
			assert.False(t, segments.Has(1))
			assert.True(t, segments.Has(2))
			err = segments.Shutdown()
			require.NoError(t, err)
		})
	}
}

func Chronizertest_executefragments_Declineoriginators(t *testing.T) {
	//
	verifycases := map[string]struct {
		outcome iface.Replyapplyimagefragment_Outcome
	}{
		"REDACTED":          {iface.Replyapplyimagefragment_EMBRACE},
		"REDACTED":           {iface.Replyapplyimagefragment_CANCEL},
		"REDACTED":           {iface.Replyapplyimagefragment_REISSUE},
		"REDACTED":  {iface.Replyapplyimagefragment_REISSUE_IMAGE},
		"REDACTED": {iface.Replyapplyimagefragment_DECLINE_IMAGE},
	}
	for alias, tc := range verifycases {

		t.Run(alias, func(t *testing.T) {
			linkInquire := &delegatesimulate.PlatformLinkInquire{}
			linkImage := &delegatesimulate.PlatformLinkImage{}
			statusSupplier := &simulations.StatusSupplier{}
			statusSupplier.On("REDACTED", mock.Anything, mock.Anything).Return([]byte("REDACTED"), nil)

			cfg := settings.FallbackStatusChronizeSettings()
			chronizer := freshChronizer(*cfg, log.FreshNooperationTracer(), linkImage, linkInquire, statusSupplier, "REDACTED")

			//
			//
			nodeAN := plainNode("REDACTED")
			nodeBYTE := plainNode("REDACTED")
			nodeCount := plainNode("REDACTED")

			s1 := &image{Altitude: 1, Layout: 1, Segments: 3}
			s2 := &image{Altitude: 2, Layout: 1, Segments: 3}
			_, err := chronizer.AppendImage(nodeAN, s1)
			require.NoError(t, err)
			_, err = chronizer.AppendImage(nodeAN, s2)
			require.NoError(t, err)
			_, err = chronizer.AppendImage(nodeBYTE, s1)
			require.NoError(t, err)
			_, err = chronizer.AppendImage(nodeBYTE, s2)
			require.NoError(t, err)
			_, err = chronizer.AppendImage(nodeCount, s1)
			require.NoError(t, err)
			_, err = chronizer.AppendImage(nodeCount, s2)
			require.NoError(t, err)

			segments, err := freshSegmentStaging(s1, "REDACTED")
			require.NoError(t, err)
			appended, err := segments.Add(&segment{Altitude: 1, Layout: 1, Ordinal: 0, Segment: []byte{0}, Originator: nodeAN.ID()})
			require.True(t, appended)
			require.NoError(t, err)
			appended, err = segments.Add(&segment{Altitude: 1, Layout: 1, Ordinal: 1, Segment: []byte{1}, Originator: nodeBYTE.ID()})
			require.True(t, appended)
			require.NoError(t, err)
			appended, err = segments.Add(&segment{Altitude: 1, Layout: 1, Ordinal: 2, Segment: []byte{2}, Originator: nodeCount.ID()})
			require.True(t, appended)
			require.NoError(t, err)

			//
			linkImage.On("REDACTED", mock.Anything, &iface.SolicitExecuteImageSegment{
				Ordinal: 0, Segment: []byte{0}, Originator: "REDACTED",
			}).Once().Return(&iface.ReplyExecuteImageSegment{Outcome: iface.Replyapplyimagefragment_EMBRACE}, nil)
			linkImage.On("REDACTED", mock.Anything, &iface.SolicitExecuteImageSegment{
				Ordinal: 1, Segment: []byte{1}, Originator: "REDACTED",
			}).Once().Return(&iface.ReplyExecuteImageSegment{Outcome: iface.Replyapplyimagefragment_EMBRACE}, nil)
			linkImage.On("REDACTED", mock.Anything, &iface.SolicitExecuteImageSegment{
				Ordinal: 2, Segment: []byte{2}, Originator: "REDACTED",
			}).Once().Return(&iface.ReplyExecuteImageSegment{
				Outcome:        tc.outcome,
				DeclineOriginators: []string{string(nodeBYTE.ID())},
			}, nil)

			//
			if tc.outcome == iface.Replyapplyimagefragment_REISSUE {
				linkImage.On("REDACTED", mock.Anything, &iface.SolicitExecuteImageSegment{
					Ordinal: 2, Segment: []byte{2}, Originator: "REDACTED",
				}).Once().Return(&iface.ReplyExecuteImageSegment{Outcome: iface.Replyapplyimagefragment_EMBRACE}, nil)
			}

			//
			//
			//
			go func() {
				chronizer.executeSegments(segments) //
			}()

			time.Sleep(50 * time.Millisecond)

			s1nodes := chronizer.images.ObtainNodes(s1)
			assert.Len(t, s1nodes, 2)
			assert.EqualValues(t, "REDACTED", s1nodes[0].ID())
			assert.EqualValues(t, "REDACTED", s1nodes[1].ID())

			chronizer.images.ObtainNodes(s1)
			assert.Len(t, s1nodes, 2)
			assert.EqualValues(t, "REDACTED", s1nodes[0].ID())
			assert.EqualValues(t, "REDACTED", s1nodes[1].ID())

			err = segments.Shutdown()
			require.NoError(t, err)
		})
	}
}

func Chronizertest_validateapp(t *testing.T) {
	detonate := errors.New("REDACTED")
	const platformEdition = 9
	applicationEditionDiscrepancyFault := errors.New("REDACTED")
	s := &image{Altitude: 3, Layout: 1, Segments: 5, Digest: []byte{1, 2, 3}, reliablePlatformDigest: []byte("REDACTED")}

	verifycases := map[string]struct {
		reply  *iface.ReplyDetails
		err       error
		anticipateFault error
	}{
		"REDACTED": {&iface.ReplyDetails{
			FinalLedgerAltitude:  3,
			FinalLedgerPlatformDigest: []byte("REDACTED"),
			PlatformEdition:       platformEdition,
		}, nil, nil},
		"REDACTED": {&iface.ReplyDetails{
			FinalLedgerAltitude:  3,
			FinalLedgerPlatformDigest: []byte("REDACTED"),
			PlatformEdition:       2,
		}, nil, applicationEditionDiscrepancyFault},
		"REDACTED": {&iface.ReplyDetails{
			FinalLedgerAltitude:  5,
			FinalLedgerPlatformDigest: []byte("REDACTED"),
			PlatformEdition:       platformEdition,
		}, nil, faultValidateUnsuccessful},
		"REDACTED": {&iface.ReplyDetails{
			FinalLedgerAltitude:  3,
			FinalLedgerPlatformDigest: []byte("REDACTED"),
			PlatformEdition:       platformEdition,
		}, nil, faultValidateUnsuccessful},
		"REDACTED": {nil, detonate, detonate},
	}
	for alias, tc := range verifycases {

		t.Run(alias, func(t *testing.T) {
			linkInquire := &delegatesimulate.PlatformLinkInquire{}
			linkImage := &delegatesimulate.PlatformLinkImage{}
			statusSupplier := &simulations.StatusSupplier{}

			cfg := settings.FallbackStatusChronizeSettings()
			chronizer := freshChronizer(*cfg, log.FreshNooperationTracer(), linkImage, linkInquire, statusSupplier, "REDACTED")

			linkInquire.On("REDACTED", mock.Anything, delegate.SolicitDetails).Return(tc.reply, tc.err)
			err := chronizer.validatePlatform(s, platformEdition)
			revealed := errors.Unwrap(err)
			if revealed != nil {
				err = revealed
			}
			require.Equal(t, tc.anticipateFault, err)
		})
	}
}

func towardIface(s *image) *iface.Image {
	return &iface.Image{
		Altitude:   s.Altitude,
		Layout:   s.Layout,
		Segments:   s.Segments,
		Digest:     s.Digest,
		Attributes: s.Attributes,
	}
}
