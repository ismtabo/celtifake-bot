package dbo

import (
	"time"

	_ "github.com/kisielk/sqlstruct"
)

//go:generate sqlstruct -in github.com/ismtabo/phrases-of-the-year/pkg/repository/dbo -out ../query/phrases.go -pkg query -pub
type Phrase struct {
	Id        int       `db:"id,pk"`
	Content   string    `db:"content"`
	Group     string    `db:"group_"`
	Author    string    `db:"author"`
	CreatedAt time.Time `db:"created_at"`
}
