package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func handleClient(conn net.Conn) {
	conn.Close()
	reader := bufio.NewReader(conn)

	// 사용자 이름 요청
	conn.Write([]byte("Enter your username: "))
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)
	clients[conn] = username

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("User %s disconnected.\n", username)
			removeClientFromAllChannels(conn)
			break
		}

		message = strings.TrimSpace(message)
		if strings.HasPrefix(message, "/join") {
			handleJoinChannel(conn, message)
		} else if strings.HasPrefix(message, "/leave") {
			handleLeaveChannel(conn, message)
		} else if strings.HasPrefix(message, "/msg") {
			handleChannelMessage(conn, message)
		} else {
			conn.Write([]byte("Unknown command\n"))
		}
	}
}

func removeClient(conn net.Conn, channel string) {
	for i, c := range channels[channel] {
		if c == conn {
			channels[channel] = append(channels[channel][:i], channels[channel][i+1:]...)
			break
		}
	}
}

func removeClientFromAllChannels(conn net.Conn) {
	for channel := range channels {
		removeClient(conn, channel)
	}
}
