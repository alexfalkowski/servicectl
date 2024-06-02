package aes

import (
	"context"

	"github.com/alexfalkowski/go-service/crypto/aes"
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

// Start for AES.
func Start(lc fx.Lifecycle, logger *zap.Logger, cfg *config.Config) {
	var (
		fn runner.StartFn
		op string
	)

	switch {
	case flags.IsSet(RotateFlag):
		fn = func(ctx context.Context) context.Context {
			k, err := aes.Generate()
			runtime.Must(err)

			os.WriteFile(string(cfg.Crypto.AES.Key), []byte(k))

			return ctx
		}
		op = "rotated key"
	case flags.IsSet(VerifyFlag):
		fn = func(ctx context.Context) context.Context {
			a, err := aes.NewAlgo(cfg.Crypto.AES)
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

	opts := &runner.Options{Lifecycle: lc, Logger: logger, Fn: fn}
	runner.Start("aes", op, opts)
}
