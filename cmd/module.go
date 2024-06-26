package cmd

import (
	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/compressor"
	"github.com/alexfalkowski/go-service/env"
	"github.com/alexfalkowski/go-service/marshaller"
	"github.com/alexfalkowski/go-service/telemetry"
	"github.com/alexfalkowski/go-service/telemetry/metrics"
	"github.com/alexfalkowski/servicectl/config"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	compressor.Module, marshaller.Module,
	telemetry.Module, metrics.Module,
	config.Module, cmd.Module,
	env.Module, fx.Provide(NewVersion),
)
