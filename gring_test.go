package gring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	r := New()
	assert.NotNil(t, r)
	assert.NotNil(t, r.nodes)
}

func TestNewFromTour(t *testing.T) {
	r := NewFromArray([]int{2, 1, 0, 4, 3})
	assert.NotNil(t, r)
	assert.NotNil(t, r.nodes)
	assert.Equal(t, []int{0, 4, 3, 2, 1}, r.tour())
}

func TestAdd(t *testing.T) {
	r := New()
	r.Add()
	assert.Equal(t, 1, len(r.nodes))

	r.Add()
	assert.Equal(t, 2, len(r.nodes))
	assert.Equal(t, []int{0, 1}, r.tour())

	r.Add()
	r.Add()
	r.Add()
	r.Add()
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5}, r.tour())
}

func TestAddWithValue(t *testing.T) {
	r := New()
	r.AddWithValue("mdoe")
	r.AddWithValue("abc")
	r.AddWithValue("sawyer")
	assert.Equal(t, []int{0, 1, 2}, r.tour())
}

func TestSetValue(t *testing.T) {
	r := NewFromArray([]int{2, 1, 0, 4, 3})
	r.SetValue(1, "hello")
}

func TestInsertAfter(t *testing.T) {
	r := NewFromArray([]int{2, 1, 0, 4, 3})
	assert.Equal(t, []int{0, 4, 3, 2, 1}, r.tour())

	r.InsertAfter(0, 3)
	assert.Equal(t, []int{4, 3, 0, 2, 1}, r.tour())

	r.InsertAfter(4, 2)
	assert.Equal(t, []int{3, 0, 2, 4, 1}, r.tour())

	r.InsertAfter(4, 2)
	assert.Equal(t, []int{3, 0, 2, 4, 1}, r.tour())

	r.InsertAfter(2, 4)
	assert.Equal(t, []int{3, 0, 4, 2, 1}, r.tour())

	r.InsertAfter(0, 2)
	assert.Equal(t, []int{3, 4, 2, 0, 1}, r.tour())

	assert.Equal(t, 5, r.Len())
}

func TestInsertBefore(t *testing.T) {
	r := NewFromArray([]int{2, 1, 0, 4, 3})
	assert.Equal(t, []int{0, 4, 3, 2, 1}, r.tour())

	r.InsertBefore(0, 3)
	assert.Equal(t, []int{4, 0, 3, 2, 1}, r.tour())

	r.InsertBefore(4, 2)
	assert.Equal(t, []int{0, 3, 4, 2, 1}, r.tour())

	r.InsertBefore(4, 2)
	assert.Equal(t, []int{0, 3, 4, 2, 1}, r.tour())

	r.InsertBefore(2, 4)
	assert.Equal(t, []int{0, 3, 2, 4, 1}, r.tour())

	r.InsertBefore(0, 2)
	assert.Equal(t, []int{3, 0, 2, 4, 1}, r.tour())

	assert.Equal(t, 5, r.Len())
}

func TestDetach(t *testing.T) {
	r := NewFromArray([]int{2, 1, 0, 4, 3})
	r.Detach(3)
	assert.Equal(t, []int{0, 4, 2, 1}, r.tour())
	r.Detach(2)
	assert.Equal(t, []int{0, 4, 1}, r.tour())
	assert.Equal(t, 3, r.Len())

	r.Detach(0)
	assert.Equal(t, []int{4, 1}, r.tour())
	assert.Equal(t, 2, r.Len())
}

func TestSwap(t *testing.T) {
	r := NewFromArray([]int{2, 1, 0, 4, 3})
	assert.Equal(t, []int{0, 4, 3, 2, 1}, r.tour())

	r.Swap(1, 3)
	assert.Equal(t, []int{0, 4, 1, 2, 3}, r.tour())

	r.Swap(0, 4)
	assert.Equal(t, []int{1, 2, 3, 4, 0}, r.tour())

	r.Swap(0, 4)
	assert.Equal(t, []int{1, 2, 3, 0, 4}, r.tour())

	assert.Equal(t, 5, r.Len())
}

func TestReverse(t *testing.T) {
	var err error
	r := NewFromArray([]int{2, 1, 0, 4, 3})
	assert.Equal(t, []int{0, 4, 3, 2, 1}, r.tour())

	err = r.Reverse()
	assert.Nil(t, err)
	assert.Equal(t, []int{0, 1, 2, 3, 4}, r.tour())

	err = r.Reverse()
	assert.Nil(t, err)
	assert.Equal(t, []int{0, 4, 3, 2, 1}, r.tour())

	assert.Equal(t, 5, r.Len())
}

func TestClone(t *testing.T) {
	r := NewFromArray([]int{2, 1, 0, 4, 3})
	assert.Equal(t, []int{0, 4, 3, 2, 1}, r.tour())

	clone := r.Clone()
	assert.Equal(t, []int{0, 4, 3, 2, 1}, clone.tour())

	r.Reverse()
	assert.Equal(t, []int{0, 1, 2, 3, 4}, r.tour())
	assert.Equal(t, []int{0, 4, 3, 2, 1}, clone.tour())

	clone.Swap(4, 2)
	assert.Equal(t, []int{0, 1, 2, 3, 4}, r.tour())
	assert.Equal(t, []int{0, 2, 3, 4, 1}, clone.tour())

	assert.Equal(t, 5, r.Len())
	assert.Equal(t, 5, clone.Len())
}

func TestIterator(t *testing.T) {
	r := NewFromArray([]int{2, 1, 0, 4, 3})
	assert.Equal(t, []int{0, 4, 3, 2, 1}, r.tour())

	r.SetValue(0, "Hi")
	r.SetValue(4, "This")
	r.SetValue(3, "is")
	r.SetValue(2, "Circular")
	r.SetValue(1, "List")

	order := make([]int, 0)
	values := make([]string, 0)
	iter, err := r.Iterator()
	assert.Nil(t, err)
	for iter.Next() {
		order = append(order, iter.Index())
		values = append(values, iter.Value().(string))
	}

	assert.Equal(t, []int{0, 4, 3, 2, 1}, order)
	assert.Equal(t, []string{"Hi", "This", "is", "Circular", "List"}, values)

	r.Reverse()
	order = make([]int, 0)
	values = make([]string, 0)
	iter, err = r.Iterator()
	for iter.Next() {
		order = append(order, iter.Index())
		values = append(values, iter.Value().(string))
	}

	assert.Equal(t, []int{0, 1, 2, 3, 4}, order)
	assert.Equal(t, []string{"Hi", "List", "Circular", "is", "This"}, values)
}

func TestIteratorSingleElement(t *testing.T) {
	r := NewFromArray([]int{0})
	assert.Equal(t, []int{0}, r.tour())

	r.SetValue(0, "Hi")

	order := make([]int, 0)
	values := make([]string, 0)
	iter, err := r.Iterator()
	assert.Nil(t, err)
	for iter.Next() {
		order = append(order, iter.Index())
		values = append(values, iter.Value().(string))
	}

	assert.Equal(t, []int{0}, order)
	assert.Equal(t, []string{"Hi"}, values)
}

func TestIteratorEmptyRing(t *testing.T) {
	r := New()
	iter, err := r.Iterator()
	assert.Nil(t, iter)
	assert.Equal(t, err, ErrEmptyRing)
}
