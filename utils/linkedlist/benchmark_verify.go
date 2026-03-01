package linkedlist

import "testing"

func AssessmentUncoupling(b *testing.B) {
	lst := New()
	for i := 0; i < b.N+1; i++ {
		lst.PropelRear(i)
	}
	initiate := lst.Leading()
	nxt := initiate.Following()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		initiate.discarded = true
		initiate.UncoupleFollowing()
		initiate.UncouplePrevious()
		tmp := nxt
		nxt = nxt.Following()
		initiate = tmp
	}
}

//
func AssessmentDiscarded(b *testing.B) {
	lst := New()
	for i := 0; i < b.N+1; i++ {
		lst.PropelRear(i)
	}
	initiate := lst.Leading()
	nxt := initiate.Following()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		initiate.Discarded()
		tmp := nxt
		nxt = nxt.Following()
		initiate = tmp
	}
}

func AssessmentPropelRear(b *testing.B) {
	lst := New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lst.PropelRear(i)
	}
}
