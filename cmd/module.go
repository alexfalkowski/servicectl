package cmd

import (
	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/compress"
	"github.com/alexfalkowski/go-service/encoding"
	"github.com/alexfalkowski/go-service/env"
	"github.com/alexfalkowski/go-service/telemetry"
	"github.com/alexfalkowski/go-service/telemetry/metrics"
	"github.com/alexfalkowski/servicectl/config"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	compress.Module, encoding.Module,
	telemetry.Module, metrics.Module,
	config.Module, cmd.Module,
	env.Module, fx.Provide(NewVersion),
)
