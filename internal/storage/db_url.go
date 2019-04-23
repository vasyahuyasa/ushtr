package storage

import (
	"time"
)

type dbUrl struct {
	ID         uint64
	url        string
	created_at time.Time
}
