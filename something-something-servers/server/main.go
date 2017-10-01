package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	Basic()
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	b := make([]byte, 1024)

	conn.SetReadDeadline(time.Now().Add(time.Second * 5))

	i, err := conn.Read(b)
	if o, e := err.(net.Error); e {
		fmt.Println("Error", o, e)
		return
	} else if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b[:i]))

	conn.Write([]byte("Hello this is server!"))
	if err := conn.Close(); err != nil {
		log.Fatal(err)
	}
}

func Basic() {
	socket, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := socket.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn)
	}
}
