package main

import (
	"fmt"
	"sort"
)

func TwoSumWithoutDict(nums []int, target int) []int {
	for i, n := range nums {
		m := target - n
		for j := 0; j < i; j++ {
			if nums[j] == m {
				fmt.Printf("%d + %d = %d\n", nums[i], nums[j], target)
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

func TwoSumII(nums []int, target int) []int {
	sort.Ints(nums)
	for l, r := 0, len(nums)-1; l < r; {
		sum := nums[l] + nums[r]
		if sum > target {
			r--
		} else if sum < target {
			l++
		} else {
			fmt.Printf("%d + %d = %d\n", nums[l], nums[r], target)
			return []int{nums[l], nums[r]}
		}
	}
	return []int{}
}

func main() {
	nums, target := []int{5, 6, 1, 3, 7}, 13
	fmt.Printf("%v\n\n", TwoSumII(nums, target))
	fmt.Printf("%v\n\n", TwoSum(nums, target))
	fmt.Printf("%v\n", TwoSumWithoutDict(nums, target))
}
