package sql

import (
	"context"
	"errors"

	"github.com/alexfalkowski/go-service/flags"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/servicectl/cmd/runner"
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
func Start(lc fx.Lifecycle, logger *zap.Logger, db *mssqlx.DBs) {
	if !flags.IsBoolSet(VerifyFlag) {
		return
	}

	fn := func(ctx context.Context) context.Context {
		err := errors.Join(db.Ping()...)
		runtime.Must(err)

		return ctx
	}

	opts := &runner.Options{Lifecycle: lc, Logger: logger, Fn: fn}
	runner.Start("pg", "verified connection", opts)
}
