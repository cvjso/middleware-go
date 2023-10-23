package main

import (
	"fmt"
	"server/tcp"
	"server/udp"
	"strings"
)

type callable func(p string, player position) position

type position struct {
	x int
	y int
}

func move(x int, y int, p *position) {
	(*p).x += x
	(*p).y += y
}

func handleMovement(param string, player *position) {
	switch param {
	case "left":
		move(-1, 0, player)
	case "right":
		move(1, 0, player)
	case "up":
		move(0, 1, player)
	case "down":
		move(0, -1, player)
	}
}

func main() {
	// tcpServer()
	udpServer()
}

func udpServer() {
	middleware := udp.NewServerRequestHandler("localhost:2000")
	player := position{0, 0}
	for {
		msg := middleware.Receive()
		parsedMsg := string(msg)
		parsedMsg = strings.Split(parsedMsg, "\n")[0]
		handleMovement(parsedMsg, &player)
		response := fmt.Sprintf("%+v\n", player)
		middleware.Send([]byte(response))
	}

}

func tcpServer() {
	middleware := tcp.NewServerRequestHandler("localhost:2000")
	player := position{0, 0}
	for {
		msg := middleware.Receive()
		parsedMsg := string(msg[:])
		handleMovement(strings.TrimSpace(parsedMsg), &player)
		response := fmt.Sprintf("%+v\n", player)
		middleware.Send([]byte(response))
	}
}
