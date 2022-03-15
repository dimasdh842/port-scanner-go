package main

import (
	"fmt"
	"net"
)

func main() {

	for i := 1; i <= 1024; i++ {
		address := fmt.Sprintf("localhost:%d", i)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			continue
		}

		conn.Close()

		fmt.Println("port ", i, " open")
	}
}
