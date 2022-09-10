package context

import (
	"context"

	"github.com/rs/zerolog"
	"gopkg.in/telebot.v3"
)

type contextKey string

const tbContextKey = contextKey("context")

func WithContext(ctx context.Context, tbCtx telebot.Context) telebot.Context {
	tbCtx.Set(string(tbContextKey), ctx)
	return tbCtx
}

func Ctx(tbCtx telebot.Context) context.Context {
	if value := tbCtx.Get(string(tbContextKey)); value != nil {
		return value.(context.Context)
	}
	zerolog.DefaultContextLogger.Fatal().Msg("Error accessing standard context. May missing calling ctx.WithContext(ctx, tgContext)")
	return nil
}
