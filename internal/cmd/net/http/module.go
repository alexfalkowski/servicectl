package http

import (
	"github.com/alexfalkowski/go-service/limiter"
	"github.com/alexfalkowski/go-service/transport/http"
	"github.com/alexfalkowski/go-service/transport/meta"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	limiter.Module,
	meta.Module,
	http.Module,
	fx.Invoke(Start),
)
