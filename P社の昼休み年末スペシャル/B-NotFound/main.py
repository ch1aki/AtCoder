s = list(input())

ord('a')
alphabet = {}

for c in s:
    if ord(c) == start:
        continue
    elif ord(c) - start == 1:
        start+=1
        continue
    else:
        print(chr(start+1))
        exit()
print("None")