package server

import (
	"fmt"
	"log"
	"net"
)

var (
	clients   = make(map[net.Conn]string)
	channels  = make(map[string][]net.Conn)
	broadcast = make(chan Message)
)

type Message struct {
	Channel string
	Content string
}

func StartServer() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	defer listener.Close()

	go handleMessages()

	fmt.Println("Server started on :8080")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}
		go handleClient(conn)
	}
}

func handleMessages() {
	for {
		msg := <-broadcast
		for _, conn := range channels[msg.Channel] {
			if _, err := conn.Write([]byte(msg.Content + "\n")); err != nil {
				log.Printf("Error sending message: %v", err)
				conn.Close()
				removeClient(conn, msg.Channel)
			}
		}
	}
}
