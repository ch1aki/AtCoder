h = int(input())

r, t = 0, 0

while h > 1:
        h /= 2
        r += t+2
        t += 2

r+=1
print(r)
