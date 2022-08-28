package tokens

import (
	"strconv"
	"strings"
)

type StrIterator struct {
	str    string
	next   int
	length int
}

func NewIterator(s string) StrIterator {
	if len(s) == 0 {
		panic("Can't iterate over an empty string.")
	}
	return StrIterator{
		str:    s,
		next:   0,
		length: len(s),
	}
}

func (it *StrIterator) Current() string {
	if it.next == 0 {
		panic("The Next function has never been called yet (can't return it.str[-1]).")
	}
	return string(it.str[it.next-1])
}

func (it *StrIterator) Next() string {
	if it.next < it.length {
		res := string(it.str[it.next])
		it.next++
		return res
	}
	return ""
}

func (it *StrIterator) CharAt(i int) string {
	return string(it.str[i])
}

func (it *StrIterator) ReadNum() (float64, error) {
	if it.next == 0 {
		panic("The Next function has never been called yet (can't return it.str[-1]).")
	}

	l := it.next - 1
	for it.next < it.length && strings.Contains("0123456789.", it.CharAt(it.next)) {
		it.Next()
	}

	res, err := strconv.ParseFloat(it.str[l:it.next], 64)
	if err != nil {
		return 0, err
	}
	return res, nil
}
