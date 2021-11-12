package tachyoncrm

import "time"

type CustomerDto struct {
	ID             string
	Registered_at  time.Time
	LastSignedInAt time.Time
}
