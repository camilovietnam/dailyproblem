# Soundex is an algorithm used to categorize phonetically, such that two names that sound alike but are spelled differently have the same representation.

# Soundex maps every name to a string consisting of one letter and three numbers, like M460.

# One version of the algorithm is as follows:

# Remove consecutive consonants with the same sound (for example, change ck -> c).
# Keep the first letter. The remaining steps only apply to the rest of the string.
# Remove all vowels, including y, w, and h.
# Replace all consonants with the following digits:
# b, f, p, v → 1
# c, g, j, k, q, s, x, z → 2
# d, t → 3
# l → 4
# m, n → 5
# r → 6
# If you don't have three numbers yet, append zeros until you do. Keep the first three numbers.
# Using this scheme, Jackson and Jaxen both map to J250.

# Implement Soundex.

import sys
import re

def soundex(word):
	# First, let's remove consecutive consonantes with same sound
	word = re.sub("([ckqs])[cskq]*", "\\1", word)
	word = re.sub("[aeiuoywh]", "", word)

	firstLetter = word[0]

	# Second, let's replace each letter for the number assigned
	word = re.sub("[bfpv]", "1", word[1:])
	word = re.sub("[cgjkqsxz]", "2", word)
	word = re.sub("[dt]", "3", word)
	word = re.sub("[l]", "4", word)
	word = re.sub("[mn]", "5", word)
	word = re.sub("[r]", "6", word)

	# Fill with zeros until we have 3 numbers
	while len(word) < 3:
		word += '0'

	# return the first letter saved and the rest of the string
	return firstLetter + word[0:3]


try:
	soundexed = soundex(sys.argv[1])
	print("The soundex of", sys.argv[1], "is", soundexed)
except IndexError:
	print("Please provide a string as the first argument")
	print("Example: python3 soundex.py MyString")
