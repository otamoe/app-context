package context

import (
	"context"

	"go.uber.org/fx"
)

func New(mctx context.Context) (fxOption fx.Option) {
	if mctx == nil {
		mctx = context.Background()
	}
	fxOption = fx.Provide(func(lc fx.Lifecycle) context.Context {
		ctx, cancel := context.WithCancel(mctx)
		lc.Append(fx.Hook{
			OnStop: func(_ context.Context) error {
				cancel()
				return nil
			},
		})
		return ctx
	})
	return
}
