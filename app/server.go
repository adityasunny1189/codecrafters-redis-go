package main

import (
	"log"
	"net"
)

const (
	HOST = "localhost"
	PORT = "6379"
	TYPE = "tcp"
)

func main() {
	log.Println("Logs from your program will appear here!")
	listen, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Panic("Failed to bind to port 6379 ", err)
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Panic("Error accepting connection: ", err)
		}
		for {
			handleConnection(conn)
			defer conn.Close()
		}
	}
}

func handleConnection(conn net.Conn) {
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		log.Panic("error occured", err)
	}
	conn.Write([]byte("+PONG\r\n"))
}
