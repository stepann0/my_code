package main

import (
	"testing"
)

func TestСlearZeroes(t *testing.T) {
	testCases := []struct {
		have Row
		want Row
	}{
		{
			Row{2, 0, 0, 4, 0, 5},
			Row{2, 4, 5, 0, 0, 0},
		},
		{
			Row{9, 0, 0, 1, 0, 0, 0, 0, 5, 5, 2, 0, 0},
			Row{9, 1, 5, 5, 2, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			Row{1, 0, 0, 9, 0, 1, 0, 0, 5},
			Row{1, 9, 1, 5, 0, 0, 0, 0, 0},
		},
		{
			Row{0, 1, 0, 3, 4},
			Row{1, 3, 4, 0, 0},
		},
		{
			Row{0, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 0},
		},
		{
			Row{5, 4, 3, 7, 2},
			Row{5, 4, 3, 7, 2},
		},
	}
	for _, tC := range testCases {
		tC.have.clearZeroes()
		if !tC.want.equal(tC.have) {
			t.Errorf("zeroes error: have %v, want %v", tC.have, tC.want)
		}
	}
}

func TestCompact(t *testing.T) {
	testCases := []struct {
		have  Row
		want  Row
		score int
	}{
		{
			Row{2, 0, 0, 4, 0, 4},
			Row{2, 8, 0, 0, 0, 0},
			8,
		},
		{
			Row{4, 2, 2, 0, 0, 0, 8, 8, 0, 0, 16, 0, 0},
			Row{4, 4, 16, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			20,
		},
		{
			Row{2, 2, 2, 2},
			Row{4, 4, 0, 0},
			8,
		},
		{
			Row{2, 4, 8, 16},
			Row{2, 4, 8, 16},
			0,
		},
		{
			Row{0, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 0},
			0,
		},
		{
			Row{4, 4, 2, 2, 8, 8, 0, 0},
			Row{8, 4, 16, 0, 0, 0, 0, 0},
			28,
		},
	}
	for _, tC := range testCases {
		score := tC.have.compact()
		if !tC.want.equal(tC.have) || score != tC.score {
			t.Errorf("compact error: have %v, want %v, score: %d", tC.have, tC.want, score)
		}
	}
}
