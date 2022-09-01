Calculator based on precendence climbing parsing algorithm.
Supported operators: `+, -, *, /, ^, unary + and -`
Supported functions: `sqrt, factorial, ln, abs, exp, sin, cos, tan, ctg, deg, rad`.
Supported constants: `pi, e, tau`.

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
expr = "((50_435 + 14_001.5)+ exp(abs(-3)- 1)) * (pi+2pi-3e)*10^-4"
ast = Parser().parse(Lexer().lex(expr))
res = P.eval(ast) # = 128887,778112198
```
```
python calc.py "(abs(-60) + 4^2)^.5 + sqrt(16/4) + sin(2pi) - ln e"

9.717797887081348
```

Reference to original article: https://www.engr.mun.ca/~theo/Misc/exp_parsing.htm