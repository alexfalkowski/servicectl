package feature

import (
	"context"

	"github.com/alexfalkowski/go-service/feature"
	"github.com/alexfalkowski/go-service/flags"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/servicectl/cmd/runner"
	"github.com/open-feature/go-sdk/openfeature"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// VerifyFlag defines wether we should verify the connection or not.
var VerifyFlag = flags.Bool()

// Params for feature.
type Params struct {
	fx.In

	Lifecycle fx.Lifecycle
	Logger    *zap.Logger
	Client    *openfeature.Client
}

// Start for feature.
func Start(params Params) {
	if !flags.IsSet(VerifyFlag) {
		return
	}

	fn := func(ctx context.Context) context.Context {
		err := feature.Ping(ctx, params.Client)
		runtime.Must(err)

		return ctx
	}

	opts := &runner.Options{
		Lifecycle: params.Lifecycle,
		Logger:    params.Logger,
		Fn:        fn,
	}

	runner.Start("feature", "verified connection", opts)
}
