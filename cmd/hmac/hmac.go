package hmac

import (
	"context"

	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/crypto/hmac"
	"github.com/alexfalkowski/go-service/flags"
	"github.com/alexfalkowski/go-service/marshaller"
	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/servicectl/cmd/runner"
	"github.com/alexfalkowski/servicectl/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// RotateFlag defines wether we should rotate keys or not.
var RotateFlag = flags.Bool()

// RunCommandParams for HMAC.
type RunCommandParams struct {
	fx.In

	Lifecycle    fx.Lifecycle
	OutputConfig *cmd.OutputConfig
	Map          *marshaller.Map
	Config       *config.Config
	Logger       *zap.Logger
}

// RunCommand for HMAC.
func RunCommand(params RunCommandParams) {
	if !flags.IsSet(RotateFlag) {
		return
	}

	fn := func(ctx context.Context, c *config.Config) context.Context {
		k, err := hmac.Generate()
		runtime.Must(err)

		c.Crypto.HMAC.Key = k

		return meta.WithAttribute(ctx, "key", meta.String(k))
	}

	p := &runner.Params{
		Lifecycle:    params.Lifecycle,
		OutputConfig: params.OutputConfig,
		Map:          params.Map,
		Config:       params.Config,
		Logger:       params.Logger,
		Fn:           fn,
	}

	runner.Run("hmac", "rotated key", p)
}