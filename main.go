package main

import (
	"os"

	sc "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/flags"
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
	flags.StringVar(ro, sc.OutputFlag,
		"output", "o", "env:AES_CONFIG_FILE", "output config location (format kind:location, default env:AES_CONFIG_FILE)")
	flags.BoolVar(ro, aes.RotateFlag, "rotate", "r", false, "rotate keys")

	return command
}
