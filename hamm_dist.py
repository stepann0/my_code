from operator import xor


def HammingDist(a: int, b: int)-> int:
    dist = 0
    while True:
        # I don't know why, but this is faster than 
        # while a != 0 or b != 0: ...
        if a == 0 and b == 0: break
        if (a^b) % 2 != 0:
            dist+=1
        a >>= 1
        b >>= 1
    return dist

# Wikipedia
def hamm_dist(a, b):
    dist = 0
    val = a^b
    while val > 0:
        val &= val-1
        dist += 1
    return dist

print(hamm_dist(1, 4))
print(hamm_dist(4, 4))
print(hamm_dist(5, 3))