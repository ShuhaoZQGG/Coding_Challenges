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
	store := models.NewStore()
	for {
		conn, err := ln.Accept()
		fmt.Println(":: Got connected, ", conn.RemoteAddr())
		if err != nil {
			fmt.Printf("::Error accepting conns: %v", err)
		}

		go handleConnection(conn, store)
	}
}

func handleConnection(conn net.Conn, store *models.Store) {

	defer conn.Close()
	intialResponse, err := utils.Serialize[models.SimpleString](*models.NewSimpleString("OK"))
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = conn.Write([]byte(intialResponse))
	if err != nil {
		fmt.Println(":: Error sending initial response to client: ", err.Error())
		return
	}

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
	fmt.Println(value)
	valueInSlices, ok := value.([]string)
	if !ok {
		return nil, fmt.Errorf("Error: expeceted a slice of strings")
	}
	fmt.Println(":: Parsed: ", valueInSlices)
	return valueInSlices, nil
}

// TODO: FIX ISSUE: It seems the response back to cli only takes effect after the next command
// response is one time slower than expected
func handleResponse(input []string, conn io.Writer, store *models.Store) {
	var response string
	var err error

	if len(input) == 0 {
		fmt.Println("No command received")
		return
	}

	// Converting the command to lower case for case-insensitive comparison
	command := strings.ToLower(input[0])

	switch command {
	case "ping":
		// Respond with "PONG" only if the command is "ping"
		response, err = utils.Serialize[models.SimpleString](*models.NewSimpleString("PONG"))
	default:
		// For any other command, respond with "OK"
		response, err = utils.Serialize[models.SimpleString](*models.NewSimpleString("OK"))
	}

	if err != nil {
		fmt.Println("Error while serializing message:", err)
		return
	}

	fmt.Println("Response:", response)
	_, err = conn.Write([]byte(response))
	if err != nil {
		fmt.Println("Error sending response:", err)
	}
}
