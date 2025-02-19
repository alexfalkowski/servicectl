package hmac

import (
	"context"

	sc "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/crypto/hmac"
	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/go-service/telemetry/logger"
	"github.com/alexfalkowski/servicectl/internal/cmd"
	"github.com/alexfalkowski/servicectl/internal/cmd/flags"
	"github.com/alexfalkowski/servicectl/internal/cmd/os"
	"github.com/alexfalkowski/servicectl/internal/cmd/runner"
	"github.com/alexfalkowski/servicectl/internal/config"
	"go.uber.org/fx"
)

// Register for hmac.
func Register(command *sc.Command) {
	flags := sc.NewFlagSet("hmac")

	flags.AddInput("")
	flags.BoolP("rotate", "r", false, "rotate key")
	flags.BoolP("verify", "v", false, "verify key")

	command.AddClient("hmac", "HMAC crypto.", flags, cmd.Module, fx.Invoke(Start))
}

// StartParams for hmac.
type StartParams struct {
	fx.In

	Set       *sc.FlagSet
	Lifecycle fx.Lifecycle
	Logger    *logger.Logger
	Generator *hmac.Generator
	Config    *config.Config
}

// Start for hmac.
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

			err = os.WriteBase64File(params.Config.Crypto.HMAC.Key, []byte(k))
			runtime.Must(err)

			return ctx
		}
		op = "rotated key"
	case flags.IsSet(params.Set, "verify"):
		fn = func(ctx context.Context) context.Context {
			a, err := hmac.NewSigner(params.Config.Crypto.HMAC)
			runtime.Must(err)

			msg := "this is a test"

			enc, err := a.Sign(msg)
			runtime.Must(err)

			err = a.Verify(enc, msg)
			runtime.Must(err)

			return meta.WithAttribute(ctx, "testMsg", meta.String(msg))
		}
		op = "verified key"
	}

	opts := &runner.Options{Lifecycle: params.Lifecycle, Logger: params.Logger, Fn: fn}
	runner.Start("hmac", op, opts)
}
