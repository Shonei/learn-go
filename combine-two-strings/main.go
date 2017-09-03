package main

import "fmt"

func main() {

	str1 := "abc"
	str2 := "def"
	str3 := "daebfc"
	b := true

	if len(str3) != len(str1 + str2) {
		b = false
	}

	for i, v := range str3 {
		var index int = int(float64(i / 2))

		if byte(v) == str1[index] || 
			byte(v) == str2[index] {
			// do nothing
		} else {
			b = false
			break
		}
	}
	fmt.Println(b)
}