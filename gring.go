package gring

import (
	"fmt"
)

type node struct {
	next  int
	prev  int
	value interface{}
}

func (n node) String() string {
	return fmt.Sprintf("[%d %d %v]", n.prev, n.next, n.value)
}

// Ring is a circular doubly linked list using array as its underlying storage.
type Ring struct {
	nodes  []*node
	length int
	head   int
}

// Creates a new empty ring.
func New() *Ring {
	r := &Ring{
		nodes:  make([]*node, 0, 8),
		length: 0,
		head:   0,
	}
	return r
}

// Creates a new ring from an existing tour.
// A tour is an array of integers that specifies the order of the nodes, e.g. [2, 1, 0, 4, 3].
// It is assumed that tour contains integers from [0,n) where n is the length of the tour.
func NewFromTour(tour []int) *Ring {
	r := New()
	if len(tour) > 0 {
		for range tour {
			n := &node{}
			r.nodes = append(r.nodes, n)
		}

		for i, v := range tour {
			next := i + 1
			if next >= len(tour) {
				next = 0
			}

			prev := i - 1
			if prev < 0 {
				prev = len(tour) - 1
			}

			r.nodes[v].next = tour[next]
			r.nodes[v].prev = tour[prev]
		}

		r.length = len(tour)
	}
	return r
}

// Adds a new node to the "end" of the ring.
// "End" of the ring is whichever node that comes before the 0th, or the head node.
// If ring is empty, adds the first node.
func (r *Ring) Add() {
	r.AddWithValue(nil)
}

// Adds a new node to the "end" of the ring with the specified value.
// "End" of the ring is whichever node that comes before the 0th, or the head node.
// If ring is empty, adds the first node.
func (r *Ring) AddWithValue(v interface{}) {
	n := &node{}
	n.value = v
	if len(r.nodes) == 0 {
		n.next = -1
		n.prev = -1
	} else if len(r.nodes) == 1 {
		n.next = 0
		n.prev = 0
		r.nodes[0].prev = 1
		r.nodes[0].next = 1
	} else {
		index := len(r.nodes)
		zprev := r.nodes[r.head].prev
		n.next = r.head
		n.prev = zprev
		r.nodes[r.head].prev = index
		r.nodes[zprev].next = index
	}

	r.nodes = append(r.nodes, n)
	r.length++
}

// Sets the value of a particular node.
func (r *Ring) SetValue(n int, v interface{}) {
	r.nodes[n].value = v
}

// Detaches node n, and inserts it after the node p, such that the end result becomes p -> n
// Returns error if ring is empty or p is detached.
func (r *Ring) InsertAfter(n, p int) error {
	if len(r.nodes) == 0 {
		return ErrEmptyRing
	}

	pnext := r.nodes[p].next
	if pnext == n {
		pnext = r.nodes[pnext].next
	}

	if r.anyIsNil(n, p, pnext) {
		return ErrInvalidOperationOnDetachedNode
	}

	r.Detach(n)

	r.nodes[n].prev = p
	r.nodes[p].next = n
	r.nodes[n].next = pnext
	r.nodes[pnext].prev = n

	r.length++

	return nil
}

// Detaches node n, and inserts it before the node p, such that the end result becomes n -> p
// Returns error if ring is empty or p is detached.
func (r *Ring) InsertBefore(n, p int) error {
	if len(r.nodes) == 0 {
		return ErrEmptyRing
	}

	pprev := r.nodes[p].prev
	if pprev == n {
		pprev = r.nodes[pprev].prev
	}

	if r.anyIsNil(n, p, pprev) {
		return ErrInvalidOperationOnDetachedNode
	}

	r.Detach(n)

	r.nodes[n].prev = pprev
	r.nodes[pprev].next = n
	r.nodes[n].next = p
	r.nodes[p].prev = n

	r.length++

	return nil
}

// Detaches a particular node from the ring, connecting its prev and next nodes together.
// Since Ring is using arrays as its underlying storage, references to detached nodes are still kept, and can be reinserted later.
func (r *Ring) Detach(n int) {
	prev := r.nodes[n].prev
	next := r.nodes[n].next
	if r.anyIsNil(prev, next) {
		// already detached
		return
	}

	r.nodes[prev].next = next
	r.nodes[next].prev = prev
	r.nodes[n].next = -1
	r.nodes[n].next = -1

	r.length--

	if n == r.head {
		r.head = next
	}
}

// Swaps two nodes in the ring.
func (r *Ring) Swap(a, b int) error {
	aprev := r.nodes[a].prev
	bprev := r.nodes[b].prev

	if r.anyIsNil(a, b, aprev, bprev) {
		return ErrInvalidOperationOnDetachedNode
	}

	r.Detach(a)
	r.Detach(b)

	if aprev == b {
		r.InsertAfter(a, bprev)
		r.InsertAfter(b, a)
	} else if bprev == a {
		r.InsertAfter(b, aprev)
		r.InsertAfter(a, b)
	} else {
		r.InsertAfter(a, bprev)
		r.InsertAfter(b, aprev)
	}

	return nil
}

// Reverses the ring direction.
func (r *Ring) Reverse() {
}

// Tours the ring and returns the node order as an array starting from the "head" node.
func (r *Ring) Tour() []int {
	tour := make([]int, 0, len(r.nodes))

	n := r.head
	tour = append(tour, n)
	n = r.nodes[n].next
	for n != r.head && n != -1 {
		tour = append(tour, n)
		n = r.nodes[n].next
	}

	return tour
}

// Returns an iterator for this ring, starting from the "head" node.
func (r *Ring) Iterator() *Iterator {
	return r.IteratorFrom(r.head)
}

// Returns an iterator for this ring, starting from the node n.
func (r *Ring) IteratorFrom(n int) *Iterator {
	return nil
}

// Duplicates the ring.
func (r *Ring) Clone() *Ring {
	return nil
}

// Get the size of the ring
func (r *Ring) Len() int {
	return r.length
}

// Check if any values is -1
func (r *Ring) anyIsNil(values ...int) bool {
	for _, v := range values {
		if v == -1 {
			return true
		}
	}
	return false
}
