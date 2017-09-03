package main

import "fmt"

func findMissing(full [10]int, parcial [9]int)(int) {
	val := 0
	for _, v := range full {
		for _, p := range parcial {
			if v == p {
				break
			} else {
				val = v
			}
		}
	}

	return val
}

func main() {
	full := [10]int{83, 98, 43, 51, 41, 95, 7, 98, 79, 47}
	less := [9]int{83, 98, 43, 51, 41, 95, 7, 98, 79}

	fmt.Println(findMissing(full, less))
}