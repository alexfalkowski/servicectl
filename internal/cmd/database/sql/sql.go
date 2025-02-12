package sql

import (
	"context"
	"errors"

	sc "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/database/sql/pg"
	"github.com/alexfalkowski/go-service/flags"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/servicectl/internal/cmd"
	"github.com/alexfalkowski/servicectl/internal/cmd/runner"
	"github.com/linxGnu/mssqlx"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Register for sql.
func Register(command *sc.Command) {
	client := command.AddClient("pg", "Postgres DB.", cmd.Module, pg.Module, fx.Invoke(start))

	flags.BoolVar(client, verify, "verify", "v", false, "verify connection")
}

var verify = flags.Bool()

//nolint:gocritic
func start(lc fx.Lifecycle, logger *zap.Logger, db *mssqlx.DBs) {
	var (
		fn runner.StartFn
		op string
	)

	switch {
	case flags.IsBoolSet(verify):
		fn = func(ctx context.Context) context.Context {
			err := errors.Join(db.Ping()...)
			runtime.Must(err)

			return ctx
		}
		op = "verified connection"
	}

	opts := &runner.Options{Lifecycle: lc, Logger: logger, Fn: fn}
	runner.Start("pg", op, opts)
}
