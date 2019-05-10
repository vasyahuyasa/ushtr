package storage

import (
	"fmt"

	_ "github.com/lib/pq"
)

type PostgreShardedStorage struct {
	shards *Shards
}

func NewStorage(shards *Shards) *PostgreShardedStorage {
	return &PostgreShardedStorage{
		shards: shards,
	}
}

func (ps *PostgreShardedStorage) Get(id uint64) (string, error) {
	return "", fmt.Errorf("not implemented")
}

func (ps *PostgreShardedStorage) Save(url string) (uint64, error) {
	return 0, fmt.Errorf("not implemented")
}
