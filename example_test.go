package hashqueue

type Cache interface {
	Put(key string, val interface{})
	Get(key string) (val interface{}, exists bool)
	Keys() map[string]struct{}
}

type lru struct {
	hq             *HashQueue
	maxCap, extCap int
}

func NewLRU(maxCap, extCap int) Cache {

	return &lru{
		hq:     New(),
		maxCap: maxCap,
		extCap: extCap,
	}
}

func (t *lru) gc() {

	if t.hq.Len() > t.extCap {

		for t.hq.Len() > t.maxCap { // while over capacity
			t.hq.PopBack()
		}
	}
}

func (t *lru) Put(key string, val interface{}) {

	t.hq.PushFront(key, val)
	t.gc() // try gc

	return
}

func (t *lru) Get(key string) (val interface{}, exists bool) {

	if val, exists = t.hq.Get(key); exists {
		t.hq.MoveToFront(key) // move to front of lru
	}

	return
}

func (t *lru) Keys() map[string]struct{} {

	keys := make(map[string]struct{})

	for _, k := range t.hq.Keys() {
		keys[string(k)] = struct{}{}
	}

	return keys
}
