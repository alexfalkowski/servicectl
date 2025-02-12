package rsa

import (
	"context"

	sc "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/crypto/rand"
	"github.com/alexfalkowski/go-service/crypto/rsa"
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

// Register for rsa.
func Register(command *sc.Command) {
	client := command.AddClient("rsa", "RSA crypto.", cmd.Module, fx.Invoke(start))

	flags.BoolVar(client, rotate, "rotate", "r", false, "rotate key")
	flags.BoolVar(client, verify, "verify", "v", false, "verify key")
}

var (
	rotate = flags.Bool()

	verify = flags.Bool()
)

func start(lc fx.Lifecycle, logger *zap.Logger, rand *rand.Generator, gen *rsa.Generator, cfg *config.Config) {
	var (
		fn runner.StartFn
		op string
	)

	switch {
	case flags.IsBoolSet(rotate):
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
	case flags.IsBoolSet(verify):
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
