import re

class EquationSolver:
    """Simple linear equations solver."""

    # Convert term to int
    def to_int(self, s: str) -> int:
        s = s.replace("x", "")
        if s == "-": # s was '-x'
            return -1
        if s == "+" or s == '': # s was '+x' or 'x'
            return 1
        return int(s)

    # Calculate all xs and coefficients 
    def parse(self, expr):
        x = 0
        coef = 0
        regexp = r"[-+]?\d*x?"
        matches = re.finditer(regexp, expr)
        for m in matches:
            s = m.group(0)
            # Escape empty matches
            if len(s) == 0:
                continue
            if "x" in s:
                x += self.to_int(s)
            else:
                coef += self.to_int(s)
        return x, coef

    # Solve equation
    def solve_equation(self, equation: str) -> str:
        left, right = equation.split("=")

        left_x, left_coef = self.parse(left)
        right_x, right_coef = self.parse(right)

        total_x = left_x - right_x
        total_coef = right_coef - left_coef

        if total_coef == 0 and total_x == 0:
            return "Infinite solutions"
        if total_x == 0 :
            return "No solution"
        return f"x={total_coef//total_x}"


print(EquationSolver().solve_equation("+34562x+2345-3330x=612331+x-2234")) # x=19