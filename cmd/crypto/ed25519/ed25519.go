package ed25519

import (
	"context"

	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/crypto/ed25519"
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
	// RotateFlag defines wether we should rotate the keys or not.
	RotateFlag = flags.Bool()

	// VerifyFlag defines wether we should verify the keys or not.
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
			pub, pri, err := ed25519.Generate()
			runtime.Must(err)

			c.Crypto.Ed25519.Public = pub
			c.Crypto.Ed25519.Private = pri

			ctx = meta.WithAttribute(ctx, "public", meta.String(pub))
			ctx = meta.WithAttribute(ctx, "private", meta.String(pri))

			return ctx
		}
		op = "rotated keys"
		oc = params.OutputConfig
	case flags.IsSet(VerifyFlag):
		fn = func(ctx context.Context, c *config.Config) context.Context {
			a, err := ed25519.NewAlgo(c.Crypto.Ed25519)
			runtime.Must(err)

			msg := "this is a test"
			enc := a.Generate(msg)

			err = a.Compare(enc, msg)
			runtime.Must(err)

			return meta.WithAttribute(ctx, "testMsg", meta.String(msg))
		}
		op = "verified keys"
	}

	opts := &runner.Options{
		Lifecycle:    params.Lifecycle,
		OutputConfig: oc,
		Map:          params.Map,
		Config:       params.Config,
		Logger:       params.Logger,
		Fn:           fn,
	}

	runner.Run("ed25519", op, opts)
}
