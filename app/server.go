package main

import (
	"io"
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
		go handleConnection(conn)
		defer conn.Close()
	}
}

func handleConnection(conn net.Conn) {
	for {
		buffer := make([]byte, 1024)
		if _, err := conn.Read(buffer); err != nil {
			if err == io.EOF {
				log.Println("terminating")
				break
			}
			log.Panic("error occured", err)
		}
		conn.Write([]byte("+PONG\r\n"))
	}
}
