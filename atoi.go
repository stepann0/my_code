package main

import (
	"fmt"
	"math"
)

func has(arr []byte, s byte) bool {
	for j := range arr {
		if s == arr[j] {
			return true
		}
	}
	return false
}

func clear(s, patt string) string {
	p := 0
	for i := 0; i < len(s); i++ {
		if s[i] != []byte(patt)[0] {
			p = i
			break
		}
	}
	return s[p:]
}

func check_sign(s string) (string, int) {
	if len(s) > 0 {
		sign := s[0]
		switch sign {
		case '-':
			return s[1:], -1
		case '+':
			return s[1:], 1
		default:
			return s, 1
		}
	} else {
		return "", 1
	}
}

func read(s string) (string, int) {
	digits := []byte("0123456789")

	s = clear(s, " ")
	var sign int
	s, sign = check_sign(s)
	s = clear(s, "0")

	num := []byte("")

	// if first symbol isnt digit
	if len(s) == 0 || !has(digits, s[0]) {
		return "", 1
	}
	// read number
	for i := 0; i < len(s); i++ {
		if has(digits, s[i]) {
			num = append(num, s[i])
		} else {
			break
		}
	}
	// convert to a string
	number := string(num)
	return number, sign
}

func atoi(s string, sign int) int {
	var number float64
	if len(s) == 0 {
		return 0
	}
	pow := len(s) - 1
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case []byte("0")[0]:
			number += 0
		case []byte("1")[0]:
			number += 1 * math.Pow(10, float64(pow))
		case []byte("2")[0]:
			number += 2 * math.Pow(10, float64(pow))
		case []byte("3")[0]:
			number += 3 * math.Pow(10, float64(pow))
		case []byte("4")[0]:
			number += 4 * math.Pow(10, float64(pow))
		case []byte("5")[0]:
			number += 5 * math.Pow(10, float64(pow))
		case []byte("6")[0]:
			number += 6 * math.Pow(10, float64(pow))
		case []byte("7")[0]:
			number += 7 * math.Pow(10, float64(pow))
		case []byte("8")[0]:
			number += 8 * math.Pow(10, float64(pow))
		case []byte("9")[0]:
			number += 9 * math.Pow(10, float64(pow))
		}
		pow--
	}
	return int(number) * sign
}

func test() bool {
	test := map[string]int{
		"   123sd  ":        123,
		"  a-2435ccc  ":     0,
		"0345kk":            345,
		"000340640":         340640,
		"4":                 4,
		"":                  0,
		"what a day! 23456": 0,
		"3405000340":        3405000340,
	}
	for t, answ := range test {
		if atoi(read(t)) != answ {
			return false
		}
	}
	return true
}

func main() {
	// fmt.Println(atoi(read(" 45607")))
	fmt.Println(test())
}
