package http

import (
	sc "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/flags"
	"github.com/alexfalkowski/go-service/limiter"
	"github.com/alexfalkowski/go-service/transport/http"
	"github.com/alexfalkowski/go-service/transport/meta"
	"github.com/alexfalkowski/servicectl/internal/cmd"
	"github.com/alexfalkowski/servicectl/internal/cmd/runner"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Register for http.
func Register(command *sc.Command) {
	client := command.AddClient("http", "HTTP Server.", cmd.Module, limiter.Module, meta.Module, http.Module, fx.Invoke(start))

	flags.BoolVar(client, verify, "verify", "v", false, "verify server")
}

var verify = flags.Bool()

//nolint:gocritic
func start(lc fx.Lifecycle, logger *zap.Logger, _ *http.Server) {
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
	runner.Start("http", op, opts)
}
