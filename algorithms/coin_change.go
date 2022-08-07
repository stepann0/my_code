package main

import (
	"fmt"
)

const Inf = 9223372036854775806 // math.MaxInt64 - 1

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
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

func NewTable(set []int, sum int) [][]int {
	m := make([][]int, len(set))
	for i := 0; i < len(m); i++ {
		m[i] = NewRow(sum)
	}
	return m
}

func Solve(coins []int, amount int) int {
	row := NewRow(amount)
	for _, coin := range coins {
		// Find minimum num of coins
		// to make change using coin
		// for each amount from 1 to amount+1
		for i := 1; i < len(row); i++ {
			if coin == i {
				row[i] = 1
			} else if coin < i {
				row[i] = min(row[i-coin]+1, row[i])
			}
		}
		fmt.Println(row)
	}

	res := row[len(row)-1]
	if res == Inf {
		return -1
	}
	return res
}

func main() {
	set := []int{3, 4, 5}
	sum := 17
	fmt.Println(Solve(set, sum))
}
