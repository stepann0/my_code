package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/exec"
)

type Field struct {
	Grid       [][]int
	rows, cols int
}

func NewField(rows, cols int) Field {
	if cols < 2 || rows < 2 {
		panic("Two rows and coloms minimum.\n")
	}

	grid := make([][]int, rows)
	for i := range grid {
		grid[i] = make([]int, cols)
	}

	F := Field{
		grid, cols, rows,
	}
	F.AddRandom()
	F.AddRandom()
	F.AddRandom()

	return F
}

func (f *Field) AddRandom() {
	empty_cells := [][2]int{} // like slice of tuples
	for i := range f.Grid {
		for j := range f.Grid[i] {
			if f.Grid[i][j] == 0 {
				empty_cells = append(empty_cells, [2]int{i, j})
			}
		}
	}

	c := empty_cells[rand.Intn(len(empty_cells))]
	vals := []int{2, 4} // you can add 2 or 4
	v := vals[rand.Intn(len(vals))]
	f.Grid[c[0]][c[1]] = v
}

func ClearZero(s []int) []int {
	res := make([]int, len(s))
	for i, j := 0, 0; i < len(s); i++ {
		if s[i] != 0 {
			res[j] = s[i]
			j++
		}
	}
	return res
}

func (f *Field) Compact(row []int) []int {
	if len(row) <= 1 {
		return row
	}

	row = ClearZero(row)
	res := make([]int, len(row))
	for i, j := 0, 0; i < len(row); {
		if row[i] == 0 {
			i++
			continue
		}

		if i < len(row)-1 && row[i] == row[i+1] {
			res[j] = row[i] + row[i+1]
			j++
			i += 2 // Escape unnesesary next element
			continue
		}
		// If not condition than just add element
		res[j] = row[i]
		j++
		i++
	}
	return res
}

// ------------------
//        LEFT
// ------------------

func (f *Field) Left() {
	for i := 0; i < f.rows; i++ {
		f.Grid[i] = f.Compact(f.Grid[i])
	}
}

// ------------------
//        RIGHT
// ------------------

func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func (f *Field) LazyCompactRight(row []int) []int {
	return reverse(f.Compact(reverse(row)))
}

func (f *Field) Right() {
	for i := 0; i < f.rows; i++ {
		f.Grid[i] = f.LazyCompactRight(f.Grid[i])
	}
}

// ------------------
//         UP
// ------------------

func (f *Field) LazyCompactUp(col_num int) {
	// Copy of colomn
	col_cp := []int{}
	for i := 0; i < f.rows; i++ {
		col_cp = append(col_cp, f.Grid[i][col_num])
	}

	col_cp = f.Compact(col_cp)

	for i := 0; i < f.rows; i++ {
		f.Grid[i][col_num] = col_cp[i]
	}
}

func (f *Field) Up() {
	for i := 0; i < f.cols; i++ {
		f.LazyCompactUp(i)
	}
}

// ------------------
//        DOWN
// ------------------

func (f *Field) LazyCompactDown(col_num int) {
	// Copy of colomn
	col_cp := []int{}
	for i := f.rows - 1; i >= 0; i-- {
		col_cp = append(col_cp, f.Grid[i][col_num])
	}
	col_cp = reverse(f.Compact(col_cp))
	for i := 0; i < f.rows; i++ {
		f.Grid[i][col_num] = col_cp[i]
	}
}

func (f *Field) Down() {
	for i := 0; i < f.cols; i++ {
		f.LazyCompactDown(i)
	}
}

// Print for debugging
func (f *Field) Print() {
	fmt.Print("\033c\033[H")
	for _, row := range f.Grid {
		for _, i := range row {
			fmt.Printf("%4d", i)
		}
		fmt.Println()
	}
	fmt.Print("\n")
}

func Colorize(cell int) string {
	if cell == 0 {
		return " · "
	}
	BACK := "\033[0m"
	colors := map[int]string{
		2:    "\033[01;38;05;16;48;05;158m",  // 2
		4:    "\033[01;38;05;15;48;05;42m",   // 4
		8:    "\033[01;38;05;15;48;05;33m",   // 8
		16:   "\033[01;38;05;15;48;05;91m",   // 16
		32:   "\033[01;38;05;91;48;05;182m",  // 32
		64:   "\033[01;38;05;17;48;05;69m",   // 64
		128:  "\033[01;38;05;52;48;05;203m",  // 128
		256:  "\033[01;38;05;15;48;05;24m",   // 256
		512:  "\033[01;38;05;233;48;05;208m", // 512
		1024: "\033[01;38;05;232;48;05;220m", // 1024
		2048: "\033[01;38;05;232;48;05;83m",  // 2048
	}
	return fmt.Sprintf("%s%3d%s", colors[cell], cell, BACK)
}

func (f *Field) Show() {
	fmt.Print("\033c\033[H")
	for _, row := range f.Grid {
		for _, i := range row {
			fmt.Print(Colorize(i))
		}
		fmt.Println()
	}
	fmt.Println()
}

// ------------------------------
// ------------ Game ------------
// ------------------------------

type Game2048 struct {
	Field *Field
	debug bool
}

func (g *Game2048) Move(key uint8) {
	const (
		UP    uint8 = 65
		DOWN  uint8 = 66
		RIGHT uint8 = 67
		LEFT  uint8 = 68
	)

	switch key {
	case UP:
		g.Field.Up()

	case DOWN:
		g.Field.Down()

	case RIGHT:
		g.Field.Right()

	case LEFT:
		g.Field.Left()
	}
	g.Field.AddRandom()

	if g.debug {
		g.Field.Print()
	} else {
		g.Field.Show()
	}
}

func (g *Game2048) Run() {
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	key := ListenKey()
	if g.debug {
		g.Field.Print()
	} else {
		g.Field.Show()
	}
	for {
		g.Move(<-key)
	}
}

func ListenKey() <-chan uint8 {
	c := make(chan uint8)
	var b []byte = make([]byte, 3)
	go func() {
		for {
			os.Stdin.Read(b)
			// If arrow key pressed
			if len(b) >= 3 && b[0] == uint8(27) && b[1] == uint8(91) {
				c <- b[2]
			}
		}
	}()
	return c
}

func main() {
	F := NewField(4, 4)
	game := Game2048{
		Field: &F,
		debug: false,
	}
	game.Run()
}
