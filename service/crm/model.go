package tachyoncrm

import "time"

type CustomerDto struct {
	ID             string
	RegisteredAt   time.Time
	LastSignedInAt time.Time
	Mail           *string
}
