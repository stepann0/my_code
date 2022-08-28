package tokens

import "fmt"

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
	Val  string
}

func NewToken(t Type, val string) Token {
	return Token{
		Type: t,
		Val:  val,
	}
}

func (i Token) String() string {
	switch {
	case i.Type == EOF:
		return "EOF"
	case i.Type == Error:
		return "error: " + i.Val
	case len(i.Val) > 10:
		return fmt.Sprintf("%s: %.10q...", i.Type, i.Val)
	}
	return fmt.Sprintf("%s: %q;", i.Type, i.Val)
}
