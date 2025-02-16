package http

import (
	sc "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/limiter"
	"github.com/alexfalkowski/go-service/transport/http"
	"github.com/alexfalkowski/go-service/transport/meta"
	"github.com/alexfalkowski/servicectl/internal/cmd"
	"github.com/alexfalkowski/servicectl/internal/cmd/flags"
	"github.com/alexfalkowski/servicectl/internal/cmd/runner"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Register for http.
func Register(command *sc.Command) {
	flags := sc.NewFlagSet("http")

	flags.AddInput("")
	flags.BoolP("verify", "v", false, "verify server")

	command.AddClient("http", "HTTP Server.", flags, cmd.Module, limiter.Module, meta.Module, http.Module, fx.Invoke(Start))
}

// StartParams for grpc.
type StartParams struct {
	fx.In

	Set       *sc.FlagSet
	Lifecycle fx.Lifecycle
	Logger    *zap.Logger
	Server    *http.Server
}

// Start for http.
//
//nolint:gocritic
func Start(params StartParams) {
	var (
		fn runner.StartFn
		op string
	)

	switch {
	case flags.IsSet(params.Set, "verify"):
		fn = runner.NoStart
		op = "started"
	}

	opts := &runner.Options{Lifecycle: params.Lifecycle, Logger: params.Logger, Fn: fn}
	runner.Start("http", op, opts)
}
