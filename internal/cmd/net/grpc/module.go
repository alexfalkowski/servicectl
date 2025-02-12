package grpc

import (
	"github.com/alexfalkowski/go-service/limiter"
	"github.com/alexfalkowski/go-service/transport/grpc"
	"github.com/alexfalkowski/go-service/transport/meta"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	limiter.Module,
	meta.Module,
	grpc.Module,
	fx.Invoke(Start),
)
