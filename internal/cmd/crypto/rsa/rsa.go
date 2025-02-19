package rsa

import (
	"context"

	sc "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/crypto/rand"
	"github.com/alexfalkowski/go-service/crypto/rsa"
	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/go-service/telemetry/logger"
	"github.com/alexfalkowski/servicectl/internal/cmd"
	cf "github.com/alexfalkowski/servicectl/internal/cmd/flags"
	"github.com/alexfalkowski/servicectl/internal/cmd/os"
	"github.com/alexfalkowski/servicectl/internal/cmd/runner"
	"github.com/alexfalkowski/servicectl/internal/config"
	"go.uber.org/fx"
)

// Register for rsa.
func Register(command *sc.Command) {
	flags := sc.NewFlagSet("rsa")

	flags.AddInput("")
	flags.BoolP("rotate", "r", false, "rotate key")
	flags.BoolP("verify", "v", false, "verify key")

	command.AddClient("rsa", "RSA crypto.", flags, cmd.Module, fx.Invoke(Start))
}

// StartParams for rsa.
type StartParams struct {
	fx.In

	Set       *sc.FlagSet
	Lifecycle fx.Lifecycle
	Logger    *logger.Logger
	Random    *rand.Generator
	Generator *rsa.Generator
	Config    *config.Config
}

// Start for rsa.
func Start(params StartParams) {
	var (
		fn runner.StartFn
		op string
	)

	switch {
	case cf.IsSet(params.Set, "rotate"):
		fn = func(ctx context.Context) context.Context {
			pub, pri, err := params.Generator.Generate()
			runtime.Must(err)

			err = os.WriteFile(params.Config.Crypto.RSA.Public, []byte(pub))
			runtime.Must(err)

			err = os.WriteFile(params.Config.Crypto.RSA.Private, []byte(pri))
			runtime.Must(err)

			return ctx
		}
		op = "rotated keys"
	case cf.IsSet(params.Set, "verify"):
		fn = func(ctx context.Context) context.Context {
			a, err := rsa.NewCipher(params.Random, params.Config.Crypto.RSA)
			runtime.Must(err)

			msg := "this is a test"
			enc, err := a.Encrypt(msg)
			runtime.Must(err)

			_, err = a.Decrypt(enc)
			runtime.Must(err)

			return meta.WithAttribute(ctx, "testMsg", meta.String(msg))
		}
		op = "verified keys"
	}

	opts := &runner.Options{Lifecycle: params.Lifecycle, Logger: params.Logger, Fn: fn}
	runner.Start("rsa", op, opts)
}
