package feature

import (
	"context"

	"github.com/alexfalkowski/go-service/feature"
	"github.com/alexfalkowski/go-service/flags"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/servicectl/cmd/runner"
	"github.com/alexfalkowski/servicectl/config"
	"github.com/open-feature/go-sdk/openfeature"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// VerifyFlag defines wether we should verify the connection or not.
var VerifyFlag = flags.Bool()

// RunParams for feature.
type RunParams struct {
	fx.In

	Lifecycle fx.Lifecycle
	Logger    *zap.Logger
	Client    *openfeature.Client
}

// Run for feature.
func Run(params RunParams) {
	if !flags.IsSet(VerifyFlag) {
		return
	}

	fn := func(ctx context.Context, _ *config.Config) context.Context {
		err := feature.Ping(ctx, params.Client)
		runtime.Must(err)

		return ctx
	}

	opts := &runner.Options{
		Lifecycle: params.Lifecycle,
		Logger:    params.Logger,
		Fn:        fn,
	}

	runner.Run("feature", "verified connection", opts)
}
