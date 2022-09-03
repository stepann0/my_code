Calculator based on precendence climbing parsing algorithm.
Supported operators: `+, -, *, /, ^, √, unary + and -`<br>
Supported functions: `sqrt, factorial, ln, abs, exp, sin, cos, tan, ctg, deg, rad`.<br>
Supported constants: `pi, e, tau`.<br>

Grammar:
```
E --> Exp(0) 
Exp(p) --> P {B Exp(q)} 
P --> U Exp(q) | "(" E ")" | v
B --> "+" | "-" | "*" |"/" | "^"
U --> "+" | "-" | "√" | F
F --> "sqrt" | "sin" | "ln" | "abs"... etc.
```
Example:
```python
L = Lexer()
P = Parser()
expr = "((50_435 + 14_001.5)+ exp(abs(-3)- 1)) * (pi+2pi-3e)*10^-4"
ast = P.parse(L.lex(expr))
res = P.eval(ast) # = 8.183938755291484
```
```
python calc.py "(abs(-60) + 4^2)^.5 + sqrt(16/4) + sin(2pi) - ln e"
9.717797887081348

python calc.py "√121+ deg(tau)- ln(e)"
370.0
```

Reference to the original article: https://www.engr.mun.ca/~theo/Misc/exp_parsing.htm