package gring

type Iterator struct {
	start   int
	current int
	r       *Ring
}

func (i *Iterator) Next() bool {
	if i.current == -1 {
		i.current = i.start
		return true
	}

	i.current = i.r.nodes[i.current].next
	return i.current != -1 && i.current != i.start
}

func (i *Iterator) Index() int {
	return i.current
}

func (i *Iterator) Value() interface{} {
	return i.r.nodes[i.current].value
}
