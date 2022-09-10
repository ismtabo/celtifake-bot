package middleware

import (
	"context"
	"fmt"

	context_ "github.com/ismtabo/phrases-of-the-year/pkg/context"
	"github.com/ismtabo/phrases-of-the-year/pkg/context/request"
	"github.com/rs/zerolog"
	"gopkg.in/telebot.v3"
)

func Context(ctx context.Context) telebot.MiddlewareFunc {
	return func(hf telebot.HandlerFunc) telebot.HandlerFunc {
		return func(tgCtx telebot.Context) error {
			return hf(context_.WithContext(ctx, tgCtx))
		}
	}
}

func Request() telebot.MiddlewareFunc {
	return func(hf telebot.HandlerFunc) telebot.HandlerFunc {
		return func(tgCtx telebot.Context) error {
			req := request.Request{
				UserID:  fmt.Sprint(tgCtx.Sender().ID),
				GroupID: fmt.Sprint(tgCtx.Chat().ID),
			}
			ctx := req.WithContext(context_.Ctx(tgCtx))
			return hf(context_.WithContext(ctx, tgCtx))
		}
	}
}

func LogOp(op string) telebot.MiddlewareFunc {
	return func(hf telebot.HandlerFunc) telebot.HandlerFunc {
		return func(tgCtx telebot.Context) error {
			ctx := context_.Ctx(tgCtx)
			ctx = zerolog.Ctx(ctx).With().Str("op", op).Logger().WithContext(ctx)
			return hf(context_.WithContext(ctx, tgCtx))
		}
	}
}
