package rsa

import (
	"context"

	"github.com/alexfalkowski/go-service/crypto/rsa"
	"github.com/alexfalkowski/go-service/flags"
	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/go-service/runtime"
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

// RunParams for AES.
type RunParams struct {
	fx.In

	Lifecycle fx.Lifecycle
	Logger    *zap.Logger
}

// Run for AES.
func Run(params RunParams) {
	var (
		fn runner.ModifyFn
		op string
	)

	switch {
	case flags.IsSet(RotateFlag):
		fn = func(ctx context.Context, c *config.Config) context.Context {
			pub, pri, err := rsa.Generate()
			runtime.Must(err)

			c.Crypto.RSA.Public = pub
			c.Crypto.RSA.Private = pri

			ctx = meta.WithAttribute(ctx, "public", meta.String(pub))
			ctx = meta.WithAttribute(ctx, "private", meta.String(pri))

			return ctx
		}
		op = "rotated keys"
	case flags.IsSet(VerifyFlag):
		fn = func(ctx context.Context, c *config.Config) context.Context {
			a, err := rsa.NewAlgo(c.Crypto.RSA)
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

	opts := &runner.Options{
		Lifecycle: params.Lifecycle,
		Logger:    params.Logger,
		Fn:        fn,
	}

	runner.Run("rsa", op, opts)
}
