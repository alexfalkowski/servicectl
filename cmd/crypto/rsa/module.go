package rsa

import (
	"github.com/alexfalkowski/go-service/crypto"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	crypto.Module,
	fx.Invoke(Start),
)
