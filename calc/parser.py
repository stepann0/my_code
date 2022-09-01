import math
import sys
from collections import namedtuple
from tree import Tree

EOF = "\u0003"
NUMBERS = "0123456789"
LETTERS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
WHITESPACE = list(" \t")
BI_OPERATORS = list("+-*/^") # operators that can be binary
UN_OPERATORS = list("-+") # operators that can be unary
FUNCTIONS = [
    "sqrt", "factorial", "ln", "abs", "exp",
    "sin", "cos", "tan", "ctg", "deg", "rad",
    ] # allowed functions

Operator = namedtuple("Operator", "prec, assoc, eval")
Operators = {
    "+": Operator(3, 'l', lambda a, b: a+b),
    "-": Operator(3, 'l', lambda a, b: a-b),
    "u-": Operator(4, 'r', lambda a: -a),
    "u+": Operator(4, 'r', lambda a: a),
    "/": Operator(5, 'l', lambda a, b: a/b),
    "*": Operator(5, 'l', lambda a, b: a*b),
    "^": Operator(6, 'r', lambda a, b: a**b),
}

Function =  namedtuple("Function", "eval")
Functions = {
    "sqrt"     : Function(math.sqrt),
    "factorial": Function(lambda a: math.factorial(int(a))),
    "ln"       : Function(math.log),
    "abs"      : Function(math.fabs),
    "exp"      : Function(math.exp),
    "sin"      : Function(math.sin),
    "cos"      : Function(math.cos),
    "tan"      : Function(math.tan),
    "ctg"      : Function(lambda a: 1/math.tan(a)),
    "deg"      : Function(math.degrees),
    "rad"      : Function(math.radians),
}

class Parser:
    """
    Parser implements a precendence climbing parsing algorithm.
    It supports operators: +, -, *, /, ^, unary + and -
    and a bunch of functions that take 1 argument: sin, cos, exp, factorial, abs etc.
    They are fully equal to the unary operators.

    Example:
    expr = "((50_435 + 14_001.5) + exp(abs(-3) - 1)) * (-1 + 2 + 3)^.5"
    ast = Parser().parse(Lexer().lex(expr))
    res = P.eval(ast) # = 128887,778112198
    """
    def __init__(self, tokens_arr=[]) -> None:
        self.iter_tokens = iter(tokens_arr)
        self.curr_tok = next(self.iter_tokens, EOF)

    def consume(self) -> None:
        self.curr_tok = next(self.iter_tokens, EOF)

    def error(self, msg: str) -> None:
        print(f"\n\033[38;05;203m{msg}\033[0m")
        exit(1)

    def expect(self, tok_type) -> None:
        if self.curr_tok == tok_type:
            self.consume()
        else:
            self.error(f"Wrong token: \"{self.curr_tok}\".")

    def unary(self, token):
        if token in UN_OPERATORS: return "u"+token
        self.error(f"Can't convert \"{token}\" to an unary operator.")

    def binary(self, token):
        if token in BI_OPERATORS: return token
        self.error(f"Can't convert \"{token}\" to an binary operator.")

    def make_leaf(self, operand):
        return operand

    def make_node(self, op, *operands):
        return Tree(op, *operands)

    def parse(self, tokens_arr) -> Tree:
        self.__init__(tokens_arr)
        t = self.Exp(0)
        self.expect(EOF)
        return t

    def Exp(self, p: int) -> Tree:
        t = self.P()
        while self.curr_tok in BI_OPERATORS and Operators[self.binary(self.curr_tok)].prec >= p:
            op = self.binary(self.curr_tok)
            self.consume()
            q = Operators[op].prec if Operators[op].assoc == 'r' else Operators[op].prec+1
            t1 = self.Exp(q)
            t = self.make_node(op, t, t1)
        return t

    def P(self) -> Tree|float:
        if self.curr_tok in UN_OPERATORS:
            op = self.unary(self.curr_tok)
            self.consume()
            q = Operators[op].prec
            t = self.Exp(q)
            return self.make_node(op, t)
        elif self.curr_tok in FUNCTIONS:
            fn = self.curr_tok
            self.consume()
            # this functions are fully equal to unary operators, so their precendece is 4
            t = self.Exp(4)
            return self.make_node(fn, t)
        elif self.curr_tok == "(":
            self.consume()
            t = self.Exp(0)
            self.expect(")")
            return t
        elif isinstance(self.curr_tok, (int, float)):
            n = self.curr_tok
            self.consume()
            return n
        else:
            self.error("Something is wrong.")

    def calc(self, operator, *operands) -> float:
        fn = None
        if operator in Functions:
            fn = Functions[operator].eval
        elif operator in Operators:
            fn = Operators[operator].eval
        if fn == None:
            self.error(f"Operator (or function) \"{operator}\" not implemented.")
        try: 
            return fn(*operands)
        except Exception:
            self.error(f"Error on {operator} or its args: {operands}.")


    def eval(self, tree) -> float:
        if tree == None or isinstance(tree, (int, float)):
            return tree
        operands = [self.eval(o) for o in tree.operands]
        return self.calc(tree.operator, *operands)