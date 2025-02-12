package main

import (
	sc "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/flags"
	"github.com/alexfalkowski/servicectl/internal/cmd"
	"github.com/alexfalkowski/servicectl/internal/cmd/crypto/aes"
	"github.com/alexfalkowski/servicectl/internal/cmd/crypto/ed25519"
	"github.com/alexfalkowski/servicectl/internal/cmd/crypto/hmac"
	"github.com/alexfalkowski/servicectl/internal/cmd/crypto/rsa"
	"github.com/alexfalkowski/servicectl/internal/cmd/database/sql"
	"github.com/alexfalkowski/servicectl/internal/cmd/hooks"
	"github.com/alexfalkowski/servicectl/internal/cmd/net/grpc"
	"github.com/alexfalkowski/servicectl/internal/cmd/net/http"
	"github.com/alexfalkowski/servicectl/internal/cmd/token"
)

func main() {
	command().ExitOnError()
}

type fn func(*sc.Command)

func command() *sc.Command {
	c := sc.New(cmd.Version)
	c.RegisterInput(c.Root(), "")

	fns := []fn{cryptoClient, databaseClient, hooksClient, netClient, tokenClient}
	for _, f := range fns {
		f(c)
	}

	return c
}

func cryptoClient(c *sc.Command) {
	ac := c.AddClient("aes", "AES crypto.", cmd.Module, aes.Module)
	flags.BoolVar(ac, aes.RotateFlag, "rotate", "r", false, "rotate key")
	flags.BoolVar(ac, aes.VerifyFlag, "verify", "v", false, "verify key")

	ah := c.AddClient("hmac", "HMAC crypto.", cmd.Module, hmac.Module)
	flags.BoolVar(ah, hmac.RotateFlag, "rotate", "r", false, "rotate key")
	flags.BoolVar(ah, hmac.VerifyFlag, "verify", "v", false, "verify key")

	ar := c.AddClient("rsa", "RSA crypto.", cmd.Module, rsa.Module)
	flags.BoolVar(ar, rsa.RotateFlag, "rotate", "r", false, "rotate keys")
	flags.BoolVar(ar, rsa.VerifyFlag, "verify", "v", false, "verify keys")

	ae := c.AddClient("ed25519", "Ed25519 crypto.", cmd.Module, ed25519.Module)
	flags.BoolVar(ae, ed25519.RotateFlag, "rotate", "r", false, "rotate keys")
	flags.BoolVar(ae, ed25519.VerifyFlag, "verify", "v", false, "verify keys")
}

func databaseClient(c *sc.Command) {
	p := c.AddClient("pg", "Postgres DB.", cmd.Module, sql.Module)
	flags.BoolVar(p, sql.VerifyFlag, "verify", "v", false, "verify connection")
}

func hooksClient(c *sc.Command) {
	h := c.AddClient("hooks", "Webhooks.", cmd.Module, hooks.Module)
	flags.BoolVar(h, hooks.RotateFlag, "rotate", "r", false, "rotate secret")
	flags.BoolVar(h, hooks.VerifyFlag, "verify", "v", false, "verify webhook")
}

func netClient(c *sc.Command) {
	h := c.AddClient("http", "HTTP Server.", cmd.Module, http.Module)
	flags.BoolVar(h, http.VerifyFlag, "verify", "v", false, "verify server")

	g := c.AddClient("grpc", "gRPC Server.", cmd.Module, grpc.Module)
	flags.BoolVar(g, grpc.VerifyFlag, "verify", "v", false, "verify server")
}

func tokenClient(c *sc.Command) {
	t := c.AddClient("token", "Security tokens.", cmd.Module, token.Module)
	flags.BoolVar(t, token.RotateFlag, "rotate", "r", false, "rotate secret")
}
