package runner

import (
	"context"
	"fmt"
	"io/fs"

	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/marshaller"
	"github.com/alexfalkowski/go-service/runtime"
	tz "github.com/alexfalkowski/go-service/telemetry/logger/zap"
	"github.com/alexfalkowski/servicectl/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// ModifyFn for cmd.
type ModifyFn func(context.Context, *config.Config) context.Context

// Params for runner.
type Params struct {
	Lifecycle    fx.Lifecycle
	OutputConfig *cmd.OutputConfig
	Map          *marshaller.Map
	Config       *config.Config
	Logger       *zap.Logger
	Fn           ModifyFn
}

// Run the cmd.
func Run(name, operation string, params *Params) {
	params.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) (err error) {
			defer func() {
				if r := recover(); r != nil {
					err = runtime.ConvertRecover(r)
				}
			}()

			ctx = writeOutConfig(ctx, params)
			msg := fmt.Sprintf("%s: successfully %s", name, operation)
			params.Logger.Info(msg, tz.Meta(ctx)...)

			return
		},
	})
}

func writeOutConfig(ctx context.Context, params *Params) context.Context {
	ctx = params.Fn(ctx, params.Config)

	m := params.Map.Get(params.OutputConfig.Kind())

	d, err := m.Marshal(params.Config)
	runtime.Must(err)

	runtime.Must(params.OutputConfig.Write(d, fs.FileMode(0o600)))

	return ctx
}
