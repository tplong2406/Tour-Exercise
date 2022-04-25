package main

import (
	"fmt"
)

/*
 * Complete the 'kangaroo' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. INTEGER x1
 *  2. INTEGER v1
 *  3. INTEGER x2
 *  4. INTEGER v2
 */

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	// Write your code here
	rs := ""

	for {
		x1 += v1
		x2 += v2
		if x1 == x2 {
			rs = "YES"
			break
		}
		if x1 > x2 {
			rs = "NO"
			break
		}
	}
	return rs
}

func main() {
	fmt.Println("Result: ", kangaroo(0, 2, 5, 3))
}
