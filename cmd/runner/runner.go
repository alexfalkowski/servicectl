package runner

import (
	"context"
	"fmt"

	"github.com/alexfalkowski/go-service/runtime"
	tz "github.com/alexfalkowski/go-service/telemetry/logger/zap"
	"github.com/alexfalkowski/servicectl/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// ModifyFn for cmd.
type ModifyFn func(context.Context, *config.Config) context.Context

// NoModify for cmd.
var NoModify = func(ctx context.Context, _ *config.Config) context.Context { return ctx }

// Options for runner.
type Options struct {
	Lifecycle fx.Lifecycle
	Logger    *zap.Logger
	Fn        ModifyFn
}

// Run the cmd.
func Run(name, operation string, opts *Options) {
	if opts.Fn == nil {
		return
	}

	opts.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) (err error) {
			defer func() {
				if r := recover(); r != nil {
					err = runtime.ConvertRecover(r)
				}
			}()

			msg := fmt.Sprintf("%s: successfully %s", name, operation)
			opts.Logger.Info(msg, tz.Meta(ctx)...)

			return
		},
	})
}
