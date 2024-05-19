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

// RunParams for grpc.
type RunParams struct {
	fx.In

	Lifecycle fx.Lifecycle
	Logger    *zap.Logger
	Server    *http.Server
}

// Run for grpc.
func Run(params RunParams) {
	if !flags.IsSet(VerifyFlag) {
		return
	}

	opts := &runner.Options{
		Lifecycle: params.Lifecycle,
		Logger:    params.Logger,
		Fn:        runner.NoModify,
	}

	runner.Run("http", "started", opts)
}
