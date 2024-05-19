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

var (
	// RotateFlag defines wether we should rotate the key or not.
	RotateFlag = flags.Bool()

	// VerifyFlag defines wether we should verify the key or not.
	VerifyFlag = flags.Bool()
)

// RunParams for HMAC.
type RunParams struct {
	fx.In

	Lifecycle    fx.Lifecycle
	OutputConfig *cmd.OutputConfig
	Map          *marshaller.Map
	Config       *config.Config
	Logger       *zap.Logger
}

// Run for HMAC.
func Run(params RunParams) {
	var (
		fn runner.ModifyFn
		op string
		oc *cmd.OutputConfig
	)

	switch {
	case flags.IsSet(RotateFlag):
		fn = func(ctx context.Context, c *config.Config) context.Context {
			k, err := hmac.Generate()
			runtime.Must(err)

			c.Crypto.HMAC.Key = k

			return meta.WithAttribute(ctx, "key", meta.String(k))
		}
		op = "rotated key"
		oc = params.OutputConfig
	case flags.IsSet(VerifyFlag):
		fn = func(ctx context.Context, c *config.Config) context.Context {
			a, err := hmac.NewAlgo(c.Crypto.HMAC)
			runtime.Must(err)

			msg := "this is a test"
			enc := a.Generate(msg)

			err = a.Compare(enc, msg)
			runtime.Must(err)

			return meta.WithAttribute(ctx, "testMsg", meta.String(msg))
		}
		op = "verified key"
	}

	opts := &runner.Options{
		Lifecycle:    params.Lifecycle,
		OutputConfig: oc,
		Map:          params.Map,
		Config:       params.Config,
		Logger:       params.Logger,
		Fn:           fn,
	}

	runner.Run("hmac", op, opts)
}
