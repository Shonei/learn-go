package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(byteConvert(156833213))
	fmt.Println(byteConvert(8101))
	fmt.Println(byteConvert(811))
	fmt.Println(byteConvert(12331, 3))
}

func byteConvert(bytes float64, options ...int) string {
	str := ""
	arr := [6]string{" bytes", " Kb", " Mb", " Gb", " Tb", " Pb"}

	if len(options) > 0 {
		str = "%." + strconv.Itoa(options[0]) + "f"
	} else {
		str = "%.2f"
	}
	iterations := 0
	for bytes > 1024 {
		bytes = bytes / 1024
		iterations++
	}

	str = str + arr[iterations]
	return fmt.Sprintf(str, bytes)
}
