package aes

import (
	"context"

	sc "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/crypto/aes"
	"github.com/alexfalkowski/go-service/crypto/rand"
	"github.com/alexfalkowski/go-service/flags"
	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/go-service/token"
	"github.com/alexfalkowski/servicectl/internal/cmd"
	"github.com/alexfalkowski/servicectl/internal/cmd/os"
	"github.com/alexfalkowski/servicectl/internal/cmd/runner"
	"github.com/alexfalkowski/servicectl/internal/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Register for aes.
func Register(command *sc.Command) {
	flags := flags.NewFlagSet("aes")

	command.RegisterInput(flags, "")
	rotate = flags.BoolP("rotate", "r", false, "rotate key")
	verify = flags.BoolP("verify", "v", false, "verify key")

	command.AddClient("aes", "AES crypto.", flags, cmd.Module, token.Module, fx.Invoke(start))
}

var (
	rotate = flags.Bool()
	verify = flags.Bool()
)

func start(lc fx.Lifecycle, logger *zap.Logger, rand *rand.Generator, gen *aes.Generator, cfg *config.Config) {
	var (
		fn runner.StartFn
		op string
	)

	switch {
	case flags.IsBoolSet(rotate):
		fn = func(ctx context.Context) context.Context {
			k, err := gen.Generate()
			runtime.Must(err)

			err = os.WriteBase64File(cfg.Crypto.AES.Key, []byte(k))
			runtime.Must(err)

			return ctx
		}
		op = "rotated key"
	case flags.IsBoolSet(verify):
		fn = func(ctx context.Context) context.Context {
			a, err := aes.NewCipher(rand, cfg.Crypto.AES)
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
