package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	delta := 1e-100
	target := math.Sqrt(x)

	for i := 0; true; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("i = %d / z = %f\n", i, z)
		if math.Abs(z-target) < delta {
			break
		}
	}
	return z
}

func main() {
	z := Sqrt(2)
	fmt.Println("Sqrt", z)
	fmt.Println("math.Sqrt", math.Sqrt(2))
	fmt.Println("diff", z-math.Sqrt(2))
}
