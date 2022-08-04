package main

func peak(arr []int) int {
	for i := 1; i < len(arr)-1; {
		if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
			// that's peak
			return i
		} else {
			i++
		}
	}
	return 1
}

func main() {
	nums := []int{1, 2, 3, 1, 7, 3}
	peak(nums)
}
