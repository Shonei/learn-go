package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	Basic()
}

func Basic() {
	conn, err := net.Dial("tcp", "localhost:8000")
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}

	// _, err = conn.Write([]byte("This is client!"))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	b := make([]byte, 1024)
	i, err := conn.Read(b)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b[:i]))
}
