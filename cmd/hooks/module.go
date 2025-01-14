package hooks

import (
	"github.com/alexfalkowski/go-service/crypto"
	"github.com/alexfalkowski/go-service/hooks"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	crypto.Module,
	hooks.Module,
	fx.Invoke(Start),
)
