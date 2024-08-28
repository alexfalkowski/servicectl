package cmd

import (
	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/compress"
	"github.com/alexfalkowski/go-service/encoding"
	"github.com/alexfalkowski/go-service/env"
	"github.com/alexfalkowski/go-service/sync"
	"github.com/alexfalkowski/go-service/telemetry"
	"github.com/alexfalkowski/servicectl/config"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	sync.Module, compress.Module, encoding.Module,
	telemetry.Module, config.Module, cmd.Module,
	env.Module, fx.Provide(NewVersion),
)
