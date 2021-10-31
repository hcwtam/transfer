package cli

import (
	"errors"
	"io"
	"log"
	"os"

	"github.com/hcwtam/transfer/internal/server"
)

const BUFFER_SIZE = 1024

func Run() {
	input, err := getArgs()

	if err != nil {
		log.Fatal(err)
	}

	command := input[0]

	switch command {
	case "send":
		send(input[1:])
	default:
		log.Fatal("argument provided is invalid")
	}
}

func getArgs() ([]string, error) {
	args := os.Args
	if len(args) == 1 {
		return nil, errors.New("you must provide arguments")
	}

	return args[1:], nil
}

func send(args []string) {
	if len(args) == 0 {
		log.Fatal("you must provide at least one file")
	}

	conn := server.Connect()

	for _, path := range args {
		// read file
		file, err := os.Open(path)

		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()

		_, err = io.Copy(conn, file)
		if err != nil {
			log.Fatal(err)
		}

		defer conn.Close()
	}
}
