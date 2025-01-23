package redis

import (
	"context"

	"github.com/alexfalkowski/go-service/flags"
	"github.com/alexfalkowski/go-service/redis"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/servicectl/cmd/runner"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// VerifyFlag defines wether we should verify the connection or not.
var VerifyFlag = flags.Bool()

// Start for redis.
//
//nolint:gocritic
func Start(lc fx.Lifecycle, logger *zap.Logger, client redis.Client) {
	var (
		fn runner.StartFn
		op string
	)

	switch {
	case flags.IsBoolSet(VerifyFlag):
		fn = func(ctx context.Context) context.Context {
			cmd := client.Ping(ctx)
			runtime.Must(cmd.Err())

			return ctx
		}
		op = "verified connection"
	}

	opts := &runner.Options{Lifecycle: lc, Logger: logger, Fn: fn}
	runner.Start("redis", op, opts)
}
