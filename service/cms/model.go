package tachyoncms

import "time"

type AggregateDto struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Data      map[string]interface{}
}
