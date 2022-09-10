package request

import (
	"context"

	"github.com/rs/zerolog"
)

type contextKey string

const requestContextKey = contextKey("request")

type Request struct {
	UserID  string
	GroupID string
}

func (r *Request) WithContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, requestContextKey, r)
}

func Ctx(ctx context.Context) *Request {
	if req := ctx.Value(requestContextKey); req != nil {
		return req.(*Request)
	}
	zerolog.Ctx(ctx).Fatal().Msg("Error accessing request in context. May forgive to do ctx.MustRequestContext(ctx)")
	return nil
}
