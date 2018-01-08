package gring_test

import (
	"fmt"
	"github.com/atedja/gring"
)

func ExampleRing_Swap() {
	ring := gring.NewFromArray([]int{2, 0, 1, 3})
	ring.Swap(0, 3)

	values := make([]int, 0)
	iter, _ := ring.Iterator()
	for iter.Next() {
		values = append(values, iter.Index())
	}

	fmt.Println(values)
	// Output:
	// [1 0 2 3]
}
