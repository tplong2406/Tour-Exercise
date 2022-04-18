package main

import (
	"fmt"
)

func UCLN(a int, b int) int {
	kQ := 0
	if a%b == 0 {
		kQ = b
	} else if b%a == 0 {
		kQ = a
	} else {
		for i := 1; i < a/2; i++ {
			if a%i == 0 && b%i == 0 {
				kQ = i
			}
		}
	}
	return kQ
}

func main() {
	var a int
	var b int
	fmt.Print("Nhap a, b: ")
	fmt.Scan(&a, &b)
	fmt.Println("Hai so ban nhap:", a, b)
	fmt.Println("Uoc chung lon nhat cua 2 so la:", UCLN(a, b))
}
