package main

import (
	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/env"
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

func command() *cmd.Command {
	command := cmd.New(env.NewVersion().String())

	aes.Register(command)
	hmac.Register(command)
	rsa.Register(command)
	ed25519.Register(command)
	sql.Register(command)
	hooks.Register(command)
	http.Register(command)
	grpc.Register(command)
	token.Register(command)

	return command
}
