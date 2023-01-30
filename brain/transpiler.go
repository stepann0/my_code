package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"unicode"
)

// printing modes
const (
	charMode = iota
	decimalMode
)

type Compiler struct {
	program     []byte
	pointerName string
	memName     string
	memSize     int
	outFile     *os.File
	indent      int
	printMode   int
}

func NewDefaultCompiler() Compiler {
	return Compiler{
		pointerName: "p",
		memName:     "mem",
		memSize:     30000,
		outFile:     os.Stdout,
		printMode:   charMode,
	}
}

func (c *Compiler) loadProgram(name string) {
	f, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}
	c.program = f
}

// TODO: refactor this function
func (c *Compiler) prologue() {
	fmt.Fprint(c.outFile, "package main\n\n")

	packages := c.checkImports()
	if im := c.imports(packages); im != "" {
		fmt.Fprintf(c.outFile, "%s\n\n", im)
	}

	has_bufio, has_os := false, false
	for _, p := range packages {
		if p == "bufio" {
			has_bufio = true
		}
		if p == "os" {
			has_os = true
		}
	}

	getchar_func := ""
	stdin := ""
	if has_bufio && has_os {
		getchar_func =
			`func getchar(reader *bufio.Reader) byte {
	b, err := reader.ReadByte()
	if err != nil {
		panic(err)
	}
	return b
}`
		stdin = "stdin := bufio.NewReader(os.Stdin)\n"
	}

	if getchar_func != "" {
		fmt.Fprintf(c.outFile, "%s\n\n", getchar_func)
	}

	func_main :=
		`func main() {
	%s := [%d]byte{}
	%s := 0
	%s
`
	fmt.Fprintf(c.outFile, func_main,
		c.memName, c.memSize, c.pointerName, stdin)

	c.indent++
}

func (c *Compiler) imports(packages []string) string {
	sort.Strings(packages)
	switch len(packages) {
	case 0:
		return ""
	case 1:
		return fmt.Sprintf("import \"%s\"", packages[0])
	default:
		imports := "import (\n"
		c.indent++
		for _, p := range packages {
			imports += fmt.Sprintf("%s\"%s\"\n", c.getIndent(), p)
		}
		c.indent--
		imports += ")"
		return imports
	}
}

func (c *Compiler) checkImports() []string {
	dot, comma := false, false
	for _, c := range c.program {
		if c == '.' {
			dot = true
		} else if c == ',' {
			comma = true
		}
	}
	if dot && comma {
		return []string{"fmt", "bufio", "os"}
	}
	if dot && (comma == false) {
		return []string{"fmt"}
	}
	if comma && (dot == false) {
		return []string{"bufio", "os"}
	}
	return []string{}
}

func (c *Compiler) incPointer(n int) {
	if n < 0 {
		c.decPointer(-n)
		return
	}

	c.writeIndent()
	if n == 1 {
		fmt.Fprintf(c.outFile, "%s++\n", c.pointerName)
		return
	}
	fmt.Fprintf(c.outFile, "%s += %d\n", c.pointerName, n)
}

func (c *Compiler) decPointer(n int) {
	c.writeIndent()
	if n == 1 {
		fmt.Fprintf(c.outFile, "%s--\n", c.pointerName)
		return
	}
	fmt.Fprintf(c.outFile, "%s -= %d\n", c.pointerName, n)
}

func (c *Compiler) add(n int) {
	if n < 0 {
		c.sub(-n)
		return
	}

	c.writeIndent()
	if n == 1 {
		fmt.Fprintf(c.outFile, "%s[%s]++\n", c.memName, c.pointerName)
		return
	}
	fmt.Fprintf(c.outFile, "%s[%s] += %d\n", c.memName, c.pointerName, n)
}

func (c *Compiler) sub(n int) {
	c.writeIndent()
	if n == 1 {
		fmt.Fprintf(c.outFile, "%s[%s]--\n", c.memName, c.pointerName)
		return
	}
	fmt.Fprintf(c.outFile, "%s[%s] -= %d\n", c.memName, c.pointerName, n)
}

func (c *Compiler) print() {
	if c.printMode == charMode {
		c.printChar()
		return
	}
	c.printDecimal()
}

func (c *Compiler) printChar() {
	c.writeIndent()
	fmt.Fprintf(c.outFile, "fmt.Printf(\"%%c\", %s[%s])\n", c.memName, c.pointerName)
}

func (c *Compiler) printDecimal() {
	c.writeIndent()
	fmt.Fprintf(c.outFile, "fmt.Printf(\"%%d \", %s[%s])\n", c.memName, c.pointerName)
}

func (c *Compiler) getChar() {
	c.writeIndent()
	fmt.Fprintf(c.outFile, "%s[%s] = getchar(stdin)\n", c.memName, c.pointerName)
}

func (c *Compiler) whileLoop() {
	c.writeIndent()
	fmt.Fprintf(c.outFile, "for %s[%s] != 0 {\n", c.memName, c.pointerName)
	c.indent++
}

func (c *Compiler) closeBrace() {
	c.indent--
	c.writeIndent()
	fmt.Fprint(c.outFile, "}\n")
}

func (c *Compiler) getIndent() string {
	return strings.Repeat("    ", c.indent)
}

func (c *Compiler) writeIndent() {
	fmt.Fprint(c.outFile, c.getIndent())
}

func (c *Compiler) group(ch byte, i int) (int, int) {
	signs := map[byte]int{
		'+': 1,
		'-': -1,
	}

	shifts := map[byte]int{
		'>': 1,
		'<': -1,
	}

	_map := signs
	if ch == '>' || ch == '<' {
		_map = shifts
	}

	res := 0
	for ; i < len(c.program); i++ {
		if n, ok := _map[c.program[i]]; ok {
			res += n
		} else {
			break
		}
	}
	return res, i
}

func (c *Compiler) compile(outFile *os.File) {
	c.outFile = outFile

	c.prologue()
	for i := 0; i < len(c.program); {
		switch c.program[i] {
		case '>', '<':
			n := 0
			for ; i < len(c.program); i++ {
				ch := c.program[i]
				if ch == '>' {
					n++
				} else if ch == '<' {
					n--
				} else if unicode.IsSpace(rune(ch)) {
					continue
				} else {
					break
				}
			}
			c.incPointer(n)
			continue // avoid i++

		case '+', '-':
			n := 0
			for ; i < len(c.program); i++ {
				ch := c.program[i]
				if ch == '+' {
					n++
				} else if ch == '-' {
					n--
				} else if unicode.IsSpace(rune(ch)) {
					continue
				} else {
					break
				}
			}
			c.add(n)
			continue // avoid i++

		case '.':
			c.print()

		case ',':
			c.getChar()

		case '[':
			c.whileLoop()

		case ']':
			c.closeBrace()
		default:
		}
		i++
	}
	c.closeBrace() // end of program
}

func main() {
	C := NewDefaultCompiler()
	C.loadProgram("program.bf")

	out, err := os.Create("out.go")
	if err != nil {
		panic(err)
	}
	C.compile(out)
}
