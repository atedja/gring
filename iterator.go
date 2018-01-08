package gring

type Iterator struct {
	start   int
	current int
	r       *Ring
}

// Advances the iterator.
// Returns true if it advances, false if it has looped back to the beginning.
func (i *Iterator) Next() bool {
	if i.current == -1 {
		i.current = i.start
		return true
	}

	i.current = i.r.nodes[i.current].next
	return i.current != -1 && i.current != i.start
}

// Returns the current node index
func (i *Iterator) Index() int {
	return i.current
}

// Returns the value of the current node
func (i *Iterator) Value() interface{} {
	return i.r.nodes[i.current].value
}
