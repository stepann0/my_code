package main

import (
	"fmt"
)

func isValid(s string) bool {
	stack := []rune{}
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, parenth := range s {
		if pair, ok := pairs[parenth]; ok { // closing parenthesis
			l := len(stack)
			if l != 0 && stack[l-1] == pair {
				stack = stack[:l-1]
			} else {
				return false
			}
		} else { // opening parenthesis
			stack = append(stack, parenth)
		}
	}

	// left unclosed parentheses
	if len(stack) != 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println(isValid("(({[({{([])}})]}))"))
}
