package main

import (
	"fmt"
	"sort"
)

/*
 * Complete the 'diagonalDifference' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) {
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	var max float32 = 0
	var min float32 = 0
	for i := 0; i < len(arr)-1; i++ {
		fmt.Printf("a[%v] = %.0d\n", i, arr[i])
		min += float32(arr[i])
		fmt.Printf("min = %.0f\n", min)
	}
	fmt.Println()
	for j := 1; j < len(arr); j++ {
		max += float32(arr[j])
		fmt.Printf("max = %.0f\n", max)
	}

}

func main() {
	a := []int32{140638725, 436257910, 953274816, 734065819, 362748590}
	plusMinus(a)
}
