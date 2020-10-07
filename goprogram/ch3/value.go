package ch3

import (
	"fmt"
	"math"
)

// bool 是%t, nan是一个float
func NaNTest() {
	nan := math.NaN()
	var b1 bool = nan == nan
	var b2 bool = nan < nan
	fmt.Printf("%f,%t,%t\n", nan, b1, b2)
}

func PrintInt() {
	fmt.Printf("%3.2d\n", 123)
	fmt.Printf("%10.3d\n", 1234)
}

func PrintFloat() {
	fmt.Printf("|%010.2f\n", 100.23)
	fmt.Printf("|%-10.5g\n", 100.23)
	fmt.Printf("|%10.5e\n", 100.23)
	fmt.Printf("%[1]g\n", 6.02211)
	fmt.Printf("%8.3f\n", 6.02211)
}

func PrintStr() {
	fmt.Printf("%10.2s\n", "abcd")
}

func Print() {
	PrintInt()
	PrintFloat()
	PrintStr()
}
