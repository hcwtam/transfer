package server

import (
	"log"
	"net"
)

func Connect() net.Conn {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

func Serve() {
	net.Listen("tcp", ":8080")
}
