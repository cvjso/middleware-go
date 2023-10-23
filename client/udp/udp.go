package udp

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

type ClientRequestHandlerUDP struct {
	serverAddr string
	conn       net.Conn
	udpAddr    *net.UDPAddr
}

func (c ClientRequestHandlerUDP) Receive() []byte {
	msg, err := bufio.NewReader(c.conn).ReadBytes('\n')
	if err != nil {
		println("Read data failed:", err.Error())
		os.Exit(1)
	}
	return msg
}

func (c ClientRequestHandlerUDP) Send(msg []byte) {
	_, err := c.conn.Write(msg)
	if err != nil {
		println("Write data failed:", err.Error())
		os.Exit(1)
	}
}

func NewClientRequestHandlerUDP(Address string) *ClientRequestHandlerUDP {
	c := ClientRequestHandlerUDP{serverAddr: Address}
	udpServer, err := net.ResolveUDPAddr("udp", c.serverAddr)
	if err != nil {
		println("ResolveUDPAddr failed:", err.Error())
		os.Exit(1)
	}
	c.udpAddr = udpServer
	conn, err := net.DialUDP("udp", nil, c.udpAddr)
	if err != nil {
		fmt.Println("Error on dialing server...")
	}
	c.conn = conn
	return &c
}
