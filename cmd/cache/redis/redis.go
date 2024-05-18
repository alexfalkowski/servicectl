package redis

import (
	"context"

	"github.com/alexfalkowski/go-service/flags"
	"github.com/alexfalkowski/go-service/redis"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/servicectl/cmd/runner"
	"github.com/alexfalkowski/servicectl/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// VerifyFlag defines wether we should verify the connection or not.
var VerifyFlag = flags.Bool()

// RunParams for redis.
type RunParams struct {
	fx.In

	Lifecycle fx.Lifecycle
	Logger    *zap.Logger
	Client    redis.Client
}

// Run for redis.
func Run(params RunParams) {
	if !flags.IsSet(VerifyFlag) {
		return
	}

	fn := func(ctx context.Context, _ *config.Config) context.Context {
		cmd := params.Client.Ping(ctx)
		runtime.Must(cmd.Err())

		return ctx
	}

	opts := &runner.Options{
		Lifecycle: params.Lifecycle,
		Logger:    params.Logger,
		Fn:        fn,
	}

	runner.Run("redis", "verified connection", opts)
}
