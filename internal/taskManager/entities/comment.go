package entities

import "time"

type Comment struct {
	Id        int64
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
