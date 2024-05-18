package sql

import (
	"context"
	"errors"

	"github.com/alexfalkowski/go-service/flags"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/servicectl/cmd/runner"
	"github.com/alexfalkowski/servicectl/config"
	"github.com/linxGnu/mssqlx"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// VerifyFlag defines wether we should verify the connection or not.
var VerifyFlag = flags.Bool()

// RunParams for sql.
type RunParams struct {
	fx.In

	Lifecycle fx.Lifecycle
	Logger    *zap.Logger
	DB        *mssqlx.DBs
}

// Run for sql.
func Run(params RunParams) {
	if !flags.IsSet(VerifyFlag) {
		return
	}

	fn := func(ctx context.Context, _ *config.Config) context.Context {
		err := errors.Join(params.DB.Ping()...)
		runtime.Must(err)

		return ctx
	}

	opts := &runner.Options{
		Lifecycle: params.Lifecycle,
		Logger:    params.Logger,
		Fn:        fn,
	}

	runner.Run("pg", "verified connection", opts)
}
