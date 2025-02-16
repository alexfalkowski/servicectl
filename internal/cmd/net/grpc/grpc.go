package grpc

import (
	sc "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/flags"
	"github.com/alexfalkowski/go-service/limiter"
	"github.com/alexfalkowski/go-service/transport/grpc"
	"github.com/alexfalkowski/go-service/transport/meta"
	"github.com/alexfalkowski/servicectl/internal/cmd"
	cf "github.com/alexfalkowski/servicectl/internal/cmd/flags"
	"github.com/alexfalkowski/servicectl/internal/cmd/runner"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Register for grpc.
func Register(command *sc.Command) {
	flags := flags.NewFlagSet("grpc")

	flags.AddInput("")
	flags.BoolP("verify", "v", false, "verify server")

	command.AddClient("grpc", "gRPC Server.", flags, cmd.Module, limiter.Module, meta.Module, grpc.Module, fx.Invoke(Start))
}

// StartParams for grpc.
type StartParams struct {
	fx.In

	Set       *flags.FlagSet
	Lifecycle fx.Lifecycle
	Logger    *zap.Logger
	Server    *grpc.Server
}

// Start for grpc.
//
//nolint:gocritic
func Start(params StartParams) {
	var (
		fn runner.StartFn
		op string
	)

	switch {
	case cf.IsSet(params.Set, "verify"):
		fn = runner.NoStart
		op = "started"
	}

	opts := &runner.Options{Lifecycle: params.Lifecycle, Logger: params.Logger, Fn: fn}
	runner.Start("grpc", op, opts)
}
