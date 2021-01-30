# Good morning! Here's your coding interview problem for today.
#
# This problem was asked by LinkedIn.
#
# You are given a string consisting of the letters x and y, such as xyxxxyxyy.
# In addition, you have an operation called flip, which changes a single x to
# y or vice versa.
#
# Determine how many times you would need to apply this operation to ensure
# that all x's come before all y's. In the preceding example, it suffices to
# flip the second and sixth characters, so you should return 2.

def flippity (string):
    string = list(string)
    flips = 0

    for i in range(len(string)):
        if string[i] == 'y':
            start = i
            for j in range(start, len(string)):
                if string[j] == 'x':
                    string[start] = 'x'
                    string[j] = 'y'
                    start += 1
                    flips += 1
    print(f"Required {flips} flips")
    return string


result = flippity('xxyxyx')
print(f"The returned string is {''.join(result)}")
