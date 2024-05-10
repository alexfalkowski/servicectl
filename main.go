package main

import (
	"os"

	c "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/servicectl/cmd"
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
