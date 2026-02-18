package kinds

//
type AutographRepositoryItem struct {
	RatifierLocation []byte
	BallotAttestOctets    []byte
}

type AutographRepository struct {
	repository map[string]AutographRepositoryItem
}

func NewAutographRepository() *AutographRepository {
	return &AutographRepository{
		repository: make(map[string]AutographRepositoryItem),
	}
}

func (sc *AutographRepository) Add(key string, item AutographRepositoryItem) {
	sc.repository[key] = item
}

func (sc *AutographRepository) Get(key string) (AutographRepositoryItem, bool) {
	item, ok := sc.repository[key]
	return item, ok
}

func (sc *AutographRepository) Len() int {
	return len(sc.repository)
}
