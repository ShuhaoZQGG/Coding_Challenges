package server

import (
	"fmt"
	"io"
	"net"
	"strings"

	"github.com/redis-mock/models"
	"github.com/redis-mock/utils"
)

type Server interface {
	Start()
	HandleConnection()
	HandleResponse()
}

type RedisServer struct {
	aof   *models.AOF
	store *models.StringStore
}

func NewRedisServer() *RedisServer {
	aof, _ := models.NewAOF()
	store := models.NewStringStore()
	return &RedisServer{
		aof:   aof,
		store: store,
	}
}

func (rs *RedisServer) Start() {
	ln, err := net.Listen("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	fmt.Println("Redis-lite server started on port 6379")
	if err != nil {
		fmt.Println("failed to create aof struct")
	}
	limitChan := make(chan struct{}, 50)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("::Error accepting conns: %v", err)
			continue
		}

		limitChan <- struct{}{}
		go rs.HandleConnection(conn, limitChan)
	}
}

func (rs *RedisServer) HandleConnection(conn net.Conn, limitChan <-chan struct{}) {

	defer conn.Close()
	defer func() { <-limitChan }()
	for {
		value, err := rs.decodeRESP(conn)
		if err != nil {
			if err == io.EOF {
				return
			} else {
				fmt.Println(":: Error reading from client: ", err.Error())
				return
			}
		}

		rs.handleResponse(conn, value)
	}
}

func (rs *RedisServer) decodeRESP(conn net.Conn) ([]string, error) {
	deserializer := new(models.RespDeserializer)
	msg := make([]byte, 1024)
	var value any
	msglen, err := conn.Read(msg)
	if err != nil {
		return nil, err
	}
	rs.aof.Write(string(msg[:msglen]))
	message := strings.TrimSpace(string(msg[:msglen]))
	value, err = deserializer.Deserialize(message)
	if err != nil {
		return nil, err
	}
	valueInSlices, ok := value.([]string)
	if !ok {
		return nil, fmt.Errorf("Error: expeceted a slice of strings")
	}
	return valueInSlices, nil
}

func (rs *RedisServer) handleResponse(conn net.Conn, input []string) {
	var response string
	var err error

	if len(input) == 0 {
		fmt.Println("No command received")
		return
	}

	// Converting the command to lower case for case-insensitive comparison
	command := strings.ToLower(input[0])
	args := input[1:]
	response, err = utils.HandleCommand(command, args, rs.store)

	if err != nil {
		fmt.Println("Error while serializing message:", err)
		return
	}
	_, err = conn.Write([]byte(response))
	if err != nil {
		fmt.Println("Error sending response:", err)
	}
}
