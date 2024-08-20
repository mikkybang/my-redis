package main

import (
	"fmt"

	// Uncomment this block to pass the first stage
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	//
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go handleRequest(conn)
	}
}

func handleRequest(c net.Conn) {
	buffer := make([]byte, 1024)
	n, readErr := c.Read(buffer)
	if readErr != nil {
		fmt.Println("Failed To read connection")
		os.Exit(1)
	}
	fmt.Printf("Received Message: %s", buffer[:n])

	_, write_err := c.Write([]byte("+PONG\r\n"))
	if write_err != nil {
		fmt.Println("failed:", write_err)
		os.Exit(1)
	}
}
