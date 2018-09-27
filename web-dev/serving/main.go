package main

import (
	"io"
	"log"
	"net"
)

var host = "localhost"
var port = ":8080"
var protocol = "tcp:"

func main() {
	li, err := net.Listen(protocol, host+port)

	if err != nil {
		return
	}
	defer li.Close()

	for {
		connection, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		io.WriteString(connection, "\nHello from TCP Server\n")
		connection.Close()
	}
}
