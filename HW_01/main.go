package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	fmt.Printf("Square root approximation of %v\n", x)
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("Iteration %v, Value = %v\n", i+1, z)
	}
	return z
}

func SqrtNewtonMethod(x float64) float64 {
	z := 1.0
	t := 0.0
	i := 0
	for {
		z, t = z-(z*z-x)/(2*z), z
		fmt.Printf("Iteration = %v\n", i+1)
		fmt.Printf("z = %v, t = %v\n", z, t)
		if math.Abs(t-z) < 1e-9 {
			break
		}
		i++
	}
	return z
}

func main() {
	ourImplementation := SqrtNewtonMethod(5)
	libraryImplementation := math.Sqrt(5)
	fmt.Printf("Guess: %v, Expected: %v, Error: %v\n",
		ourImplementation, libraryImplementation,
		math.Abs(ourImplementation-libraryImplementation))
}
