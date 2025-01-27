package token

import (
	"context"

	"github.com/alexfalkowski/go-service/crypto/rand"
	"github.com/alexfalkowski/go-service/env"
	"github.com/alexfalkowski/go-service/flags"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/go-service/token"
	"github.com/alexfalkowski/servicectl/cmd/os"
	"github.com/alexfalkowski/servicectl/cmd/runner"
	"github.com/alexfalkowski/servicectl/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// RotateFlag defines wether we should rotate the key or not.
var RotateFlag = flags.Bool()

// Start for token.
//
//nolint:gocritic
func Start(lc fx.Lifecycle, logger *zap.Logger, rand *rand.Generator, cfg *config.Config, name env.Name) {
	var (
		fn runner.StartFn
		op string
	)

	switch {
	case flags.IsBoolSet(RotateFlag):
		fn = func(ctx context.Context) context.Context {
			switch cfg.Token.Kind {
			case "key":
				k, err := rand.GenerateString(64)
				runtime.Must(err)

				err = os.WriteBase64File(cfg.Token.Secret, []byte(k))
				runtime.Must(err)
			case "token":
				k, err := token.Generate(name, rand)
				runtime.Must(err)

				err = os.WriteFile(cfg.Token.Secret, []byte(k))
				runtime.Must(err)
			}

			return ctx
		}
		op = "rotated " + cfg.Token.Kind
	}

	opts := &runner.Options{Lifecycle: lc, Logger: logger, Fn: fn}
	runner.Start("token", op, opts)
}
