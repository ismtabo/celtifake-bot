package repository

import (
	"context"

	"github.com/ismtabo/phrases-of-the-year/pkg/model"
)

//go:generate mockgen -destination ./mocks/phrases.go -package mocks . PhrasesRepository
type PhrasesRepository interface {
	CreatePhrase(ctx context.Context, phrase *model.Phrase) (*model.Phrase, error)
	GetPhrases(ctx context.Context, match string) ([]*model.Phrase, error)
}
