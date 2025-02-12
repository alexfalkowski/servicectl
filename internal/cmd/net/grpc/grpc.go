package grpc

import (
	"github.com/alexfalkowski/go-service/flags"
	"github.com/alexfalkowski/go-service/transport/grpc"
	"github.com/alexfalkowski/servicectl/internal/cmd/runner"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// VerifyFlag defines wether we should verify or not.
//

var VerifyFlag = flags.Bool()

// Start for grpc.
//
//nolint:gocritic
func Start(lc fx.Lifecycle, logger *zap.Logger, _ *grpc.Server) {
	var (
		fn runner.StartFn
		op string
	)

	switch {
	case flags.IsBoolSet(VerifyFlag):
		fn = runner.NoStart
		op = "started"
	}

	opts := &runner.Options{Lifecycle: lc, Logger: logger, Fn: fn}
	runner.Start("grpc", op, opts)
}
