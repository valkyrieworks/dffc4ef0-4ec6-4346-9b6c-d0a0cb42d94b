package statusconnect

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	iface "github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/utils/log"
	engineconnect "github.com/valkyrieworks/utils/align"
	"github.com/valkyrieworks/p2p"
	netpeersims "github.com/valkyrieworks/p2p/simulations"
	cometstatus "github.com/valkyrieworks/schema/consensuscore/status"
	statusproto "github.com/valkyrieworks/schema/consensuscore/statusconnect"
	cometrelease "github.com/valkyrieworks/schema/consensuscore/release"
	"github.com/valkyrieworks/gateway"
	gatewaysims "github.com/valkyrieworks/gateway/simulations"
	sm "github.com/valkyrieworks/status"
	"github.com/valkyrieworks/statusconnect/simulations"
	"github.com/valkyrieworks/kinds"
	"github.com/valkyrieworks/release"
)

const verifyApplicationRelease = 9

//
func configureProposalAligner() (*aligner, *gatewaysims.ApplicationLinkMirror) {
	linkInquire := &gatewaysims.ApplicationLinkInquire{}
	linkMirror := &gatewaysims.ApplicationLinkMirror{}
	statusSource := &simulations.StatusSource{}
	statusSource.On("REDACTED", mock.Anything, mock.Anything).Return([]byte("REDACTED"), nil)
	cfg := settings.StandardStatusAlignSettings()
	aligner := newAligner(*cfg, log.NewNoopTracer(), linkMirror, linkInquire, statusSource, "REDACTED")

	return aligner, linkMirror
}

//
func basicNode(id string) *netpeersims.Node {
	node := &netpeersims.Node{}
	node.On("REDACTED").Return(p2p.ID(id))
	return node
}

func Verifyaligner_Alignany(t *testing.T) {
	status := sm.Status{
		LedgerUID: "REDACTED",
		Release: cometstatus.Release{
			Agreement: cometrelease.Agreement{
				Ledger: release.LedgerProtocol,
				App:   verifyApplicationRelease,
			},
			Software: release.TMCoreSemaphoreRev,
		},

		FinalLedgerLevel: 1,
		FinalLedgerUID:     kinds.LedgerUID{Digest: []byte("REDACTED")},
		FinalLedgerTime:   time.Now(),
		FinalOutcomesDigest: []byte("REDACTED"),
		ApplicationDigest:         []byte("REDACTED"),

		FinalRatifiers: &kinds.RatifierAssign{Recommender: &kinds.Ratifier{Location: []byte("REDACTED")}},
		Ratifiers:     &kinds.RatifierAssign{Recommender: &kinds.Ratifier{Location: []byte("REDACTED")}},
		FollowingRatifiers: &kinds.RatifierAssign{Recommender: &kinds.Ratifier{Location: []byte("REDACTED")}},

		AgreementOptions:                  *kinds.StandardAgreementOptions(),
		FinalLevelAgreementOptionsModified: 1,
	}
	endorse := &kinds.Endorse{LedgerUID: kinds.LedgerUID{Digest: []byte("REDACTED")}}

	segments := []*segment{
		{Level: 1, Layout: 1, Ordinal: 0, Segment: []byte{1, 1, 0}},
		{Level: 1, Layout: 1, Ordinal: 1, Segment: []byte{1, 1, 1}},
		{Level: 1, Layout: 1, Ordinal: 2, Segment: []byte{1, 1, 2}},
	}
	s := &mirror{Level: 1, Layout: 1, Segments: 3, Digest: []byte{1, 2, 3}}

	statusSource := &simulations.StatusSource{}
	statusSource.On("REDACTED", mock.Anything, uint64(1)).Return(status.ApplicationDigest, nil)
	statusSource.On("REDACTED", mock.Anything, uint64(2)).Return([]byte("REDACTED"), nil)
	statusSource.On("REDACTED", mock.Anything, uint64(1)).Return(endorse, nil)
	statusSource.On("REDACTED", mock.Anything, uint64(1)).Return(status, nil)
	linkMirror := &gatewaysims.ApplicationLinkMirror{}
	linkInquire := &gatewaysims.ApplicationLinkInquire{}

	cfg := settings.StandardStatusAlignSettings()
	aligner := newAligner(*cfg, log.NewNoopTracer(), linkMirror, linkInquire, statusSource, "REDACTED")

	//
	_, err := aligner.AppendSegment(&segment{Level: 1, Layout: 1, Ordinal: 0, Segment: []byte{1}})
	require.Error(t, err)

	//
	nodeA := &netpeersims.Node{}
	nodeA.On("REDACTED").Return(p2p.ID("REDACTED"))
	nodeA.On("REDACTED", mock.MatchedBy(func(i any) bool {
		e, ok := i.(p2p.Packet)
		if !ok {
			return false
		}
		req, ok := e.Signal.(*statusproto.MirrorsQuery)
		return ok && e.StreamUID == MirrorStream && req != nil
	})).Return(true)
	aligner.AppendNode(nodeA)
	nodeA.AssertExpectations(t)

	nodeBYTE := &netpeersims.Node{}
	nodeBYTE.On("REDACTED").Return(p2p.ID("REDACTED"))
	nodeBYTE.On("REDACTED", mock.MatchedBy(func(i any) bool {
		e, ok := i.(p2p.Packet)
		if !ok {
			return false
		}
		req, ok := e.Signal.(*statusproto.MirrorsQuery)
		return ok && e.StreamUID == MirrorStream && req != nil
	})).Return(true)
	aligner.AppendNode(nodeBYTE)
	nodeBYTE.AssertExpectations(t)

	//
	//
	isNew, err := aligner.AppendMirror(nodeA, s)
	require.NoError(t, err)
	assert.True(t, isNew)

	isNew, err = aligner.AppendMirror(nodeBYTE, s)
	require.NoError(t, err)
	assert.False(t, isNew)

	isNew, err = aligner.AppendMirror(nodeBYTE, &mirror{Level: 2, Layout: 2, Segments: 3, Digest: []byte{1}})
	require.NoError(t, err)
	assert.True(t, isNew)

	//
	//
	linkMirror.On("REDACTED", mock.Anything, &iface.QueryProposalMirror{
		Mirror: &iface.Mirror{
			Level: 2,
			Layout: 2,
			Segments: 3,
			Digest:   []byte{1},
		},
		ApplicationDigest: []byte("REDACTED"),
	}).Return(&iface.ReplyProposalMirror{Outcome: iface.Replymirrorsnapshot_DECLINE_LAYOUT}, nil)
	linkMirror.On("REDACTED", mock.Anything, &iface.QueryProposalMirror{
		Mirror: &iface.Mirror{
			Level:   s.Level,
			Layout:   s.Layout,
			Segments:   s.Segments,
			Digest:     s.Digest,
			Metainfo: s.Metainfo,
		},
		ApplicationDigest: []byte("REDACTED"),
	}).Times(2).Return(&iface.ReplyProposalMirror{Outcome: iface.Replymirrorsnapshot_ALLOW}, nil)

	segmentQueries := make(map[uint32]int)
	segmentQueriesMutex := engineconnect.Lock{}
	onSegmentQuery := func(args mock.Arguments) {
		e, ok := args[0].(p2p.Packet)
		require.True(t, ok)
		msg := e.Signal.(*statusproto.SegmentQuery)
		require.EqualValues(t, 1, msg.Level)
		require.EqualValues(t, 1, msg.Layout)
		require.LessOrEqual(t, msg.Ordinal, uint32(len(segments)))

		appended, err := aligner.AppendSegment(segments[msg.Ordinal])
		require.NoError(t, err)
		assert.True(t, appended)

		segmentQueriesMutex.Lock()
		segmentQueries[msg.Ordinal]++
		segmentQueriesMutex.Unlock()
	}
	nodeA.On("REDACTED", mock.MatchedBy(func(i any) bool {
		e, ok := i.(p2p.Packet)
		return ok && e.StreamUID == SegmentStream
	})).Maybe().Run(onSegmentQuery).Return(true)
	nodeBYTE.On("REDACTED", mock.MatchedBy(func(i any) bool {
		e, ok := i.(p2p.Packet)
		return ok && e.StreamUID == SegmentStream
	})).Maybe().Run(onSegmentQuery).Return(true)

	//
	//
	//
	linkMirror.On("REDACTED", mock.Anything, &iface.QueryExecuteMirrorSegment{
		Ordinal: 2, Segment: []byte{1, 1, 2},
	}).Once().Run(func(args mock.Arguments) { time.Sleep(2 * time.Second) }).Return(
		&iface.ReplyExecuteMirrorSegment{
			Outcome:        iface.Replyexecutemirrorsegment_REPROCESS_MIRROR,
			ReacquireSegments: []uint32{1},
		}, nil)

	linkMirror.On("REDACTED", mock.Anything, &iface.QueryExecuteMirrorSegment{
		Ordinal: 0, Segment: []byte{1, 1, 0},
	}).Times(2).Return(&iface.ReplyExecuteMirrorSegment{Outcome: iface.Replyexecutemirrorsegment_ALLOW}, nil)
	linkMirror.On("REDACTED", mock.Anything, &iface.QueryExecuteMirrorSegment{
		Ordinal: 1, Segment: []byte{1, 1, 1},
	}).Times(2).Return(&iface.ReplyExecuteMirrorSegment{Outcome: iface.Replyexecutemirrorsegment_ALLOW}, nil)
	linkMirror.On("REDACTED", mock.Anything, &iface.QueryExecuteMirrorSegment{
		Ordinal: 2, Segment: []byte{1, 1, 2},
	}).Once().Return(&iface.ReplyExecuteMirrorSegment{Outcome: iface.Replyexecutemirrorsegment_ALLOW}, nil)
	linkInquire.On("REDACTED", mock.Anything, gateway.QueryDetails).Return(&iface.ReplyDetails{
		ApplicationRelease:       verifyApplicationRelease,
		FinalLedgerLevel:  1,
		FinalLedgerApplicationDigest: []byte("REDACTED"),
	}, nil)

	newStatus, finalEndorse, err := aligner.AlignAny(0, func() {})
	require.NoError(t, err)

	time.Sleep(50 * time.Millisecond) //

	segmentQueriesMutex.Lock()
	assert.Equal(t, map[uint32]int{0: 1, 1: 2, 2: 1}, segmentQueries)
	segmentQueriesMutex.Unlock()

	anticipateStatus := status

	assert.Equal(t, anticipateStatus, newStatus)
	assert.Equal(t, endorse, finalEndorse)

	linkMirror.AssertExpectations(t)
	linkInquire.AssertExpectations(t)
	nodeA.AssertExpectations(t)
	nodeBYTE.AssertExpectations(t)
}

func Verifyaligner_Alignany_nomirrors(t *testing.T) {
	aligner, _ := configureProposalAligner()
	_, _, err := aligner.AlignAny(0, func() {})
	assert.Equal(t, errNoMirrors, err)
}

func Verifyaligner_Alignany_cancel(t *testing.T) {
	aligner, linkMirror := configureProposalAligner()

	s := &mirror{Level: 1, Layout: 1, Segments: 3, Digest: []byte{1, 2, 3}}
	_, err := aligner.AppendMirror(basicNode("REDACTED"), s)
	require.NoError(t, err)
	linkMirror.On("REDACTED", mock.Anything, &iface.QueryProposalMirror{
		Mirror: toIface(s), ApplicationDigest: []byte("REDACTED"),
	}).Once().Return(&iface.ReplyProposalMirror{Outcome: iface.Replymirrorsnapshot_CANCEL}, nil)

	_, _, err = aligner.AlignAny(0, func() {})
	assert.Equal(t, errCancel, err)
	linkMirror.AssertExpectations(t)
}

func Verifyaligner_Alignany_decline(t *testing.T) {
	aligner, linkMirror := configureProposalAligner()

	//
	s22 := &mirror{Level: 2, Layout: 2, Segments: 3, Digest: []byte{1, 2, 3}}
	s12 := &mirror{Level: 1, Layout: 2, Segments: 3, Digest: []byte{1, 2, 3}}
	s11 := &mirror{Level: 1, Layout: 1, Segments: 3, Digest: []byte{1, 2, 3}}
	_, err := aligner.AppendMirror(basicNode("REDACTED"), s22)
	require.NoError(t, err)
	_, err = aligner.AppendMirror(basicNode("REDACTED"), s12)
	require.NoError(t, err)
	_, err = aligner.AppendMirror(basicNode("REDACTED"), s11)
	require.NoError(t, err)

	linkMirror.On("REDACTED", mock.Anything, &iface.QueryProposalMirror{
		Mirror: toIface(s22), ApplicationDigest: []byte("REDACTED"),
	}).Once().Return(&iface.ReplyProposalMirror{Outcome: iface.Replymirrorsnapshot_DECLINE}, nil)

	linkMirror.On("REDACTED", mock.Anything, &iface.QueryProposalMirror{
		Mirror: toIface(s12), ApplicationDigest: []byte("REDACTED"),
	}).Once().Return(&iface.ReplyProposalMirror{Outcome: iface.Replymirrorsnapshot_DECLINE}, nil)

	linkMirror.On("REDACTED", mock.Anything, &iface.QueryProposalMirror{
		Mirror: toIface(s11), ApplicationDigest: []byte("REDACTED"),
	}).Once().Return(&iface.ReplyProposalMirror{Outcome: iface.Replymirrorsnapshot_DECLINE}, nil)

	_, _, err = aligner.AlignAny(0, func() {})
	assert.Equal(t, errNoMirrors, err)
	linkMirror.AssertExpectations(t)
}

func Verifyaligner_Alignany_decline_layout(t *testing.T) {
	aligner, linkMirror := configureProposalAligner()

	//
	s22 := &mirror{Level: 2, Layout: 2, Segments: 3, Digest: []byte{1, 2, 3}}
	s12 := &mirror{Level: 1, Layout: 2, Segments: 3, Digest: []byte{1, 2, 3}}
	s11 := &mirror{Level: 1, Layout: 1, Segments: 3, Digest: []byte{1, 2, 3}}
	_, err := aligner.AppendMirror(basicNode("REDACTED"), s22)
	require.NoError(t, err)
	_, err = aligner.AppendMirror(basicNode("REDACTED"), s12)
	require.NoError(t, err)
	_, err = aligner.AppendMirror(basicNode("REDACTED"), s11)
	require.NoError(t, err)

	linkMirror.On("REDACTED", mock.Anything, &iface.QueryProposalMirror{
		Mirror: toIface(s22), ApplicationDigest: []byte("REDACTED"),
	}).Once().Return(&iface.ReplyProposalMirror{Outcome: iface.Replymirrorsnapshot_DECLINE_LAYOUT}, nil)

	linkMirror.On("REDACTED", mock.Anything, &iface.QueryProposalMirror{
		Mirror: toIface(s11), ApplicationDigest: []byte("REDACTED"),
	}).Once().Return(&iface.ReplyProposalMirror{Outcome: iface.Replymirrorsnapshot_CANCEL}, nil)

	_, _, err = aligner.AlignAny(0, func() {})
	assert.Equal(t, errCancel, err)
	linkMirror.AssertExpectations(t)
}

func Verifyaligner_Alignany_decline_emitter(t *testing.T) {
	aligner, linkMirror := configureProposalAligner()

	nodeA := basicNode("REDACTED")
	nodeBYTE := basicNode("REDACTED")
	nodeC := basicNode("REDACTED")

	//
	//
	//
	sa := &mirror{Level: 1, Layout: 1, Segments: 3, Digest: []byte{1, 2, 3}}
	sb := &mirror{Level: 2, Layout: 1, Segments: 3, Digest: []byte{1, 2, 3}}
	sc := &mirror{Level: 3, Layout: 1, Segments: 3, Digest: []byte{1, 2, 3}}
	sbc := &mirror{Level: 4, Layout: 1, Segments: 3, Digest: []byte{1, 2, 3}}
	_, err := aligner.AppendMirror(nodeA, sa)
	require.NoError(t, err)
	_, err = aligner.AppendMirror(nodeBYTE, sb)
	require.NoError(t, err)
	_, err = aligner.AppendMirror(nodeC, sc)
	require.NoError(t, err)
	_, err = aligner.AppendMirror(nodeBYTE, sbc)
	require.NoError(t, err)
	_, err = aligner.AppendMirror(nodeC, sbc)
	require.NoError(t, err)

	linkMirror.On("REDACTED", mock.Anything, &iface.QueryProposalMirror{
		Mirror: toIface(sbc), ApplicationDigest: []byte("REDACTED"),
	}).Once().Return(&iface.ReplyProposalMirror{Outcome: iface.Replymirrorsnapshot_DECLINE_EMITTER}, nil)

	linkMirror.On("REDACTED", mock.Anything, &iface.QueryProposalMirror{
		Mirror: toIface(sa), ApplicationDigest: []byte("REDACTED"),
	}).Once().Return(&iface.ReplyProposalMirror{Outcome: iface.Replymirrorsnapshot_DECLINE}, nil)

	_, _, err = aligner.AlignAny(0, func() {})
	assert.Equal(t, errNoMirrors, err)
	linkMirror.AssertExpectations(t)
}

func Verifyaligner_Alignany_ifaceerror(t *testing.T) {
	aligner, linkMirror := configureProposalAligner()

	errBoom := errors.New("REDACTED")
	s := &mirror{Level: 1, Layout: 1, Segments: 3, Digest: []byte{1, 2, 3}}
	_, err := aligner.AppendMirror(basicNode("REDACTED"), s)
	require.NoError(t, err)
	linkMirror.On("REDACTED", mock.Anything, &iface.QueryProposalMirror{
		Mirror: toIface(s), ApplicationDigest: []byte("REDACTED"),
	}).Once().Return(nil, errBoom)

	_, _, err = aligner.AlignAny(0, func() {})
	assert.True(t, errors.Is(err, errBoom))
	linkMirror.AssertExpectations(t)
}

func Verifyaligner_mirrorsnapshot(t *testing.T) {
	unclearErr := errors.New("REDACTED")
	boom := errors.New("REDACTED")

	verifyscenarios := map[string]struct {
		outcome    iface.Replymirrorsnapshot_Outcome
		err       error
		anticipateErr error
	}{
		"REDACTED":           {iface.Replymirrorsnapshot_ALLOW, nil, nil},
		"REDACTED":            {iface.Replymirrorsnapshot_CANCEL, nil, errCancel},
		"REDACTED":           {iface.Replymirrorsnapshot_DECLINE, nil, errDeclineMirror},
		"REDACTED":    {iface.Replymirrorsnapshot_DECLINE_LAYOUT, nil, errDeclineLayout},
		"REDACTED":    {iface.Replymirrorsnapshot_DECLINE_EMITTER, nil, errDeclineEmitter},
		"REDACTED":          {iface.Replymirrorsnapshot_UNCLEAR, nil, unclearErr},
		"REDACTED":            {0, boom, boom},
		"REDACTED": {9, nil, unclearErr},
	}
	for label, tc := range verifyscenarios {

		t.Run(label, func(t *testing.T) {
			aligner, linkMirror := configureProposalAligner()
			s := &mirror{Level: 1, Layout: 1, Segments: 3, Digest: []byte{1, 2, 3}, validatedApplicationDigest: []byte("REDACTED")}
			linkMirror.On("REDACTED", mock.Anything, &iface.QueryProposalMirror{
				Mirror: toIface(s),
				ApplicationDigest:  []byte("REDACTED"),
			}).Return(&iface.ReplyProposalMirror{Outcome: tc.outcome}, tc.err)
			err := aligner.proposalMirror(s)
			if tc.anticipateErr == unclearErr {
				require.Error(t, err)
			} else {
				exposed := errors.Unwrap(err)
				if exposed != nil {
					err = exposed
				}
				assert.Equal(t, tc.anticipateErr, err)
			}
		})
	}
}

func Verifyaligner_executesegments_Outcomes(t *testing.T) {
	unclearErr := errors.New("REDACTED")
	boom := errors.New("REDACTED")

	verifyscenarios := map[string]struct {
		outcome    iface.Replyexecutemirrorsegment_Outcome
		err       error
		anticipateErr error
	}{
		"REDACTED":           {iface.Replyexecutemirrorsegment_ALLOW, nil, nil},
		"REDACTED":            {iface.Replyexecutemirrorsegment_CANCEL, nil, errCancel},
		"REDACTED":            {iface.Replyexecutemirrorsegment_REPROCESS, nil, nil},
		"REDACTED":   {iface.Replyexecutemirrorsegment_REPROCESS_MIRROR, nil, errReprocessMirror},
		"REDACTED":  {iface.Replyexecutemirrorsegment_DECLINE_MIRROR, nil, errDeclineMirror},
		"REDACTED":          {iface.Replyexecutemirrorsegment_UNCLEAR, nil, unclearErr},
		"REDACTED":            {0, boom, boom},
		"REDACTED": {9, nil, unclearErr},
	}
	for label, tc := range verifyscenarios {

		t.Run(label, func(t *testing.T) {
			linkInquire := &gatewaysims.ApplicationLinkInquire{}
			linkMirror := &gatewaysims.ApplicationLinkMirror{}
			statusSource := &simulations.StatusSource{}
			statusSource.On("REDACTED", mock.Anything, mock.Anything).Return([]byte("REDACTED"), nil)

			cfg := settings.StandardStatusAlignSettings()
			aligner := newAligner(*cfg, log.NewNoopTracer(), linkMirror, linkInquire, statusSource, "REDACTED")

			content := []byte{1, 2, 3}
			segments, err := newSegmentBuffer(&mirror{Level: 1, Layout: 1, Segments: 1}, "REDACTED")
			require.NoError(t, err)
			_, err = segments.Add(&segment{Level: 1, Layout: 1, Ordinal: 0, Segment: content})
			require.NoError(t, err)

			linkMirror.On("REDACTED", mock.Anything, &iface.QueryExecuteMirrorSegment{
				Ordinal: 0, Segment: content,
			}).Once().Return(&iface.ReplyExecuteMirrorSegment{Outcome: tc.outcome}, tc.err)
			if tc.outcome == iface.Replyexecutemirrorsegment_REPROCESS {
				linkMirror.On("REDACTED", mock.Anything, &iface.QueryExecuteMirrorSegment{
					Ordinal: 0, Segment: content,
				}).Once().Return(&iface.ReplyExecuteMirrorSegment{
					Outcome: iface.Replyexecutemirrorsegment_ALLOW,
				}, nil)
			}

			err = aligner.executeSegments(segments)
			if tc.anticipateErr == unclearErr {
				require.Error(t, err)
			} else {
				exposed := errors.Unwrap(err)
				if exposed != nil {
					err = exposed
				}
				assert.Equal(t, tc.anticipateErr, err)
			}
			linkMirror.AssertExpectations(t)
		})
	}
}

func Verifyaligner_executesegments_Reclaimsegments(t *testing.T) {
	//
	verifyscenarios := map[string]struct {
		outcome iface.Replyexecutemirrorsegment_Outcome
	}{
		"REDACTED":          {iface.Replyexecutemirrorsegment_ALLOW},
		"REDACTED":           {iface.Replyexecutemirrorsegment_CANCEL},
		"REDACTED":           {iface.Replyexecutemirrorsegment_REPROCESS},
		"REDACTED":  {iface.Replyexecutemirrorsegment_REPROCESS_MIRROR},
		"REDACTED": {iface.Replyexecutemirrorsegment_DECLINE_MIRROR},
	}
	for label, tc := range verifyscenarios {

		t.Run(label, func(t *testing.T) {
			linkInquire := &gatewaysims.ApplicationLinkInquire{}
			linkMirror := &gatewaysims.ApplicationLinkMirror{}
			statusSource := &simulations.StatusSource{}
			statusSource.On("REDACTED", mock.Anything, mock.Anything).Return([]byte("REDACTED"), nil)

			cfg := settings.StandardStatusAlignSettings()
			aligner := newAligner(*cfg, log.NewNoopTracer(), linkMirror, linkInquire, statusSource, "REDACTED")

			segments, err := newSegmentBuffer(&mirror{Level: 1, Layout: 1, Segments: 3}, "REDACTED")
			require.NoError(t, err)
			appended, err := segments.Add(&segment{Level: 1, Layout: 1, Ordinal: 0, Segment: []byte{0}})
			require.True(t, appended)
			require.NoError(t, err)
			appended, err = segments.Add(&segment{Level: 1, Layout: 1, Ordinal: 1, Segment: []byte{1}})
			require.True(t, appended)
			require.NoError(t, err)
			appended, err = segments.Add(&segment{Level: 1, Layout: 1, Ordinal: 2, Segment: []byte{2}})
			require.True(t, appended)
			require.NoError(t, err)

			//
			linkMirror.On("REDACTED", mock.Anything, &iface.QueryExecuteMirrorSegment{
				Ordinal: 0, Segment: []byte{0},
			}).Once().Return(&iface.ReplyExecuteMirrorSegment{Outcome: iface.Replyexecutemirrorsegment_ALLOW}, nil)
			linkMirror.On("REDACTED", mock.Anything, &iface.QueryExecuteMirrorSegment{
				Ordinal: 1, Segment: []byte{1},
			}).Once().Return(&iface.ReplyExecuteMirrorSegment{Outcome: iface.Replyexecutemirrorsegment_ALLOW}, nil)
			linkMirror.On("REDACTED", mock.Anything, &iface.QueryExecuteMirrorSegment{
				Ordinal: 2, Segment: []byte{2},
			}).Once().Return(&iface.ReplyExecuteMirrorSegment{
				Outcome:        tc.outcome,
				ReacquireSegments: []uint32{1},
			}, nil)

			//
			//
			//
			go func() {
				aligner.executeSegments(segments) //
			}()

			time.Sleep(50 * time.Millisecond)
			assert.True(t, segments.Has(0))
			assert.False(t, segments.Has(1))
			assert.True(t, segments.Has(2))
			err = segments.End()
			require.NoError(t, err)
		})
	}
}

func Verifyaligner_executesegments_Declineemitters(t *testing.T) {
	//
	verifyscenarios := map[string]struct {
		outcome iface.Replyexecutemirrorsegment_Outcome
	}{
		"REDACTED":          {iface.Replyexecutemirrorsegment_ALLOW},
		"REDACTED":           {iface.Replyexecutemirrorsegment_CANCEL},
		"REDACTED":           {iface.Replyexecutemirrorsegment_REPROCESS},
		"REDACTED":  {iface.Replyexecutemirrorsegment_REPROCESS_MIRROR},
		"REDACTED": {iface.Replyexecutemirrorsegment_DECLINE_MIRROR},
	}
	for label, tc := range verifyscenarios {

		t.Run(label, func(t *testing.T) {
			linkInquire := &gatewaysims.ApplicationLinkInquire{}
			linkMirror := &gatewaysims.ApplicationLinkMirror{}
			statusSource := &simulations.StatusSource{}
			statusSource.On("REDACTED", mock.Anything, mock.Anything).Return([]byte("REDACTED"), nil)

			cfg := settings.StandardStatusAlignSettings()
			aligner := newAligner(*cfg, log.NewNoopTracer(), linkMirror, linkInquire, statusSource, "REDACTED")

			//
			//
			nodeA := basicNode("REDACTED")
			nodeBYTE := basicNode("REDACTED")
			nodeC := basicNode("REDACTED")

			s1 := &mirror{Level: 1, Layout: 1, Segments: 3}
			s2 := &mirror{Level: 2, Layout: 1, Segments: 3}
			_, err := aligner.AppendMirror(nodeA, s1)
			require.NoError(t, err)
			_, err = aligner.AppendMirror(nodeA, s2)
			require.NoError(t, err)
			_, err = aligner.AppendMirror(nodeBYTE, s1)
			require.NoError(t, err)
			_, err = aligner.AppendMirror(nodeBYTE, s2)
			require.NoError(t, err)
			_, err = aligner.AppendMirror(nodeC, s1)
			require.NoError(t, err)
			_, err = aligner.AppendMirror(nodeC, s2)
			require.NoError(t, err)

			segments, err := newSegmentBuffer(s1, "REDACTED")
			require.NoError(t, err)
			appended, err := segments.Add(&segment{Level: 1, Layout: 1, Ordinal: 0, Segment: []byte{0}, Emitter: nodeA.ID()})
			require.True(t, appended)
			require.NoError(t, err)
			appended, err = segments.Add(&segment{Level: 1, Layout: 1, Ordinal: 1, Segment: []byte{1}, Emitter: nodeBYTE.ID()})
			require.True(t, appended)
			require.NoError(t, err)
			appended, err = segments.Add(&segment{Level: 1, Layout: 1, Ordinal: 2, Segment: []byte{2}, Emitter: nodeC.ID()})
			require.True(t, appended)
			require.NoError(t, err)

			//
			linkMirror.On("REDACTED", mock.Anything, &iface.QueryExecuteMirrorSegment{
				Ordinal: 0, Segment: []byte{0}, Emitter: "REDACTED",
			}).Once().Return(&iface.ReplyExecuteMirrorSegment{Outcome: iface.Replyexecutemirrorsegment_ALLOW}, nil)
			linkMirror.On("REDACTED", mock.Anything, &iface.QueryExecuteMirrorSegment{
				Ordinal: 1, Segment: []byte{1}, Emitter: "REDACTED",
			}).Once().Return(&iface.ReplyExecuteMirrorSegment{Outcome: iface.Replyexecutemirrorsegment_ALLOW}, nil)
			linkMirror.On("REDACTED", mock.Anything, &iface.QueryExecuteMirrorSegment{
				Ordinal: 2, Segment: []byte{2}, Emitter: "REDACTED",
			}).Once().Return(&iface.ReplyExecuteMirrorSegment{
				Outcome:        tc.outcome,
				DeclineEmitters: []string{string(nodeBYTE.ID())},
			}, nil)

			//
			if tc.outcome == iface.Replyexecutemirrorsegment_REPROCESS {
				linkMirror.On("REDACTED", mock.Anything, &iface.QueryExecuteMirrorSegment{
					Ordinal: 2, Segment: []byte{2}, Emitter: "REDACTED",
				}).Once().Return(&iface.ReplyExecuteMirrorSegment{Outcome: iface.Replyexecutemirrorsegment_ALLOW}, nil)
			}

			//
			//
			//
			go func() {
				aligner.executeSegments(segments) //
			}()

			time.Sleep(50 * time.Millisecond)

			s1nodes := aligner.mirrors.FetchNodes(s1)
			assert.Len(t, s1nodes, 2)
			assert.EqualValues(t, "REDACTED", s1nodes[0].ID())
			assert.EqualValues(t, "REDACTED", s1nodes[1].ID())

			aligner.mirrors.FetchNodes(s1)
			assert.Len(t, s1nodes, 2)
			assert.EqualValues(t, "REDACTED", s1nodes[0].ID())
			assert.EqualValues(t, "REDACTED", s1nodes[1].ID())

			err = segments.End()
			require.NoError(t, err)
		})
	}
}

func Verifyaligner_validateapplication(t *testing.T) {
	boom := errors.New("REDACTED")
	const applicationRelease = 9
	applicationReleaseDiscrepancyErr := errors.New("REDACTED")
	s := &mirror{Level: 3, Layout: 1, Segments: 5, Digest: []byte{1, 2, 3}, validatedApplicationDigest: []byte("REDACTED")}

	verifyscenarios := map[string]struct {
		reply  *iface.ReplyDetails
		err       error
		anticipateErr error
	}{
		"REDACTED": {&iface.ReplyDetails{
			FinalLedgerLevel:  3,
			FinalLedgerApplicationDigest: []byte("REDACTED"),
			ApplicationRelease:       applicationRelease,
		}, nil, nil},
		"REDACTED": {&iface.ReplyDetails{
			FinalLedgerLevel:  3,
			FinalLedgerApplicationDigest: []byte("REDACTED"),
			ApplicationRelease:       2,
		}, nil, applicationReleaseDiscrepancyErr},
		"REDACTED": {&iface.ReplyDetails{
			FinalLedgerLevel:  5,
			FinalLedgerApplicationDigest: []byte("REDACTED"),
			ApplicationRelease:       applicationRelease,
		}, nil, errValidateErrored},
		"REDACTED": {&iface.ReplyDetails{
			FinalLedgerLevel:  3,
			FinalLedgerApplicationDigest: []byte("REDACTED"),
			ApplicationRelease:       applicationRelease,
		}, nil, errValidateErrored},
		"REDACTED": {nil, boom, boom},
	}
	for label, tc := range verifyscenarios {

		t.Run(label, func(t *testing.T) {
			linkInquire := &gatewaysims.ApplicationLinkInquire{}
			linkMirror := &gatewaysims.ApplicationLinkMirror{}
			statusSource := &simulations.StatusSource{}

			cfg := settings.StandardStatusAlignSettings()
			aligner := newAligner(*cfg, log.NewNoopTracer(), linkMirror, linkInquire, statusSource, "REDACTED")

			linkInquire.On("REDACTED", mock.Anything, gateway.QueryDetails).Return(tc.reply, tc.err)
			err := aligner.validateApplication(s, applicationRelease)
			exposed := errors.Unwrap(err)
			if exposed != nil {
				err = exposed
			}
			require.Equal(t, tc.anticipateErr, err)
		})
	}
}

func toIface(s *mirror) *iface.Mirror {
	return &iface.Mirror{
		Level:   s.Level,
		Layout:   s.Layout,
		Segments:   s.Segments,
		Digest:     s.Digest,
		Metainfo: s.Metainfo,
	}
}
