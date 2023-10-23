package tcp

import (
	"bufio"
	"fmt"
	"net"
)

type ServerRequestHandler struct {
	Address    string
	Listener   net.Listener
	Connection net.Conn
}

func (s ServerRequestHandler) connectionIsEmpty() bool {
	return s.Connection == nil
}

func (s ServerRequestHandler) Send(msg []byte) {
	if !s.connectionIsEmpty() {
		s.Connection.Write(msg)
	}
}

func (s *ServerRequestHandler) Receive() []byte {
	if s.connectionIsEmpty() {
		l, err := s.Listener.Accept()
		if err != nil {
			fmt.Println("Error on listening")
			fmt.Println(err.Error())

		}
		s.Connection = l
	}
	data, err := bufio.NewReader(s.Connection).ReadBytes('\n')
	if err != nil {
		fmt.Println("Error on receive")
		fmt.Println(err.Error())
		fmt.Println("Closing connection...")
		s.Connection = nil
	}
	return data
}

func NewServerRequestHandler(Address string) *ServerRequestHandler {
	s := ServerRequestHandler{Address: Address}
	l, err := net.Listen("tcp", s.Address)
	if err != nil {
		fmt.Println("Error on listening")
		fmt.Println(err.Error())
	}
	s.Listener = l
	return &s
}
