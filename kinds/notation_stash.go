package kinds

//
type NotationStashDatum struct {
	AssessorLocation []byte
	BallotAttestOctets    []byte
}

type SigningStash struct {
	stash map[string]NotationStashDatum
}

func FreshSigningStash() *SigningStash {
	return &SigningStash{
		stash: make(map[string]NotationStashDatum),
	}
}

func (sc *SigningStash) Add(key string, datum NotationStashDatum) {
	sc.stash[key] = datum
}

func (sc *SigningStash) Get(key string) (NotationStashDatum, bool) {
	datum, ok := sc.stash[key]
	return datum, ok
}

func (sc *SigningStash) Len() int {
	return len(sc.stash)
}
