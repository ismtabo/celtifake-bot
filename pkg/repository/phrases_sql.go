package repository

import (
	"context"
	"database/sql"

	"github.com/andreyvit/sqlexpr"
	"github.com/ismtabo/phrases-of-the-year/pkg/context/request"
	"github.com/ismtabo/phrases-of-the-year/pkg/model"
	"github.com/ismtabo/phrases-of-the-year/pkg/repository/dbo"
	"github.com/ismtabo/phrases-of-the-year/pkg/repository/query"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
)

type sqlitePhrasesRepository struct {
	db *sql.DB
}

func NewSqlitePhrasesRepository(db *sql.DB) PhrasesRepository {
	return sqlitePhrasesRepository{db: db}
}

func (r sqlitePhrasesRepository) CreatePhrase(ctx context.Context, phrase *model.Phrase) (*model.Phrase, error) {
	dboPhrase := &dbo.Phrase{}
	if err := copier.Copy(dboPhrase, phrase); err != nil {
		return nil, errors.Wrap(err, "error mapping phrase")
	}
	s := sqlexpr.Insert{Table: query.PhrasesTable}
	s.Set(query.PhraseContent, dboPhrase.Content)
	s.Set(query.PhraseGroup, dboPhrase.Group)
	s.Set(query.PhraseAuthor, dboPhrase.Author)
	sql, args := sqlexpr.Build(s)
	if _, err := r.db.ExecContext(ctx, sql, args...); err != nil {
		return nil, errors.Wrap(err, "error inserting phrase")
	}
	return phrase, nil
}

func (r sqlitePhrasesRepository) GetPhrases(ctx context.Context, match string) ([]*model.Phrase, error) {
	req := request.Ctx(ctx)
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
	sql, args := sqlexpr.Build(s)
	rows, err := r.db.QueryContext(ctx, sql, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "error getting phrase by '%s'", match)
	}
	rawPhrases, err := query.ScanAllPhrases(rows, query.AllFacet)
	if err != nil {
		return nil, errors.Wrap(err, "error scanning rows")
	}
	phrases := []*model.Phrase{}
	if err := copier.Copy(&phrases, &rawPhrases); err != nil {
		return nil, errors.Wrap(err, "error mapping rows")
	}
	return phrases, nil
}
