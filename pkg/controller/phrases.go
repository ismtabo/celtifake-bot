package controller

import (
	"context"

	tele "gopkg.in/telebot.v3"
)

//go:generate mockgen -destination ./mocks/telebot.go -package mocks gopkg.in/telebot.v3 Context
//go:generate mockgen -destination ./mocks/phrases.go -package mocks . TelegramBotApiController
type TelegramBotApiController interface {
	Start(ctx context.Context, context tele.Context) error
	Help(ctx context.Context, context tele.Context) error
	New(ctx context.Context, context tele.Context) error
	Search(ctx context.Context, context tele.Context) error
}
