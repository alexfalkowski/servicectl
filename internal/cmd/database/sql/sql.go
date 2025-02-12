package sql

import (
	"context"
	"errors"

	"github.com/alexfalkowski/go-service/flags"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/servicectl/internal/cmd/runner"
	"github.com/linxGnu/mssqlx"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// VerifyFlag defines wether we should verify the connection or not.
var VerifyFlag = flags.Bool()

// Params for sql.
type Params struct {
	fx.In

	Lifecycle fx.Lifecycle
	Logger    *zap.Logger
	DB        *mssqlx.DBs
}

// Start for sql.
//
//nolint:gocritic
func Start(lc fx.Lifecycle, logger *zap.Logger, db *mssqlx.DBs) {
	var (
		fn runner.StartFn
		op string
	)

	switch {
	case flags.IsBoolSet(VerifyFlag):
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
