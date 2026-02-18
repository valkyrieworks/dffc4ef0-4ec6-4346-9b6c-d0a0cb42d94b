package cmap

import (
	engineconnect "github.com/valkyrieworks/utils/align"
)

//
type CIndex struct {
	m map[string]any
	l engineconnect.Lock
}

func NewCIndex() *CIndex {
	return &CIndex{
		m: make(map[string]any),
	}
}

func (cm *CIndex) Set(key string, item any) {
	cm.l.Lock()
	cm.m[key] = item
	cm.l.Unlock()
}

func (cm *CIndex) Get(key string) any {
	cm.l.Lock()
	val := cm.m[key]
	cm.l.Unlock()
	return val
}

func (cm *CIndex) Has(key string) bool {
	cm.l.Lock()
	_, ok := cm.m[key]
	cm.l.Unlock()
	return ok
}

func (cm *CIndex) Erase(key string) {
	cm.l.Lock()
	delete(cm.m, key)
	cm.l.Unlock()
}

func (cm *CIndex) Volume() int {
	cm.l.Lock()
	volume := len(cm.m)
	cm.l.Unlock()
	return volume
}

func (cm *CIndex) Flush() {
	cm.l.Lock()
	cm.m = make(map[string]any)
	cm.l.Unlock()
}

func (cm *CIndex) Keys() []string {
	cm.l.Lock()

	keys := make([]string, 0, len(cm.m))
	for k := range cm.m {
		keys = append(keys, k)
	}
	cm.l.Unlock()
	return keys
}

func (cm *CIndex) Items() []any {
	cm.l.Lock()
	items := make([]any, 0, len(cm.m))
	for _, v := range cm.m {
		items = append(items, v)
	}
	cm.l.Unlock()
	return items
}
