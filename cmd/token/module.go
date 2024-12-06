package token

import (
	"github.com/alexfalkowski/go-service/crypto"
	"github.com/alexfalkowski/go-service/token"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	crypto.Module,
	token.Module,
	fx.Invoke(Start),
)
