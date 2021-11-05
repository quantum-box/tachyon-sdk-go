package tachyonid

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid"
)

func NewUlID() string {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	return id.String()
}
