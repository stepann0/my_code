package main

import (
	"fmt"
	"math"
	"strconv"

	. "stepa.basic.calc/tokens"
	"stepa.basic.calc/utils"
)

func ShuntingYard(lex []Token) []Token {
	op_stack := utils.Stack[Token]{}
	queue := utils.Queue[Token]{}

	for _, tok := range lex {
		if tok.Type == Number {
			queue.Push(tok)
		} else if tok.Type == OperatorType {
			o1 := tok.Val
			if !op_stack.Empty() {
				for i := len(op_stack) - 1; i >= 0 && op_stack[i].Type != LeftParen; i-- {
					o2 := op_stack[i].Val
					prec1, prec2 := Operators[o1].Prec, Operators[o2].Prec
					assoc1 := Operators[o1].Assoc
					if prec2 > prec1 || prec2 == prec1 && assoc1 == 'l' {
						queue.Push(op_stack.Pop())
					}
				}
			}
			op_stack.Push(tok)
		} else if tok.Type == LeftParen {
			op_stack.Push(tok)
		} else if tok.Type == RightParen {
			if !op_stack.Empty() {
				for i := len(op_stack) - 1; i >= 0 && op_stack[i].Type != LeftParen; i-- {
					queue.Push(op_stack.Pop())
				}
			}
			utils.RmLeftPar(&op_stack)
		}
	}

	for !op_stack.Empty() {
		queue.Push(op_stack.Pop())
	}
	return queue
}

func EvaluateRPN(rpn []Token) float64 {
	stack := utils.Stack[float64]{}

	for _, tok := range rpn {
		if tok.Type == Number {
			n, _ := strconv.ParseFloat(tok.Val, 64)
			stack.Push(n)
		} else if !stack.Empty() && tok.Type == OperatorType {
			// if binary operator
			if Operators[tok.Val].Operands == 2 {
				operand1, operand2 := stack.Pop(), stack.Pop()
				stack.Push(EvalBinary(tok, operand1, operand2))
			}
		}
	}
	return stack.Top()
}

func EvalBinary(operator Token, o1, o2 float64) float64 {
	switch operator.Val {
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

func main() {
	expr := "3.82^(10.5-.5*5) / (10^2*3 - 100)" // = 226,713252563
	res := EvaluateRPN(ShuntingYard(Lex(expr)))
	fmt.Println(res)
}
