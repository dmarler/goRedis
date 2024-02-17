package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	fmt.Println("Starting server...")
	l, err := net.Listen("tcp", ":6969")
	if err != nil {
		fmt.Println("You suck...", err)
		return
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Really suck...", err)
		return
	}

	fmt.Println("Server started...")

	defer conn.Close()

	for {
		buf := make([]byte, 1024)

		_, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading from client: ", err)
			os.Exit(1)
		}

		conn.Write([]byte("+OK\r\n"))
	}
}
