package main

import (
	"fmt"
	"sort"
)

func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < len(nums)-2; i++ {
		// Skip same nums[i] values
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		for l, r := i+1, len(nums)-1; l < r; {
			sum := nums[l] + nums[r]
			if sum > -nums[i] {
				r-- // decrease sum
			} else if sum < -nums[i] {
				l++ // increase sum
			} else {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				// Move left pointer and skip same values
				for l = l + 1; nums[l-1] == nums[l] && l < r; {
					l++
				}
			}
		}
	}
	return res
}

func main() {
	fmt.Println(ThreeSum([]int{-1, 0, 1, 2, -1, -4}))
}
