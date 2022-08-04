package main

import (
	"fmt"
	"sort"
)

func main() {
	B := 100
	A := []int{20, 20, 20, 20, 20, 30}
	sort.Ints(A)

	n := 0
	for _, a := range A {
		if B >= a {
			B -= a
			n++
		}
	}
	fmt.Println(n)
}
