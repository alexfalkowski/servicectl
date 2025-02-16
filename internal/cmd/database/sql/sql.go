package sql

import (
	"context"
	"errors"

	sc "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/database/sql/pg"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/servicectl/internal/cmd"
	"github.com/alexfalkowski/servicectl/internal/cmd/flags"
	"github.com/alexfalkowski/servicectl/internal/cmd/runner"
	"github.com/linxGnu/mssqlx"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Register for sql.
func Register(command *sc.Command) {
	flags := sc.NewFlagSet("pg")

	flags.AddInput("")
	flags.BoolP("verify", "v", false, "verify connection")

	command.AddClient("pg", "Postgres DB.", flags, cmd.Module, pg.Module, fx.Invoke(Start))
}

// StartParams for sql.
type StartParams struct {
	fx.In

	Set       *sc.FlagSet
	Lifecycle fx.Lifecycle
	Logger    *zap.Logger
	DB        *mssqlx.DBs
}

// Start
//
//nolint:gocritic
func Start(params StartParams) {
	var (
		fn runner.StartFn
		op string
	)

	switch {
	case flags.IsSet(params.Set, "verify"):
		fn = func(ctx context.Context) context.Context {
			err := errors.Join(params.DB.Ping()...)
			runtime.Must(err)

			return ctx
		}
		op = "verified connection"
	}

	opts := &runner.Options{Lifecycle: params.Lifecycle, Logger: params.Logger, Fn: fn}
	runner.Start("pg", op, opts)
}
