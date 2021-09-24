package main

import "fmt"

func powerSeries(a int) (int, int) {
	return a * a, a * a * a
}

func tuples_print(){
	fmt.Println(powerSeries(5))
}

