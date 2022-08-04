package main

import "fmt"

func quickSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	index := partition(arr, start, end)
	quickSort(arr, start, index)
	quickSort(arr, index+1, end)
}

func partition(arr []int, start, end int) int {
	pivot := arr[(start+end)/2]
	i, j := start, end

	for {
		for arr[i] < pivot {
			i++
		}
		for arr[j] > pivot {
			j--
		}
		if i >= j {
			return j
		}
		i++
		j--
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func main() {
	arr := []int{4, 7, 2, 8, -3, -2, 4, -1, 0, 8}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
