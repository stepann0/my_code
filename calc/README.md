Calculator based on precendence climbing parsing algorithm.
Supported operators: +, -, *, /, ^, unary + and -
Supported functions: sqrt, factorial, ln, abs, exp, sin, cos, tan, ctg, deg, rad.

Grammar:
```
E --> Exp(0) 
Exp(p) --> P {B Exp(q)} 
P --> U Exp(q) | "(" E ")" | v
B --> "+" | "-"  | "*" |"/" | "^"
U --> "+" | "-" | F
F --> "sqrt" | "sin" | "ln" | "abs"... etc.
```
Example:
```python
expr = "((50_435 + 14_001.5) + exp(abs(-3) - 1)) * (-1 + 2 + 3)^.5"
ast = Parser().parse(Lexer().lex(expr))
res = P.eval(ast) # = 128887,778112198
```
```
python calc.py "(abs -60 + 4)^.5 + sqrt(16)"

12.0
```