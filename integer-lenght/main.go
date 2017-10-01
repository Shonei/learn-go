package main

import "fmt"

func main() {
	var value int
	value = 593485703962409
	fmt.Println(intLenght(value))

}

func intLenght(val int) int {
	var i int
	for i = 1; val > 10; i++ {
		val = val / 10
	}

	return i
}
