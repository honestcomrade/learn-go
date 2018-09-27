package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

var port = ":8080"
var protocol = "tcp"

func main() {
	fmt.Println("Creating listener at", protocol, port)
	li, err := net.Listen(protocol, port)

	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Println("Connection TIMEOUT")
	}
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I head you say: %s\n", ln)
	}
	defer conn.Close()
	fmt.Println("Connection Closed")
}
