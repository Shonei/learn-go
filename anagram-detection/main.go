package main

import (
	"fmt"
)

func main() {
	str := "Acda"
	anagram := "c0saA"
	t := '0'
	fmt.Println(t)
	// fmt.Println(findAnagrams(str, anagram))
	for _, val := range str {
		t = t ^ val
	}

	for _, val := range anagram {
		t = t ^ val
	}

	fmt.Println(t)
}

func findAnagrams(str, anagram string) int {
	end := len(str) - 1
	anagramCount := 0

	for i := len(anagram) - 1; i < end; i++ {
		if areAnagrams(str[i-len(anagram)+1:i+1], anagram) {
			anagramCount++
		}
	}
	return anagramCount
}

func areAnagrams(str, check string) bool {
	hits := 0

	for _, val := range str {
		for i, v := range check {
			if val == v {
				//  I could not figure out what should happen in this if statement
				check = check[:i] + check[i+1:]
				hits++
				break
			}
		}
	}

	return hits == len(str)
}
