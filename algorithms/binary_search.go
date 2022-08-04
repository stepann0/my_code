package main

import (
	"fmt"
)

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Integer interface {
	Signed | Unsigned
}

type Float interface {
	~float32 | ~float64
}

type Ordered interface {
	Integer | Float | ~string
}

// Function returns index of searched element or -1
func BinarySearch[T Ordered](arr []T, search T) int {
	var left int = 0
	var right int = len(arr) - 1

	for left <= right {
		mid := (left + right) / 2
		mid_val := arr[mid]
		if mid_val > search {
			right = mid - 1
		} else if mid_val < search {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func BinarySearchRecursive[T Ordered](arr []T, search T) bool {
	mid := (len(arr) - 1) / 2
	mid_val := arr[mid]
	if len(arr) < 1 {
		return false
	}

	if mid_val > search {
		return BinarySearchRecursive(arr[:mid], search)
	} else if mid_val < search {
		return BinarySearchRecursive(arr[mid+1:], search)
	} else {
		return true
	}
}

func main() {
	arr := []string{"asdf", "bsdf", "csdf", "dsdf", "esdf", "fsdf", "gsdf", "hsdf"}
	res := BinarySearch(arr, "csdf")
	fmt.Println(res)
}
