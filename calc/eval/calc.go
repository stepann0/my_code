package main

import (
	"fmt"
	"math"
	"strconv"

	. "stepa.basic.calc/tokens"
	. "stepa.basic.calc/utils"
)

func ShuntingYard(lex []Token) []Token {
	op_stack := Stack[Token]{}
	queue := Queue[Token]{}

	for _, tok := range lex {
		if tok.Type == OperatorType {
			o1 := tok.Text
			if !op_stack.Empty() {
				for i := len(op_stack) - 1; i >= 0 && op_stack[i].Type != LeftParen; i-- {
					o2 := op_stack[i].Text
					prec1, prec2 := Operators[o1].Prec, Operators[o2].Prec
					assoc1 := Operators[o1].Assoc
					if prec2 > prec1 || prec2 == prec1 && assoc1 == 'l' {
						queue.Push(op_stack.Pop())
					}
				}
			}
			op_stack.Push(tok)

		} else if tok.Type == Number || tok.Type == LeftParen {
			queue.Push(tok)
		} else if tok.Type == RightParen {
			if !op_stack.Empty() {
				for i := len(op_stack) - 1; i >= 0 && op_stack[i].Type != LeftParen; i-- {
					queue.Push(op_stack.Pop())
				}
			}
			rmLeftPar(&op_stack)
		}
	}

	for !op_stack.Empty() {
		queue.Push(op_stack.Pop())
	}
	fmt.Println("queue: ", queue)
	return queue
}

func rmLeftPar(s *Stack[Token]) {
	for !s.Empty() && s.Top().Type == LeftParen {
		s.Pop()
	}
}

func binary(op Token) bool {
	switch op.Text {
	case "+", "-", "*", "/", "%", "^":
		return true
	case "!":
		return false
	}
	panic("Unkown operator.")
}

func evalBinary(operator Token, o1, o2 float64) float64 {
	switch operator.Text {
	case "+":
		return o2 + o1
	case "-":
		return o2 - o1
	case "*":
		return o2 * o1
	case "/":
		return o2 / o1
	case "^":
		return math.Pow(o2, o1)
	}
	panic("Unkown operator.")
}

func EvaluateRPN(rpn []Token) float64 {
	stack := Stack[float64]{}

	for _, tok := range rpn {
		if tok.Type == Number {
			n, _ := strconv.ParseFloat(tok.Text, 64)
			stack.Push(n)
		} else if !stack.Empty() && tok.Type == OperatorType {
			if binary(tok) {
				operand1, operand2 := stack.Pop(), stack.Pop()
				stack.Push(evalBinary(tok, operand1, operand2))
			}
		}
	}
	return stack.Top()
}

// 2 4 + 4 1 - / 5 2 * * 1 0 - -
func main() {
	e := EvaluateRPN(ShuntingYard(Lex("(2+4)/(4-1)*(5*2)-(1-0)")))
	fmt.Println(e)
}
