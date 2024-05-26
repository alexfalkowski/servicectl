package main

import (
	"os"

	sc "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/flags"
	"github.com/alexfalkowski/servicectl/cmd"
	"github.com/alexfalkowski/servicectl/cmd/cache/redis"
	"github.com/alexfalkowski/servicectl/cmd/crypto/aes"
	"github.com/alexfalkowski/servicectl/cmd/crypto/ed25519"
	"github.com/alexfalkowski/servicectl/cmd/crypto/hmac"
	"github.com/alexfalkowski/servicectl/cmd/crypto/rsa"
	"github.com/alexfalkowski/servicectl/cmd/database/sql"
	cf "github.com/alexfalkowski/servicectl/cmd/feature"
	ch "github.com/alexfalkowski/servicectl/cmd/hooks"
	"github.com/alexfalkowski/servicectl/cmd/net/grpc"
	"github.com/alexfalkowski/servicectl/cmd/net/http"
	"github.com/alexfalkowski/servicectl/cmd/security/token"
)

func main() {
	if err := command().Run(); err != nil {
		os.Exit(1)
	}
}

type fn func(*sc.Command)

func command() *sc.Command {
	c := sc.New(cmd.Version)
	c.RegisterInput("")

	fns := []fn{cache, crypto, database, feature, hooks, net, tkn}
	for _, f := range fns {
		f(c)
	}

	return c
}

func cache(c *sc.Command) {
	r := c.AddClientCommand("redis", "Redis cache.", cmd.Module, redis.Module)
	flags.BoolVar(r, redis.VerifyFlag, "verify", "v", false, "verify connection")
}

func crypto(c *sc.Command) {
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
}

func database(c *sc.Command) {
	p := c.AddClientCommand("pg", "Postgres DB.", cmd.Module, sql.Module)
	flags.BoolVar(p, sql.VerifyFlag, "verify", "v", false, "verify connection")
}

func feature(c *sc.Command) {
	f := c.AddClientCommand("feature", "Feature flags.", cmd.Module, cf.Module)
	flags.BoolVar(f, cf.VerifyFlag, "verify", "v", false, "verify connection")
}

func hooks(c *sc.Command) {
	h := c.AddClientCommand("hooks", "Webhooks.", cmd.Module, ch.Module)
	flags.BoolVar(h, ch.RotateFlag, "rotate", "r", false, "rotate secret")
	flags.BoolVar(h, ch.VerifyFlag, "verify", "v", false, "verify webhook")
}

func net(c *sc.Command) {
	h := c.AddClientCommand("http", "HTTP Server.", cmd.Module, http.Module)
	flags.BoolVar(h, http.VerifyFlag, "verify", "v", false, "verify server")

	g := c.AddClientCommand("grpc", "gRPC Server.", cmd.Module, grpc.Module)
	flags.BoolVar(g, grpc.VerifyFlag, "verify", "v", false, "verify server")
}

func tkn(c *sc.Command) {
	s := c.AddClientCommand("token", "Token.", cmd.Module, token.Module)
	flags.BoolVar(s, token.RotateFlag, "rotate", "r", false, "rotate key")
	flags.BoolVar(s, token.VerifyFlag, "verify", "v", false, "verify key")
}
