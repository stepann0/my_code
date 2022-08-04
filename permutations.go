package main

import "fmt"

// Initially k = length(arr)
// Heap's algorithm
func generate[T any](arr []T, k int, store *[][]T) {
	if k == 0 {
		// length(arr) == 0
		fmt.Println("Source array can't be empty.")
		return
	}

	if k == 1 {
		fmt.Println(arr)
		*store = append(*store, arr)
		return
	}

	generate(arr, k-1, store)
	for i := 0; i < k-1; i++ {
		if k%2 == 0 {
			arr[i], arr[k-1] = arr[k-1], arr[i]
		} else {
			arr[0], arr[k-1] = arr[k-1], arr[0]
		}
		generate(arr, k-1, store)
	}
}

func permutations[T any](arr []T) [][]T {
	store := [][]T{}
	generate(arr, len(arr), &store)
	return store
}

func main() {
	s := []int{3, 4, 5, 6, 7}
	fmt.Println(permutations(s))
}
