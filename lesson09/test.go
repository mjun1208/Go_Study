// 2026-11-05
// Slice (copy)

package main

import "fmt"

func YO() {
	a := []int{10, 20, 30}

	b := make([]int, len(a))
	copy(b, a)

	b[0] = 999

	fmt.Println(a) // [10 20 30]
	fmt.Println(b) // [999 20 30]
}
