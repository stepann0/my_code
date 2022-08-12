package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

type Row []int

func (row *Row) Compact() int {
	length := len(*row)
	if length <= 1 {
		return 0
	}

	row.ClearZeroes()
	score := 0
	res := make([]int, length)
	for i, j := 0, 0; i < length; {
		if (*row)[i] == 0 {
			i++
			continue
		}

		if i < length-1 && (*row)[i] == (*row)[i+1] {
			sum := (*row)[i] * 2
			res[j] = sum
			score += sum
			j++
			i += 2 // Skip unnesesary next element
			continue
		}
		// If not condition than just add element
		res[j] = (*row)[i]
		j++
		i++
	}
	*row = res
	return score
}

func (row *Row) ClearZeroes() {
	res := make([]int, len(*row))
	for i, j := 0, 0; i < len(*row); i++ {
		if (*row)[i] != 0 {
			res[j] = (*row)[i]
			j++
		}
	}
	*row = res
}

func (row *Row) Reverse() {
	for i, j := 0, len(*row)-1; i < j; i, j = i+1, j-1 {
		(*row)[i], (*row)[j] = (*row)[j], (*row)[i]
	}
}

func (row *Row) Colorize(colors *map[int]string) string {
	s := ""
	for _, cell := range *row {
		if cell == 0 {
			s += (*colors)[0]
			continue
		}
		s += fmt.Sprintf("%s%3d%s", (*colors)[cell], cell, "\033[0m")
	}
	return s
}

type Field struct {
	Grid       []Row
	rows, cols int
}

func NewField(rows, cols int) Field {
	if cols < 2 || rows < 2 {
		panic("Two rows and coloms minimum.\n")
	}

	grid := make([]Row, rows)
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

	cell := empty_cells[rand.Intn(len(empty_cells))] // choose random cell

	// 90% probability for 2, and 10% for 4
	if v := rand.Float64(); v < 0.9 {
		f.Grid[cell[0]][cell[1]] = 2
	} else {
		f.Grid[cell[0]][cell[1]] = 4
	}
}

func (f1 *Field) Equal(f2 *Field) bool {
	if f1.cols != f2.cols || f1.rows != f2.rows {
		return false
	}
	for i := range f1.Grid {
		for j := range f1.Grid[i] {
			if f1.Grid[i][j] != f2.Grid[i][j] {
				return false
			}
		}
	}
	return true
}

func (f *Field) Copy() Field {
	grid := make([]Row, f.rows)
	for i := range grid {
		grid[i] = make(Row, f.cols)
		for j := range f.Grid[i] {
			grid[i][j] = f.Grid[i][j]
		}
	}
	return Field{
		Grid: grid,
		cols: f.cols,
		rows: f.rows,
	}
}

// ------------------
//        LEFT
// ------------------

func (f *Field) Left() int {
	score := 0
	for i := 0; i < f.rows; i++ {
		score += f.Grid[i].Compact()
	}
	return score
}

// ------------------
//        RIGHT
// ------------------

func (f *Field) LazyCompactRight(i int) int {
	f.Grid[i].Reverse()
	score := f.Grid[i].Compact()
	f.Grid[i].Reverse()
	return score
}

func (f *Field) Right() int {
	score := 0
	for i := 0; i < f.rows; i++ {
		score += f.LazyCompactRight(i)
	}
	return score
}

// ------------------
//         UP
// ------------------

func (f *Field) LazyCompactUp(col_num int) int {
	// Copy of colomn
	col_cp := make(Row, f.rows)
	for i := range col_cp {
		col_cp[i] = f.Grid[i][col_num]
	}

	score := col_cp.Compact()
	for i := 0; i < f.rows; i++ {
		f.Grid[i][col_num] = col_cp[i]
	}
	return score
}

func (f *Field) Up() int {
	score := 0
	for i := 0; i < f.cols; i++ {
		score += f.LazyCompactUp(i)
	}
	return score
}

// ------------------
//        DOWN
// ------------------

func (f *Field) LazyCompactDown(col_num int) int {
	// Reversed copy of colomn
	col_cp := Row{}
	for i := f.rows - 1; i >= 0; i-- {
		col_cp = append(col_cp, f.Grid[i][col_num])
	}

	score := col_cp.Compact()
	col_cp.Reverse()
	for i := 0; i < f.rows; i++ {
		f.Grid[i][col_num] = col_cp[i]
	}
	return score
}

func (f *Field) Down() int {
	score := 0
	for i := 0; i < f.cols; i++ {
		score += f.LazyCompactDown(i)
	}
	return score
}

// Print for debugging
func (f *Field) Print() {
	for _, row := range f.Grid {
		for _, i := range row {
			fmt.Printf("%4d", i)
		}
		fmt.Println()
	}
	fmt.Println()
}

// ------------------------------
// ------------ Game ------------
// ------------------------------

type Game2048 struct {
	Field       *Field
	Score       int
	ColorScheme *map[int]string
}

func (g *Game2048) Move(key uint8) {
	const (
		UP    uint8 = 65
		DOWN  uint8 = 66
		RIGHT uint8 = 67
		LEFT  uint8 = 68
	)
	old_field := g.Field.Copy()
	switch key {
	case UP:
		g.Score += g.Field.Up()

	case DOWN:
		g.Score += g.Field.Down()

	case RIGHT:
		g.Score += g.Field.Right()

	case LEFT:
		g.Score += g.Field.Left()
	}

	g.Show(g.Field)
	if old_field.Equal(g.Field) {
		return
	}

	// Add delay before adding random cell
	time.Sleep(time.Millisecond * 180)
	g.Field.AddRandom()
	g.Show(g.Field)
}

func (g *Game2048) Show(f *Field) {
	// Print game field
	fmt.Println("\033c\033[H") // clear terminal
	for i := range f.Grid {
		fmt.Printf(" %s\n", f.Grid[i].Colorize(g.ColorScheme))
	}
	// Print score. g.ColorScheme[-1] is a color for score
	fmt.Printf("\n%s %s%d %s\n\n", (*g.ColorScheme)[-1], "Score: ", g.Score, "\033[0m")
}

func (g *Game2048) Run() {
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	key := ListenKey()
	g.Show(g.Field)
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
	rand.Seed(time.Now().UnixNano())
	F := NewField(4, 4)
	game := Game2048{
		Field: &F,
		ColorScheme: &map[int]string{
			0:    " · ",                          // 0
			2:    "\033[01;38;05;16;48;05;158m",  // 2
			4:    "\033[01;38;05;15;48;05;42m",   // 4
			8:    "\033[01;38;05;15;48;05;33m",   // 8
			16:   "\033[01;38;05;15;48;05;98m",   // 16
			32:   "\033[01;38;05;91;48;05;182m",  // 32
			64:   "\033[01;38;05;15;48;05;102m",  // 64
			128:  "\033[01;38;05;52;48;05;203m",  // 128
			256:  "\033[01;38;05;15;48;05;24m",   // 256
			512:  "\033[01;38;05;233;48;05;208m", // 512
			1024: "\033[01;38;05;232;48;05;220m", // 1024
			2048: "\033[01;38;05;232;48;05;83m",  // 2048
			-1:   "\033[01;38;05;232;48;05;194m", // for score
		},
	}
	game.Run()
}
