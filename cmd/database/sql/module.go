package sql

import (
	"github.com/alexfalkowski/go-service/database/sql/pg"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	pg.Module,
	fx.Invoke(Run),
)
