package client

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

// RunCommandParams for client.
type RunCommandParams struct {
	fx.In

	Lifecycle fx.Lifecycle
	Logger    *zap.Logger
}

// RunCommand for client.
func RunCommand(params RunCommandParams) {
	params.Lifecycle.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			params.Logger.Info("awesome client")

			return nil
		},
	})
}
