package cmd

import (
	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/servicectl/cmd/aes"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	aes.Module,
	cmd.Module,
	fx.Provide(NewVersion),
)
