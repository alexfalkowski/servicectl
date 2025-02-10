package main

import (
	sc "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/flags"
	"github.com/alexfalkowski/servicectl/cmd"
	"github.com/alexfalkowski/servicectl/cmd/crypto/aes"
	"github.com/alexfalkowski/servicectl/cmd/crypto/ed25519"
	"github.com/alexfalkowski/servicectl/cmd/crypto/hmac"
	"github.com/alexfalkowski/servicectl/cmd/crypto/rsa"
	"github.com/alexfalkowski/servicectl/cmd/database/sql"
	ch "github.com/alexfalkowski/servicectl/cmd/hooks"
	"github.com/alexfalkowski/servicectl/cmd/net/grpc"
	"github.com/alexfalkowski/servicectl/cmd/net/http"
	"github.com/alexfalkowski/servicectl/cmd/token"
)

func main() {
	command().ExitOnError()
}

type fn func(*sc.Command)

func command() *sc.Command {
	c := sc.New(cmd.Version)
	c.RegisterInput(c.Root(), "")

	fns := []fn{crypto, database, hooks, net, tkn}
	for _, f := range fns {
		f(c)
	}

	return c
}

func crypto(c *sc.Command) {
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

func database(c *sc.Command) {
	p := c.AddClient("pg", "Postgres DB.", cmd.Module, sql.Module)
	flags.BoolVar(p, sql.VerifyFlag, "verify", "v", false, "verify connection")
}

func hooks(c *sc.Command) {
	h := c.AddClient("hooks", "Webhooks.", cmd.Module, ch.Module)
	flags.BoolVar(h, ch.RotateFlag, "rotate", "r", false, "rotate secret")
	flags.BoolVar(h, ch.VerifyFlag, "verify", "v", false, "verify webhook")
}

func net(c *sc.Command) {
	h := c.AddClient("http", "HTTP Server.", cmd.Module, http.Module)
	flags.BoolVar(h, http.VerifyFlag, "verify", "v", false, "verify server")

	g := c.AddClient("grpc", "gRPC Server.", cmd.Module, grpc.Module)
	flags.BoolVar(g, grpc.VerifyFlag, "verify", "v", false, "verify server")
}

func tkn(c *sc.Command) {
	t := c.AddClient("token", "Security tokens.", cmd.Module, token.Module)
	flags.BoolVar(t, token.RotateFlag, "rotate", "r", false, "rotate secret")
}
