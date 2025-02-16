package ed25519

import (
	"context"

	sc "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/crypto/ed25519"
	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/servicectl/internal/cmd"
	"github.com/alexfalkowski/servicectl/internal/cmd/flags"
	"github.com/alexfalkowski/servicectl/internal/cmd/os"
	"github.com/alexfalkowski/servicectl/internal/cmd/runner"
	"github.com/alexfalkowski/servicectl/internal/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Register for ed25519.
func Register(command *sc.Command) {
	flags := sc.NewFlagSet("ed25519")

	flags.AddInput("")
	flags.BoolP("rotate", "r", false, "rotate key")
	flags.BoolP("verify", "v", false, "verify key")

	command.AddClient("ed25519", "Ed25519 crypto.", flags, cmd.Module, fx.Invoke(Start))
}

// StartParams for ed25519.
type StartParams struct {
	fx.In

	Set       *sc.FlagSet
	Lifecycle fx.Lifecycle
	Logger    *zap.Logger
	Generator *ed25519.Generator
	Config    *config.Config
}

// Start for ed25519.
func Start(params StartParams) {
	var (
		fn runner.StartFn
		op string
	)

	switch {
	case flags.IsSet(params.Set, "rotate"):
		fn = func(ctx context.Context) context.Context {
			pub, pri, err := params.Generator.Generate()
			runtime.Must(err)

			err = os.WriteFile(params.Config.Crypto.Ed25519.Public, []byte(pub))
			runtime.Must(err)

			err = os.WriteFile(params.Config.Crypto.Ed25519.Private, []byte(pri))
			runtime.Must(err)

			return ctx
		}
		op = "rotated keys"
	case flags.IsSet(params.Set, "verify"):
		fn = func(ctx context.Context) context.Context {
			a, err := ed25519.NewSigner(params.Config.Crypto.Ed25519)
			runtime.Must(err)

			msg := "this is a test"

			enc, err := a.Sign(msg)
			runtime.Must(err)

			err = a.Verify(enc, msg)
			runtime.Must(err)

			return meta.WithAttribute(ctx, "testMsg", meta.String(msg))
		}
		op = "verified keys"
	}

	opts := &runner.Options{Lifecycle: params.Lifecycle, Logger: params.Logger, Fn: fn}
	runner.Start("ed25519", op, opts)
}
