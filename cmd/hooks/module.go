package hooks

import (
	"github.com/alexfalkowski/go-service/hooks"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	hooks.Module,
	fx.Invoke(Run),
)
