package utils

import (
	"stepa.basic.calc/tokens"
)

func RmLeftPar(s *Stack[tokens.Token]) {
	for !s.Empty() && s.Top().Type == tokens.LeftParen {
		s.Pop()
	}
}
