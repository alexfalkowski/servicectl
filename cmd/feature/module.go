package feature

import (
	"github.com/alexfalkowski/go-service/feature"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	feature.Module,
	fx.Invoke(Run),
)
