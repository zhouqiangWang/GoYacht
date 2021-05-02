package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	fmt.Println("Hello, world")

	fmt.Println("Now is ", time.Now())

	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(math.Pi)

	fmt.Println(add(123, 34))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(sqrt(2))
}

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

const epsilon float64 = 1e-14

func sqrt(x float64) float64 {
	z := 1.0
	for ; math.Abs(z*z-x) > epsilon; z -= (z*z - x) / (2 * z) {
	}
	return z
}
