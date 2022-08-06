package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
)

type Field struct {
	Grid       [][]int
	rows, cols int
}

func NewField(rows, cols int) (Field, error) {
	if cols < 2 && rows < 2 {
		return Field{[][]int{}, 0, 0},
			fmt.Errorf("Two rows and coloms minimum.\n")
	}

	grid := [][]int{}
	for i := 0; i < rows; i++ {
		grid = append(grid, []int{})
		for j := 0; j < cols; j++ {
			grid[i] = append(grid[i], 0)
		}
	}

	F := Field{
		grid, cols, rows,
	}
	F.AddRandom()
	F.AddRandom()
	F.AddRandom()

	return F, nil
}

func (f *Field) AddRandom() {
	zero_indexes := [][2]int{}
	for i := 0; i < len(f.Grid); i++ {
		for j := 0; j < len(f.Grid[i]); j++ {
			if f.Grid[i][j] == 0 {
				zero_indexes = append(zero_indexes, [2]int{i, j})
			}
		}
	}

	r := zero_indexes[rand.Intn(len(zero_indexes))]
	vals := []int{2, 4}
	v := vals[rand.Intn(len(vals))]
	f.Grid[r[0]][r[1]] = v
}

func ClearZero(s []int) []int {
	res := []int{}
	for _, v := range s {
		if v != 0 {
			res = append(res, v)
		}
	}

	for len(res) < len(s) {
		res = append(res, 0)
	}
	return res
}
func (f *Field) Compact(row []int) []int {
	if len(row) <= 1 {
		return row
	}

	row = ClearZero(row)
	res := []int{}
	for i := 0; i < len(row); {
		if row[i] == 0 {
			i++
			continue
		}

		if i < len(row)-1 && row[i] == row[i+1] {
			res = append(res, row[i]+row[i+1])
			i += 2 // Escape unnesesary next element
			continue
		}
		// If not condition than just add element
		res = append(res, row[i])
		i++
	}

	for len(res) < len(row) {
		res = append(res, 0)
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
	for i := 0; i < len(f.Grid); i++ {
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
	fmt.Print("\n\n")
}

func Colorize(cell int) string {
	const (
		// <BACKGROUND>_<FOREGROUND> format
		BACK         = "\033[0m"
		BLANK        = "\033[01;48;05;253m"           // for 0
		WHITE_BLACK  = "\033[01;38;05;232;48;05;15m"  // for 2
		LGREY_DGREY  = "\033[01;38;05;238;48;05;252m" // for 4
		YELLOW_BROWN = "\033[01;38;05;172;48;05;227m" // for 8
		ORANGE_WHITE = "\033[01;38;05;15;48;05;214m"  // for 16
		RED_WHITE    = "\033[01;38;05;15;48;05;166m"  // for 32

		VL_GREEN_DGREEN = "\033[01;38;05;22;48;05;188m"
		SL_GREEN_DGREEN = "\033[01;38;05;22;48;05;121m"
		GREEN_DRGEEN    = "\033[01;38;05;16;48;05;83m"
		LBLUE_DBLUE     = "\033[01;38;05;16;48;05;86m"
		BLUE_WHITE      = "\033[01;38;05;194;48;05;27m"

		A = "\033[01;38;05;234;48;05;158m"
		B = "\033[01;38;05;15;48;05;42m"
		C = "\033[01;38;05;15;48;05;62m"
		D = "\033[01;38;05;15;48;05;93m"
		E = "\033[01;38;05;15;48;05;205m"
		F = "\033[01;38;05;232;48;05;190m"
		G = "\033[01;38;05;15;48;05;214m"
		H = "\033[01;38;05;15;48;05;196m"
	)

	var color string
	switch cell {
	case 0:
		// TODO correct spaces in blank cells
		// return fmt.Sprintf("%s    %s", BLANK, BACK)
		return "   "
	case 2:
		color = A
	case 4:
		color = B
	case 8:
		color = C
	case 16:
		color = D
	case 32:
		color = E
	case 64:
		color = F
	case 128:
		color = G
	case 256:
		color = H
	}
	return fmt.Sprintf("%s%3d%s", color, cell, BACK)
}
func (f *Field) Show() {
	fmt.Print("\033c\033[H")
	for _, row := range f.Grid {
		for _, i := range row {
			fmt.Print(Colorize(i))
		}
		fmt.Println()
	}
	fmt.Print("\n\n")
}

// ------------------------------
// ------------ Game ------------
// ------------------------------
type Game2048 struct {
	Field *Field
	debug bool
}

func (g *Game2048) Move(key uint8) {
	const UP uint8 = 65
	const DOWN uint8 = 66
	const RIGHT uint8 = 67
	const LEFT uint8 = 68

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
			// If arrow
			if len(b) >= 3 && b[0] == uint8(27) && b[1] == uint8(91) {
				c <- b[2]
			}
		}
	}()
	return c
}

func main() {
	F, err := NewField(4, 4)
	if err != nil {
		panic(err)
	}
	game := Game2048{
		Field: &F,
		debug: false,
	}
	game.Run()
}
