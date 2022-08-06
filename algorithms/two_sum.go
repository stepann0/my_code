package main

import "fmt"

func TwoSumWithoutDict(nums []int, target int) []int {
	for i, n := range nums {
		m := target - n
		for j := 0; j < i; j++ {
			if nums[j] == m {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func TwoSum(nums []int, target int) []int {
	dict := map[int]int{}
	for i1, n := range nums {
		m := target - n
		if i2, has := dict[m]; has {
			fmt.Printf("%d + %d = %d\n", nums[i2], nums[i1], target)
			return []int{i2, i1}
		} else {
			dict[n] = i1
		}
	}
	return []int{}
}

func main() {
	fmt.Println(TwoSum([]int{3, 5, 1, 0, 7}, 100))
}
