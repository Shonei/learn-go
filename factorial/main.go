package main

import (
	"fmt"
)

func main() {
	fmt.Println(factorial(5))
}

func factorial(fac int) int {
	result := 1
	for ; fac > 1; fac-- {
		result *= fac
	}
	return result
}
