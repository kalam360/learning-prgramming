package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	for  n := 0; n < 10; n++ {
		z -= (z*z - x)/(2*z)
		fmt.Println(z)
	}
	return z
}

func test_loops_functions() {
	fmt.Println(Sqrt(2))
}