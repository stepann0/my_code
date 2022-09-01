class Tree():
    def __init__(self, operator: str="_", *operands) -> None:
        self.operator = operator
        self.operands = operands

    def print_(self, prefix=""):
        ident = "  "
        branch = "│ "
        tee = "├─"
        last = "└─"

        res = ""
        for i, ch in enumerate(self.operands):
            pointer = last
            extension = ident
            if i < len(self.operands)-1:
                extension = branch
                pointer = tee

            if isinstance(ch, (int, float)):
                res += f"{prefix}{pointer}{ch}\n"
            elif isinstance(ch, Tree):
                res += f"{prefix}{pointer}{ch.operator}\n"
                res += ch.print_(prefix+extension)
        return res

    def __str__(self) -> str:
        return self.operator + "\n" + self.print_()