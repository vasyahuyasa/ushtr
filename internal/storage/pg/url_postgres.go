package pg

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/vasyahuyasa/ushtr/internal/storage"

	_ "github.com/lib/pq"
)

const (
	saveQuery = `INSERT INTO ushtr_%d.urls (url) VALUES ($1) RETURNING id`
	getQuery  = `SELECT url from ushtr_%d.urls WHERE id = $1`
)

// ShardedStorage is sharded storage based on instagram sharding algorithm
// https://instagram-engineering.com/sharding-ids-at-instagram-1cf5a71e5a5c
type ShardedStorage struct {
	shards *Shards
}

func NewStorage(shards *Shards) *ShardedStorage {
	return &ShardedStorage{
		shards: shards,
	}
}

func (ss *ShardedStorage) Get(id uint64) (string, error) {
	n := shardID(id)
	shard, ok := ss.shards.getShard(n)
	if !ok {
		log.Printf("shard %d is not defined", n)
		return "", storage.ErrNotFound
	}

	var url string
	err := shard.db.QueryRow(getQueryForShard(n), id).Scan(&url)
	if err == sql.ErrNoRows {
		return "", storage.ErrNotFound
	}
	if err != nil {
		return "", fmt.Errorf("can not query id %d from storage: %v", id, err)
	}

	return url, nil
}

func (ss *ShardedStorage) Save(url string) (uint64, error) {
	shard := ss.shards.getRandomShard()
	query := saveQueryForShard(shard.id)

	var id uint64
	row := shard.db.QueryRow(query, url)
	err := row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("can not save url to postgre sharded storage: %v", err)
	}

	return id, nil
}

func saveQueryForShard(shardID int) string {
	return fmt.Sprintf(saveQuery, shardID)
}

func getQueryForShard(shardID int) string {
	return fmt.Sprintf(getQuery, shardID)
}

// shardId extracts shard id from normal id
func shardID(id uint64) int {
	return int(id & 0x3ff)
}
