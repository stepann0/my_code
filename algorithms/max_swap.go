package main

import (
	"fmt"
	"strconv"
)

// Find right-most max digit in s.
func max_digit(s []byte) int {
	if len(s) == 0 {
		panic("Can't find max digit in an empty string.")
	}

	max_i := 0
	for i := range s {
		if s[i] >= s[max_i] && i > max_i {
			max_i = i
		}
	}
	return max_i
}

// Swaps first digit and right-most max digit in s or
// recursively do that for s[1:]
func swap(s []byte) []byte {
	if len(s) == 1 {
		return s
	}

	max_i := max_digit(s)
	if s[0] == s[max_i] {
		return append([]byte{s[0]}, swap(s[1:])...)
	}
	s[0], s[max_i] = s[max_i], s[0]
	return s
}

// Wrap function for swap().
func MaximumSwap(num int) int {
	s := []byte(strconv.Itoa(num))
	res, _ := strconv.Atoi(string(swap(s)))
	return res
}

// Iterative brute force solution for maximum swap problem.
func BruteForce(num int) int {
	if num < 10 {
		return num
	}
	s := []byte(strconv.Itoa(num))
	max := num
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if s[j] <= s[i] {
				continue
			}
			// trying to swap
			s[i], s[j] = s[j], s[i]
			if n, _ := strconv.Atoi(string(s)); n > max {
				max = n
			}
			// back to source slice
			s[j], s[i] = s[i], s[j]
		}
	}
	return max
}

func main() {
	for _, num := range []int{9125, 18, 2736, 987654321, 54312, 19997, 115, 7} {
		fmt.Printf("%d: %d\n", num, MaximumSwap(num))
	}
}
