package main

import "fmt"

func main() {
	a := []int32{}
	a = append(a, 1, 2, 3, 4, 6, 7, 8)
	fmt.Println(a)
	// a = append(a[:5], append([]int32{5}, a[3:]...)...)
	a = append(a, 0)
	copy(a[5:], a[4:])
	a[4] = 5
	fmt.Println(a)
	// a = append(a[:4], append([]int32{9, 8, 7}, a[4:]...)...)
	// fmt.Println(a)

}
