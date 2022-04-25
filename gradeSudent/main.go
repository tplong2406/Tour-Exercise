package main

import (
	"fmt"
)

/*
 * Complete the 'gradingStudents' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY grades as parameter.
 */

func gradingStudents(grades []int32) []int32 {
	// Write your code here
	rS := grades
	for i := range grades {

		// switch rS[i] {
		// case 98, 99:
		// 	rS[i] = 100
		// case 93, 94:
		// 	rS[i] = 95
		// case 88, 89:
		// 	rS[i] = 90
		// case 83, 84:
		// 	rS[i] = 85
		// case 78, 79:
		// 	rS[i] = 80
		// case 73, 74:
		// 	rS[i] = 75
		// case 68, 69:
		// 	rS[i] = 70
		// case 63, 64:
		// 	rS[i] = 65
		// case 58, 59:
		// 	rS[i] = 60
		// case 53, 54:
		// 	rS[i] = 55
		// case 48, 49:
		// 	rS[i] = 50
		// case 43, 44:
		// 	rS[i] = 45
		// case 38, 39:
		// 	rS[i] = 40

		if rS[i]%5 == 3 {
			rS[i] += 2
		}
		if rS[i]%5 == 4 {
			rS[i]++
		}
	}
	return rS
}

func main() {
	arr := []int32{73, 67, 38, 33}
	fmt.Println(arr)
	fmt.Println(gradingStudents(arr))
}
