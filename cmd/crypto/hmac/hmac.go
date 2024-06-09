package hmac

import (
	"context"

	"github.com/alexfalkowski/go-service/crypto/hmac"
	"github.com/alexfalkowski/go-service/flags"
	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/servicectl/cmd/os"
	"github.com/alexfalkowski/servicectl/cmd/runner"
	"github.com/alexfalkowski/servicectl/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var (
	// RotateFlag defines wether we should rotate the key or not.
	RotateFlag = flags.Bool()

	// VerifyFlag defines wether we should verify the key or not.
	VerifyFlag = flags.Bool()
)

// Start for HMAC.
func Start(lc fx.Lifecycle, logger *zap.Logger, cfg *config.Config) {
	var (
		fn runner.StartFn
		op string
	)

	switch {
	case flags.IsSet(RotateFlag):
		fn = func(ctx context.Context) context.Context {
			k, err := hmac.Generate()
			runtime.Must(err)

			os.WriteFile(string(cfg.Crypto.HMAC.Key), []byte(k))

			return ctx
		}
		op = "rotated key"
	case flags.IsSet(VerifyFlag):
		fn = func(ctx context.Context) context.Context {
			a, err := hmac.NewAlgo(cfg.Crypto.HMAC)
			runtime.Must(err)

			msg := "this is a test"
			enc := a.Sign(msg)

			err = a.Verify(enc, msg)
			runtime.Must(err)

			return meta.WithAttribute(ctx, "testMsg", meta.String(msg))
		}
		op = "verified key"
	}

	opts := &runner.Options{Lifecycle: lc, Logger: logger, Fn: fn}
	runner.Start("hmac", op, opts)
}
