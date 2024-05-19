package aes

import (
	"context"

	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/crypto/aes"
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

// RunParams for AES.
type RunParams struct {
	fx.In

	Lifecycle    fx.Lifecycle
	OutputConfig *cmd.OutputConfig
	Map          *marshaller.Map
	Config       *config.Config
	Logger       *zap.Logger
}

// Run for AES.
func Run(params RunParams) {
	var (
		fn runner.ModifyFn
		op string
		oc *cmd.OutputConfig
	)

	switch {
	case flags.IsSet(RotateFlag):
		fn = func(ctx context.Context, c *config.Config) context.Context {
			k, err := aes.Generate()
			runtime.Must(err)

			c.Crypto.AES.Key = k

			return meta.WithAttribute(ctx, "key", meta.String(k))
		}
		op = "rotated key"
		oc = params.OutputConfig
	case flags.IsSet(VerifyFlag):
		fn = func(ctx context.Context, c *config.Config) context.Context {
			a, err := aes.NewAlgo(c.Crypto.AES)
			runtime.Must(err)

			msg := "this is a test"
			enc, err := a.Encrypt(msg)
			runtime.Must(err)

			_, err = a.Decrypt(enc)
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

	runner.Run("aes", op, opts)
}
