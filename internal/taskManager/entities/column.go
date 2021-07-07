package entities

import "time"

type Column struct {
	Id        int64
	Name      string
	Tasks     []*Task
	Position  uint32
	CreatedAt time.Time
	UpdatedAt time.Time
}
