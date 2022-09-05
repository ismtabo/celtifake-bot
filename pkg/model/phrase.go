package model

import "time"

type Phrase struct {
	Id        uint
	Content   string
	Group     string
	Author    string
	CreatedAt time.Time
}
