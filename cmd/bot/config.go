package main

import (
	"strings"
)

type Config struct {
	Telegram struct {
		Token Secret `yml:"host" envconfig:"TELEGRAM_TOKEN"`
	} `yml:"telegram"`
	Log struct {
		Level string `yml:"level" envconfig:"LOG_LEVEL"`
	} `yml:"log"`
	Database struct {
		Host     string `yml:"host" envconfig:"DATABASE_HOST"`
		Port     int    `yml:"port" envconfig:"DATABASE_PORT"`
		User     string `yml:"user" envconfig:"DATABASE_USER"`
		Password Secret `yml:"password" envconfig:"DATABASE_PASS"`
		Name     string `yml:"name" envconfig:"DATABASE_NAME"`
	} `yml:"database"`
}

type Secret string

func (s Secret) String() string {
	return strings.Repeat("*", len(s))
}

func (s Secret) Raw() string {
	return string(s)
}
