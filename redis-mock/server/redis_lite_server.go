package server

import (
	"fmt"
	"io"
	"net"
	"strings"

	"github.com/redis-mock/models"
	"github.com/redis-mock/utils"
)

func StartServer() {
	ln, err := net.Listen("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	fmt.Println("Redis-lite server started on port 6379")
	store := models.NewStringStore()
	limitChan := make(chan struct{}, 100)
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("::Error accepting conns: %v", err)
			continue
		}

		limitChan <- struct{}{}
		go handleConnection(conn, store, limitChan)
	}
}

func handleConnection(conn net.Conn, store *models.StringStore, limitChan <-chan struct{}) {

	defer conn.Close()
	defer func() { <-limitChan }()
	for {
		value, err := decodeRESP(conn)
		if err != nil {
			if err == io.EOF {
				return
			} else {
				fmt.Println(":: Error reading from client: ", err.Error())
				return
			}
		}

		handleResponse(value, conn, store)
	}
}

func decodeRESP(conn io.Reader) ([]string, error) {
	msg := make([]byte, 1024)
	var value any
	msglen, err := conn.Read(msg)
	if err != nil {
		return nil, err
	}
	message := strings.TrimSpace(string(msg[:msglen]))
	value, err = utils.Deserialize(message)
	if err != nil {
		return nil, err
	}
	valueInSlices, ok := value.([]string)
	if !ok {
		return nil, fmt.Errorf("Error: expeceted a slice of strings")
	}
	return valueInSlices, nil
}

func handleResponse(input []string, conn io.Writer, store *models.StringStore) {
	var response string
	var err error

	if len(input) == 0 {
		fmt.Println("No command received")
		return
	}

	// Converting the command to lower case for case-insensitive comparison
	command := strings.ToLower(input[0])
	args := input[1:]
	response, err = utils.HandleCommand(command, args, store)

	if err != nil {
		fmt.Println("Error while serializing message:", err)
		return
	}
	_, err = conn.Write([]byte(response))
	if err != nil {
		fmt.Println("Error sending response:", err)
	}
}
