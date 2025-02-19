package hooks

import (
	"context"
	"net/http"
	"strconv"
	"time"

	sc "github.com/alexfalkowski/go-service/cmd"
	h "github.com/alexfalkowski/go-service/hooks"
	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/go-service/telemetry/logger"
	"github.com/alexfalkowski/servicectl/internal/cmd"
	"github.com/alexfalkowski/servicectl/internal/cmd/flags"
	"github.com/alexfalkowski/servicectl/internal/cmd/os"
	"github.com/alexfalkowski/servicectl/internal/cmd/runner"
	"github.com/alexfalkowski/servicectl/internal/config"
	hooks "github.com/standard-webhooks/standard-webhooks/libraries/go"
	"go.uber.org/fx"
)

// Register for hooks.
func Register(command *sc.Command) {
	flags := sc.NewFlagSet("hooks")

	flags.AddInput("")
	flags.BoolP("rotate", "r", false, "rotate secret")
	flags.BoolP("verify", "v", false, "verify webhook")

	command.AddClient("hooks", "Webhooks.", flags, cmd.Module, h.Module, fx.Invoke(Start))
}

// StartParams for hooks.
type StartParams struct {
	fx.In

	Set       *sc.FlagSet
	Lifecycle fx.Lifecycle
	Logger    *logger.Logger
	Generator *h.Generator
	Hook      *hooks.Webhook
	Config    *config.Config
}

// Start for hooks.
func Start(params StartParams) {
	var (
		fn runner.StartFn
		op string
	)

	switch {
	case flags.IsSet(params.Set, "rotate"):
		fn = func(ctx context.Context) context.Context {
			s, err := params.Generator.Generate()
			runtime.Must(err)

			err = os.WriteFile(params.Config.Hooks.Secret, []byte(s))
			runtime.Must(err)

			return ctx
		}
		op = "rotated secret"
	case flags.IsSet(params.Set, "verify"):
		fn = func(ctx context.Context) context.Context {
			id, ts, p := "test", time.Now(), []byte("test")

			sig, err := params.Hook.Sign(id, ts, p)
			runtime.Must(err)

			h := http.Header{}
			h.Add(hooks.HeaderWebhookID, id)
			h.Add(hooks.HeaderWebhookSignature, sig)
			h.Add(hooks.HeaderWebhookTimestamp, strconv.FormatInt(ts.Unix(), 10))

			err = params.Hook.Verify(p, h)
			runtime.Must(err)

			return meta.WithAttribute(ctx, "testMsg", meta.String(string(p)))
		}
		op = "verified"
	}

	opts := &runner.Options{Lifecycle: params.Lifecycle, Logger: params.Logger, Fn: fn}
	runner.Start("hooks", op, opts)
}
