package main

import (
	"fmt"
)

var pairs map[string]string = map[string]string{
	")": "(",
	"}": "{",
	"]": "[",
}

func isClosePar(p string) bool {
	//  == instead strings.Contains() is for speed
	if p == ")" || p == "]" || p == "}" {
		return true
	}
	return false
}

func isValid(s string) bool {
	stack := []string{}

	for _, v := range s {
		par := string(v)

		if isClosePar(par) {
			l := len(stack)
			if l != 0 && stack[l-1] == pairs[par] {
				stack = stack[:l-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, par)
		}
	}

	// unclosed parenthes remain
	if len(stack) != 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println(isValid("(]"))
}
