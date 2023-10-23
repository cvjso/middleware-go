package main

import (
	"bufio"
	"client/tcp"
	"client/udp"
	"fmt"
	"os"
	"strings"
)

func main() {
	// clientTCP()
	clientUDP()
}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	text, _ := reader.ReadString('\n')
	return text
}

func clientUDP() {
	middleware := udp.NewClientRequestHandlerUDP("localhost:2000")
	for {
		text := readInput()
		middleware.Send([]byte(text))
		msg := middleware.Receive()
		parsedMsg := string(msg[:])
		fmt.Print("-> " + parsedMsg)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client closing...")
			return
		}
	}
}

func clientTCP() {
	middleware := tcp.NewClientRequestHandlerTCP("localhost:2000")
	for {
		text := readInput()
		middleware.Send([]byte(text))

		rawMessage := middleware.Receive()
		message := string(rawMessage[:])
		fmt.Print("-> " + message)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client closing...")
			return
		}
	}
}
