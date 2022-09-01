EOF          = "\u0003"
NUMBERS      = "0123456789"
LETTERS      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
WHITESPACE   = list(" \t")
BI_OPERATORS = list("+-*/^") # operators that can be binary
UN_OPERATORS = list("-+") # operators that can be unary

class Lexer:
    def __init__(self, expr: str="") -> None:
        self.expr = iter(expr)
        self.curr_char = next(self.expr, None)
        self.tokens_arr = []

    def consume(self) -> None:
        self.curr_char = next(self.expr, None)

    def read_num(self) -> float:
        res = ""
        while self.curr_char and self.curr_char in NUMBERS+"._":
            res += self.curr_char
            self.consume()
        try:
            return float(res)
        except ValueError:
            self.error(f"Can't convert {res} to float.")

    def read_func(self) -> str:
        fn = ""
        while self.curr_char and self.curr_char in LETTERS+NUMBERS+"_":
            fn += self.curr_char
            self.consume()
        return fn

    def read_token(self):
        while self.curr_char:
            # ignore whitespace chars
            if self.curr_char in WHITESPACE:
                self.consume()
                
            # parse numbers
            elif self.curr_char in NUMBERS+".":
                yield self.read_num()
                
            # process math operators
            elif self.curr_char in BI_OPERATORS+UN_OPERATORS:
                tok = self.curr_char
                self.consume()
                yield tok
            
            elif self.curr_char in "()":
                tok = self.curr_char
                self.consume()
                yield tok
            
            elif self.curr_char in LETTERS:
                yield self.read_func()
            else:
                raise SyntaxError(self.expr)

    def lex(self, expr: str):
        self.__init__(expr)
        tok = next(self.read_token(), None)
        while tok != None:
            self.tokens_arr.append(tok)
            tok = next(self.read_token(), None)
        return self.tokens_arr