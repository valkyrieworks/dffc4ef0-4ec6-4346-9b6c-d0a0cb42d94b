package componentindex

import (
	commitchronize "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
)

//
type CNIndex struct {
	m map[string]any
	l commitchronize.Exclusion
}

func FreshCNIndex() *CNIndex {
	return &CNIndex{
		m: make(map[string]any),
	}
}

func (cm *CNIndex) Set(key string, datum any) {
	cm.l.Lock()
	cm.m[key] = datum
	cm.l.Unlock()
}

func (cm *CNIndex) Get(key string) any {
	cm.l.Lock()
	val := cm.m[key]
	cm.l.Unlock()
	return val
}

func (cm *CNIndex) Has(key string) bool {
	cm.l.Lock()
	_, ok := cm.m[key]
	cm.l.Unlock()
	return ok
}

func (cm *CNIndex) Erase(key string) {
	cm.l.Lock()
	delete(cm.m, key)
	cm.l.Unlock()
}

func (cm *CNIndex) Extent() int {
	cm.l.Lock()
	extent := len(cm.m)
	cm.l.Unlock()
	return extent
}

func (cm *CNIndex) Flush() {
	cm.l.Lock()
	cm.m = make(map[string]any)
	cm.l.Unlock()
}

func (cm *CNIndex) Tokens() []string {
	cm.l.Lock()

	tokens := make([]string, 0, len(cm.m))
	for k := range cm.m {
		tokens = append(tokens, k)
	}
	cm.l.Unlock()
	return tokens
}

func (cm *CNIndex) Items() []any {
	cm.l.Lock()
	elements := make([]any, 0, len(cm.m))
	for _, v := range cm.m {
		elements = append(elements, v)
	}
	cm.l.Unlock()
	return elements
}
