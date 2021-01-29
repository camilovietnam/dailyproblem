# Given integers M and N, write a program that counts how
# many positive integer pairs (a, b) satisfy the following conditions:
#
# a + b = M
# a XOR b = N

import sys
import math


def integers(M, N):
    for i in range(1, math.ceil(M/2) + 1):
        if (i ^ M-i) == N:
            print("Found: (" + str(i) + ", " + str(M - i) + ")")


M = int(sys.argv[1])
N = int(sys.argv[2])

integers(M, N)
