package main

import (
	"github.com/redis-mock/server"
)

func main() {
	server := server.NewRedisServer()
	server.Start()
}
