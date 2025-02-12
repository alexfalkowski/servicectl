package grpc

import (
	sc "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/flags"
	"github.com/alexfalkowski/go-service/limiter"
	"github.com/alexfalkowski/go-service/transport/grpc"
	"github.com/alexfalkowski/go-service/transport/meta"
	"github.com/alexfalkowski/servicectl/internal/cmd"
	"github.com/alexfalkowski/servicectl/internal/cmd/runner"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Register for grpc.
func Register(command *sc.Command) {
	client := command.AddClient("grpc", "gRPC Server.", cmd.Module, limiter.Module, meta.Module, grpc.Module, fx.Invoke(start))

	flags.BoolVar(client, verify, "verify", "v", false, "verify server")
}

var verify = flags.Bool()

//nolint:gocritic
func start(lc fx.Lifecycle, logger *zap.Logger, _ *grpc.Server) {
	var (
		fn runner.StartFn
		op string
	)

	switch {
	case flags.IsBoolSet(verify):
		fn = runner.NoStart
		op = "started"
	}

	opts := &runner.Options{Lifecycle: lc, Logger: logger, Fn: fn}
	runner.Start("grpc", op, opts)
}
