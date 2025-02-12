package token

import (
	"context"

	sc "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/crypto/rand"
	"github.com/alexfalkowski/go-service/env"
	"github.com/alexfalkowski/go-service/flags"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/go-service/token"
	"github.com/alexfalkowski/servicectl/internal/cmd"
	"github.com/alexfalkowski/servicectl/internal/cmd/os"
	"github.com/alexfalkowski/servicectl/internal/cmd/runner"
	"github.com/alexfalkowski/servicectl/internal/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Register for token.
func Register(command *sc.Command) {
	client := command.AddClient("token", "Security tokens.", cmd.Module, token.Module, fx.Invoke(start))

	flags.BoolVar(client, rotate, "rotate", "r", false, "rotate secret")
}

var rotate = flags.Bool()

//nolint:gocritic
func start(lc fx.Lifecycle, logger *zap.Logger, rand *rand.Generator, cfg *config.Config, name env.Name) {
	var (
		fn runner.StartFn
		op string
	)

	switch {
	case flags.IsBoolSet(rotate):
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
