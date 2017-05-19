package hashqueue

import (
	"fmt"
	"github.com/google/uuid"
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

	for i := 0; i < 100; i++ {
		h.PushBack(Key(uuid.New().String()), uuid.New().String())
	}

	h.Sort(func(l, r *Element) bool {
		return l.Value.(string) < r.Value.(string)
	})

	for e := h.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v => %v\n", e.Key, e.Value)
	}
}
