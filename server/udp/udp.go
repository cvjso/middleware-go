package udp

import (
	"fmt"
	"log"
	"net"
)

type ServerRequestHandler struct {
	address    string
	listener   net.Listener
	connection net.PacketConn
	addr       net.Addr
}

func (s ServerRequestHandler) Send(msg []byte) {
	s.connection.WriteTo(msg, s.addr)
}

func (s *ServerRequestHandler) Receive() []byte {
	buf := make([]byte, 1024)
	_, addr, err := s.connection.ReadFrom(buf)
	if err != nil {
		log.Fatal(err)
	}
	s.addr = addr
	return buf
}

func NewServerRequestHandler(Address string) *ServerRequestHandler {
	s := ServerRequestHandler{address: Address}
	c, err := net.ListenPacket("udp", Address)
	if err != nil {
		fmt.Println("Error on listening address")
		fmt.Println(err.Error())
	}
	s.connection = c
	return &s

}
