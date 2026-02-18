package statereplication

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	atci "github.com/valkyrieworks/atci/kinds"
	"github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/utils/log"
	ctsync "github.com/valkyrieworks/utils/alignment"
	"github.com/valkyrieworks/p2p"
	netmocks "github.com/valkyrieworks/p2p/simulations"
	ctstatus "github.com/valkyrieworks/schema/consensuscore/status"
	sschema "github.com/valkyrieworks/schema/consensuscore/statereplication"
	ctrelease "github.com/valkyrieworks/schema/consensuscore/release"
	"github.com/valkyrieworks/gateway"
	gatwaymocks "github.com/valkyrieworks/gateway/simulations"
	sm "github.com/valkyrieworks/status"
	"github.com/valkyrieworks/statereplication/simulations"
	"github.com/valkyrieworks/kinds"
	"github.com/valkyrieworks/release"
)

const testAppVersion = 9

//
func setupOfferSyncer() (*syncer, *gatwaymocks.AppConnSnapshot) {
	connQuery := &gatwaymocks.AppConnQuery{}
	connSnapshot := &gatwaymocks.AppConnSnapshot{}
	stateProvider := &simulations.StateProvider{}
	stateProvider.On("REDACTED", mock.Anything, mock.Anything).Return([]byte("REDACTED"), nil)
	cfg := settings.DefaultStateSyncConfig()
	syncer := newSyncer(*cfg, log.NewNopLogger(), connSnapshot, connQuery, stateProvider, "REDACTED")

	return syncer, connSnapshot
}

//
func simplePeer(id string) *netmocks.Peer {
	peer := &netmocks.Peer{}
	peer.On("REDACTED").Return(p2p.ID(id))
	return peer
}

func TestSyncer_SyncAny(t *testing.T) {
	state := sm.State{
		ChainID: "REDACTED",
		Version: ctstatus.Version{
			Consensus: ctrelease.Consensus{
				Block: release.BlockProtocol,
				App:   testAppVersion,
			},
			Software: release.TMCoreSemVer,
		},

		LastBlockHeight: 1,
		LastBlockID:     kinds.BlockID{Hash: []byte("REDACTED")},
		LastBlockTime:   time.Now(),
		LastResultsHash: []byte("REDACTED"),
		AppHash:         []byte("REDACTED"),

		LastValidators: &kinds.ValidatorSet{Proposer: &kinds.Validator{Address: []byte("REDACTED")}},
		Validators:     &kinds.ValidatorSet{Proposer: &kinds.Validator{Address: []byte("REDACTED")}},
		NextValidators: &kinds.ValidatorSet{Proposer: &kinds.Validator{Address: []byte("REDACTED")}},

		ConsensusParams:                  *kinds.DefaultConsensusParams(),
		LastHeightConsensusParamsChanged: 1,
	}
	commit := &kinds.Commit{BlockID: kinds.BlockID{Hash: []byte("REDACTED")}}

	chunks := []*chunk{
		{Height: 1, Format: 1, Index: 0, Chunk: []byte{1, 1, 0}},
		{Height: 1, Format: 1, Index: 1, Chunk: []byte{1, 1, 1}},
		{Height: 1, Format: 1, Index: 2, Chunk: []byte{1, 1, 2}},
	}
	s := &snapshot{Height: 1, Format: 1, Chunks: 3, Hash: []byte{1, 2, 3}}

	stateProvider := &simulations.StateProvider{}
	stateProvider.On("REDACTED", mock.Anything, uint64(1)).Return(state.AppHash, nil)
	stateProvider.On("REDACTED", mock.Anything, uint64(2)).Return([]byte("REDACTED"), nil)
	stateProvider.On("REDACTED", mock.Anything, uint64(1)).Return(commit, nil)
	stateProvider.On("REDACTED", mock.Anything, uint64(1)).Return(state, nil)
	connSnapshot := &gatwaymocks.AppConnSnapshot{}
	connQuery := &gatwaymocks.AppConnQuery{}

	cfg := settings.DefaultStateSyncConfig()
	syncer := newSyncer(*cfg, log.NewNopLogger(), connSnapshot, connQuery, stateProvider, "REDACTED")

	//
	_, err := syncer.AddChunk(&chunk{Height: 1, Format: 1, Index: 0, Chunk: []byte{1}})
	require.Error(t, err)

	//
	peerA := &netmocks.Peer{}
	peerA.On("REDACTED").Return(p2p.ID("REDACTED"))
	peerA.On("REDACTED", mock.MatchedBy(func(i any) bool {
		e, ok := i.(p2p.Envelope)
		if !ok {
			return false
		}
		req, ok := e.Message.(*sschema.SnapshotsRequest)
		return ok && e.ChannelID == SnapshotChannel && req != nil
	})).Return(true)
	syncer.AddPeer(peerA)
	peerA.AssertExpectations(t)

	peerB := &netmocks.Peer{}
	peerB.On("REDACTED").Return(p2p.ID("REDACTED"))
	peerB.On("REDACTED", mock.MatchedBy(func(i any) bool {
		e, ok := i.(p2p.Envelope)
		if !ok {
			return false
		}
		req, ok := e.Message.(*sschema.SnapshotsRequest)
		return ok && e.ChannelID == SnapshotChannel && req != nil
	})).Return(true)
	syncer.AddPeer(peerB)
	peerB.AssertExpectations(t)

	//
	//
	isNew, err := syncer.AddSnapshot(peerA, s)
	require.NoError(t, err)
	assert.True(t, isNew)

	isNew, err = syncer.AddSnapshot(peerB, s)
	require.NoError(t, err)
	assert.False(t, isNew)

	isNew, err = syncer.AddSnapshot(peerB, &snapshot{Height: 2, Format: 2, Chunks: 3, Hash: []byte{1}})
	require.NoError(t, err)
	assert.True(t, isNew)

	//
	//
	connSnapshot.On("REDACTED", mock.Anything, &atci.RequestOfferSnapshot{
		Snapshot: &atci.Snapshot{
			Height: 2,
			Format: 2,
			Chunks: 3,
			Hash:   []byte{1},
		},
		AppHash: []byte("REDACTED"),
	}).Return(&atci.ResponseOfferSnapshot{Result: atci.ResponseOfferSnapshot_REJECT_FORMAT}, nil)
	connSnapshot.On("REDACTED", mock.Anything, &atci.RequestOfferSnapshot{
		Snapshot: &atci.Snapshot{
			Height:   s.Height,
			Format:   s.Format,
			Chunks:   s.Chunks,
			Hash:     s.Hash,
			Metadata: s.Metadata,
		},
		AppHash: []byte("REDACTED"),
	}).Times(2).Return(&atci.ResponseOfferSnapshot{Result: atci.ResponseOfferSnapshot_ACCEPT}, nil)

	chunkRequests := make(map[uint32]int)
	chunkRequestsMtx := ctsync.Mutex{}
	onChunkRequest := func(args mock.Arguments) {
		e, ok := args[0].(p2p.Envelope)
		require.True(t, ok)
		msg := e.Message.(*sschema.ChunkRequest)
		require.EqualValues(t, 1, msg.Height)
		require.EqualValues(t, 1, msg.Format)
		require.LessOrEqual(t, msg.Index, uint32(len(chunks)))

		added, err := syncer.AddChunk(chunks[msg.Index])
		require.NoError(t, err)
		assert.True(t, added)

		chunkRequestsMtx.Lock()
		chunkRequests[msg.Index]++
		chunkRequestsMtx.Unlock()
	}
	peerA.On("REDACTED", mock.MatchedBy(func(i any) bool {
		e, ok := i.(p2p.Envelope)
		return ok && e.ChannelID == ChunkChannel
	})).Maybe().Run(onChunkRequest).Return(true)
	peerB.On("REDACTED", mock.MatchedBy(func(i any) bool {
		e, ok := i.(p2p.Envelope)
		return ok && e.ChannelID == ChunkChannel
	})).Maybe().Run(onChunkRequest).Return(true)

	//
	//
	//
	connSnapshot.On("REDACTED", mock.Anything, &atci.RequestApplySnapshotChunk{
		Index: 2, Chunk: []byte{1, 1, 2},
	}).Once().Run(func(args mock.Arguments) { time.Sleep(2 * time.Second) }).Return(
		&atci.ResponseApplySnapshotChunk{
			Result:        atci.ResponseApplySnapshotChunk_RETRY_SNAPSHOT,
			RefetchChunks: []uint32{1},
		}, nil)

	connSnapshot.On("REDACTED", mock.Anything, &atci.RequestApplySnapshotChunk{
		Index: 0, Chunk: []byte{1, 1, 0},
	}).Times(2).Return(&atci.ResponseApplySnapshotChunk{Result: atci.ResponseApplySnapshotChunk_ACCEPT}, nil)
	connSnapshot.On("REDACTED", mock.Anything, &atci.RequestApplySnapshotChunk{
		Index: 1, Chunk: []byte{1, 1, 1},
	}).Times(2).Return(&atci.ResponseApplySnapshotChunk{Result: atci.ResponseApplySnapshotChunk_ACCEPT}, nil)
	connSnapshot.On("REDACTED", mock.Anything, &atci.RequestApplySnapshotChunk{
		Index: 2, Chunk: []byte{1, 1, 2},
	}).Once().Return(&atci.ResponseApplySnapshotChunk{Result: atci.ResponseApplySnapshotChunk_ACCEPT}, nil)
	connQuery.On("REDACTED", mock.Anything, gateway.RequestInfo).Return(&atci.ResponseInfo{
		AppVersion:       testAppVersion,
		LastBlockHeight:  1,
		LastBlockAppHash: []byte("REDACTED"),
	}, nil)

	newState, lastCommit, err := syncer.SyncAny(0, func() {})
	require.NoError(t, err)

	time.Sleep(50 * time.Millisecond) //

	chunkRequestsMtx.Lock()
	assert.Equal(t, map[uint32]int{0: 1, 1: 2, 2: 1}, chunkRequests)
	chunkRequestsMtx.Unlock()

	expectState := state

	assert.Equal(t, expectState, newState)
	assert.Equal(t, commit, lastCommit)

	connSnapshot.AssertExpectations(t)
	connQuery.AssertExpectations(t)
	peerA.AssertExpectations(t)
	peerB.AssertExpectations(t)
}

func TestSyncer_SyncAny_noSnapshots(t *testing.T) {
	syncer, _ := setupOfferSyncer()
	_, _, err := syncer.SyncAny(0, func() {})
	assert.Equal(t, errNoSnapshots, err)
}

func TestSyncer_SyncAny_abort(t *testing.T) {
	syncer, connSnapshot := setupOfferSyncer()

	s := &snapshot{Height: 1, Format: 1, Chunks: 3, Hash: []byte{1, 2, 3}}
	_, err := syncer.AddSnapshot(simplePeer("REDACTED"), s)
	require.NoError(t, err)
	connSnapshot.On("REDACTED", mock.Anything, &atci.RequestOfferSnapshot{
		Snapshot: toABCI(s), AppHash: []byte("REDACTED"),
	}).Once().Return(&atci.ResponseOfferSnapshot{Result: atci.ResponseOfferSnapshot_ABORT}, nil)

	_, _, err = syncer.SyncAny(0, func() {})
	assert.Equal(t, errAbort, err)
	connSnapshot.AssertExpectations(t)
}

func TestSyncer_SyncAny_reject(t *testing.T) {
	syncer, connSnapshot := setupOfferSyncer()

	//
	s22 := &snapshot{Height: 2, Format: 2, Chunks: 3, Hash: []byte{1, 2, 3}}
	s12 := &snapshot{Height: 1, Format: 2, Chunks: 3, Hash: []byte{1, 2, 3}}
	s11 := &snapshot{Height: 1, Format: 1, Chunks: 3, Hash: []byte{1, 2, 3}}
	_, err := syncer.AddSnapshot(simplePeer("REDACTED"), s22)
	require.NoError(t, err)
	_, err = syncer.AddSnapshot(simplePeer("REDACTED"), s12)
	require.NoError(t, err)
	_, err = syncer.AddSnapshot(simplePeer("REDACTED"), s11)
	require.NoError(t, err)

	connSnapshot.On("REDACTED", mock.Anything, &atci.RequestOfferSnapshot{
		Snapshot: toABCI(s22), AppHash: []byte("REDACTED"),
	}).Once().Return(&atci.ResponseOfferSnapshot{Result: atci.ResponseOfferSnapshot_REJECT}, nil)

	connSnapshot.On("REDACTED", mock.Anything, &atci.RequestOfferSnapshot{
		Snapshot: toABCI(s12), AppHash: []byte("REDACTED"),
	}).Once().Return(&atci.ResponseOfferSnapshot{Result: atci.ResponseOfferSnapshot_REJECT}, nil)

	connSnapshot.On("REDACTED", mock.Anything, &atci.RequestOfferSnapshot{
		Snapshot: toABCI(s11), AppHash: []byte("REDACTED"),
	}).Once().Return(&atci.ResponseOfferSnapshot{Result: atci.ResponseOfferSnapshot_REJECT}, nil)

	_, _, err = syncer.SyncAny(0, func() {})
	assert.Equal(t, errNoSnapshots, err)
	connSnapshot.AssertExpectations(t)
}

func TestSyncer_SyncAny_reject_format(t *testing.T) {
	syncer, connSnapshot := setupOfferSyncer()

	//
	s22 := &snapshot{Height: 2, Format: 2, Chunks: 3, Hash: []byte{1, 2, 3}}
	s12 := &snapshot{Height: 1, Format: 2, Chunks: 3, Hash: []byte{1, 2, 3}}
	s11 := &snapshot{Height: 1, Format: 1, Chunks: 3, Hash: []byte{1, 2, 3}}
	_, err := syncer.AddSnapshot(simplePeer("REDACTED"), s22)
	require.NoError(t, err)
	_, err = syncer.AddSnapshot(simplePeer("REDACTED"), s12)
	require.NoError(t, err)
	_, err = syncer.AddSnapshot(simplePeer("REDACTED"), s11)
	require.NoError(t, err)

	connSnapshot.On("REDACTED", mock.Anything, &atci.RequestOfferSnapshot{
		Snapshot: toABCI(s22), AppHash: []byte("REDACTED"),
	}).Once().Return(&atci.ResponseOfferSnapshot{Result: atci.ResponseOfferSnapshot_REJECT_FORMAT}, nil)

	connSnapshot.On("REDACTED", mock.Anything, &atci.RequestOfferSnapshot{
		Snapshot: toABCI(s11), AppHash: []byte("REDACTED"),
	}).Once().Return(&atci.ResponseOfferSnapshot{Result: atci.ResponseOfferSnapshot_ABORT}, nil)

	_, _, err = syncer.SyncAny(0, func() {})
	assert.Equal(t, errAbort, err)
	connSnapshot.AssertExpectations(t)
}

func TestSyncer_SyncAny_reject_sender(t *testing.T) {
	syncer, connSnapshot := setupOfferSyncer()

	peerA := simplePeer("REDACTED")
	peerB := simplePeer("REDACTED")
	peerC := simplePeer("REDACTED")

	//
	//
	//
	sa := &snapshot{Height: 1, Format: 1, Chunks: 3, Hash: []byte{1, 2, 3}}
	sb := &snapshot{Height: 2, Format: 1, Chunks: 3, Hash: []byte{1, 2, 3}}
	sc := &snapshot{Height: 3, Format: 1, Chunks: 3, Hash: []byte{1, 2, 3}}
	sbc := &snapshot{Height: 4, Format: 1, Chunks: 3, Hash: []byte{1, 2, 3}}
	_, err := syncer.AddSnapshot(peerA, sa)
	require.NoError(t, err)
	_, err = syncer.AddSnapshot(peerB, sb)
	require.NoError(t, err)
	_, err = syncer.AddSnapshot(peerC, sc)
	require.NoError(t, err)
	_, err = syncer.AddSnapshot(peerB, sbc)
	require.NoError(t, err)
	_, err = syncer.AddSnapshot(peerC, sbc)
	require.NoError(t, err)

	connSnapshot.On("REDACTED", mock.Anything, &atci.RequestOfferSnapshot{
		Snapshot: toABCI(sbc), AppHash: []byte("REDACTED"),
	}).Once().Return(&atci.ResponseOfferSnapshot{Result: atci.ResponseOfferSnapshot_REJECT_SENDER}, nil)

	connSnapshot.On("REDACTED", mock.Anything, &atci.RequestOfferSnapshot{
		Snapshot: toABCI(sa), AppHash: []byte("REDACTED"),
	}).Once().Return(&atci.ResponseOfferSnapshot{Result: atci.ResponseOfferSnapshot_REJECT}, nil)

	_, _, err = syncer.SyncAny(0, func() {})
	assert.Equal(t, errNoSnapshots, err)
	connSnapshot.AssertExpectations(t)
}

func TestSyncer_SyncAny_abciError(t *testing.T) {
	syncer, connSnapshot := setupOfferSyncer()

	errBoom := errors.New("REDACTED")
	s := &snapshot{Height: 1, Format: 1, Chunks: 3, Hash: []byte{1, 2, 3}}
	_, err := syncer.AddSnapshot(simplePeer("REDACTED"), s)
	require.NoError(t, err)
	connSnapshot.On("REDACTED", mock.Anything, &atci.RequestOfferSnapshot{
		Snapshot: toABCI(s), AppHash: []byte("REDACTED"),
	}).Once().Return(nil, errBoom)

	_, _, err = syncer.SyncAny(0, func() {})
	assert.True(t, errors.Is(err, errBoom))
	connSnapshot.AssertExpectations(t)
}

func TestSyncer_offerSnapshot(t *testing.T) {
	unknownErr := errors.New("REDACTED")
	boom := errors.New("REDACTED")

	testcases := map[string]struct {
		result    atci.ResponseOfferSnapshot_Result
		err       error
		expectErr error
	}{
		"REDACTED":           {atci.ResponseOfferSnapshot_ACCEPT, nil, nil},
		"REDACTED":            {atci.ResponseOfferSnapshot_ABORT, nil, errAbort},
		"REDACTED":           {atci.ResponseOfferSnapshot_REJECT, nil, errRejectSnapshot},
		"REDACTED":    {atci.ResponseOfferSnapshot_REJECT_FORMAT, nil, errRejectFormat},
		"REDACTED":    {atci.ResponseOfferSnapshot_REJECT_SENDER, nil, errRejectSender},
		"REDACTED":          {atci.ResponseOfferSnapshot_UNKNOWN, nil, unknownErr},
		"REDACTED":            {0, boom, boom},
		"REDACTED": {9, nil, unknownErr},
	}
	for name, tc := range testcases {

		t.Run(name, func(t *testing.T) {
			syncer, connSnapshot := setupOfferSyncer()
			s := &snapshot{Height: 1, Format: 1, Chunks: 3, Hash: []byte{1, 2, 3}, trustedAppHash: []byte("REDACTED")}
			connSnapshot.On("REDACTED", mock.Anything, &atci.RequestOfferSnapshot{
				Snapshot: toABCI(s),
				AppHash:  []byte("REDACTED"),
			}).Return(&atci.ResponseOfferSnapshot{Result: tc.result}, tc.err)
			err := syncer.offerSnapshot(s)
			if tc.expectErr == unknownErr {
				require.Error(t, err)
			} else {
				unwrapped := errors.Unwrap(err)
				if unwrapped != nil {
					err = unwrapped
				}
				assert.Equal(t, tc.expectErr, err)
			}
		})
	}
}

func TestSyncer_applyChunks_Results(t *testing.T) {
	unknownErr := errors.New("REDACTED")
	boom := errors.New("REDACTED")

	testcases := map[string]struct {
		result    atci.ResponseApplySnapshotChunk_Result
		err       error
		expectErr error
	}{
		"REDACTED":           {atci.ResponseApplySnapshotChunk_ACCEPT, nil, nil},
		"REDACTED":            {atci.ResponseApplySnapshotChunk_ABORT, nil, errAbort},
		"REDACTED":            {atci.ResponseApplySnapshotChunk_RETRY, nil, nil},
		"REDACTED":   {atci.ResponseApplySnapshotChunk_RETRY_SNAPSHOT, nil, errRetrySnapshot},
		"REDACTED":  {atci.ResponseApplySnapshotChunk_REJECT_SNAPSHOT, nil, errRejectSnapshot},
		"REDACTED":          {atci.ResponseApplySnapshotChunk_UNKNOWN, nil, unknownErr},
		"REDACTED":            {0, boom, boom},
		"REDACTED": {9, nil, unknownErr},
	}
	for name, tc := range testcases {

		t.Run(name, func(t *testing.T) {
			connQuery := &gatwaymocks.AppConnQuery{}
			connSnapshot := &gatwaymocks.AppConnSnapshot{}
			stateProvider := &simulations.StateProvider{}
			stateProvider.On("REDACTED", mock.Anything, mock.Anything).Return([]byte("REDACTED"), nil)

			cfg := settings.DefaultStateSyncConfig()
			syncer := newSyncer(*cfg, log.NewNopLogger(), connSnapshot, connQuery, stateProvider, "REDACTED")

			body := []byte{1, 2, 3}
			chunks, err := newChunkQueue(&snapshot{Height: 1, Format: 1, Chunks: 1}, "REDACTED")
			require.NoError(t, err)
			_, err = chunks.Add(&chunk{Height: 1, Format: 1, Index: 0, Chunk: body})
			require.NoError(t, err)

			connSnapshot.On("REDACTED", mock.Anything, &atci.RequestApplySnapshotChunk{
				Index: 0, Chunk: body,
			}).Once().Return(&atci.ResponseApplySnapshotChunk{Result: tc.result}, tc.err)
			if tc.result == atci.ResponseApplySnapshotChunk_RETRY {
				connSnapshot.On("REDACTED", mock.Anything, &atci.RequestApplySnapshotChunk{
					Index: 0, Chunk: body,
				}).Once().Return(&atci.ResponseApplySnapshotChunk{
					Result: atci.ResponseApplySnapshotChunk_ACCEPT,
				}, nil)
			}

			err = syncer.applyChunks(chunks)
			if tc.expectErr == unknownErr {
				require.Error(t, err)
			} else {
				unwrapped := errors.Unwrap(err)
				if unwrapped != nil {
					err = unwrapped
				}
				assert.Equal(t, tc.expectErr, err)
			}
			connSnapshot.AssertExpectations(t)
		})
	}
}

func TestSyncer_applyChunks_RefetchChunks(t *testing.T) {
	//
	testcases := map[string]struct {
		result atci.ResponseApplySnapshotChunk_Result
	}{
		"REDACTED":          {atci.ResponseApplySnapshotChunk_ACCEPT},
		"REDACTED":           {atci.ResponseApplySnapshotChunk_ABORT},
		"REDACTED":           {atci.ResponseApplySnapshotChunk_RETRY},
		"REDACTED":  {atci.ResponseApplySnapshotChunk_RETRY_SNAPSHOT},
		"REDACTED": {atci.ResponseApplySnapshotChunk_REJECT_SNAPSHOT},
	}
	for name, tc := range testcases {

		t.Run(name, func(t *testing.T) {
			connQuery := &gatwaymocks.AppConnQuery{}
			connSnapshot := &gatwaymocks.AppConnSnapshot{}
			stateProvider := &simulations.StateProvider{}
			stateProvider.On("REDACTED", mock.Anything, mock.Anything).Return([]byte("REDACTED"), nil)

			cfg := settings.DefaultStateSyncConfig()
			syncer := newSyncer(*cfg, log.NewNopLogger(), connSnapshot, connQuery, stateProvider, "REDACTED")

			chunks, err := newChunkQueue(&snapshot{Height: 1, Format: 1, Chunks: 3}, "REDACTED")
			require.NoError(t, err)
			added, err := chunks.Add(&chunk{Height: 1, Format: 1, Index: 0, Chunk: []byte{0}})
			require.True(t, added)
			require.NoError(t, err)
			added, err = chunks.Add(&chunk{Height: 1, Format: 1, Index: 1, Chunk: []byte{1}})
			require.True(t, added)
			require.NoError(t, err)
			added, err = chunks.Add(&chunk{Height: 1, Format: 1, Index: 2, Chunk: []byte{2}})
			require.True(t, added)
			require.NoError(t, err)

			//
			connSnapshot.On("REDACTED", mock.Anything, &atci.RequestApplySnapshotChunk{
				Index: 0, Chunk: []byte{0},
			}).Once().Return(&atci.ResponseApplySnapshotChunk{Result: atci.ResponseApplySnapshotChunk_ACCEPT}, nil)
			connSnapshot.On("REDACTED", mock.Anything, &atci.RequestApplySnapshotChunk{
				Index: 1, Chunk: []byte{1},
			}).Once().Return(&atci.ResponseApplySnapshotChunk{Result: atci.ResponseApplySnapshotChunk_ACCEPT}, nil)
			connSnapshot.On("REDACTED", mock.Anything, &atci.RequestApplySnapshotChunk{
				Index: 2, Chunk: []byte{2},
			}).Once().Return(&atci.ResponseApplySnapshotChunk{
				Result:        tc.result,
				RefetchChunks: []uint32{1},
			}, nil)

			//
			//
			//
			go func() {
				syncer.applyChunks(chunks) //
			}()

			time.Sleep(50 * time.Millisecond)
			assert.True(t, chunks.Has(0))
			assert.False(t, chunks.Has(1))
			assert.True(t, chunks.Has(2))
			err = chunks.Close()
			require.NoError(t, err)
		})
	}
}

func TestSyncer_applyChunks_RejectSenders(t *testing.T) {
	//
	testcases := map[string]struct {
		result atci.ResponseApplySnapshotChunk_Result
	}{
		"REDACTED":          {atci.ResponseApplySnapshotChunk_ACCEPT},
		"REDACTED":           {atci.ResponseApplySnapshotChunk_ABORT},
		"REDACTED":           {atci.ResponseApplySnapshotChunk_RETRY},
		"REDACTED":  {atci.ResponseApplySnapshotChunk_RETRY_SNAPSHOT},
		"REDACTED": {atci.ResponseApplySnapshotChunk_REJECT_SNAPSHOT},
	}
	for name, tc := range testcases {

		t.Run(name, func(t *testing.T) {
			connQuery := &gatwaymocks.AppConnQuery{}
			connSnapshot := &gatwaymocks.AppConnSnapshot{}
			stateProvider := &simulations.StateProvider{}
			stateProvider.On("REDACTED", mock.Anything, mock.Anything).Return([]byte("REDACTED"), nil)

			cfg := settings.DefaultStateSyncConfig()
			syncer := newSyncer(*cfg, log.NewNopLogger(), connSnapshot, connQuery, stateProvider, "REDACTED")

			//
			//
			peerA := simplePeer("REDACTED")
			peerB := simplePeer("REDACTED")
			peerC := simplePeer("REDACTED")

			s1 := &snapshot{Height: 1, Format: 1, Chunks: 3}
			s2 := &snapshot{Height: 2, Format: 1, Chunks: 3}
			_, err := syncer.AddSnapshot(peerA, s1)
			require.NoError(t, err)
			_, err = syncer.AddSnapshot(peerA, s2)
			require.NoError(t, err)
			_, err = syncer.AddSnapshot(peerB, s1)
			require.NoError(t, err)
			_, err = syncer.AddSnapshot(peerB, s2)
			require.NoError(t, err)
			_, err = syncer.AddSnapshot(peerC, s1)
			require.NoError(t, err)
			_, err = syncer.AddSnapshot(peerC, s2)
			require.NoError(t, err)

			chunks, err := newChunkQueue(s1, "REDACTED")
			require.NoError(t, err)
			added, err := chunks.Add(&chunk{Height: 1, Format: 1, Index: 0, Chunk: []byte{0}, Sender: peerA.ID()})
			require.True(t, added)
			require.NoError(t, err)
			added, err = chunks.Add(&chunk{Height: 1, Format: 1, Index: 1, Chunk: []byte{1}, Sender: peerB.ID()})
			require.True(t, added)
			require.NoError(t, err)
			added, err = chunks.Add(&chunk{Height: 1, Format: 1, Index: 2, Chunk: []byte{2}, Sender: peerC.ID()})
			require.True(t, added)
			require.NoError(t, err)

			//
			connSnapshot.On("REDACTED", mock.Anything, &atci.RequestApplySnapshotChunk{
				Index: 0, Chunk: []byte{0}, Sender: "REDACTED",
			}).Once().Return(&atci.ResponseApplySnapshotChunk{Result: atci.ResponseApplySnapshotChunk_ACCEPT}, nil)
			connSnapshot.On("REDACTED", mock.Anything, &atci.RequestApplySnapshotChunk{
				Index: 1, Chunk: []byte{1}, Sender: "REDACTED",
			}).Once().Return(&atci.ResponseApplySnapshotChunk{Result: atci.ResponseApplySnapshotChunk_ACCEPT}, nil)
			connSnapshot.On("REDACTED", mock.Anything, &atci.RequestApplySnapshotChunk{
				Index: 2, Chunk: []byte{2}, Sender: "REDACTED",
			}).Once().Return(&atci.ResponseApplySnapshotChunk{
				Result:        tc.result,
				RejectSenders: []string{string(peerB.ID())},
			}, nil)

			//
			if tc.result == atci.ResponseApplySnapshotChunk_RETRY {
				connSnapshot.On("REDACTED", mock.Anything, &atci.RequestApplySnapshotChunk{
					Index: 2, Chunk: []byte{2}, Sender: "REDACTED",
				}).Once().Return(&atci.ResponseApplySnapshotChunk{Result: atci.ResponseApplySnapshotChunk_ACCEPT}, nil)
			}

			//
			//
			//
			go func() {
				syncer.applyChunks(chunks) //
			}()

			time.Sleep(50 * time.Millisecond)

			s1peers := syncer.snapshots.GetPeers(s1)
			assert.Len(t, s1peers, 2)
			assert.EqualValues(t, "REDACTED", s1peers[0].ID())
			assert.EqualValues(t, "REDACTED", s1peers[1].ID())

			syncer.snapshots.GetPeers(s1)
			assert.Len(t, s1peers, 2)
			assert.EqualValues(t, "REDACTED", s1peers[0].ID())
			assert.EqualValues(t, "REDACTED", s1peers[1].ID())

			err = chunks.Close()
			require.NoError(t, err)
		})
	}
}

func TestSyncer_verifyApp(t *testing.T) {
	boom := errors.New("REDACTED")
	const appVersion = 9
	appVersionMismatchErr := errors.New("REDACTED")
	s := &snapshot{Height: 3, Format: 1, Chunks: 5, Hash: []byte{1, 2, 3}, trustedAppHash: []byte("REDACTED")}

	testcases := map[string]struct {
		response  *atci.ResponseInfo
		err       error
		expectErr error
	}{
		"REDACTED": {&atci.ResponseInfo{
			LastBlockHeight:  3,
			LastBlockAppHash: []byte("REDACTED"),
			AppVersion:       appVersion,
		}, nil, nil},
		"REDACTED": {&atci.ResponseInfo{
			LastBlockHeight:  3,
			LastBlockAppHash: []byte("REDACTED"),
			AppVersion:       2,
		}, nil, appVersionMismatchErr},
		"REDACTED": {&atci.ResponseInfo{
			LastBlockHeight:  5,
			LastBlockAppHash: []byte("REDACTED"),
			AppVersion:       appVersion,
		}, nil, errVerifyFailed},
		"REDACTED": {&atci.ResponseInfo{
			LastBlockHeight:  3,
			LastBlockAppHash: []byte("REDACTED"),
			AppVersion:       appVersion,
		}, nil, errVerifyFailed},
		"REDACTED": {nil, boom, boom},
	}
	for name, tc := range testcases {

		t.Run(name, func(t *testing.T) {
			connQuery := &gatwaymocks.AppConnQuery{}
			connSnapshot := &gatwaymocks.AppConnSnapshot{}
			stateProvider := &simulations.StateProvider{}

			cfg := settings.DefaultStateSyncConfig()
			syncer := newSyncer(*cfg, log.NewNopLogger(), connSnapshot, connQuery, stateProvider, "REDACTED")

			connQuery.On("REDACTED", mock.Anything, gateway.RequestInfo).Return(tc.response, tc.err)
			err := syncer.verifyApp(s, appVersion)
			unwrapped := errors.Unwrap(err)
			if unwrapped != nil {
				err = unwrapped
			}
			require.Equal(t, tc.expectErr, err)
		})
	}
}

func toABCI(s *snapshot) *atci.Snapshot {
	return &atci.Snapshot{
		Height:   s.Height,
		Format:   s.Format,
		Chunks:   s.Chunks,
		Hash:     s.Hash,
		Metadata: s.Metadata,
	}
}
