package storage

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

const (
	maxSlot = 1023
	minSlot = 0
)

type Shard struct {
	from int
	to   int
	db   *sql.DB
}

type Shards struct {
	list []Shard
}

func MakeShard(from, to int, host string, port int, user, password, database string) (Shard, error) {
	if from > to {
		return Shard{}, errors.New("\"from\" can not be more than \"to\"")
	}

	if from < minSlot || from > maxSlot || to < minSlot || to > maxSlot {
		return Shard{}, fmt.Errorf("slot range must be between %d and %d", minSlot, maxSlot)
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, database)
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		return Shard{}, errors.Wrapf(err, "can not make shard for range %d...%d", from, to)
	}

	return Shard{
		from: from,
		to:   to,
		db:   db,
	}, nil
}

func (s Shard) overlap(s2 Shard) bool {
	if s.from < s2.from {
		return s2.from < s.to
	}

	return s2.to > s.from
}

func (sm *Shards) AddShard(s Shard) error {
	for _, shard := range sm.list {
		if s.overlap(shard) {
			return fmt.Errorf("shard %d...%d overlap shard %d...%d")
		}
	}
	sm.list = append(sm.list, s)
	return nil
}

func (sm *Shards) getShard(slot int) (bool, Shard) {
	for _, shard := range sm.list {
		if slot >= shard.from && slot <= shard.to {
			return true, shard
		}
	}
	return false, Shard{}
}
