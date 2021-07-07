package entities

import "time"

type Board struct {
	Id          int64
	Name        string
	Description string
	Columns     []*Column
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
