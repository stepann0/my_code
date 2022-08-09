package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
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

	cell := empty_cells[rand.Intn(len(empty_cells))] // choose random cell

	// 90% probability for 2, and 10% for 4
	if v := rand.Float64(); v < 0.9 {
		f.Grid[cell[0]][cell[1]] = 2
	} else {
		f.Grid[cell[0]][cell[1]] = 4
	}
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

func (f *Field) Compact(row []int) ([]int, int) {
	if len(row) <= 1 {
		return row, 0
	}

	row = ClearZero(row)
	score := 0
	res := make([]int, len(row))
	for i, j := 0, 0; i < len(row); {
		if row[i] == 0 {
			i++
			continue
		}

		if i < len(row)-1 && row[i] == row[i+1] {
			sum := row[i] * 2
			res[j] = sum
			score += sum
			j++
			i += 2 // Skip unnesesary next element
			continue
		}
		// If not condition than just add element
		res[j] = row[i]
		j++
		i++
	}
	return res, score
}

// ------------------
//        LEFT
// ------------------

func (f *Field) Left() int {
	score := 0
	for i := 0; i < f.rows; i++ {
		var s int
		f.Grid[i], s = f.Compact(f.Grid[i])
		score += s
	}
	return score
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

func (f *Field) LazyCompactRight(row []int) ([]int, int) {
	row, score := f.Compact(reverse(row))
	return reverse(row), score
}

func (f *Field) Right() int {
	score := 0
	for i := 0; i < f.rows; i++ {
		var s int
		f.Grid[i], s = f.LazyCompactRight(f.Grid[i])
		score += s
	}
	return score
}

// ------------------
//         UP
// ------------------

func (f *Field) LazyCompactUp(col_num int) int {
	// Copy of colomn
	col_cp := make([]int, f.rows)
	for i := range col_cp {
		col_cp[i] = f.Grid[i][col_num]
	}

	col_cp, score := f.Compact(col_cp)

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
	col_cp := []int{}
	for i := f.rows - 1; i >= 0; i-- {
		col_cp = append(col_cp, f.Grid[i][col_num])
	}
	col_cp, score := f.Compact(col_cp)
	col_cp = reverse(col_cp)
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
	// fmt.Print("\033c\033[H")
	for _, row := range f.Grid {
		for _, i := range row {
			fmt.Printf("%4d", i)
		}
		fmt.Println()
	}
	fmt.Print("\n")
}

// ------------------------------
// ------------ Game ------------
// ------------------------------

type Game2048 struct {
	Field       *Field
	Score       int
	ColorScheme map[int]string
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
		g.Score += g.Field.Up()

	case DOWN:
		g.Score += g.Field.Down()

	case RIGHT:
		g.Score += g.Field.Right()

	case LEFT:
		g.Score += g.Field.Left()
	}
	g.Show(g.Field)

	// Add delay before adding random cell
	time.Sleep(time.Millisecond * 160)
	g.Field.AddRandom()
	g.Show(g.Field)
}

func (g *Game2048) Show(f *Field) {
	// Print game field
	fmt.Print("\033c\033[H") // clear terminal
	for _, row := range f.Grid {
		for _, r := range row {
			fmt.Print(g.Colorize(r))
		}
		fmt.Println()
	}
	// Print score
	fmt.Printf("\n%s%s%d%s\n\n", "\033[01;38;05;17;48;05;15m", "Score: ", g.Score, "\033[0m")
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

func (g *Game2048) Colorize(cell int) string {
	if cell == 0 {
		return " · "
	}
	return fmt.Sprintf("%s%3d%s", g.ColorScheme[cell], cell, "\033[0m")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	F := NewField(5, 5)
	game := Game2048{
		Field: &F,
		ColorScheme: map[int]string{
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
		},
	}
	game.Run()
}
