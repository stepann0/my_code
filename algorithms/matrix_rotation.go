package main

import "fmt"

type Matrix struct {
	Matrix     [][]int
	Rows, Cols int
}

func NewMatrix(rows, cols int) Matrix {
	m := make([][]int, rows)
	for i := range m {
		m[i] = make([]int, cols)
	}
	return Matrix{
		Matrix: m,
		Rows:   rows,
		Cols:   cols,
	}
}

func (m *Matrix) Transpose() {
	new_matrix := NewMatrix(m.Cols, m.Rows)
	for i := range new_matrix.Matrix {
		for j := range new_matrix.Matrix[i] {
			new_matrix.Matrix[i][j] = m.Matrix[j][i]
		}
	}
	*m = new_matrix
}

func (m *Matrix) Reflect() {
	for k := range m.Matrix {
		for i, j := 0, len(m.Matrix[k])-1; i < j; i, j = i+1, j-1 {
			m.Matrix[k][i], m.Matrix[k][j] = m.Matrix[k][j], m.Matrix[k][i]
		}
	}
}

func (m *Matrix) RotateMatrix() {
	m.Transpose()
	m.Reflect()
}

func main() {
	m := NewMatrix(3, 2)
	m.Matrix = [][]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	m.RotateMatrix()
	fmt.Println(m)
}

// 1 2
// 3 4
// 5 6
