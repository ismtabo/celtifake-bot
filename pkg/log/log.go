package log

import (
	"github.com/rs/zerolog"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func NewLogger(log *zerolog.Logger) tgbotapi.BotLogger {
	return &logger{log}
}

type logger struct {
	logger *zerolog.Logger
}

func (l logger) Println(v ...interface{}) {
	l.logger.Print(v...)
}
func (l logger) Printf(format string, v ...interface{}) {
	l.logger.Printf(format, v...)
}
