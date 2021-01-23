# Find the intersection of two strings
import sys

def longest_intersection(string1, string2):
    longest = ''

    for i in range(0, len(string1)):
        for j in range(i+1, len(string1)+1):
            candidate = string1[i:j]
            if (candidate in string2) and (len(candidate) > len(longest)):
                longest = candidate

    return longest


if len(sys.argv) <= 2:
	print('[!] Please provide at least two arguments')
else:
	result = longest_intersection(sys.argv[1], sys.argv[2])
	print('The longest intersection is:', result)