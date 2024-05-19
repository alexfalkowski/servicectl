package hooks

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/alexfalkowski/go-service/flags"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/servicectl/cmd/runner"
	"github.com/alexfalkowski/servicectl/config"
	hooks "github.com/standard-webhooks/standard-webhooks/libraries/go"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// SignFlag defines wether we should sign or not.
var SignFlag = flags.Bool()

// RunParams for hooks.
type RunParams struct {
	fx.In

	Lifecycle fx.Lifecycle
	Logger    *zap.Logger
	Webhook   *hooks.Webhook
}

// Run for hooks.
func Run(params RunParams) {
	if !flags.IsSet(SignFlag) {
		return
	}

	fn := func(ctx context.Context, _ *config.Config) context.Context {
		id, ts, p := "test", time.Now(), []byte("test")

		sig, err := params.Webhook.Sign(id, ts, p)
		runtime.Must(err)

		h := http.Header{}
		h.Add(hooks.HeaderWebhookID, id)
		h.Add(hooks.HeaderWebhookSignature, sig)
		h.Add(hooks.HeaderWebhookTimestamp, strconv.FormatInt(ts.Unix(), 10))

		err = params.Webhook.Verify(p, h)
		runtime.Must(err)

		return ctx
	}

	opts := &runner.Options{
		Lifecycle: params.Lifecycle,
		Logger:    params.Logger,
		Fn:        fn,
	}

	runner.Run("hooks", "signed", opts)
}
