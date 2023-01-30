package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

type Interpreter struct {
	memory    [30000]byte
	pointer   int
	program   []byte
	jumpTable map[int]int // for loops
}

func NewInterpreter(prog string) Interpreter {
	I := Interpreter{
		jumpTable: map[int]int{},
	}
	I.loadProgram(prog)
	I.buildJumpTable()
	return I
}

func (I *Interpreter) currCell() byte {
	return I.memory[I.pointer]
}

func (I *Interpreter) incPointer() {
	I.pointer++
}

func (I *Interpreter) decPointer() {
	I.pointer--
}

func (I *Interpreter) add() {
	if I.pointer < 0 || I.pointer >= len(I.memory) {
		panic("you ran out of program's memory")
	}
	I.memory[I.pointer]++
}

func (I *Interpreter) sub() {
	if I.pointer < 0 || I.pointer >= len(I.memory) {
		panic("you ran out of program's memory")
	}
	I.memory[I.pointer]--
}

func (I *Interpreter) Execute() {
	stdin_reader := bufio.NewReader(os.Stdin)

	for i := 0; i < len(I.program); {
		switch I.program[i] {
		case '>':
			I.incPointer()
		case '<':
			I.decPointer()
		case '+':
			I.add()
		case '-':
			I.sub()
		case '.':
			fmt.Printf("%c", I.currCell())
		case ',':
			I.memory[I.pointer] = getChar(stdin_reader)
		case '[':
			if I.currCell() == 0 {
				i = I.jumpTable[i]
			}
		case ']':
			if I.currCell() != 0 {
				i = I.jumpTable[i]
			}
		default:
		}
		i++
	}
}

func (I *Interpreter) loadProgram(name string) {
	f, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}
	I.program = f
}

func (I *Interpreter) buildJumpTable() {
	stack := []int{}
	for i := 0; i < len(I.program); i++ {
		switch I.program[i] {
		case '[':
			stack = append(stack, i)
		case ']':
			if len(stack) == 0 {
				panic("loop brackets mismatching")
			}
			open_br := stack[len(stack)-1]
			I.jumpTable[open_br] = i
			I.jumpTable[i] = open_br
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) != 0 {
		panic("loop brackets mismatching")
	}
}

func getChar(reader *bufio.Reader) byte {
	b, err := reader.ReadByte()
	if err != nil {
		panic(err)
	}
	return b
}

func main() {
	I := NewInterpreter("program.bf")
	I.Execute()
}
