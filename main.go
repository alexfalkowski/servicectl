package main

import (
	"os"

	sc "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/flags"
	"github.com/alexfalkowski/servicectl/cmd"
	"github.com/alexfalkowski/servicectl/cmd/aes"
	"github.com/alexfalkowski/servicectl/cmd/ed25519"
	"github.com/alexfalkowski/servicectl/cmd/hmac"
	"github.com/alexfalkowski/servicectl/cmd/rsa"
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
	flags.BoolVar(ac, aes.RotateFlag, "rotate", "r", false, "rotate key")
	flags.BoolVar(ac, aes.VerifyFlag, "verify", "v", false, "verify key")

	ah := c.AddClientCommand("hmac", "HMAC crypto.", cmd.Module, hmac.Module)
	flags.BoolVar(ah, hmac.RotateFlag, "rotate", "r", false, "rotate key")
	flags.BoolVar(ah, hmac.VerifyFlag, "verify", "v", false, "verify key")

	ar := c.AddClientCommand("rsa", "RSA crypto.", cmd.Module, rsa.Module)
	flags.BoolVar(ar, rsa.RotateFlag, "rotate", "r", false, "rotate keys")
	flags.BoolVar(ar, rsa.VerifyFlag, "verify", "v", false, "verify keys")

	ae := c.AddClientCommand("ed25519", "Ed25519 crypto.", cmd.Module, ed25519.Module)
	flags.BoolVar(ae, ed25519.RotateFlag, "rotate", "r", false, "rotate keys")
	flags.BoolVar(ae, ed25519.VerifyFlag, "verify", "v", false, "verify keys")

	return c
}
