package server

import (
	"fmt"
	"net"

	"github.com/redis-mock/utils"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	utils.HandleScanner(conn)
}

func StartServer() {
	ln, err := net.Listen("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	fmt.Println("Redis-lite server started on port 6379")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}
