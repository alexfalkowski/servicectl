package cmd

import (
	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/module"
	"github.com/alexfalkowski/go-service/telemetry"
	"github.com/alexfalkowski/servicectl/internal/config"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	module.Module, telemetry.Module,
	config.Module, cmd.Module,
)
