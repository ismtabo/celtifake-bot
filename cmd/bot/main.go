package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ismtabo/phrases-of-the-year/pkg/cfg"
	context_ "github.com/ismtabo/phrases-of-the-year/pkg/context"
	"github.com/ismtabo/phrases-of-the-year/pkg/controller"
	"github.com/ismtabo/phrases-of-the-year/pkg/middleware"
	"github.com/ismtabo/phrases-of-the-year/pkg/repository"
	"github.com/ismtabo/phrases-of-the-year/pkg/service"
	"github.com/rs/zerolog"

	_ "github.com/lib/pq"
	tele "gopkg.in/telebot.v3"
)

func main() {
	log := getLogger()
	var config Config
	if err := cfg.Load("config.yml", &config); err != nil {
		log.Fatal().Err(err).Msg("Error loading configuration.")
	}
	log.Debug().Msgf("Configuration loaded: %+v", config)
	if err := configLogger(&config); err != nil {
		log.Fatal().Err(err).Msg("Error configuring the logger.")
	}

	pref := tele.Settings{
		Token:  config.Telegram.Token.Raw(),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal().Err(err).Msg("Error configuring the telegram bot.")
		return
	}

	db := getSqliteConn(&config, log)
	repo := repository.NewPostgresPhrasesRepository(db)
	svc := service.NewPhrasesService(repo)
	ctrl := controller.NewTelegramApiBotController(bot, svc)

	ctx := context.Background()
	bot.Use(middleware.Context(ctx))
	bot.Use(middleware.Request())

	bot.Handle("/start", func(tgCtx tele.Context) error {
		ctx := context_.Ctx(tgCtx)
		return ctrl.Start(ctx, tgCtx)
	}, middleware.LogOp("start"))
	bot.Handle("/help", func(tgCtx tele.Context) error {
		ctx := context_.Ctx(tgCtx)
		return ctrl.Help(ctx, tgCtx)
	}, middleware.LogOp("help"))
	bot.Handle("/new", func(tgCtx tele.Context) error {
		ctx := context_.Ctx(tgCtx)
		return ctrl.New(ctx, tgCtx)
	}, middleware.LogOp("new"))
	bot.Handle("/search", func(tgCtx tele.Context) error {
		ctx := context_.Ctx(tgCtx)
		return ctrl.Search(ctx, tgCtx)
	}, middleware.LogOp("search"))
	bot.OnError = func(err error, context tele.Context) {
		log.Err(err).Msgf("error handling message %+v", context)
		if err := context.Send("Something bad occurs"); err != nil {
			log.Err(err).Msgf("error sending error message %+v", context)
		}
	}

	log.Info().Msg("Starting bot..")
	bot.Start()
}

func getLogger() *zerolog.Logger {
	zerolog.TimeFieldFormat = "2006-01-02T15:04:05.000Z07:00"
	zerolog.TimestampFieldName = "time"
	zerolog.LevelFieldName = "lvl"
	zerolog.MessageFieldName = "msg"
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()
	return &log
}

func configLogger(config *Config) error {
	lvl, err := zerolog.ParseLevel(strings.ToLower(config.Log.Level))
	if err != nil {
		return err
	}
	zerolog.SetGlobalLevel(lvl)
	return nil
}

func getSqliteConn(config *Config, log *zerolog.Logger) *sql.DB {
	uri := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", config.Database.User, config.Database.Password.Raw(), config.Database.Host, config.Database.Port, config.Database.Name)
	db, err := sql.Open("postgres", uri)
	if err != nil {
		log.Fatal().Err(err).Msg("Error connecting to database.")
	}
	return db
}
