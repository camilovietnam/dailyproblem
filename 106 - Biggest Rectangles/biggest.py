# You are given a histogram consisting of rectangles of different heights. These
# heights are represented in an input list, such that [1, 3, 2, 5] corresponds to
# the following diagram:

#       x
#       x  
#   x   x
#   x x x
# x x x x
# Determine the area of the largest rectangle that can be formed only from the 
# bars of the histogram. For the diagram above, for example, this would be six,
# representing the 2 x 3 area at the bottom right.

import sys

def toTheLeft(colId, columns):
    currentCol = colId-1
    toTheLeft = 0

    while currentCol >= 0:
        if columns[currentCol] >= columns[colId]:
            toTheLeft += 1
        currentCol -= 1

    return toTheLeft

def toTheRight(colId, columns):
    currentCol = colId + 1
    toTheRight = 0

    while currentCol < len(columns):
        if columns[currentCol] >= columns[colId]:
            toTheRight += 1
        currentCol += 1
    
    return toTheRight

def findMax(columns): 
    max = [0,0]

    for colId in range(len(columns)):
        toLeft = toTheLeft(colId, columns)
        toRight = toTheRight(colId, columns)

        sumCols = toLeft + toRight + 1
        if sumCols*columns[colId] > (max[0] * max[1]):
            max = [sumCols, columns[colId]]

    print('the max rectangle is', max[0], 'x', max[1])

findMax(eval(sys.argv[1]))
