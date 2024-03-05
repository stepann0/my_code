package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strings"
	"time"

	"golang.org/x/term"
)

type Row []int

// сдвигает строку в лево по правилам 2048 и возврацает заработанные очки
// {4, 4, 2, 2, 8, 8, 0, 0} -> {8, 4, 16, 0, 0, 0, 0, 0} = 28
func (row *Row) compact() int {
	length := len(*row)
	if length <= 1 {
		return 0
	}

	row.clearZeroes()
	score := 0
	res := make(Row, length)
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
	(*row) = res
	return score
}

// передвигает все 0 в конец: {2, 0, 0, 4, 0, 4} -> {2, 4, 4, 0, 0, 0}
func (row Row) clearZeroes() {
	sort.SliceStable(row, func(i, j int) bool {
		return row[j] == 0
	})
}

// переворачивает строку
func (row Row) reverse() {
	for i, j := 0, len(row)-1; i < j; i, j = i+1, j-1 {
		row[i], row[j] = row[j], row[i]
	}
}

func (row1 Row) equal(row2 Row) bool {
	if len(row1) != len(row2) {
		return false
	}
	for i := range row1 {
		if row1[i] != row2[i] {
			return false
		}
	}
	return true
}

func (row Row) colorize(colors map[int]string) string {
	var s strings.Builder
	for _, cell := range row {
		if cell == 0 {
			fmt.Fprint(&s, colors[0])
			continue
		}
		fmt.Fprintf(&s, "%s%3d%s", colors[cell], cell, "\033[0m")
	}
	return s.String()
}

type Field struct {
	Grid       []Row
	rows, cols int
}

func NewField(rows, cols int) Field {
	if cols < 2 || rows < 2 {
		panic("two rows and coloms minimum.\n")
	}

	grid := make([]Row, rows)
	for i := range grid {
		grid[i] = make([]int, cols)
	}

	F := Field{
		grid, cols, rows,
	}
	F.addRandom()
	F.addRandom()
	F.addRandom()

	return F
}

func (f *Field) addRandom() {
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

func (f1 *Field) equal(f2 *Field) bool {
	for i := range f1.Grid {
		if !f1.Grid[i].equal(f2.Grid[i]) {
			return false
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
		score += f.Grid[i].compact()
	}
	return score
}

// ------------------
//        RIGHT
// ------------------

func (f *Field) compactRight(i int) int {
	f.Grid[i].reverse()
	score := f.Grid[i].compact()
	f.Grid[i].reverse()
	return score
}

func (f *Field) Right() int {
	score := 0
	for i := 0; i < f.rows; i++ {
		score += f.compactRight(i)
	}
	return score
}

// ------------------
//         UP
// ------------------

func (f *Field) compactUp(col_num int) int {
	// Copy of colomn
	col_cp := make(Row, f.rows)
	for i := range col_cp {
		col_cp[i] = f.Grid[i][col_num]
	}

	score := col_cp.compact()
	for i := 0; i < f.rows; i++ {
		f.Grid[i][col_num] = col_cp[i]
	}
	return score
}

func (f *Field) Up() int {
	score := 0
	for i := 0; i < f.cols; i++ {
		score += f.compactUp(i)
	}
	return score
}

// ------------------
//        DOWN
// ------------------

func (f *Field) compactDown(col_num int) int {
	// reversed copy of colomn
	col_cp := Row{}
	for i := f.rows - 1; i >= 0; i-- {
		col_cp = append(col_cp, f.Grid[i][col_num])
	}

	score := col_cp.compact()
	col_cp.reverse()
	for i := 0; i < f.rows; i++ {
		f.Grid[i][col_num] = col_cp[i]
	}
	return score
}

func (f *Field) Down() int {
	score := 0
	for i := 0; i < f.cols; i++ {
		score += f.compactDown(i)
	}
	return score
}

// Print for debugging
func (f Field) String() string {
	s := ""
	for i := range f.Grid {
		for j := range f.Grid[i] {
			s += fmt.Sprintf("%5d", f.Grid[i][j])
		}
		s += "\n"
	}
	s += fmt.Sprintf("%d×%d\n", f.rows, f.cols)
	return s
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
	old_field := g.Field.Copy()
	switch key {
	case keyUp:
		g.Score += g.Field.Up()
	case keyDown:
		g.Score += g.Field.Down()
	case keyRight:
		g.Score += g.Field.Right()
	case keyLeft:
		g.Score += g.Field.Left()
	}

	g.Show(g.Field)
	if old_field.equal(g.Field) {
		return
	}

	// Add delay before adding random cell
	time.Sleep(time.Millisecond * 180)
	g.Field.addRandom()
	g.Show(g.Field)
}

func (g *Game2048) Show(f *Field) {
	// Print game field
	fmt.Println("\033c\033[H") // clear terminal
	for i := range f.Grid {
		fmt.Printf(" %s\n\r", f.Grid[i].colorize(g.ColorScheme))
	}
	// Print score. g.ColorScheme[-1] is a color for score
	fmt.Printf("\n%s %s%d %s\n\r", g.ColorScheme[-1], "Score: ", g.Score, "\033[0m")
}

func (g *Game2048) Run() {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	keys := listenKey()
	g.Show(g.Field)
	for {
		key := <-keys
		switch key {
		case keyCtrlC, keyCtrlD:
			os.Exit(0)
		case keyUp, keyDown, keyLeft, keyRight:
			g.Move(key)
		}
	}
}

const (
	keyUnknown uint8 = 0
	keyCtrlC         = 3
	keyCtrlD         = 4
	keyEnter         = '\r'
	keyEscape        = '\033'
	keyUp            = iota
	keyDown
	keyLeft
	keyRight
)

func bytesToKey(buf []byte) uint8 {
	switch len(buf) {
	case 1:
		if buf[0] == 3 || buf[0] == 4 || buf[0] == '\r' || buf[0] == '\033' {
			return uint8(buf[0]) // ^C, ^D, enter, ESC
		}
	case 3:
		if buf[0] == 27 && buf[1] == 91 {
			switch buf[2] {
			case 'A':
				return keyUp
			case 'B':
				return keyDown
			case 'C':
				return keyRight
			case 'D':
				return keyLeft
			}
		}
	}
	return keyUnknown
}

func listenKey() <-chan uint8 {
	c := make(chan uint8)
	go func() {
		buf := make([]byte, 16)
		for {
			n, err := os.Stdin.Read(buf)
			if err != nil {
				panic(err)
			}
			c <- bytesToKey(buf[:n])
		}
	}()
	return c
}

var ColorScheme = map[int]string{
	0:    " · ",
	2:    rgbBg(16, 48, 5, 158),
	4:    rgbBg(15, 48, 5, 42),
	8:    rgbBg(15, 48, 5, 33),
	16:   rgbBg(15, 48, 5, 98),
	32:   rgbBg(91, 48, 5, 182),
	64:   rgbBg(15, 48, 5, 102),
	128:  rgbBg(52, 48, 5, 203),
	256:  rgbBg(15, 48, 5, 24),
	512:  rgbBg(233, 48, 5, 208),
	1024: rgbBg(232, 48, 5, 220),
	2048: rgbBg(232, 48, 5, 83),
	-1:   rgbBg(232, 48, 5, 194), // for score
}

func rgbBg(c, r, g, b int) string {
	return fmt.Sprintf("%c[01;38;05;%d;%d;%d;%dm", keyEscape, c, r, g, b)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	F := NewField(4, 4)
	game := Game2048{
		Field:       &F,
		ColorScheme: ColorScheme,
	}
	game.Run()
}
