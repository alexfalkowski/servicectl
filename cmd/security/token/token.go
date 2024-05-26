package token

import (
	"context"
	"encoding/base64"

	"github.com/alexfalkowski/go-service/crypto/argon2"
	"github.com/alexfalkowski/go-service/crypto/rand"
	"github.com/alexfalkowski/go-service/flags"
	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/go-service/security/token"
	"github.com/alexfalkowski/servicectl/cmd/runner"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var (
	// RotateFlag defines wether we should rotate the key or not.
	RotateFlag = flags.Bool()

	// VerifyFlag defines wether we should verify the key or not.
	VerifyFlag = flags.Bool()
)

// Start for tokens.
func Start(lc fx.Lifecycle, logger *zap.Logger, algo argon2.Algo, tkn token.Tokenizer) {
	var (
		fn runner.StartFn
		op string
	)

	switch {
	case flags.IsSet(RotateFlag):
		fn = func(ctx context.Context) context.Context {
			d, err := rand.GenerateString(16)
			runtime.Must(err)

			k := base64.StdEncoding.EncodeToString([]byte(d))

			h, err := algo.Generate(d)
			runtime.Must(err)

			ctx = meta.WithAttribute(ctx, "key", meta.String(k))
			ctx = meta.WithAttribute(ctx, "hash", meta.String(h))

			return ctx
		}
		op = "rotated key and hash"
	case flags.IsSet(VerifyFlag):
		fn = func(ctx context.Context) context.Context {
			ctx, d, err := tkn.Generate(ctx)
			runtime.Must(err)

			ctx, err = tkn.Verify(ctx, d)
			runtime.Must(err)

			ctx = meta.WithAttribute(ctx, "key", meta.String(d))

			return ctx
		}
		op = "verified key"
	}

	opts := &runner.Options{Lifecycle: lc, Logger: logger, Fn: fn}
	runner.Start("token", op, opts)
}
