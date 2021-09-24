package main

import (
	"fmt"
	"math"
)

func squareRoot(x float64) string {
	if x < 0 {
		return squareRoot(-x) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v:= math.Pow(x,n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func test_sqrt() {
	fmt.Println(squareRoot(2), squareRoot(-4))
	fmt.Println(pow(3,2,10))
	fmt.Println(pow(3,3,20))
}
