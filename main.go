package main

import (
	"os"

	"github.com/alexfalkowski/go-client-template/cmd"
	c "github.com/alexfalkowski/go-service/cmd"
)

func main() {
	if err := command().Run(); err != nil {
		os.Exit(1)
	}
}

func command() *c.Command {
	command := c.New(cmd.Version)

	command.AddClient(cmd.ClientOptions...)

	return command
}
