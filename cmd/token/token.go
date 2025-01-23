package token

import (
	"context"

	"github.com/alexfalkowski/go-service/crypto/rand"
	"github.com/alexfalkowski/go-service/flags"
	"github.com/alexfalkowski/go-service/runtime"
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
func Start(lc fx.Lifecycle, logger *zap.Logger, rand *rand.Generator, cfg *config.Config) {
	var (
		fn runner.StartFn
		op string
	)

	switch {
	case flags.IsBoolSet(RotateFlag):
		fn = func(ctx context.Context) context.Context {
			k, err := rand.GenerateString(64)
			runtime.Must(err)

			err = os.WriteBase64File(cfg.Token.Key, []byte(k))
			runtime.Must(err)

			return ctx
		}
		op = "rotated key"
	}

	opts := &runner.Options{Lifecycle: lc, Logger: logger, Fn: fn}
	runner.Start("token", op, opts)
}
