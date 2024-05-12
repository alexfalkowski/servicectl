package main

import (
	"os"

	sc "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/flags"
	"github.com/alexfalkowski/servicectl/cmd"
	"github.com/alexfalkowski/servicectl/cmd/aes"
	"github.com/alexfalkowski/servicectl/cmd/hmac"
)

func main() {
	if err := command().Run(); err != nil {
		os.Exit(1)
	}
}

func command() *sc.Command {
	c := sc.New(cmd.Version)
	c.RegisterInput("")
	c.RegisterOutput("")

	ac := c.AddClientCommand("aes", "AES crypto.", cmd.Module, aes.Module)
	flags.BoolVar(ac, aes.RotateFlag, "rotate", "r", false, "rotate keys")
	flags.BoolVar(ac, aes.VerifyFlag, "verify", "v", false, "verify keys")

	ah := c.AddClientCommand("hmac", "HMAC crypto.", cmd.Module, hmac.Module)
	flags.BoolVar(ah, hmac.RotateFlag, "rotate", "r", false, "rotate keys")
	flags.BoolVar(ah, hmac.VerifyFlag, "verify", "v", false, "verify keys")

	return c
}
