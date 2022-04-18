package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	temp := 0.0
	for i := 0; i < 10; i++ {
		z, temp = z-(z*z-x)/(2*z), z
		fmt.Printf("z = %f,temp =  %f\n", z, temp)
		if math.Abs(z-temp) < 1e-10 {
			break
		}

	}
	return z
}

func main() {
	fmt.Println(Sqrt(4))
	fmt.Println(math.Sqrt(4))
}
