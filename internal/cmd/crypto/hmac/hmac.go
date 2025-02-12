package hmac

import (
	"context"

	sc "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/crypto/hmac"
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

// Register for hmac.
func Register(command *sc.Command) {
	client := command.AddClient("hmac", "HMAC crypto.", cmd.Module, fx.Invoke(start))

	flags.BoolVar(client, rotate, "rotate", "r", false, "rotate key")
	flags.BoolVar(client, verify, "verify", "v", false, "verify key")
}

var (
	rotate = flags.Bool()

	verify = flags.Bool()
)

func start(lc fx.Lifecycle, logger *zap.Logger, gen *hmac.Generator, cfg *config.Config) {
	var (
		fn runner.StartFn
		op string
	)

	switch {
	case flags.IsBoolSet(rotate):
		fn = func(ctx context.Context) context.Context {
			k, err := gen.Generate()
			runtime.Must(err)

			err = os.WriteBase64File(cfg.Crypto.HMAC.Key, []byte(k))
			runtime.Must(err)

			return ctx
		}
		op = "rotated key"
	case flags.IsBoolSet(verify):
		fn = func(ctx context.Context) context.Context {
			a, err := hmac.NewSigner(cfg.Crypto.HMAC)
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

	opts := &runner.Options{Lifecycle: lc, Logger: logger, Fn: fn}
	runner.Start("hmac", op, opts)
}
