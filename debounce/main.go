package main

import (
	"fmt"
	"time"
)

func main() {
	f := timeFunction(func() {
		fmt.Println(time.Now())
	}, 2)

	f()
	f()
}

func timeFunction(f func(), sec int) func() {
	timer := time.Now().Local().Add(-(time.Duration(sec) * time.Second))

	return func() {
		if int(time.Since(timer)) > sec {
			f()
			timer = time.Now()
		} else {
			timer = time.Now()
		}
	}
}
