package controller

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	ctrl_mocks "github.com/ismtabo/phrases-of-the-year/pkg/controller/mocks"
	"github.com/ismtabo/phrases-of-the-year/pkg/model"
	svc_mocks "github.com/ismtabo/phrases-of-the-year/pkg/service/mocks"
	"gopkg.in/telebot.v3"
)

func TestStart(t *testing.T) {
	ctrl_ := gomock.NewController(t)
	mockSvc := svc_mocks.NewMockPhrasesService(ctrl_)
	mockContext := ctrl_mocks.NewMockContext(ctrl_)

	expectedBotUsername := "bot"
	bot := &telebot.Bot{Me: &telebot.User{Username: expectedBotUsername}}
	ctrl := NewTelegramApiBotController(bot, mockSvc)

	mockContext.EXPECT().Send(fmt.Sprintf("Welcome to @%s", expectedBotUsername)).Return(nil)

	if err := ctrl.Start(context.Background(), mockContext); err != nil {
		t.Fatalf("error was not expected while handling start: %s", err)
	}
}

func TestHelp(t *testing.T) {
	ctrl_ := gomock.NewController(t)
	mockSvc := svc_mocks.NewMockPhrasesService(ctrl_)
	mockContext := ctrl_mocks.NewMockContext(ctrl_)

	bot := &telebot.Bot{}
	ctrl := NewTelegramApiBotController(bot, mockSvc)

	mockContext.EXPECT().Send("I understand /new and /search.").Return(nil)

	if err := ctrl.Help(context.Background(), mockContext); err != nil {
		t.Fatalf("error was not expected while handling start: %s", err)
	}
}

func TestNew_Username(t *testing.T) {
	ctrl_ := gomock.NewController(t)
	mockSvc := svc_mocks.NewMockPhrasesService(ctrl_)
	mockContext := ctrl_mocks.NewMockContext(ctrl_)

	bot := &telebot.Bot{}
	ctrl := NewTelegramApiBotController(bot, mockSvc)

	mockContext.EXPECT().Text().Return("/new phrase")
	mockContext.EXPECT().Sender().Return(&telebot.User{Username: "username"})
	mockContext.EXPECT().Chat().Return(&telebot.Chat{ID: 1})
	mockSvc.EXPECT().CreatePhrase(gomock.Any(), &model.Phrase{Content: "phrase", Author: "@username", Group: "1"})
	mockContext.EXPECT().Send("Successfully added phrase to collection").Return(nil)

	if err := ctrl.New(context.Background(), mockContext); err != nil {
		t.Fatalf("error was not expected while handling start: %s", err)
	}
}

func TestNew_FullName(t *testing.T) {
	ctrl_ := gomock.NewController(t)
	mockSvc := svc_mocks.NewMockPhrasesService(ctrl_)
	mockContext := ctrl_mocks.NewMockContext(ctrl_)

	bot := &telebot.Bot{}
	ctrl := NewTelegramApiBotController(bot, mockSvc)

	mockContext.EXPECT().Text().Return("/new phrase")
	mockContext.EXPECT().Sender().Return(&telebot.User{FirstName: "firstname", LastName: "lastname"})
	mockContext.EXPECT().Chat().Return(&telebot.Chat{ID: 1})
	mockSvc.EXPECT().CreatePhrase(gomock.Any(), &model.Phrase{Content: "phrase", Author: "firstname lastname", Group: "1"})
	mockContext.EXPECT().Send("Successfully added phrase to collection").Return(nil)

	if err := ctrl.New(context.Background(), mockContext); err != nil {
		t.Fatalf("error was not expected while handling start: %s", err)
	}
}

func TestNew_EmptyContent(t *testing.T) {
	ctrl_ := gomock.NewController(t)
	mockSvc := svc_mocks.NewMockPhrasesService(ctrl_)
	mockContext := ctrl_mocks.NewMockContext(ctrl_)

	bot := &telebot.Bot{}
	ctrl := NewTelegramApiBotController(bot, mockSvc)

	mockContext.EXPECT().Text().Return("/new")
	mockContext.EXPECT().Send("Missing phrase. Usage: /new <phrase content> (compatible with multiline messages)").Return(nil)

	if err := ctrl.New(context.Background(), mockContext); err != nil {
		t.Fatalf("error was not expected while handling start: %s", err)
	}
}

func TestNew_MultilineContext(t *testing.T) {
	ctrl_ := gomock.NewController(t)
	mockSvc := svc_mocks.NewMockPhrasesService(ctrl_)
	mockContext := ctrl_mocks.NewMockContext(ctrl_)

	bot := &telebot.Bot{}
	ctrl := NewTelegramApiBotController(bot, mockSvc)

	expectedContent := `multine
	phrase`
	mockContext.EXPECT().Text().Return(fmt.Sprintf("/new %s", expectedContent))
	mockContext.EXPECT().Sender().Return(&telebot.User{FirstName: "first", LastName: "last"})
	mockContext.EXPECT().Chat().Return(&telebot.Chat{})
	mockSvc.EXPECT().CreatePhrase(gomock.Any(), &model.Phrase{Content: expectedContent, Author: "first last", Group: "0"})
	mockContext.EXPECT().Send("Successfully added phrase to collection").Return(nil)

	if err := ctrl.New(context.Background(), mockContext); err != nil {
		t.Fatalf("error was not expected while handling start: %s", err)
	}
}

func TestSearch(t *testing.T) {
	ctrl_ := gomock.NewController(t)
	mockSvc := svc_mocks.NewMockPhrasesService(ctrl_)
	mockContext := ctrl_mocks.NewMockContext(ctrl_)

	bot := &telebot.Bot{}
	ctrl := NewTelegramApiBotController(bot, mockSvc)

	expectedMatch := "word"
	expectedPhrases := []*model.Phrase{
		{Content: "content", Author: "author", CreatedAt: time.Now()},
	}
	s := &strings.Builder{}
	if err := phrasesTmpl.Execute(s, expectedPhrases); err != nil {
		t.Fatalf("error was not expected while generating message: %s", err)
	}
	expectedMessage := s.String()
	mockContext.EXPECT().Data().Return(expectedMatch)
	mockSvc.EXPECT().GetPhrases(gomock.Any(), expectedMatch).Return(expectedPhrases, nil)
	mockContext.EXPECT().Send(expectedMessage, gomock.Any()).Return(nil)

	if err := ctrl.Search(context.Background(), mockContext); err != nil {
		t.Fatalf("error was not expected while handling start: %s", err)
	}
}

func TestSearch_EmptyMatch(t *testing.T) {
	ctrl_ := gomock.NewController(t)
	mockSvc := svc_mocks.NewMockPhrasesService(ctrl_)
	mockContext := ctrl_mocks.NewMockContext(ctrl_)

	bot := &telebot.Bot{}
	ctrl := NewTelegramApiBotController(bot, mockSvc)

	mockContext.EXPECT().Data().Return("")
	mockContext.EXPECT().Send("Missing searching terms. Usage: /search <search terms...>").Return(nil)

	if err := ctrl.Search(context.Background(), mockContext); err != nil {
		t.Fatalf("error was not expected while handling start: %s", err)
	}
}

func TestSearch_EmptyResults(t *testing.T) {
	ctrl_ := gomock.NewController(t)
	mockSvc := svc_mocks.NewMockPhrasesService(ctrl_)
	mockContext := ctrl_mocks.NewMockContext(ctrl_)

	bot := &telebot.Bot{}
	ctrl := NewTelegramApiBotController(bot, mockSvc)

	expectedMatch := "word"
	expectedPhrases := []*model.Phrase{}
	mockContext.EXPECT().Data().Return(expectedMatch)
	mockSvc.EXPECT().GetPhrases(gomock.Any(), expectedMatch).Return(expectedPhrases, nil)
	mockContext.EXPECT().Send("Not phrases found").Return(nil)

	if err := ctrl.Search(context.Background(), mockContext); err != nil {
		t.Fatalf("error was not expected while handling start: %s", err)
	}
}
