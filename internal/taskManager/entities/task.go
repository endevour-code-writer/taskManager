package entities

import "time"

type Task struct {
	Id          int64
	Name        string
	Description string
	Status      string
	Comments    []*Comment
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
