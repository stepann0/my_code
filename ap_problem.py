# from itertools import combinations


# comb = combinations([-1, -1, -1, -4], 3)
# mults = []
# for c in comb:
#     print(c)
#     mult = 1
#     for i in c: mult *= i
#     print(mult)
#     mults.append(mult)
# print(max(mults))

#
# # ищем 3 max+
# # если нет 3 max+, то 2 min- и 1 max+
# # если нет 1 max+, то ищем 3 min-
#
#
#

nums = list(range(50000))

def first_missing_int(nums):
    min_int = 1
    while True:
        if min_int in nums:
            min_int += 1
        else: return min_int
print(first_missing_int(nums))
