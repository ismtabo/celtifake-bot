package service

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ismtabo/phrases-of-the-year/pkg/model"
	"github.com/ismtabo/phrases-of-the-year/pkg/repository/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCreatePhrase(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockPhrasesRepository(ctrl)
	svc := NewPhrasesService(mockRepo)

	expectedPhrase := &model.Phrase{
		Content: "content",
		Author:  "author",
		Group:   "group",
	}
	mockRepo.EXPECT().CreatePhrase(gomock.Any(), expectedPhrase).Return(expectedPhrase, nil)

	phrase, err := svc.CreatePhrase(context.Background(), expectedPhrase)
	if err != nil {
		t.Fatalf("error was not expected while creating phrase: %s", err)
	}

	assert.Equal(t, expectedPhrase, phrase)
}

func TestGetPhrases(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockPhrasesRepository(ctrl)
	svc := NewPhrasesService(mockRepo)

	expectedMatch := "match"
	expectedPhrases := []*model.Phrase{}
	mockRepo.EXPECT().GetPhrases(gomock.Any(), expectedMatch).Return(expectedPhrases, nil)

	phrases, err := svc.GetPhrases(context.Background(), expectedMatch)
	if err != nil {
		t.Fatalf("error was not expected while creating phrase: %s", err)
	}

	assert.Equal(t, expectedPhrases, phrases)
}
