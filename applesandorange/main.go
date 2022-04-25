package main

import "fmt"

/*
 * Complete the 'countApplesAndOranges' function below.
 *
 * The function accepts following parameters:
 *  1. INTEGER s
 *  2. INTEGER t
 *  3. INTEGER a
 *  4. INTEGER b
 *  5. INTEGER_ARRAY apples
 *  6. INTEGER_ARRAY oranges
 */

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	// Write your code here
	countApples := 0
	countOranges := 0
	for i := range apples {
		// apples[i]+=a
		if apples[i]+a >= s && apples[i]+a <= t {
			countApples++
		}
	}
	for j := range oranges {
		// oranges[j] += b
		if oranges[j]+b >= s && oranges[j]+b <= t {
			countOranges++
		}
	}
	fmt.Println(countApples)
	fmt.Println(countOranges)
}

func main() {
	var s, t, a, b int32 = 7, 11, 5, 15
	apples := []int32{-2, 2, 1}
	oranges := []int32{5, -6}
	countApplesAndOranges(s, t, a, b, apples, oranges)
}
