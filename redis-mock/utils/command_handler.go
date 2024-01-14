package utils

import (
	"github.com/redis-mock/models"
)

func HandleCommand(command string, args []string, store *models.Store) (string, error) {
	var response string
	var err error
	switch command {
	case "ping":
		response, err = HandlePing(args, store)
	case "set":
		response, err = HandleSet(args, store)
	case "get":
		response, err = HandleGet(args, store)
	case "exists":
		response, err = HandleExists(args, store)
	case "del":
		response, err = HandleDelete(args, store)
	default:
		// For any other command, respond with "OK"
		response, err = HandleDefault(args, store)
	}

	return response, err
}
