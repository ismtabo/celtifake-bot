package controller

import (
	"context"
	"fmt"
	"strings"
	"text/template"
	"time"

	"github.com/ismtabo/phrases-of-the-year/pkg/model"
	"github.com/ismtabo/phrases-of-the-year/pkg/service"
	"github.com/pkg/errors"
	tele "gopkg.in/telebot.v3"
)

const phraseMsg = `*Found phrases:*
{{ range . -}}
{{.Content}}

_\(Added by {{.Author}} on {{.CreatedAt | format "02/01/2006" }}\)_
\-\-\-\-\-\-\-\-\-\-\-
{{ end -}}
`

var (
	funcMap = template.FuncMap{
		"format": func(format string, t time.Time) string {
			return t.Format(format)
		},
	}
	phrasesTmpl = template.Must(template.New("phrases").Funcs(funcMap).Parse(phraseMsg))
)

type telegramApiBotImpl struct {
	bot *tele.Bot
	svc service.PhrasesService
}

func NewTelegramApiBotController(bot *tele.Bot, svc service.PhrasesService) TelegramBotApiController {
	return &telegramApiBotImpl{bot: bot, svc: svc}
}

func (t telegramApiBotImpl) Start(ctx context.Context, context tele.Context) error {
	return context.Send(fmt.Sprintf("Welcome to @%s", t.bot.Me.Username))
}

func (t telegramApiBotImpl) Help(ctx context.Context, context tele.Context) error {
	return context.Send("I understand /new and /search.")
}

func (t telegramApiBotImpl) New(ctx context.Context, context tele.Context) error {
	content := strings.TrimSpace(strings.Join(strings.Split(context.Text(), " ")[1:], " "))
	if content == "" {
		return context.Send("Missing phrase. Usage: /new <phrase content> (compatible with multiline messages)")
	}
	author := ""
	if sender := context.Sender(); sender.Username != "" {
		author = fmt.Sprintf("@%s", sender.Username)
	} else {
		author = fmt.Sprintf("%s %s", sender.FirstName, sender.LastName)
	}
	phrase := &model.Phrase{
		Content: content,
		Author:  author,
		Group:   fmt.Sprint(context.Chat().ID),
	}
	_, err := t.svc.CreatePhrase(ctx, phrase)
	if err != nil {
		return err
	}
	return context.Send("Successfully added phrase to collection")
}

func (t telegramApiBotImpl) Search(ctx context.Context, context tele.Context) error {
	match := context.Data()
	if match == "" {
		return context.Send("Missing searching terms. Usage: /search <search terms...>")
	}
	phrases, err := t.svc.GetPhrases(ctx, match)
	if err != nil {
		return err
	}
	if len(phrases) == 0 {
		return context.Send("Not phrases found")
	}
	b := &strings.Builder{}
	if err := phrasesTmpl.Execute(b, phrases); err != nil {
		return errors.Wrap(err, "error while generating message")
	}
	message := b.String()
	return context.Send(message, &tele.SendOptions{
		ParseMode: tele.ModeMarkdownV2,
	})
}
