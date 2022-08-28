package tokens

import (
	"fmt"
	"strings"
)

type Type int

const (
	EOF Type = iota
	Error
	LeftParen
	RightParen
	OperatorType
	Number
	Function
)

func (t Type) String() string {
	switch t {
	case EOF:
		return "EOF"
	case Error:
		return "err"
	case LeftParen:
		return "("
	case RightParen:
		return ")"
	case OperatorType:
		return "op"
	case Number:
		return "num"
	}
	return ""
}

type Token struct {
	Type Type
	Val  any
	Text string
}

func NewToken(t Type, val any, text string) Token {
	return Token{
		Type: t,
		Val:  val,
		Text: text,
	}
}

func (i Token) String() string {
	switch {
	case i.Type == EOF:
		return "EOF"
	case i.Type == Error:
		return "error: " + i.Text
	case len(i.Text) > 10:
		return fmt.Sprintf("%s: %.10q...", i.Type, i.Text)
	}
	return fmt.Sprintf("%s: %q;", i.Type, i.Text)
}

func Lex(expr string) []Token {
	tokens := []Token{}
	it := NewIterator(expr)
	tok := it.Next()
	for tok != "" {
		if tok == " " || tok == "\t" {
			tok = it.Next()
			continue
		}
		if strings.Contains("0123456789.", tok) {
			n, err := it.ReadNum()
			if err != nil {
				panic(err)
			}
			tokens = append(tokens, NewToken(Number, n, fmt.Sprintf("%v", n)))
		} else if strings.Contains("+-*/%^!", tok) {
			tokens = append(tokens, NewToken(OperatorType, Operators[tok], tok))
		} else if tok == "(" {
			tokens = append(tokens, NewToken(LeftParen, nil, tok))
		} else if tok == ")" {
			tokens = append(tokens, NewToken(RightParen, nil, tok))
		} else {
			panic("Unknown symbol.")
		}
		tok = it.Next()
	}
	return tokens
}
