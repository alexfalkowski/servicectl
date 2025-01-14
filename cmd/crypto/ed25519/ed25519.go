package ed25519

import (
	"context"

	"github.com/alexfalkowski/go-service/crypto/ed25519"
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
	// RotateFlag defines wether we should rotate the keys or not.
	RotateFlag = flags.Bool()

	// VerifyFlag defines wether we should verify the keys or not.
	VerifyFlag = flags.Bool()
)

// Start for AES.
func Start(lc fx.Lifecycle, logger *zap.Logger, gen *ed25519.Generator, cfg *config.Config) {
	var (
		fn runner.StartFn
		op string
	)

	switch {
	case flags.IsBoolSet(RotateFlag):
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
	case flags.IsBoolSet(VerifyFlag):
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
