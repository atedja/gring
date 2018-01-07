package gring

type Iterator struct {
	current int
	r       *Ring
}

func (i *Iterator) Next() bool {
	return true
}
