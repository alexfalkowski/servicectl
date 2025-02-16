package ed25519

import (
	"context"

	sc "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/crypto/ed25519"
	"github.com/alexfalkowski/go-service/flags"
	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/servicectl/internal/cmd"
	"github.com/alexfalkowski/servicectl/internal/cmd/os"
	"github.com/alexfalkowski/servicectl/internal/cmd/runner"
	"github.com/alexfalkowski/servicectl/internal/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Register for ed25519.
func Register(command *sc.Command) {
	flags := flags.NewFlagSet("ed25519")

	command.RegisterInput(flags, "")
	rotate = flags.BoolP("rotate", "r", false, "rotate key")
	verify = flags.BoolP("verify", "v", false, "verify key")

	command.AddClient("ed25519", "Ed25519 crypto.", flags, cmd.Module, fx.Invoke(start))
}

var (
	rotate = flags.Bool()
	verify = flags.Bool()
)

func start(lc fx.Lifecycle, logger *zap.Logger, gen *ed25519.Generator, cfg *config.Config) {
	var (
		fn runner.StartFn
		op string
	)

	switch {
	case flags.IsBoolSet(rotate):
		fn = func(ctx context.Context) context.Context {
			pub, pri, err := gen.Generate()
			runtime.Must(err)

			err = os.WriteFile(cfg.Crypto.Ed25519.Public, []byte(pub))
			runtime.Must(err)

			err = os.WriteFile(cfg.Crypto.Ed25519.Private, []byte(pri))
			runtime.Must(err)

			return ctx
		}
		op = "rotated keys"
	case flags.IsBoolSet(verify):
		fn = func(ctx context.Context) context.Context {
			a, err := ed25519.NewSigner(cfg.Crypto.Ed25519)
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

	opts := &runner.Options{Lifecycle: lc, Logger: logger, Fn: fn}
	runner.Start("ed25519", op, opts)
}
