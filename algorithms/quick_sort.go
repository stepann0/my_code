package main

import "fmt"

// types that supports >, >=, <, <= and == operators
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

func QuickSort[T Ordered](arr []T, order, start, end int) {
	if start >= end {
		return
	}
	index := partition(arr, order, start, end)
	QuickSort(arr, order, start, index-1)
	QuickSort(arr, order, index+1, end)
}

func partition[T Ordered](arr []T, order, start, end int) int {
	pivot := arr[end]
	i := start - 1
	if order >= 0 { // ascending order
		for j := start; j <= end-1; j++ {
			if arr[j] <= pivot {
				i++
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	} else { // descending order
		for j := start; j <= end-1; j++ {
			if arr[j] >= pivot {
				i++
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	arr[i+1], arr[end] = arr[end], arr[i+1]
	return i + 1
}

// Sorts arr using Quick Sort algorithm.
// Set order to 1 for ascending order and -1 for descending.
func Sort[T Ordered](arr []T, order int) {
	QuickSort(arr, order, 0, len(arr)-1)
}

func main() {
	arr := []int{-576, -7, 82, 4, -8, 9, 1, 103, 67, -985, 0, 4}
	// arr := []string{"a", "q", "y", "A", "O", "t", "e", "g", "N", "m", "K", "p"}
	Sort(arr, -1)
	fmt.Println(arr)
}
