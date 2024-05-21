package hooks

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/alexfalkowski/go-service/flags"
	h "github.com/alexfalkowski/go-service/hooks"
	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/servicectl/cmd/runner"
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

// Start for hooks.
func Start(lc fx.Lifecycle, logger *zap.Logger, hook *hooks.Webhook) {
	var (
		fn runner.StartFn
		op string
	)

	switch {
	case flags.IsSet(RotateFlag):
		fn = func(ctx context.Context) context.Context {
			s, err := h.Generate()
			runtime.Must(err)

			return meta.WithAttribute(ctx, "key", meta.String(s))
		}
		op = "rotated secret"
	case flags.IsSet(VerifyFlag):
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

			return meta.WithAttribute(ctx, "testMsg", meta.String(p))
		}
		op = "verified"
	}

	opts := &runner.Options{Lifecycle: lc, Logger: logger, Fn: fn}
	runner.Start("hooks", op, opts)
}
