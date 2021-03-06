package hashqueue

import (
	"fmt"
	"github.com/google/uuid"
	"strings"
	"testing"
)

func TestHashQueue(t *testing.T) {

	h := New()

	/*
		h.PushBack("foo", "foo")
		h.PushFront("bar", "bar")
		h.InsertAfter("qux", "qux", "bar")
		h.InsertBefore("baz", "baz", "foo")

		for e := h.Front(); e != nil; e = e.Next() {
			fmt.Printf("%v => %v\n", e.Key, e.Value)
		}

		fmt.Printf("--\n")
		h.MoveToFront("qux")
		h.MoveToBack("bar")
		h.MoveBefore("baz", "bar")
		h.MoveAfter("foo", "qux")

		for e := h.Front(); e != nil; e = e.Next() {
			fmt.Printf("%v => %v\n", e.Key, e.Value)
		}

		fmt.Printf("--\n")

		h.Sort(func(l, r *Element) bool {
			return l.Key < r.Key
		})
		for e := h.Front(); e != nil; e = e.Next() {
			fmt.Printf("%v => %v\n", e.Key, e.Value)
		}
	*/

	/*
		for i := 0; i < 100; i++ {
			h.PushBack(uuid.New().String(), uuid.New().String())
		}

		h.Sort(func(l, r *Element) bool {
			return l.Value.(string) < r.Value.(string)
		})
	*/

	for i := 0; i < 10000; i++ {

		key, val := uuid.New().String(), uuid.New().String()

		h.Put(key, val, func(k string, v Value) bool {
			return strings.Compare(val, v.(string)) < 0
		})
	}

	var lastKey string
	var failed bool

	h.Range(func(key string, val Value) bool {

		fmt.Printf("%v => %v\n", key, val)

		// if key > lastKey {
		if val.(string) > lastKey {

			lastKey = val.(string)
			return true
		}

		failed = true
		return false
	})

	if failed {
		fmt.Printf("FAILED\n")
	} else {
		fmt.Printf("PASS\n")
	}

}
