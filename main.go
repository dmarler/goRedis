package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Starting server...")
	l, err := net.Listen("tcp", ":6969")
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Server started...")

	defer conn.Close()

	for {
		resp := NewResp(conn)
		value, err := resp.Read()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(value)

		writer := NewWriter(conn)
		writer.Write(Value{typ: "string", str: "OK"})
	}
}
