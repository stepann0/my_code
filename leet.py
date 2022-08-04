s = "la3(55(3al-oossoo-asdf"
max_len = 0
for i, l in enumerate(s):
    for j, lt in enumerate(s[i:], start=i):
        if lt == l:
            subs = s[i:j+1]
            if subs == subs[::-1] and len(subs) > max_len:
                answ = subs
                max_len = len(subs)

print(answ)
