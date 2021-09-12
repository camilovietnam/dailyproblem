// *********************************************************
// This problem was asked by Square.
//
// Assume you have access to a function toss_biased() which
// returns 0 or 1 with a probability that's not 50-50 (but
// also not 0-100 or 100-0). You do not know the bias of the
// coin.
//
// Write a function to simulate an unbiased coin toss.
// *********************************************************

// Solution: if we call the possibility of getting a 0 with
// the biased coin "p0", and the possibility of getting a 1
// "p1", the possible combinations when we throw the coin
// twice could be:
//
// 0 followed by 0 = p0 * p0
// 0 followed by 1 = p0 * p1
// 1 followed by 0 = p1 * p0
// 1 followed by 1 = p1 * p1
//
// We can see that the possibilities for 0 followed by 1 and
// 1 followed by 0 are the same (p0*p1=p1*p0). Therefore, if
// we ignore the other two, we will have a system with a 50/50 bias

package main

import (
    "math/rand"
    "fmt"
    "time"
    )

func main () {
    // First we toss one thousand biased coins
    // Because the bias we have is 2:8, we should get around 200 of one result,
    // and 800 of the other result.
    oneThousandBiasedTosses()

    // Second, we toss one thousand unbiased coins
    // We should get around 500 of each result
    oneThousandUnbiasedTosses()
}

// This is my simulation of a biased toss, using an array of integers that contains
// more of one value than the other. The ratio is 2:8 ones to zeros
func toss_biased() int {
    source := rand.NewSource(time.Now().UnixNano())
    rand1 := rand.New(source)

    var biases = []int {1,1,1,1,1,1,1,1,0,0}

    var i = rand1.Intn(10)
    return biases[i]
}

// This is the unbiased function
func toss_unbiased() int {
    var firstToss, secondToss int

    for firstToss == secondToss {
        firstToss = toss_biased()
        secondToss = toss_biased()
    }

    return secondToss
}

func oneThousandBiasedTosses() {
    fmt.Println("Result of biased tosses (should be around 2:8):")
    oneThousandTosses(toss_biased)
}

func oneThousandUnbiasedTosses() {
    fmt.Println("Result of unbiased tosses (should be around 5:5):")
    oneThousandTosses(toss_unbiased)
}

func oneThousandTosses(callback func() int) {
    var tossOne, tossZero int

    for i := 0; i < 1000; i++ {
        if callback() == 0 {
            tossZero++
        } else {
            tossOne++
        }
    }

    fmt.Println("Number of zeros: ", tossZero)
    fmt.Println("Number of ones: ", tossOne)
}

func abs(i int) int {
    if i < 0 {
        return -1*i
    }

    return i
}