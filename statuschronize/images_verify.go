package statuschronize

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	netmocks "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p/simulations"
)

func Verifyimage_Token(t *testing.T) {
	verifycases := map[string]struct {
		alter func(*image)
	}{
		"REDACTED":      {func(s *image) { s.Altitude = 9 }},
		"REDACTED":      {func(s *image) { s.Layout = 9 }},
		"REDACTED": {func(s *image) { s.Segments = 9 }},
		"REDACTED":        {func(s *image) { s.Digest = []byte{9} }},
		"REDACTED":     {func(s *image) { s.Attributes = nil }},
	}
	for alias, tc := range verifycases {

		t.Run(alias, func(t *testing.T) {
			s := image{
				Altitude:   3,
				Layout:   1,
				Segments:   7,
				Digest:     []byte{1, 2, 3},
				Attributes: []byte{255},
			}
			prior := s.Key()
			tc.alter(&s)
			subsequent := s.Key()
			assert.NotEqual(t, prior, subsequent)
		})
	}
}

func Verifyimagehub_Append(t *testing.T) {
	node := &netmocks.Node{}
	node.On("REDACTED").Return(p2p.ID("REDACTED"))

	//
	hub := freshImageHub()
	appended, err := hub.Add(node, &image{
		Altitude: 1,
		Layout: 1,
		Segments: 1,
		Digest:   []byte{1},
	})
	require.NoError(t, err)
	assert.True(t, appended)

	//
	anotherNode := &netmocks.Node{}
	anotherNode.On("REDACTED").Return(p2p.ID("REDACTED"))
	appended, err = hub.Add(node, &image{
		Altitude: 1,
		Layout: 1,
		Segments: 1,
		Digest:   []byte{1},
	})
	require.NoError(t, err)
	assert.False(t, appended)

	//
	image := hub.Optimal()
	require.NotNil(t, image)
}

func Verifyimagehub_Getnode(t *testing.T) {
	hub := freshImageHub()

	s := &image{Altitude: 1, Layout: 1, Segments: 1, Digest: []byte{1}}
	nodeAN := &netmocks.Node{}
	nodeAN.On("REDACTED").Return(p2p.ID("REDACTED"))
	nodeBYTE := &netmocks.Node{}
	nodeBYTE.On("REDACTED").Return(p2p.ID("REDACTED"))

	_, err := hub.Add(nodeAN, s)
	require.NoError(t, err)
	_, err = hub.Add(nodeBYTE, s)
	require.NoError(t, err)
	_, err = hub.Add(nodeAN, &image{Altitude: 2, Layout: 1, Segments: 1, Digest: []byte{1}})
	require.NoError(t, err)

	//
	observedAN := false
	observedBYTE := false
	for !observedAN || !observedBYTE {
		node := hub.ObtainNode(s)
		switch node.ID() {
		case p2p.ID("REDACTED"):
			observedAN = true
		case p2p.ID("REDACTED"):
			observedBYTE = true
		}
	}

	//
	node := hub.ObtainNode(&image{Altitude: 9, Layout: 9})
	assert.Nil(t, node)
}

func Verifyimagehub_Getnodes(t *testing.T) {
	hub := freshImageHub()

	s := &image{Altitude: 1, Layout: 1, Segments: 1, Digest: []byte{1}}
	nodeAN := &netmocks.Node{}
	nodeAN.On("REDACTED").Return(p2p.ID("REDACTED"))
	nodeBYTE := &netmocks.Node{}
	nodeBYTE.On("REDACTED").Return(p2p.ID("REDACTED"))

	_, err := hub.Add(nodeAN, s)
	require.NoError(t, err)
	_, err = hub.Add(nodeBYTE, s)
	require.NoError(t, err)
	_, err = hub.Add(nodeAN, &image{Altitude: 2, Layout: 1, Segments: 1, Digest: []byte{2}})
	require.NoError(t, err)

	nodes := hub.ObtainNodes(s)
	assert.Len(t, nodes, 2)
	assert.EqualValues(t, "REDACTED", nodes[0].ID())
	assert.EqualValues(t, "REDACTED", nodes[1].ID())
}

func Verifyimagehub_Ordered_Optimal(t *testing.T) {
	hub := freshImageHub()

	//
	//
	//
	anticipateImages := []struct {
		image *image
		nodes    []string
	}{
		{&image{Altitude: 2, Layout: 2, Segments: 4, Digest: []byte{1, 3}}, []string{"REDACTED", "REDACTED", "REDACTED"}},
		{&image{Altitude: 2, Layout: 2, Segments: 5, Digest: []byte{1, 2}}, []string{"REDACTED"}},
		{&image{Altitude: 2, Layout: 1, Segments: 3, Digest: []byte{1, 2}}, []string{"REDACTED", "REDACTED"}},
		{&image{Altitude: 1, Layout: 2, Segments: 5, Digest: []byte{1, 2}}, []string{"REDACTED", "REDACTED"}},
		{&image{Altitude: 1, Layout: 1, Segments: 4, Digest: []byte{1, 2}}, []string{"REDACTED", "REDACTED", "REDACTED"}},
	}

	//
	for i := len(anticipateImages) - 1; i >= 0; i-- {
		for _, nodeUUID := range anticipateImages[i].nodes {
			node := &netmocks.Node{}
			node.On("REDACTED").Return(p2p.ID(nodeUUID))
			_, err := hub.Add(node, anticipateImages[i].image)
			require.NoError(t, err)
		}
	}

	//
	ordered := hub.Ordered()
	assert.Len(t, ordered, len(anticipateImages))
	for i := range ordered {
		assert.Equal(t, anticipateImages[i].image, ordered[i])
	}

	//
	for i := range anticipateImages {
		image := anticipateImages[i].image
		require.Equal(t, image, hub.Optimal())
		hub.Decline(image)
	}
	assert.Nil(t, hub.Optimal())
}

func Verifyimagehub_Decline(t *testing.T) {
	hub := freshImageHub()
	node := &netmocks.Node{}
	node.On("REDACTED").Return(p2p.ID("REDACTED"))

	images := []*image{
		{Altitude: 2, Layout: 2, Segments: 1, Digest: []byte{1, 2}},
		{Altitude: 2, Layout: 1, Segments: 1, Digest: []byte{1, 2}},
		{Altitude: 1, Layout: 2, Segments: 1, Digest: []byte{1, 2}},
		{Altitude: 1, Layout: 1, Segments: 1, Digest: []byte{1, 2}},
	}
	for _, s := range images {
		_, err := hub.Add(node, s)
		require.NoError(t, err)
	}

	hub.Decline(images[0])
	assert.Equal(t, images[1:], hub.Ordered())

	appended, err := hub.Add(node, images[0])
	require.NoError(t, err)
	assert.False(t, appended)

	appended, err = hub.Add(node, &image{Altitude: 3, Layout: 3, Segments: 1, Digest: []byte{1}})
	require.NoError(t, err)
	assert.True(t, appended)
}

func Verifyimagehub_Denyformat(t *testing.T) {
	hub := freshImageHub()
	node := &netmocks.Node{}
	node.On("REDACTED").Return(p2p.ID("REDACTED"))

	images := []*image{
		{Altitude: 2, Layout: 2, Segments: 1, Digest: []byte{1, 2}},
		{Altitude: 2, Layout: 1, Segments: 1, Digest: []byte{1, 2}},
		{Altitude: 1, Layout: 2, Segments: 1, Digest: []byte{1, 2}},
		{Altitude: 1, Layout: 1, Segments: 1, Digest: []byte{1, 2}},
	}
	for _, s := range images {
		_, err := hub.Add(node, s)
		require.NoError(t, err)
	}

	hub.DeclineLayout(1)
	assert.Equal(t, []*image{images[0], images[2]}, hub.Ordered())

	appended, err := hub.Add(node, &image{Altitude: 3, Layout: 1, Segments: 1, Digest: []byte{1}})
	require.NoError(t, err)
	assert.False(t, appended)
	assert.Equal(t, []*image{images[0], images[2]}, hub.Ordered())

	appended, err = hub.Add(node, &image{Altitude: 3, Layout: 3, Segments: 1, Digest: []byte{1}})
	require.NoError(t, err)
	assert.True(t, appended)
}

func Verifyimagehub_Denynode(t *testing.T) {
	hub := freshImageHub()

	nodeAN := &netmocks.Node{}
	nodeAN.On("REDACTED").Return(p2p.ID("REDACTED"))
	nodeBYTE := &netmocks.Node{}
	nodeBYTE.On("REDACTED").Return(p2p.ID("REDACTED"))

	s1 := &image{Altitude: 1, Layout: 1, Segments: 1, Digest: []byte{1}}
	s2 := &image{Altitude: 2, Layout: 1, Segments: 1, Digest: []byte{2}}
	s3 := &image{Altitude: 3, Layout: 1, Segments: 1, Digest: []byte{2}}

	_, err := hub.Add(nodeAN, s1)
	require.NoError(t, err)
	_, err = hub.Add(nodeAN, s2)
	require.NoError(t, err)

	_, err = hub.Add(nodeBYTE, s2)
	require.NoError(t, err)
	_, err = hub.Add(nodeBYTE, s3)
	require.NoError(t, err)

	hub.DeclineNode(nodeAN.ID())

	assert.Empty(t, hub.ObtainNodes(s1))

	nodes2 := hub.ObtainNodes(s2)
	assert.Len(t, nodes2, 1)
	assert.EqualValues(t, "REDACTED", nodes2[0].ID())

	nodes3 := hub.ObtainNodes(s2)
	assert.Len(t, nodes3, 1)
	assert.EqualValues(t, "REDACTED", nodes3[0].ID())

	//
	_, err = hub.Add(nodeAN, s1)
	require.NoError(t, err)
	assert.Empty(t, hub.ObtainNodes(s1))
}

func Verifyimagehub_Removenode(t *testing.T) {
	hub := freshImageHub()

	nodeAN := &netmocks.Node{}
	nodeAN.On("REDACTED").Return(p2p.ID("REDACTED"))
	nodeBYTE := &netmocks.Node{}
	nodeBYTE.On("REDACTED").Return(p2p.ID("REDACTED"))

	s1 := &image{Altitude: 1, Layout: 1, Segments: 1, Digest: []byte{1}}
	s2 := &image{Altitude: 2, Layout: 1, Segments: 1, Digest: []byte{2}}

	_, err := hub.Add(nodeAN, s1)
	require.NoError(t, err)
	_, err = hub.Add(nodeAN, s2)
	require.NoError(t, err)
	_, err = hub.Add(nodeBYTE, s1)
	require.NoError(t, err)

	hub.DiscardNode(nodeAN.ID())

	nodes1 := hub.ObtainNodes(s1)
	assert.Len(t, nodes1, 1)
	assert.EqualValues(t, "REDACTED", nodes1[0].ID())

	nodes2 := hub.ObtainNodes(s2)
	assert.Empty(t, nodes2)

	//
	_, err = hub.Add(nodeAN, s1)
	require.NoError(t, err)
	nodes1 = hub.ObtainNodes(s1)
	assert.Len(t, nodes1, 2)
	assert.EqualValues(t, "REDACTED", nodes1[0].ID())
	assert.EqualValues(t, "REDACTED", nodes1[1].ID())
}
