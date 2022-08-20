def decode_string(s: str) -> str:
    stack = []
    for i in s:
        if i != ']':
            stack.append(i)
        else:
            substr = ""
            while len(stack) > 0 and stack[-1] != '[':
                substr = stack.pop() + substr
            stack.pop()

            num = ""
            while len(stack) > 0 and stack[-1] in "0123456789":
                num = stack.pop()+num
            
            if substr == "":
                continue

            if num == "":
                stack.append(substr)
                continue

            stack.append(int(num)*substr)
    return ''.join(stack)

print(decode_string("3[z]2[2[y]pq4[2[jk]e1[f]]]ef"))