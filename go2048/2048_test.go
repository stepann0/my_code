package main

import (
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		in  Row
		out Row
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
		tC.in.clearZeroes()
		if !tC.out.equal(tC.in) {
			t.Errorf("error: have %v, want %v", tC.in, tC.out)
		}
	}
}
