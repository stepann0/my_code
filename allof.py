import random

a = 62
cond = True
array = []

pick = 300
size_ar = 300

if size_ar - pick <= 1: 
    for i in range(size_ar):
        array.append(random.randint(0, pick))
        for j in array[:i]:
            cond = cond and (j != array[i])
        while (cond == False) and (i > 0):
            array[i] = random.randint(0, pick)
            cond = True
            for k in array[:i]:
                cond = cond and (k != array[i])
else:
    print("It's incredible!")

print(array)
for i in array:
    isin = 0
    for j in array:
        if (i == j):
            isin += 1
    if isin >= 2:
        print("alg doesnt work!")
        break