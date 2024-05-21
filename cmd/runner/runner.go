package runner

import (
	"context"
	"fmt"

	"github.com/alexfalkowski/go-service/cmd"
	tz "github.com/alexfalkowski/go-service/telemetry/logger/zap"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// StartFn for cmd.
type StartFn func(context.Context) context.Context

// NoStart for cmd.
var NoStart = func(ctx context.Context) context.Context { return ctx }

// Options for runner.
type Options struct {
	Lifecycle fx.Lifecycle
	Logger    *zap.Logger
	Fn        StartFn
}

// Start the cmd.
func Start(name, operation string, opts *Options) {
	if opts.Fn == nil {
		return
	}

	cmd.Start(opts.Lifecycle, func(ctx context.Context) {
		ctx = opts.Fn(ctx)
		msg := fmt.Sprintf("%s: successfully %s", name, operation)

		opts.Logger.Info(msg, tz.Meta(ctx)...)
	})
}
