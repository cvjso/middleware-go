package main

import (
	"fmt"
	"net"
)

func client() {

}

func clientRequestHandler(typeCon string, addr string) {
	conn, err := net.Dial(typeCon, addr)
	if err != nil {
		fmt.Printf(err.Error())
	}
	content := []byte{}
	conn.Read(content)
	fmt.Printf(string(content))
}
