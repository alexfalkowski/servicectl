package hooks

import (
	"context"
	"net/http"
	"strconv"
	"time"

	sc "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/flags"
	h "github.com/alexfalkowski/go-service/hooks"
	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/servicectl/internal/cmd"
	"github.com/alexfalkowski/servicectl/internal/cmd/os"
	"github.com/alexfalkowski/servicectl/internal/cmd/runner"
	"github.com/alexfalkowski/servicectl/internal/config"
	hooks "github.com/standard-webhooks/standard-webhooks/libraries/go"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Register for hooks.
func Register(command *sc.Command) {
	flags := flags.NewFlagSet("hooks")

	command.RegisterInput(flags, "")
	rotate = flags.BoolP("rotate", "r", false, "rotate secret")
	verify = flags.BoolP("verify", "v", false, "verify webhook")

	command.AddClient("hooks", "Webhooks.", flags, cmd.Module, h.Module, fx.Invoke(start))
}

var (
	rotate = flags.Bool()
	verify = flags.Bool()
)

func start(lc fx.Lifecycle, logger *zap.Logger, gen *h.Generator, hook *hooks.Webhook, cfg *config.Config) {
	var (
		fn runner.StartFn
		op string
	)

	switch {
	case flags.IsBoolSet(rotate):
		fn = func(ctx context.Context) context.Context {
			s, err := gen.Generate()
			runtime.Must(err)

			err = os.WriteFile(cfg.Hooks.Secret, []byte(s))
			runtime.Must(err)

			return ctx
		}
		op = "rotated secret"
	case flags.IsBoolSet(verify):
		fn = func(ctx context.Context) context.Context {
			id, ts, p := "test", time.Now(), []byte("test")

			sig, err := hook.Sign(id, ts, p)
			runtime.Must(err)

			h := http.Header{}
			h.Add(hooks.HeaderWebhookID, id)
			h.Add(hooks.HeaderWebhookSignature, sig)
			h.Add(hooks.HeaderWebhookTimestamp, strconv.FormatInt(ts.Unix(), 10))

			err = hook.Verify(p, h)
			runtime.Must(err)

			return meta.WithAttribute(ctx, "testMsg", meta.String(string(p)))
		}
		op = "verified"
	}

	opts := &runner.Options{Lifecycle: lc, Logger: logger, Fn: fn}
	runner.Start("hooks", op, opts)
}
