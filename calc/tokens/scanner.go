package tokens

import (
	"fmt"
	"strings"
)

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
			tokens = append(tokens, NewToken(Number, fmt.Sprintf("%v", n)))
		} else if strings.Contains("+-*/%^!", tok) {
			tokens = append(tokens, NewToken(OperatorType, tok))
		} else if tok == "(" {
			tokens = append(tokens, NewToken(LeftParen, tok))
		} else if tok == ")" {
			tokens = append(tokens, NewToken(RightParen, tok))
		} else {
			panic("Unknown symbol.")
		}
		tok = it.Next()
	}
	return tokens
}
