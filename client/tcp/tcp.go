package tcp

import (
	"bufio"
	"fmt"
	"net"
)

type ClientRequestHandlerTCP struct {
	serverAddr string
	conn       net.Conn
}

func (c ClientRequestHandlerTCP) Send(msg []byte) {
	if c.conn == nil {
		fmt.Println("Closing connection...")
		return
	}
	c.conn.Write(msg)
}

func (c ClientRequestHandlerTCP) Receive() []byte {
	msg, err := bufio.NewReader(c.conn).ReadBytes('\n')
	if err != nil {
		fmt.Println("Error on receive")
	}
	return msg
}

func NewClientRequestHandlerTCP(Address string) *ClientRequestHandlerTCP {
	c := ClientRequestHandlerTCP{serverAddr: Address}
	conn, err := net.Dial("tcp", c.serverAddr)
	if err != nil {
		fmt.Println("Error on dialing server...")
	}
	c.conn = conn
	return &c
}
