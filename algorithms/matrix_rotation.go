package main

import "fmt"

// Generics
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Integer interface {
	Signed | Unsigned
}

type Float interface {
	~float32 | ~float64
}

type Number interface {
	Integer | Float
}

type Matrix[T Number] struct {
	Rows, Cols int
	Matrix     [][]T
}

func NewEmpty[T Number](rows, cols int) *Matrix[T] {
	var el T
	m := make([][]T, rows)
	for i := 0; i < rows; i++ {
		m = append(m, []T{})
		for j := 0; j < cols; j++ {
			m[i] = append(m[i], el)
		}
	}
	return &Matrix[T]{
		Rows:   rows,
		Cols:   cols,
		Matrix: m,
	}
}

func New[T Number](arr [][]T) (*Matrix[T], error) {
	rows := len(arr)
	if rows == 0 {
		return &Matrix[T]{
			Rows: 0, Cols: 0,
			Matrix: [][]T{},
		}, nil
	}
	cols := len(arr[0])

	for i := 0; i < rows; i++ {
		if len(arr[i]) != cols {
			return nil, fmt.Errorf("Не все строки одинаковой длины.")
		}
	}
	return &Matrix[T]{
		Rows:   rows,
		Cols:   cols,
		Matrix: arr,
	}, nil
}

func (a *Matrix[T]) Mul(b *Matrix[T]) (*Matrix[T], error) {
	if a.Cols != b.Rows {
		return nil, fmt.Errorf("a.Cols != b.Rows!")
	}
	res := NewEmpty[T](a.Rows, b.Cols)

	for i := 0; i < res.Rows; i++ {
		for j := 0; j < res.Cols; j++ {
			for k := 0; k < b.Rows; k++ {
				res.Matrix[i][j] += a.Matrix[i][k] * b.Matrix[k][j]
			}
		}
	}
	return res, nil
}

func Vector[T Number](v []T) *Matrix[T] {
	return &Matrix[T]{
		Rows:   1,
		Cols:   len(v),
		Matrix: [][]T{v},
	}
}

func (a *Matrix[T]) Rotate(b *Matrix[T]) (*Matrix[T], error) {

}

func main() {
	a, _ := New([][]int{})

	b, _ := New([][]int{})
	fmt.Println(a.Mul(b))
}
