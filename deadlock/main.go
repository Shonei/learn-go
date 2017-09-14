package main

import "time"

func main() {
	c1 := make(chan int, 2)
	c2 := make(chan int, 2)

	go func() {
		for i := range c1 {
			println("G1 got", i)
			c2 <- i // waits here to write to c2
			// can't write to c2 as c2 is blocked a reading
		}
	}()

	go func() {
		for i := range c2 {
			println("G2 got", i)
			c1 <- i // stoprs here and waits to write to c1
			// can't write to c1 because c1 is already full
			// because we send 2 elements to it
		}
	}()

	c1 <- 1
	c1 <- 1

	time.Sleep(5)

}
