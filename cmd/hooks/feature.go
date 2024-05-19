package hooks

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/flags"
	h "github.com/alexfalkowski/go-service/hooks"
	"github.com/alexfalkowski/go-service/marshaller"
	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/servicectl/cmd/runner"
	"github.com/alexfalkowski/servicectl/config"
	hooks "github.com/standard-webhooks/standard-webhooks/libraries/go"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var (
	// RotateFlag defines wether we should rotate the secret or not.
	RotateFlag = flags.Bool()

	// VerifyFlag defines wether we should verify the hook or not.
	VerifyFlag = flags.Bool()
)

// RunParams for hooks.
type RunParams struct {
	fx.In

	Lifecycle    fx.Lifecycle
	OutputConfig *cmd.OutputConfig
	Map          *marshaller.Map
	Config       *config.Config
	Logger       *zap.Logger
	Webhook      *hooks.Webhook
}

// Run for hooks.
func Run(params RunParams) {
	var (
		fn runner.ModifyFn
		op string
		oc *cmd.OutputConfig
	)

	switch {
	case flags.IsSet(RotateFlag):
		fn = func(ctx context.Context, c *config.Config) context.Context {
			s, err := h.Generate()
			runtime.Must(err)

			c.Hooks.Secret = s

			return meta.WithAttribute(ctx, "key", meta.String(s))
		}
		op = "rotated secret"
		oc = params.OutputConfig
	case flags.IsSet(VerifyFlag):
		fn = func(ctx context.Context, _ *config.Config) context.Context {
			id, ts, p := "test", time.Now(), []byte("test")

			sig, err := params.Webhook.Sign(id, ts, p)
			runtime.Must(err)

			h := http.Header{}
			h.Add(hooks.HeaderWebhookID, id)
			h.Add(hooks.HeaderWebhookSignature, sig)
			h.Add(hooks.HeaderWebhookTimestamp, strconv.FormatInt(ts.Unix(), 10))

			err = params.Webhook.Verify(p, h)
			runtime.Must(err)

			return meta.WithAttribute(ctx, "testMsg", meta.String(p))
		}
		op = "verified"
	}

	opts := &runner.Options{
		Lifecycle:    params.Lifecycle,
		OutputConfig: oc,
		Map:          params.Map,
		Config:       params.Config,
		Logger:       params.Logger,
		Fn:           fn,
	}

	runner.Run("hooks", op, opts)
}
