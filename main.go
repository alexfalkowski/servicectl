package main

import (
	"os"

	sc "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/servicectl/cmd"
	"github.com/alexfalkowski/servicectl/cmd/aes"
)

func main() {
	if err := command().Run(); err != nil {
		os.Exit(1)
	}
}

func command() *sc.Command {
	command := sc.New(cmd.Version)

	ro := command.AddClientCommand("aes", "AES crypto.", cmd.Module)
	sc.StringVar(ro, sc.OutputFlag, "output", "o", "env:AES_CONFIG_FILE", "output config location (format kind:location, default env:AES_CONFIG_FILE)")
	sc.BoolVar(ro, aes.Rotate, "rotate", "r", false, "rotate keys")

	return command
}
