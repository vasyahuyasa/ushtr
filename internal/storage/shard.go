package storage

import (
	"database/sql"
	"fmt"
	"math/rand"

	"github.com/pkg/errors"
)

const (
	maxSlot = 1023
	minSlot = 0
)

type Shard struct {
	id   int
	db   *sql.DB
	host string
	port int
}

type Shards struct {
	ids  []int
	list map[int]Shard
}

func MakeShard(id int, host string, port int, user, password, database string) (Shard, error) {
	if id < minSlot || id > maxSlot {
		return Shard{}, fmt.Errorf("slot must be between %d and %d", minSlot, maxSlot)
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, database)
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		return Shard{}, errors.Wrapf(err, "can not make shard for %d", id)
	}

	return Shard{
		id:   id,
		db:   db,
		host: host,
		port: port,
	}, nil
}

func (s Shard) overlap(s2 Shard) bool {
	return s.id == s2.id
}

func (s Shard) String() string {
	return fmt.Sprintf("%d - %s:%d", s.id, s.host, s.port)
}

func (sm *Shards) AddShard(s Shard) error {
	if sm.list == nil {
		sm.list = map[int]Shard{}
	}
	for _, shard := range sm.list {
		if s.overlap(shard) {
			return fmt.Errorf("shard %d alredy defined", s.id)
		}
	}
	sm.list[s.id] = s
	sm.ids = append(sm.ids, s.id)
	return nil
}

func (sm *Shards) getShard(shard int) (Shard, bool) {
	s, ok := sm.list[shard]
	return s, ok
}

func (sm *Shards) getRandomShard() Shard {
	n := rand.Intn(len(sm.ids))
	return sm.list[n]
}
