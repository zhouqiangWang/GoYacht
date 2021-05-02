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
}

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}
