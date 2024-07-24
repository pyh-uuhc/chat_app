package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	go readMessage(conn)

	reader := bufio.NewReader(os.Stdin)
	for {
		message, _ := reader.ReadString('\n')
		fmt.Fprint(conn, message)
	}
}

func readMessage(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Disconnected from server")
			return
		}
		fmt.Printf(message)
	}
}
