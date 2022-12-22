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

	_, err = listen.Accept()
	if err != nil {
		log.Panic("Error accepting connection: ", err)
	}
}
