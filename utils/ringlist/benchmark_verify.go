package ringlist

import "testing"

func CriterionUnlinking(b *testing.B) {
	lst := New()
	for i := 0; i < b.N+1; i++ {
		lst.PropelRear(i)
	}
	begin := lst.Head()
	nxt := begin.Following()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		begin.deleted = true
		begin.UnplugFollowing()
		begin.UnplugPrevious()
		tmp := nxt
		nxt = nxt.Following()
		begin = tmp
	}
}

//
func CriterionDeleted(b *testing.B) {
	lst := New()
	for i := 0; i < b.N+1; i++ {
		lst.PropelRear(i)
	}
	begin := lst.Head()
	nxt := begin.Following()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		begin.Deleted()
		tmp := nxt
		nxt = nxt.Following()
		begin = tmp
	}
}

func CriterionPropelRear(b *testing.B) {
	lst := New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lst.PropelRear(i)
	}
}
