package cli

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

type cli struct {
}

func (c *cli) getArgs() ([]string, error) {
	args := os.Args
	if len(args) == 1 {
		return nil, errors.New("you must provide arguments")
	}

	return args[1:], nil
}

func Run() {
	cli := cli{}

	input, err := cli.getArgs()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(strings.Join(input, " "))
}
