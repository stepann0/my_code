package main

import (
	"fmt"
	"strconv"
	"strings"
)

// функция dfs добавляет неправильное значение subset в store
// поэтому нужен bugfix, который восстанавливает правильное значение
// из строкового представления subset
func bugfix(s string) []int {
	if len(s) <= 2 {
		return []int{}
	}

	s = s[1 : len(s)-1]
	res := []int{}
	for _, i := range strings.Split(s, " ") {
		conv, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		res = append(res, conv)
	}
	return res
}

func dfs(nums []int, subset []int, i int, store *[][]int) {
	if i == len(nums) {
		fmt.Printf("I wanna add %v\n", subset)
		*store = append(*store, subset)
		fmt.Printf("but I add %v\n", (*store)[len(*store)-1])
		return
	}
	// не добавляем элемент в subset
	dfs(nums, subset, i+1, store)

	// добавляем элемент в subset
	subset = append(subset, nums[i])
	dfs(nums, subset, i+1, store)
}

func subsets(nums []int) [][]int {
	store := [][]int{}
	dfs(nums, []int{}, 0, &store)
	return store
}

func Has(store [][]int, seq []int) bool {
	for i := 0; i < len(store); i++ {
		if Equal((store)[i], seq) {
			return true
		}
	}
	return false
}

func Equal(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func main() {
	// subsets([]int{9, 0, 3, 5, 7})
	subs := subsets([]int{4, 0, 7, 2, 1, 6})
	fmt.Println(subs)
}
