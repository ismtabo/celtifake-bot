package repository

import (
	"context"
	"database/sql/driver"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/andreyvit/sqlexpr"
	"github.com/ismtabo/phrases-of-the-year/pkg/context/request"
	"github.com/ismtabo/phrases-of-the-year/pkg/model"
	"github.com/ismtabo/phrases-of-the-year/pkg/repository/dbo"
	"github.com/ismtabo/phrases-of-the-year/pkg/repository/query"
	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	sqlexpr.SetDialect(sqlexpr.SQLiteDialect)
}

func TestCreatePhrase(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	repo := NewPostgresPhrasesRepository(db)

	expectedPhrase := &model.Phrase{
		Content: "content",
		Author:  "author",
		Group:   "group",
	}

	dboPhrase := &dbo.Phrase{}
	if err := copier.Copy(dboPhrase, expectedPhrase); err != nil {
		t.Fatalf("error was expected while mapping phrase: %s", err)
	}

	s := sqlexpr.Insert{Table: query.PhrasesTable}
	s.Set(query.PhraseContent, dboPhrase.Content)
	s.Set(query.PhraseGroup, dboPhrase.Group)
	s.Set(query.PhraseAuthor, dboPhrase.Author)
	s.Set(query.PhraseCreatedAt, dboPhrase.CreatedAt)
	sql, args := sqlexpr.Build(s)
	execArgs := []driver.Value{}
	for _, arg := range args {
		value := driver.Value(arg)
		execArgs = append(execArgs, value)
	}

	mock.ExpectExec(sql).
		WithArgs(execArgs...)

	// now we execute our method
	phrase, err := repo.CreatePhrase(context.Background(), expectedPhrase)
	if err != nil {
		t.Fatalf("error was not expected while creating phrase: %s", err)
	}

	assert.Equal(t, expectedPhrase, phrase)

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: %s", err)
	}
}

func TestCreatePhrase_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	repo := NewPostgresPhrasesRepository(db)

	expectedError := errors.New("expected error")

	expectedPhrase := &model.Phrase{
		Content: "content",
		Author:  "author",
		Group:   "group",
	}

	dboPhrase := &dbo.Phrase{}
	if err := copier.Copy(dboPhrase, expectedPhrase); err != nil {
		t.Fatalf("error was expected while mapping phrase: %s", err)
	}

	s := sqlexpr.Insert{Table: query.PhrasesTable}
	s.Set(query.PhraseContent, dboPhrase.Content)
	s.Set(query.PhraseGroup, dboPhrase.Group)
	s.Set(query.PhraseAuthor, dboPhrase.Author)
	s.Set(query.PhraseCreatedAt, dboPhrase.CreatedAt)
	sql, args := sqlexpr.Build(s)
	execArgs := []driver.Value{}
	for _, arg := range args {
		value := driver.Value(arg)
		execArgs = append(execArgs, value)
	}

	mock.ExpectExec(sql).
		WithArgs(execArgs...).
		WillReturnError(expectedError)

	// now we execute our method
	_, err = repo.CreatePhrase(context.Background(), expectedPhrase)
	if err == nil {
		t.Fatalf("error was expected while invalid creating phrase: %s", err)
	}

	assert.ErrorIs(t, err, expectedError)

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetPhrase_Empty(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	repo := NewPostgresPhrasesRepository(db)

	req := request.Request{GroupID: "1"}
	match := "value"

	s := sqlexpr.Select{From: query.PhrasesTable}
	query.AddPhraseFields(&s, query.AllFacet)
	s.AddWhere(
		sqlexpr.Op(
			sqlexpr.Func("to_tsvector", sqlexpr.Column(query.PhraseContent)),
			"@@",
			sqlexpr.Func("to_tsquery", sqlexpr.Value(match)),
		),
		sqlexpr.Eq(sqlexpr.Column(query.PhraseGroup), req.GroupID),
	)
	sql, _ := sqlexpr.Build(s)

	rows := mock.NewRows([]string{"id", "content"})
	mock.ExpectQuery(sql).WithArgs(match).WillReturnRows(rows)

	// now we execute our method
	ctx := req.WithContext(context.Background())
	phrases, err := repo.GetPhrases(ctx, "phrase")
	if err != nil {
		t.Fatalf("error was not expected while getting phrases: %s", err)
	}

	assert.Len(t, phrases, 0)

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetPhrase_Results(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	repo := NewPostgresPhrasesRepository(db)

	req := request.Request{GroupID: "1"}
	match := "value"
	expectedPhrase := &model.Phrase{
		Id:        1,
		Content:   "content",
		Author:    "author",
		Group:     "group",
		CreatedAt: time.Now(),
	}

	s := sqlexpr.Select{From: query.PhrasesTable}
	query.AddPhraseFields(&s, query.AllFacet)
	s.AddWhere(
		sqlexpr.Op(
			sqlexpr.Func("to_tsvector", sqlexpr.Column(query.PhraseContent)),
			"@@",
			sqlexpr.Func("to_tsquery", sqlexpr.Value(match)),
		),
		sqlexpr.Eq(sqlexpr.Column(query.PhraseGroup), req.GroupID),
	)
	sql, _ := sqlexpr.Build(s)

	rows := mock.
		NewRows([]string{"id", "content", "author", "group", "created_at"}).
		AddRow(1, "value", "author", "group", time.Now())
	mock.ExpectQuery(sql).WithArgs(match).WillReturnRows(rows)

	// now we execute our method
	ctx := req.WithContext(context.Background())
	phrases, err := repo.GetPhrases(ctx, "phrase")
	if err != nil {
		t.Fatalf("error was not expected while getting phrases: %s", err)
	}

	assert.Len(t, phrases, 1)
	assert.Contains(t, phrases, expectedPhrase)

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: %s", err)
	}
}
