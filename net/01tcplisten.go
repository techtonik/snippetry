// transpilery:primary
package main

import (
	"fmt"
	"net"
)

func main() {
	listenon := "localhost:11111"
	fmt.Println("Listening on: " + listenon)
	acceptor, _ := net.Listen("tcp", listenon)
	conn, _ := acceptor.Accept()
	fmt.Printf("New connection from: %s\n", conn.RemoteAddr())
}
