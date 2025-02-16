package token

import (
	"context"

	sc "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/crypto/rand"
	"github.com/alexfalkowski/go-service/env"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/go-service/token"
	"github.com/alexfalkowski/servicectl/internal/cmd"
	"github.com/alexfalkowski/servicectl/internal/cmd/flags"
	"github.com/alexfalkowski/servicectl/internal/cmd/os"
	"github.com/alexfalkowski/servicectl/internal/cmd/runner"
	"github.com/alexfalkowski/servicectl/internal/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Register for token.
func Register(command *sc.Command) {
	flags := sc.NewFlagSet("token")

	flags.AddInput("")
	flags.BoolP("rotate", "r", false, "rotate secret")

	command.AddClient("token", "Security tokens.", flags, cmd.Module, token.Module, fx.Invoke(Start))
}

// StartParams for token.
type StartParams struct {
	fx.In
	Lifecycle fx.Lifecycle
	Set       *sc.FlagSet
	Logger    *zap.Logger
	Generator *rand.Generator
	Config    *config.Config
	Name      env.Name
}

// Start for token.
//
//nolint:gocritic
func Start(params StartParams) {
	var (
		fn runner.StartFn
		op string
	)

	switch {
	case flags.IsSet(params.Set, "rotate"):
		fn = func(ctx context.Context) context.Context {
			switch params.Config.Token.Kind {
			case "key":
				k, err := params.Generator.GenerateString(64)
				runtime.Must(err)

				err = os.WriteBase64File(params.Config.Token.Secret, []byte(k))
				runtime.Must(err)
			case "token":
				k, err := token.Generate(params.Name, params.Generator)
				runtime.Must(err)

				err = os.WriteFile(params.Config.Token.Secret, []byte(k))
				runtime.Must(err)
			}

			return ctx
		}
		op = "rotated " + params.Config.Token.Kind
	}

	opts := &runner.Options{Lifecycle: params.Lifecycle, Logger: params.Logger, Fn: fn}
	runner.Start("token", op, opts)
}
