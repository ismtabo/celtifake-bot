package service

import (
	"context"

	"github.com/ismtabo/phrases-of-the-year/pkg/model"
	"github.com/ismtabo/phrases-of-the-year/pkg/repository"
)

type phrasesService struct {
	repo repository.PhrasesRepository
}

func NewPhrasesService(repo repository.PhrasesRepository) PhrasesService {
	return &phrasesService{repo: repo}
}

func (s phrasesService) CreatePhrase(ctx context.Context, phrase *model.Phrase) (*model.Phrase, error) {
	return s.repo.CreatePhrase(ctx, phrase)
}
func (s phrasesService) GetPhrases(ctx context.Context, match string) ([]*model.Phrase, error) {
	return s.repo.GetPhrases(ctx, match)
}
