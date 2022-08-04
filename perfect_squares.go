package main

import (
	"fmt"
)

// Same as coin_change.go exept GetPerfectSet() function
const Inf = 9223372036854775806 // math.MaxInt64 - 1

func GetPerfectSet(n int) []int {
	p_set := []int{}
	root := 1
	p_num := root * root
	for n >= p_num {
		n -= p_num
		p_set = append(p_set, p_num)
		root++
		p_num = root * root
	}
	return p_set
}

func NewRow(length int) []int {
	row := make([]int, length+1)
	if length >= 1 {
		for i := 1; i < len(row); i++ {
			row[i] = Inf
		}
	}
	return row
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func numSquares(n int) int {
	set := GetPerfectSet(n * n)
	row := NewRow(n)
	for _, coin := range set {
		// Find minimum num of set
		// to make change using $coin
		// for each n from 1 to n+1
		for i := 1; i < len(row); i++ {
			if coin == i {
				row[i] = 1
			} else if coin < i {
				row[i] = min(row[i-coin]+1, row[i])
			}
		}
	}

	res := row[len(row)-1]
	if res == Inf {
		return -1
	}
	return res
}

func main() {
	fmt.Println(GetPerfectSet(13))
}
