package statereplication

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/p2p"
	netmocks "github.com/valkyrieworks/p2p/simulations"
)

func TestSnapshot_Key(t *testing.T) {
	testcases := map[string]struct {
		modify func(*snapshot)
	}{
		"REDACTED":      {func(s *snapshot) { s.Height = 9 }},
		"REDACTED":      {func(s *snapshot) { s.Format = 9 }},
		"REDACTED": {func(s *snapshot) { s.Chunks = 9 }},
		"REDACTED":        {func(s *snapshot) { s.Hash = []byte{9} }},
		"REDACTED":     {func(s *snapshot) { s.Metadata = nil }},
	}
	for name, tc := range testcases {

		t.Run(name, func(t *testing.T) {
			s := snapshot{
				Height:   3,
				Format:   1,
				Chunks:   7,
				Hash:     []byte{1, 2, 3},
				Metadata: []byte{255},
			}
			before := s.Key()
			tc.modify(&s)
			after := s.Key()
			assert.NotEqual(t, before, after)
		})
	}
}

func TestSnapshotPool_Add(t *testing.T) {
	peer := &netmocks.Peer{}
	peer.On("REDACTED").Return(p2p.ID("REDACTED"))

	//
	pool := newSnapshotPool()
	added, err := pool.Add(peer, &snapshot{
		Height: 1,
		Format: 1,
		Chunks: 1,
		Hash:   []byte{1},
	})
	require.NoError(t, err)
	assert.True(t, added)

	//
	otherPeer := &netmocks.Peer{}
	otherPeer.On("REDACTED").Return(p2p.ID("REDACTED"))
	added, err = pool.Add(peer, &snapshot{
		Height: 1,
		Format: 1,
		Chunks: 1,
		Hash:   []byte{1},
	})
	require.NoError(t, err)
	assert.False(t, added)

	//
	snapshot := pool.Best()
	require.NotNil(t, snapshot)
}

func TestSnapshotPool_GetPeer(t *testing.T) {
	pool := newSnapshotPool()

	s := &snapshot{Height: 1, Format: 1, Chunks: 1, Hash: []byte{1}}
	peerA := &netmocks.Peer{}
	peerA.On("REDACTED").Return(p2p.ID("REDACTED"))
	peerB := &netmocks.Peer{}
	peerB.On("REDACTED").Return(p2p.ID("REDACTED"))

	_, err := pool.Add(peerA, s)
	require.NoError(t, err)
	_, err = pool.Add(peerB, s)
	require.NoError(t, err)
	_, err = pool.Add(peerA, &snapshot{Height: 2, Format: 1, Chunks: 1, Hash: []byte{1}})
	require.NoError(t, err)

	//
	seenA := false
	seenB := false
	for !seenA || !seenB {
		peer := pool.GetPeer(s)
		switch peer.ID() {
		case p2p.ID("REDACTED"):
			seenA = true
		case p2p.ID("REDACTED"):
			seenB = true
		}
	}

	//
	peer := pool.GetPeer(&snapshot{Height: 9, Format: 9})
	assert.Nil(t, peer)
}

func TestSnapshotPool_GetPeers(t *testing.T) {
	pool := newSnapshotPool()

	s := &snapshot{Height: 1, Format: 1, Chunks: 1, Hash: []byte{1}}
	peerA := &netmocks.Peer{}
	peerA.On("REDACTED").Return(p2p.ID("REDACTED"))
	peerB := &netmocks.Peer{}
	peerB.On("REDACTED").Return(p2p.ID("REDACTED"))

	_, err := pool.Add(peerA, s)
	require.NoError(t, err)
	_, err = pool.Add(peerB, s)
	require.NoError(t, err)
	_, err = pool.Add(peerA, &snapshot{Height: 2, Format: 1, Chunks: 1, Hash: []byte{2}})
	require.NoError(t, err)

	peers := pool.GetPeers(s)
	assert.Len(t, peers, 2)
	assert.EqualValues(t, "REDACTED", peers[0].ID())
	assert.EqualValues(t, "REDACTED", peers[1].ID())
}

func TestSnapshotPool_Ranked_Best(t *testing.T) {
	pool := newSnapshotPool()

	//
	//
	//
	expectSnapshots := []struct {
		snapshot *snapshot
		peers    []string
	}{
		{&snapshot{Height: 2, Format: 2, Chunks: 4, Hash: []byte{1, 3}}, []string{"REDACTED", "REDACTED", "REDACTED"}},
		{&snapshot{Height: 2, Format: 2, Chunks: 5, Hash: []byte{1, 2}}, []string{"REDACTED"}},
		{&snapshot{Height: 2, Format: 1, Chunks: 3, Hash: []byte{1, 2}}, []string{"REDACTED", "REDACTED"}},
		{&snapshot{Height: 1, Format: 2, Chunks: 5, Hash: []byte{1, 2}}, []string{"REDACTED", "REDACTED"}},
		{&snapshot{Height: 1, Format: 1, Chunks: 4, Hash: []byte{1, 2}}, []string{"REDACTED", "REDACTED", "REDACTED"}},
	}

	//
	for i := len(expectSnapshots) - 1; i >= 0; i-- {
		for _, peerID := range expectSnapshots[i].peers {
			peer := &netmocks.Peer{}
			peer.On("REDACTED").Return(p2p.ID(peerID))
			_, err := pool.Add(peer, expectSnapshots[i].snapshot)
			require.NoError(t, err)
		}
	}

	//
	ranked := pool.Ranked()
	assert.Len(t, ranked, len(expectSnapshots))
	for i := range ranked {
		assert.Equal(t, expectSnapshots[i].snapshot, ranked[i])
	}

	//
	for i := range expectSnapshots {
		snapshot := expectSnapshots[i].snapshot
		require.Equal(t, snapshot, pool.Best())
		pool.Reject(snapshot)
	}
	assert.Nil(t, pool.Best())
}

func TestSnapshotPool_Reject(t *testing.T) {
	pool := newSnapshotPool()
	peer := &netmocks.Peer{}
	peer.On("REDACTED").Return(p2p.ID("REDACTED"))

	snapshots := []*snapshot{
		{Height: 2, Format: 2, Chunks: 1, Hash: []byte{1, 2}},
		{Height: 2, Format: 1, Chunks: 1, Hash: []byte{1, 2}},
		{Height: 1, Format: 2, Chunks: 1, Hash: []byte{1, 2}},
		{Height: 1, Format: 1, Chunks: 1, Hash: []byte{1, 2}},
	}
	for _, s := range snapshots {
		_, err := pool.Add(peer, s)
		require.NoError(t, err)
	}

	pool.Reject(snapshots[0])
	assert.Equal(t, snapshots[1:], pool.Ranked())

	added, err := pool.Add(peer, snapshots[0])
	require.NoError(t, err)
	assert.False(t, added)

	added, err = pool.Add(peer, &snapshot{Height: 3, Format: 3, Chunks: 1, Hash: []byte{1}})
	require.NoError(t, err)
	assert.True(t, added)
}

func TestSnapshotPool_RejectFormat(t *testing.T) {
	pool := newSnapshotPool()
	peer := &netmocks.Peer{}
	peer.On("REDACTED").Return(p2p.ID("REDACTED"))

	snapshots := []*snapshot{
		{Height: 2, Format: 2, Chunks: 1, Hash: []byte{1, 2}},
		{Height: 2, Format: 1, Chunks: 1, Hash: []byte{1, 2}},
		{Height: 1, Format: 2, Chunks: 1, Hash: []byte{1, 2}},
		{Height: 1, Format: 1, Chunks: 1, Hash: []byte{1, 2}},
	}
	for _, s := range snapshots {
		_, err := pool.Add(peer, s)
		require.NoError(t, err)
	}

	pool.RejectFormat(1)
	assert.Equal(t, []*snapshot{snapshots[0], snapshots[2]}, pool.Ranked())

	added, err := pool.Add(peer, &snapshot{Height: 3, Format: 1, Chunks: 1, Hash: []byte{1}})
	require.NoError(t, err)
	assert.False(t, added)
	assert.Equal(t, []*snapshot{snapshots[0], snapshots[2]}, pool.Ranked())

	added, err = pool.Add(peer, &snapshot{Height: 3, Format: 3, Chunks: 1, Hash: []byte{1}})
	require.NoError(t, err)
	assert.True(t, added)
}

func TestSnapshotPool_RejectPeer(t *testing.T) {
	pool := newSnapshotPool()

	peerA := &netmocks.Peer{}
	peerA.On("REDACTED").Return(p2p.ID("REDACTED"))
	peerB := &netmocks.Peer{}
	peerB.On("REDACTED").Return(p2p.ID("REDACTED"))

	s1 := &snapshot{Height: 1, Format: 1, Chunks: 1, Hash: []byte{1}}
	s2 := &snapshot{Height: 2, Format: 1, Chunks: 1, Hash: []byte{2}}
	s3 := &snapshot{Height: 3, Format: 1, Chunks: 1, Hash: []byte{2}}

	_, err := pool.Add(peerA, s1)
	require.NoError(t, err)
	_, err = pool.Add(peerA, s2)
	require.NoError(t, err)

	_, err = pool.Add(peerB, s2)
	require.NoError(t, err)
	_, err = pool.Add(peerB, s3)
	require.NoError(t, err)

	pool.RejectPeer(peerA.ID())

	assert.Empty(t, pool.GetPeers(s1))

	peers2 := pool.GetPeers(s2)
	assert.Len(t, peers2, 1)
	assert.EqualValues(t, "REDACTED", peers2[0].ID())

	peers3 := pool.GetPeers(s2)
	assert.Len(t, peers3, 1)
	assert.EqualValues(t, "REDACTED", peers3[0].ID())

	//
	_, err = pool.Add(peerA, s1)
	require.NoError(t, err)
	assert.Empty(t, pool.GetPeers(s1))
}

func TestSnapshotPool_RemovePeer(t *testing.T) {
	pool := newSnapshotPool()

	peerA := &netmocks.Peer{}
	peerA.On("REDACTED").Return(p2p.ID("REDACTED"))
	peerB := &netmocks.Peer{}
	peerB.On("REDACTED").Return(p2p.ID("REDACTED"))

	s1 := &snapshot{Height: 1, Format: 1, Chunks: 1, Hash: []byte{1}}
	s2 := &snapshot{Height: 2, Format: 1, Chunks: 1, Hash: []byte{2}}

	_, err := pool.Add(peerA, s1)
	require.NoError(t, err)
	_, err = pool.Add(peerA, s2)
	require.NoError(t, err)
	_, err = pool.Add(peerB, s1)
	require.NoError(t, err)

	pool.RemovePeer(peerA.ID())

	peers1 := pool.GetPeers(s1)
	assert.Len(t, peers1, 1)
	assert.EqualValues(t, "REDACTED", peers1[0].ID())

	peers2 := pool.GetPeers(s2)
	assert.Empty(t, peers2)

	//
	_, err = pool.Add(peerA, s1)
	require.NoError(t, err)
	peers1 = pool.GetPeers(s1)
	assert.Len(t, peers1, 2)
	assert.EqualValues(t, "REDACTED", peers1[0].ID())
	assert.EqualValues(t, "REDACTED", peers1[1].ID())
}
