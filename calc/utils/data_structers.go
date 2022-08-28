package utils

type Stack[T any] []T

func (s *Stack[T]) Push(val T) {
	*s = append(*s, val)
}

func (s *Stack[T]) Pop() T {
	l := len(*s)
	if l == 0 {
		panic("Stack is empty.")
	}
	val := (*s)[l-1]
	*s = (*s)[:l-1]
	return val
}

func (s *Stack[T]) Top() T {
	l := len(*s)
	if l == 0 {
		panic("Stack is empty.")
	}
	return (*s)[l-1]
}

func (s *Stack[T]) Empty() bool {
	return len(*s) == 0
}

type Queue[T any] []T

func (q *Queue[T]) Push(val T) {
	*q = append(*q, val)
}

func (q *Queue[T]) Pop() T {
	if len(*q) == 0 {
		panic("Queue is empty.")
	}
	val := (*q)[0]
	*q = (*q)[1:]
	return val
}

func (q *Queue[T]) Empty() bool {
	return len(*q) == 0
}
