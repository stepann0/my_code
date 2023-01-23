import unittest
from parser import Parser
from lexer import Lexer


class TestCalc(unittest.TestCase):
    L = Lexer()
    P = Parser()

    def calc(self, expr):
        return self.P.eval(self.P.parse(self.L.lex(expr)))

    def test_calc(self):
        self.assertAlmostEqual(self.calc("((50_435+14_001.5)+exp(abs(-3)-1))*(-1+2+3)^.5"), 128887.7781, places=4)
        self.assertAlmostEqual(self.calc("+55^ (-1.01)+ (-201)"), -200.982532382, places=4)
        self.assertAlmostEqual(self.calc("(2^3^.5)^3"), 36.660445757, places=4)
        self.assertAlmostEqual(self.calc("√(√(√(√(1+2*3+4*(5+6*7+8)*9))))"), 1.607449607, places=4)
        self.assertAlmostEqual(self.calc("abs(-----5)"), 5, places=4)
        self.assertAlmostEqual(self.calc("factorial (10) / factorial 9* factorial 8/ factorial 7* factorial 6"), 57600, places=4)
        self.assertAlmostEqual(self.calc("2^3^2^-(ln(e))"), 3.321997085, places=4)
        self.assertAlmostEqual(self.calc("pi^(1/2pi)^(1/4pi)^(1/8pi)^(1/16pi)^2e"), 13.971709543, places=4)

if __name__ == '__main__':
    unittest.main()