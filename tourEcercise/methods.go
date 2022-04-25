package main

import (
	"fmt"
	"math"
)

type number struct {
	x float64
	y float64
}

func (v number) Abs() float64 {
	return math.Sqrt(math.Pow(v.x, 2) + math.Pow(v.y, 2))
}

func Abs(v number) float64 {
	return math.Sqrt(math.Pow(v.x, 2) + math.Pow(v.y, 2))
}

func main() {
	v := number{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(Abs(v))
}
