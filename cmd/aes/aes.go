package aes

import (
	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/crypto/aes"
	"github.com/alexfalkowski/go-service/marshaller"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/servicectl/cmd/runner"
	"github.com/alexfalkowski/servicectl/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Rotate the keys.
var Rotate = cmd.Bool()

// RunCommandParams for AES.
type RunCommandParams struct {
	fx.In

	Lifecycle    fx.Lifecycle
	OutputConfig *cmd.OutputConfig
	Factory      *marshaller.Factory
	Config       *config.Config
	Logger       *zap.Logger
}

// RunCommand for AES.
func RunCommand(params RunCommandParams) {
	if !*Rotate {
		return
	}

	fn := func(c *config.Config, l *zap.Logger) {
		k, err := aes.Generate()
		runtime.Must(err)

		c.Crypto.AES.Key = k

		l.Info("rotated aes key")
	}

	runner.Run(&runner.Params{
		Lifecycle:    params.Lifecycle,
		OutputConfig: params.OutputConfig,
		Factory:      params.Factory,
		Config:       params.Config,
		Logger:       params.Logger,
		Fn:           fn,
	})
}
