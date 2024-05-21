package http

import (
	"github.com/alexfalkowski/go-service/flags"
	"github.com/alexfalkowski/go-service/transport/http"
	"github.com/alexfalkowski/servicectl/cmd/runner"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// VerifyFlag defines wether we should verify or not.
var VerifyFlag = flags.Bool()

// Start for grpc.
func Start(lc fx.Lifecycle, logger *zap.Logger, _ *http.Server) {
	if !flags.IsSet(VerifyFlag) {
		return
	}

	opts := &runner.Options{Lifecycle: lc, Logger: logger, Fn: runner.NoStart}
	runner.Start("http", "started", opts)
}
