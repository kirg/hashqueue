package hashqueue

import (
	"container/list"
	"sort"
)

type (
	// Key for the hash-queue
	Key string

	// Value for the key
	Value interface{}

	// Element in the hash-queue
	Element struct {
		Key      Key
		Value    Value
		listElem *list.Element
	}

	// HashQueue object
	HashQueue struct {
		l *list.List
		m map[Key]*Element
	}
)

// New returns an empty, initialized hash-queue.
func New() *HashQueue {
	return new(HashQueue).Init()
}

// Init initializes/clears a hash-queue.
func (h *HashQueue) Init() *HashQueue {

	if h.l == nil {
		h.l = list.New()
	} else {
		h.l.Init()
	}

	h.m = make(map[Key]*Element)

	return h
}

// Len returns the number of elements in the queue.
func (h *HashQueue) Len() int {
	return len(h.m)
}

// Front returns the first element of queue, or nil.
func (h *HashQueue) Front() *Element {

	if listElem := h.l.Front(); listElem != nil {
		return listElem.Value.(*Element)
	}

	return nil
}

// Back returns the last element of queue, or nil.
func (h *HashQueue) Back() *Element {

	if listElem := h.l.Back(); listElem != nil {
		return listElem.Value.(*Element)
	}

	return nil
}

// Remove removes the element corresponding to the key; it returns the removed
// element.
func (h *HashQueue) Remove(key Key) *Element {

	if e, ok := h.m[key]; ok {

		// remove from list and clear out the 'listElem' from 'e'
		h.l.Remove(e.listElem).(*Element).listElem = nil
		delete(h.m, key) // also remove from map

		return e
	}

	return nil
}

// Swap swaps the elements corresponding to the given keys.
func (h *HashQueue) Swap(l, r Key) {

	// if the keys are the same, no-op
	if l == r {
		return
	}

	// if either of the keys aren't found, ignore
	el, er := h.m[l], h.m[r]
	if el == nil || er == nil {
		return
	}

	// find the elements prev to l and r
	elprev, erprev := el.Prev(), er.Prev()

	switch {
	case erprev == el: // l immediately precedes r
		h.MoveAfter(l, r)

	case elprev == er: // r immediately precedes l
		h.MoveAfter(r, l)

	case elprev == nil: // l is the first element
		h.MoveAfter(l, erprev.Key)
		h.MoveToFront(r)

	case erprev == nil: // r is the first element
		h.MoveToFront(l)
		h.MoveAfter(r, elprev.Key)

	default:
		h.MoveAfter(l, erprev.Key)
		h.MoveAfter(r, elprev.Key)
	}

	return
}

// PushFront inserts a new element with the given key and value, at the front
// of the queue and returns the newly added entry.
func (h *HashQueue) PushFront(key Key, val Value) *Element {

	// if an element already exists in the list against the given key,
	// update it's value and move it to the front
	if e, ok := h.m[key]; ok {

		e.Value = val
		h.l.MoveToFront(e.listElem)

		return e
	}

	// else, push a new element to the list and update map
	e := &Element{Key: key, Value: val}
	e.listElem = h.l.PushFront(e)
	h.m[key] = e

	return e
}

// PopFront pops out the first element from the queue; will panic if called
// on empty list.
func (h *HashQueue) PopFront() *Element {
	return h.Remove(h.Front().Key)
}

// PushBack inserts a new element with the given key and value, at the back
// of the queue and returns the newly added entry.
func (h *HashQueue) PushBack(key Key, val Value) *Element {

	// if an element already exists in the list against the given key,
	// update it's value and move it to the back
	if e, ok := h.m[key]; ok {

		e.Value = val
		h.l.MoveToBack(e.listElem)

		return e
	}

	// else, push a new element to the list and update map
	e := &Element{Key: key, Value: val}
	e.listElem = h.l.PushBack(e)
	h.m[key] = e

	return e
}

// PopBack pops out the last element from the queue
func (h *HashQueue) PopBack() *Element {
	return h.Remove(h.Back().Key)
}

// InsertBefore inserts a new element with the given key and value, before the
// given key in the queue, and returns the newly added entry. If an entry already
// exists correponding to the key, then it updates the value and moves the entry,
// if needed.
func (h *HashQueue) InsertBefore(key Key, val Value, mark Key) *Element {

	// if an element already exists in the list against the given key,
	// update it's value and move it into position
	if e, ok := h.m[key]; ok {

		// replace 'Value' on existing element
		e.Value = val
		h.l.MoveBefore(e.listElem, h.m[mark].listElem)

		return e
	}

	// else, insert a new element to the list and update map
	e := &Element{Key: key, Value: val}
	e.listElem = h.l.InsertBefore(e, h.m[mark].listElem)
	h.m[key] = e

	return e
}

// InsertAfter inserts a new element with the given key and value, after the
// given key in the queue, and returns the newly added entry. If an entry already
// exists corresponding to the key, then it updates the value and moves the entry,
// if needed.
func (h *HashQueue) InsertAfter(key Key, val Value, mark Key) *Element {

	// if an element already exists in the list against the given key,
	// update it's value and move it into position
	if e, ok := h.m[key]; ok {

		// replace 'Value' on existing element
		e.Value = val
		h.l.MoveAfter(e.listElem, h.m[mark].listElem)

		return e
	}

	// else, insert a new element to the list and update map
	e := &Element{Key: key, Value: val}
	e.listElem = h.l.InsertAfter(e, h.m[mark].listElem)
	h.m[key] = e

	return e
}

// MoveToFront moves the element corresponding to the given key to the front
// of the queue.
func (h *HashQueue) MoveToFront(key Key) {
	h.l.MoveToFront(h.m[key].listElem)
}

// MoveToBack moves the element corresponding to the given key to the back
// of the queue.
func (h *HashQueue) MoveToBack(key Key) {
	h.l.MoveToBack(h.m[key].listElem)
}

// MoveBefore moves the element corresponding to the given key to before the
// mark key, in the queue.
func (h *HashQueue) MoveBefore(key, mark Key) {
	h.l.MoveBefore(h.m[key].listElem, h.m[mark].listElem)
}

// MoveAfter moves the element corresponding to the given key to after the
// mark key, in the queue.
func (h *HashQueue) MoveAfter(key, mark Key) {
	h.l.MoveAfter(h.m[key].listElem, h.m[mark].listElem)
}

// PushBackHashQueue copies all the elements from the given queue to the back
// of this queue.
func (h *HashQueue) PushBackHashQueue(other *HashQueue) {

	for e := other.Front(); e != nil; e = e.Next() {
		h.PushBack(e.Key, e.Value)
	}
}

// PushFrontHashQueue copies all the elements from the given queue to the front
// of this queue.
func (h *HashQueue) PushFrontHashQueue(other *HashQueue) {

	for e := other.Back(); e != nil; e = e.Prev() {
		other.PushFront(e.Key, e.Value)
	}
}

// Get retrieves the value corresponding to the key
func (h *HashQueue) Get(key Key) (val Value, ok bool) {

	if e, ok := h.m[key]; ok {
		return e.Value, true
	}

	return nil, false
}

// Seek seeks to the element corresponding to the key
func (h *HashQueue) Seek(key Key) *Element {
	return h.m[key]
}

// Keys returns all the keys in order
func (h *HashQueue) Keys() []Key {

	keys := make([]Key, 0, h.Len())

	for listElem := h.l.Front(); listElem != nil; listElem = listElem.Next() {
		keys = append(keys, listElem.Value.(*Element).Key)
	}

	return keys
}

// Next returns the next element in the queue or nil.
func (e *Element) Next() *Element {

	if n := e.listElem.Next(); n != nil {
		return n.Value.(*Element)
	}

	return nil
}

// Prev returns the previous element in the queue or nil.
func (e *Element) Prev() *Element {

	if p := e.listElem.Prev(); p != nil {
		return p.Value.(*Element)
	}

	return nil
}

// Less is the function that is used to determine the sort order
type Less func(l, r *Element) bool

// keySorter implements sort.Interface
type keySorter struct {
	h    *HashQueue
	keys []Key
	less Less
}

func (h *HashQueue) sorter(less Less) sort.Interface {
	return &keySorter{
		h:    h,
		keys: h.Keys(),
		less: less,
	}
}

// Len implements the sort.Interface::Len method
func (s *keySorter) Len() int {
	return s.h.Len()
}

// Less implements the sort.Interface::Less method
func (s *keySorter) Less(i, j int) bool {
	return s.less(s.h.Seek(s.keys[i]), s.h.Seek(s.keys[j]))
}

// Swap implements the sort.Interface::Swap method
func (s *keySorter) Swap(i, j int) {
	s.h.Swap(s.keys[i], s.keys[j])              // swap elements
	s.keys[i], s.keys[j] = s.keys[j], s.keys[i] // swap keys
}

// Sort sorts all of the elements in the list using the given sorting function
func (h *HashQueue) Sort(less Less) {

	sort.Sort(h.sorter(less))
}
