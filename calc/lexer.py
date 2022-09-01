EOF          = "\u0003"
NUMBERS      = "0123456789"
LETTERS      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
WHITESPACE   = list(" \t")
OPERATORS = list("+-*/^")
CONSTS = ["pi", "e", "tau"]


class Lexer:
    def __init__(self, expr: str="") -> None:
        self.expr = iter(expr)
        self.curr_char = next(self.expr, EOF)
        self.tokens_arr = []

    def consume(self) -> None:
        self.curr_char = next(self.expr, EOF)

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

    def lex(self, expr=""):
        self.__init__(expr)
        while self.curr_char != EOF:
            # ignore whitespace chars
            if self.curr_char in WHITESPACE:
                self.consume()
                
            # parse numbers
            elif self.curr_char in NUMBERS+".":
                self.tokens_arr.append(self.read_num())
                
            # process math operators
            elif self.curr_char in OPERATORS:
                tok = self.curr_char
                self.consume()
                self.tokens_arr.append(tok)
            
            elif self.curr_char in "()":
                tok = self.curr_char
                self.consume()
                self.tokens_arr.append(tok)
            
            elif self.curr_char in LETTERS:
                name = self.read_func()
                if name in CONSTS\
                    and len(self.tokens_arr)>0\
                    and isinstance(self.tokens_arr[-1], (int, float)):

                    self.tokens_arr.append("*")
                self.tokens_arr.append(name)
            else:
                self.error(f"Unknown symbol \"{self.curr_char}\".")
        return self.tokens_arr

    def error(self, msg: str) -> None:
        print(f"\n\033[38;05;203m{msg}\033[0m")
        exit(1)