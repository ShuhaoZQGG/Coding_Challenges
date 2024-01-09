package utils

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func HandleScanner(conn net.Conn) {
	var args []string
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		input := scanner.Text()
		fmt.Println(input)
		args = strings.Split(input, " ")
	}
	HandleInputs(args, conn)
}
