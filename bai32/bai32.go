package main

import (
	"fmt"
	"math"
)

var canh struct {
	a float32
	b float32
	c float32
}

func kiemTraTG(a float32, b float32, c float32) bool {
	if a+b > c && a+c > b && b+c > a {
		return true
	}
	return false
}

func chuVi(a float32, b float32, c float32) float32 {
	return a + b + c
}

func dienTich(a float32, b float32, c float32) float32 {
	var p = chuVi(a, b, c) / 2
	return float32(math.Sqrt(float64(p * (p - a) * (p - b) * (p - c))))
}

func main() {
	fmt.Print("Nhap 3 canh tam giac: ")
	fmt.Scan(&canh.a, &canh.b, &canh.c)
	if kiemTraTG(canh.a, canh.b, canh.c) {

	} else {
		fmt.Println("Khong phai la tam giac!!!")
	}
}
