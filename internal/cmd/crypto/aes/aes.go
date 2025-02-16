package aes

import (
	"context"

	sc "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/crypto/aes"
	"github.com/alexfalkowski/go-service/crypto/rand"
	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/go-service/token"
	"github.com/alexfalkowski/servicectl/internal/cmd"
	"github.com/alexfalkowski/servicectl/internal/cmd/flags"
	"github.com/alexfalkowski/servicectl/internal/cmd/os"
	"github.com/alexfalkowski/servicectl/internal/cmd/runner"
	"github.com/alexfalkowski/servicectl/internal/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Register for aes.
func Register(command *sc.Command) {
	flags := sc.NewFlagSet("aes")

	flags.AddInput("")
	flags.BoolP("rotate", "r", false, "rotate key")
	flags.BoolP("verify", "v", false, "verify key")

	command.AddClient("aes", "AES crypto.", flags, cmd.Module, token.Module, fx.Invoke(Start))
}

// StartParams for aes.
type StartParams struct {
	fx.In

	Set       *sc.FlagSet
	Lifecycle fx.Lifecycle
	Logger    *zap.Logger
	Random    *rand.Generator
	Generator *aes.Generator
	Config    *config.Config
}

// Start for aes.
func Start(params StartParams) {
	var (
		fn runner.StartFn
		op string
	)

	switch {
	case flags.IsSet(params.Set, "rotate"):
		fn = func(ctx context.Context) context.Context {
			k, err := params.Generator.Generate()
			runtime.Must(err)

			err = os.WriteBase64File(params.Config.Crypto.AES.Key, []byte(k))
			runtime.Must(err)

			return ctx
		}
		op = "rotated key"
	case flags.IsSet(params.Set, "verify"):
		fn = func(ctx context.Context) context.Context {
			a, err := aes.NewCipher(params.Random, params.Config.Crypto.AES)
			runtime.Must(err)

			msg := "this is a test"
			enc, err := a.Encrypt(msg)
			runtime.Must(err)

			_, err = a.Decrypt(enc)
			runtime.Must(err)

			return meta.WithAttribute(ctx, "testMsg", meta.String(msg))
		}
		op = "verified key"
	}

	opts := &runner.Options{Lifecycle: params.Lifecycle, Logger: params.Logger, Fn: fn}
	runner.Start("aes", op, opts)
}
