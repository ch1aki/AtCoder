import sys

sys.setrecursionlimit(100000)

def fnc(n):
    if n < 2:
        return 1
    return n * fnc(n-2)

def cnt(s):
    r = 0
    for c in s.reversed():
        if c != "0":
            break
        r+=1
    return r


n = int(input())
r = fnc(n)
print(cnt(r))
