package statusconnect

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/p2p"
	netpeersims "github.com/valkyrieworks/p2p/simulations"
)

func Verifymirror_Key(t *testing.T) {
	verifyscenarios := map[string]struct {
		adjust func(*mirror)
	}{
		"REDACTED":      {func(s *mirror) { s.Level = 9 }},
		"REDACTED":      {func(s *mirror) { s.Layout = 9 }},
		"REDACTED": {func(s *mirror) { s.Segments = 9 }},
		"REDACTED":        {func(s *mirror) { s.Digest = []byte{9} }},
		"REDACTED":     {func(s *mirror) { s.Metainfo = nil }},
	}
	for label, tc := range verifyscenarios {

		t.Run(label, func(t *testing.T) {
			s := mirror{
				Level:   3,
				Layout:   1,
				Segments:   7,
				Digest:     []byte{1, 2, 3},
				Metainfo: []byte{255},
			}
			prior := s.Key()
			tc.adjust(&s)
			after := s.Key()
			assert.NotEqual(t, prior, after)
		})
	}
}

func Verifymirrordepot_Append(t *testing.T) {
	node := &netpeersims.Node{}
	node.On("REDACTED").Return(p2p.ID("REDACTED"))

	//
	depository := newMirrorDepository()
	appended, err := depository.Add(node, &mirror{
		Level: 1,
		Layout: 1,
		Segments: 1,
		Digest:   []byte{1},
	})
	require.NoError(t, err)
	assert.True(t, appended)

	//
	anotherNode := &netpeersims.Node{}
	anotherNode.On("REDACTED").Return(p2p.ID("REDACTED"))
	appended, err = depository.Add(node, &mirror{
		Level: 1,
		Layout: 1,
		Segments: 1,
		Digest:   []byte{1},
	})
	require.NoError(t, err)
	assert.False(t, appended)

	//
	mirror := depository.Optimal()
	require.NotNil(t, mirror)
}

func Verifymirrordepot_Getnode(t *testing.T) {
	depository := newMirrorDepository()

	s := &mirror{Level: 1, Layout: 1, Segments: 1, Digest: []byte{1}}
	nodeA := &netpeersims.Node{}
	nodeA.On("REDACTED").Return(p2p.ID("REDACTED"))
	nodeBYTE := &netpeersims.Node{}
	nodeBYTE.On("REDACTED").Return(p2p.ID("REDACTED"))

	_, err := depository.Add(nodeA, s)
	require.NoError(t, err)
	_, err = depository.Add(nodeBYTE, s)
	require.NoError(t, err)
	_, err = depository.Add(nodeA, &mirror{Level: 2, Layout: 1, Segments: 1, Digest: []byte{1}})
	require.NoError(t, err)

	//
	viewedA := false
	viewedBYTE := false
	for !viewedA || !viewedBYTE {
		node := depository.FetchNode(s)
		switch node.ID() {
		case p2p.ID("REDACTED"):
			viewedA = true
		case p2p.ID("REDACTED"):
			viewedBYTE = true
		}
	}

	//
	node := depository.FetchNode(&mirror{Level: 9, Layout: 9})
	assert.Nil(t, node)
}

func Verifymirrordepot_Getnodes(t *testing.T) {
	depository := newMirrorDepository()

	s := &mirror{Level: 1, Layout: 1, Segments: 1, Digest: []byte{1}}
	nodeA := &netpeersims.Node{}
	nodeA.On("REDACTED").Return(p2p.ID("REDACTED"))
	nodeBYTE := &netpeersims.Node{}
	nodeBYTE.On("REDACTED").Return(p2p.ID("REDACTED"))

	_, err := depository.Add(nodeA, s)
	require.NoError(t, err)
	_, err = depository.Add(nodeBYTE, s)
	require.NoError(t, err)
	_, err = depository.Add(nodeA, &mirror{Level: 2, Layout: 1, Segments: 1, Digest: []byte{2}})
	require.NoError(t, err)

	nodes := depository.FetchNodes(s)
	assert.Len(t, nodes, 2)
	assert.EqualValues(t, "REDACTED", nodes[0].ID())
	assert.EqualValues(t, "REDACTED", nodes[1].ID())
}

func Verifymirrordepot_Rated_Optimal(t *testing.T) {
	depository := newMirrorDepository()

	//
	//
	//
	anticipateMirrors := []struct {
		mirror *mirror
		nodes    []string
	}{
		{&mirror{Level: 2, Layout: 2, Segments: 4, Digest: []byte{1, 3}}, []string{"REDACTED", "REDACTED", "REDACTED"}},
		{&mirror{Level: 2, Layout: 2, Segments: 5, Digest: []byte{1, 2}}, []string{"REDACTED"}},
		{&mirror{Level: 2, Layout: 1, Segments: 3, Digest: []byte{1, 2}}, []string{"REDACTED", "REDACTED"}},
		{&mirror{Level: 1, Layout: 2, Segments: 5, Digest: []byte{1, 2}}, []string{"REDACTED", "REDACTED"}},
		{&mirror{Level: 1, Layout: 1, Segments: 4, Digest: []byte{1, 2}}, []string{"REDACTED", "REDACTED", "REDACTED"}},
	}

	//
	for i := len(anticipateMirrors) - 1; i >= 0; i-- {
		for _, nodeUID := range anticipateMirrors[i].nodes {
			node := &netpeersims.Node{}
			node.On("REDACTED").Return(p2p.ID(nodeUID))
			_, err := depository.Add(node, anticipateMirrors[i].mirror)
			require.NoError(t, err)
		}
	}

	//
	rated := depository.Rated()
	assert.Len(t, rated, len(anticipateMirrors))
	for i := range rated {
		assert.Equal(t, anticipateMirrors[i].mirror, rated[i])
	}

	//
	for i := range anticipateMirrors {
		mirror := anticipateMirrors[i].mirror
		require.Equal(t, mirror, depository.Optimal())
		depository.Decline(mirror)
	}
	assert.Nil(t, depository.Optimal())
}

func Verifymirrordepot_Decline(t *testing.T) {
	depository := newMirrorDepository()
	node := &netpeersims.Node{}
	node.On("REDACTED").Return(p2p.ID("REDACTED"))

	mirrors := []*mirror{
		{Level: 2, Layout: 2, Segments: 1, Digest: []byte{1, 2}},
		{Level: 2, Layout: 1, Segments: 1, Digest: []byte{1, 2}},
		{Level: 1, Layout: 2, Segments: 1, Digest: []byte{1, 2}},
		{Level: 1, Layout: 1, Segments: 1, Digest: []byte{1, 2}},
	}
	for _, s := range mirrors {
		_, err := depository.Add(node, s)
		require.NoError(t, err)
	}

	depository.Decline(mirrors[0])
	assert.Equal(t, mirrors[1:], depository.Rated())

	appended, err := depository.Add(node, mirrors[0])
	require.NoError(t, err)
	assert.False(t, appended)

	appended, err = depository.Add(node, &mirror{Level: 3, Layout: 3, Segments: 1, Digest: []byte{1}})
	require.NoError(t, err)
	assert.True(t, appended)
}

func Verifymirrordepot_Denyformat(t *testing.T) {
	depository := newMirrorDepository()
	node := &netpeersims.Node{}
	node.On("REDACTED").Return(p2p.ID("REDACTED"))

	mirrors := []*mirror{
		{Level: 2, Layout: 2, Segments: 1, Digest: []byte{1, 2}},
		{Level: 2, Layout: 1, Segments: 1, Digest: []byte{1, 2}},
		{Level: 1, Layout: 2, Segments: 1, Digest: []byte{1, 2}},
		{Level: 1, Layout: 1, Segments: 1, Digest: []byte{1, 2}},
	}
	for _, s := range mirrors {
		_, err := depository.Add(node, s)
		require.NoError(t, err)
	}

	depository.DeclineLayout(1)
	assert.Equal(t, []*mirror{mirrors[0], mirrors[2]}, depository.Rated())

	appended, err := depository.Add(node, &mirror{Level: 3, Layout: 1, Segments: 1, Digest: []byte{1}})
	require.NoError(t, err)
	assert.False(t, appended)
	assert.Equal(t, []*mirror{mirrors[0], mirrors[2]}, depository.Rated())

	appended, err = depository.Add(node, &mirror{Level: 3, Layout: 3, Segments: 1, Digest: []byte{1}})
	require.NoError(t, err)
	assert.True(t, appended)
}

func Verifymirrordepot_Denynode(t *testing.T) {
	depository := newMirrorDepository()

	nodeA := &netpeersims.Node{}
	nodeA.On("REDACTED").Return(p2p.ID("REDACTED"))
	nodeBYTE := &netpeersims.Node{}
	nodeBYTE.On("REDACTED").Return(p2p.ID("REDACTED"))

	s1 := &mirror{Level: 1, Layout: 1, Segments: 1, Digest: []byte{1}}
	s2 := &mirror{Level: 2, Layout: 1, Segments: 1, Digest: []byte{2}}
	s3 := &mirror{Level: 3, Layout: 1, Segments: 1, Digest: []byte{2}}

	_, err := depository.Add(nodeA, s1)
	require.NoError(t, err)
	_, err = depository.Add(nodeA, s2)
	require.NoError(t, err)

	_, err = depository.Add(nodeBYTE, s2)
	require.NoError(t, err)
	_, err = depository.Add(nodeBYTE, s3)
	require.NoError(t, err)

	depository.DeclineNode(nodeA.ID())

	assert.Empty(t, depository.FetchNodes(s1))

	nodestwo := depository.FetchNodes(s2)
	assert.Len(t, nodestwo, 1)
	assert.EqualValues(t, "REDACTED", nodestwo[0].ID())

	nodethree := depository.FetchNodes(s2)
	assert.Len(t, nodethree, 1)
	assert.EqualValues(t, "REDACTED", nodethree[0].ID())

	//
	_, err = depository.Add(nodeA, s1)
	require.NoError(t, err)
	assert.Empty(t, depository.FetchNodes(s1))
}

func Verifymirrordepot_Deletenode(t *testing.T) {
	depository := newMirrorDepository()

	nodeA := &netpeersims.Node{}
	nodeA.On("REDACTED").Return(p2p.ID("REDACTED"))
	nodeBYTE := &netpeersims.Node{}
	nodeBYTE.On("REDACTED").Return(p2p.ID("REDACTED"))

	s1 := &mirror{Level: 1, Layout: 1, Segments: 1, Digest: []byte{1}}
	s2 := &mirror{Level: 2, Layout: 1, Segments: 1, Digest: []byte{2}}

	_, err := depository.Add(nodeA, s1)
	require.NoError(t, err)
	_, err = depository.Add(nodeA, s2)
	require.NoError(t, err)
	_, err = depository.Add(nodeBYTE, s1)
	require.NoError(t, err)

	depository.DeleteNode(nodeA.ID())

	nodesone := depository.FetchNodes(s1)
	assert.Len(t, nodesone, 1)
	assert.EqualValues(t, "REDACTED", nodesone[0].ID())

	nodestwo := depository.FetchNodes(s2)
	assert.Empty(t, nodestwo)

	//
	_, err = depository.Add(nodeA, s1)
	require.NoError(t, err)
	nodesone = depository.FetchNodes(s1)
	assert.Len(t, nodesone, 2)
	assert.EqualValues(t, "REDACTED", nodesone[0].ID())
	assert.EqualValues(t, "REDACTED", nodesone[1].ID())
}
