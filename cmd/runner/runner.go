package runner

import (
	"context"
	"io/fs"

	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/marshaller"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/servicectl/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// ModifyFn for cmd.
type ModifyFn func(*config.Config, *zap.Logger)

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
func Run(params *Params) {
	params.Lifecycle.Append(fx.Hook{
		OnStart: func(_ context.Context) (err error) {
			defer func() {
				if r := recover(); r != nil {
					err = runtime.ConvertRecover(r)
				}
			}()

			writeOutConfig(params)

			return
		},
	})
}

func writeOutConfig(params *Params) {
	params.Fn(params.Config, params.Logger)

	m := params.Map.Get(params.OutputConfig.Kind())

	d, err := m.Marshal(params.Config)
	runtime.Must(err)

	runtime.Must(params.OutputConfig.Write(d, fs.FileMode(0o600)))
}
