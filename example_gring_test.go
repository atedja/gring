package gring_test

import (
	"fmt"
	"github.com/atedja/gring"
)

func ExampleRing() {
	ring := gring.New()
	ring.AddWithValue(1)
	ring.AddWithValue(10)
	ring.AddWithValue(100)

	values := make([]int, 0)
	iter, _ := ring.Iterator()
	for iter.Next() {
		values = append(values, iter.Value().(int))
	}

	fmt.Println(values)
	// Output:
	// [1 10 100]
}
