package server

import (
	"net"
	"strings"
)

func handleJoinChannel(conn net.Conn, message string) {
	parts := strings.Split(message, " ")
	if len(parts) != 2 {
		conn.Write([]byte("Usage: /join <channel>\n"))
		return
	}
	channel := parts[1]
	channels[channel] = append(channels[channel], conn)
	conn.Write([]byte("Joined channel " + channel + "\n"))
}

func handleLeaveChannel(conn net.Conn, message string) {
	parts := strings.Split(message, " ")
	if len(parts) != 2 {
		conn.Write([]byte("Usage: /leave <channel>\n"))
		return
	}
	channel := parts[1]
	removeClient(conn, channel)
	conn.Write([]byte("Left channel " + channel + "\n"))
}

func handleChannelMessage(conn net.Conn, message string) {
	parts := strings.SplitN(message, " ", 3)
	if len(parts) != 3 {
		conn.Write([]byte("Usage: /msg <channel> <message>\n"))
		return
	}
	channel := parts[1]
	content := parts[2]
	broadcast <- Message{Channel: channel, Content: clients[conn] + ": " + content}
}
