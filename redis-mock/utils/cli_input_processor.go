package utils

import (
	"fmt"
	"net"
	"strings"

	"github.com/redis-mock/models"
)

func HandleInputs(args []string, conn net.Conn) {
	store := models.NewStore()

	switch strings.ToUpper(args[0]) {
	case "SET":
		if len(args) != 3 {
			fmt.Fprintln(conn, "ERROR: SET command requires two arguments")
		} else {
			store.Set(args[1], args[2])
			fmt.Fprintln(conn, "OK")
		}
	case "GET":
		if len(args) != 2 {
			fmt.Fprintln(conn, "ERROR: GET command requires one argument")
		} else {
			if value, ok := store.Get(args[1]); ok {
				fmt.Fprintln(conn, value)
			} else {
				fmt.Fprintln(conn, "nil")
			}
		}
	default:
		err, _ := Serialize[models.SimpleError](*models.NewSimpleError(fmt.Sprintf("Error: Unknown Command %v", args)))
		fmt.Println(err)
		fmt.Fprintln(conn, err)
	}
}
