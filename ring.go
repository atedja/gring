package gring

import ()

type node struct {
	next  int
	prev  int
	value interface{}
}

// Ring is a circular doubly linked list using array as its underlying storage.
// Nodes in the ring can be detached, reinserted, or swapped.
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

// Creates a new ring from a tour.
// A tour is an array of integers that specifies the order of the nodes, e.g. [2, 1, 0, 4, 3].
// It is assumed that the array contains integers from [0,n) where n is the length of the array,
// and each integer occurs only once.
func NewFromArray(tour []int) *Ring {
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
// "End" of the ring is whichever node that comes before the 0th or the head node.
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

// Returns the value of a particular node
func (r *Ring) Value(n int) interface{} {
	return r.nodes[n].value
}

// Detaches node n, and inserts it after the node p, such that the end result becomes p -> n.
// Returns error if ring is empty or p is a detached node.
func (r *Ring) InsertAfter(n, p int) error {
	if len(r.nodes) == 0 {
		return ErrEmptyRing
	}

	pnext := r.nodes[p].next
	if pnext == n {
		pnext = r.nodes[pnext].next
	}

	if anyIsInvalid(n, p, pnext) {
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

// Detaches node n, and inserts it before the node p, such that the end result becomes n -> p.
// Returns error if ring is empty or p is a detached node.
func (r *Ring) InsertBefore(n, p int) error {
	if len(r.nodes) == 0 {
		return ErrEmptyRing
	}

	pprev := r.nodes[p].prev
	if pprev == n {
		pprev = r.nodes[pprev].prev
	}

	if anyIsInvalid(n, p, pprev) {
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
	if anyIsInvalid(prev, next) {
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

	if anyIsInvalid(a, b, aprev, bprev) {
		return ErrInvalidOperationOnDetachedNode
	}

	r.Detach(a)
	r.Detach(b)

	var err error
	if aprev == b {
		err = r.InsertAfter(a, bprev)
		err = r.InsertAfter(b, a)
	} else if bprev == a {
		err = r.InsertAfter(b, aprev)
		err = r.InsertAfter(a, b)
	} else {
		err = r.InsertAfter(a, bprev)
		err = r.InsertAfter(b, aprev)
	}

	return err
}

// Given a node and a target node, sets node's next to target, while still maintaining the loop.
// This will reverse part of the tour.
func (r *Ring) TwoOptSwap(n, target int) error {
	oldNext := r.nodes[n].next
	oldTargetNext := r.nodes[target].next

	if anyIsInvalid(n, target, oldNext, oldTargetNext) {
		return ErrInvalidOperationOnDetachedNode
	}

	// disconnect
	r.nodes[oldNext].prev = -1
	r.nodes[oldTargetNext].prev = -1

	// connect with new one
	r.nodes[n].next = target

	// reverse the direction. this will loop until we hit oldNext
	var old = n
	var current = target
	for current != -1 {
		var oldPrev = r.nodes[current].prev
		r.nodes[current].prev = old
		r.nodes[current].next = oldPrev

		old = current
		current = oldPrev
	}

	r.nodes[oldNext].next = oldTargetNext
	r.nodes[oldTargetNext].prev = oldNext

	return nil
}

// Reverses the ring direction.
func (r *Ring) Reverse() error {
	if len(r.nodes) == 0 {
		return ErrEmptyRing
	}

	current := r.head
	prev := r.nodes[current].prev

	// stupid golang has no do-while. have to hack it in with the first bool
	first := true
	for first || (current != r.head && current != -1) {
		first = false
		r.nodes[current].prev = r.nodes[current].next
		r.nodes[current].next = prev

		current = prev
		prev = r.nodes[current].prev
	}

	return nil
}

// Tours the ring and returns the node order as an array starting from the "head" node.
func (r *Ring) tour() []int {
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

// Returns an iterator for this ring.
// Iterator usually starts from the 0th node, although not guaranteed.
// Returns an error if Ring is empty.
func (r *Ring) Iterator() (*Iterator, error) {
	return r.iteratorFrom(r.head)
}

// Returns an iterator for this ring, starting from the node n.
func (r *Ring) iteratorFrom(n int) (*Iterator, error) {
	if len(r.nodes) == 0 {
		return nil, ErrEmptyRing
	}

	return &Iterator{start: n, current: -1, r: r}, nil
}

// Duplicates the ring.
func (r *Ring) Clone() *Ring {
	clone := &Ring{
		nodes:  make([]*node, len(r.nodes)),
		length: r.length,
		head:   r.head,
	}

	for i, n := range r.nodes {
		clone.nodes[i] = &node{n.next, n.prev, n.value}
	}

	return clone
}

// Gets the size of the ring
func (r *Ring) Len() int {
	return r.length
}

// Checks if any values is -1
func anyIsInvalid(values ...int) bool {
	for _, v := range values {
		if v == -1 {
			return true
		}
	}
	return false
}
