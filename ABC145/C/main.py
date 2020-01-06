#!/usr/bin/env python3

import math
import itertools

def distance(p1, p2):
    return math.sqrt((p1[0]-p2[0])**2 + (p1[1]-p2[1])**2)

n = int(input())
points = {i:map(int, input().split()) for i in range(n)}

total_distance = 0

for cource in itertools.permutations(range(n)):
    for i in range(1, n):
        total_distance += distance(points[i-1], points[i])

print(points)
print(total_distance/math.factorial(n))