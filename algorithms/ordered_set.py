from typing import Callable, Iterable
from collections import OrderedDict

def check_pair(t: tuple) -> None:
    if len(t) != 2:
        raise Exception(f"{t}: Only binary sets allowed.")


class Set():
    """Set partially implements ordered set. Only what is required in PairSet class."""

    def __init__(self, I: Iterable=[], check: Callable=None) -> None:
        # (i, None) is because OrderedDict 
        # must be in form like [(key1, val1), (key1, val1)...]
        if check != None:
            self._set = OrderedDict([(i, None) for i in I if check(i) == None])
        else:
            self._set = OrderedDict([(i, None) for i in I])
            
    def add(self, el) -> None:
        self._set[el] = None

    def delete(self, el) -> None:
        try:
            del self._set[el]
        except KeyError:
            return
            
    def __str__(self) -> str:
        k = ', '.join([str(i) for i in self._set.keys()])
        return f"{{ {k} }}"


class PairSet(Set):
    """PairSet solves a few small discrete math tasks from university."""

    def __init__(self, I: Iterable = []) -> None:
        super().__init__(I, check=check_pair)

    def comp(self, s):
        if isinstance(s, PairSet):
            return self.comp(s._set)

        res = PairSet()
        for i in self._set:
            for j in s:
                if i[1] == j[0]:
                    res.add((i[0], j[1]))
        return res
        
    def inverse(self):
        return PairSet([(i[1], i[0]) for i in self._set])

    def proj1(self) -> list:
        return Set([i[0] for i in self._set])
    
    def proj2(self) -> list:
        return Set([i[1] for i in self._set])



a, b, c, d, e, f = "a", "b", "c", "d", "e", "f"

# Задание с дискретной математики
P = PairSet([(a, c), (c, b), (b, b), (b, c)])
T = PairSet([(a, b), (c, a), (c, c), (b, a), (b, c), (b, b)])
print(P.inverse())
print(P.inverse().comp(P))
print(P.comp(T))

A = Set([a, c, c, d, a, a, d, d, d])
print(A)

