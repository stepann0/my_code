import sys
from lexer import Lexer
from parser import Parser

def main():
    if len(sys.argv) < 2:
        return
    expr = sys.argv[1]
    L = Lexer()
    P = Parser()
    tree = P.parse(L.lex(expr))
    res = P.eval(tree)
    print(f"\n{res}")

if __name__ == "__main__":
    main()