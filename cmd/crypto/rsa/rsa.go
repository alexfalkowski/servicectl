package rsa

import (
	"context"

	"github.com/alexfalkowski/go-service/crypto/rand"
	"github.com/alexfalkowski/go-service/crypto/rsa"
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
func Start(lc fx.Lifecycle, logger *zap.Logger, rand *rand.Generator, gen *rsa.Generator, cfg *config.Config) {
	var (
		fn runner.StartFn
		op string
	)

	switch {
	case flags.IsBoolSet(RotateFlag):
		fn = func(ctx context.Context) context.Context {
			pub, pri, err := gen.Generate()
			runtime.Must(err)

			err = os.WriteFile(cfg.Crypto.RSA.Public, []byte(pub))
			runtime.Must(err)

			err = os.WriteFile(cfg.Crypto.RSA.Private, []byte(pri))
			runtime.Must(err)

			return ctx
		}
		op = "rotated keys"
	case flags.IsBoolSet(VerifyFlag):
		fn = func(ctx context.Context) context.Context {
			a, err := rsa.NewCipher(rand, cfg.Crypto.RSA)
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

	opts := &runner.Options{Lifecycle: lc, Logger: logger, Fn: fn}
	runner.Start("rsa", op, opts)
}
