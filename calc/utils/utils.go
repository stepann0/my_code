package utils

import (
	"fmt"
	"math"
	"strconv"
)

func is_operator(s string) bool {
	switch s {
	case "%", "*", "+", "-", "/", "^":
		return true
	}
	return false
}

func is_num(s string) bool {
	if len(s) != 1 {
		return false
	}
	r := rune(s[0])
	if r >= '0' && r <= '9' {
		return true
	}
	return false
}

func eval_binary(operator rune, o1, o2 float64) float64 {
	switch operator {
	case '+':
		return o2 + o1
	case '-':
		return o2 - o1
	case '*':
		return o2 * o1
	case '/':
		return o2 / o1
	case '^':
		return math.Pow(o2, o1)
	}
	panic(fmt.Errorf("Wrong operator."))
}

func rmLeftPar(s *Stack[rune]) {
	for !s.Empty() {
		if s.Top() == '(' {
			s.Pop()
		} else {
			return
		}
	}
}

func rune2float(num rune) float64 {
	res, _ := strconv.ParseFloat(string(num), 64)
	return res
}

func runes2str(r []rune) string {
	return string(r)
}
