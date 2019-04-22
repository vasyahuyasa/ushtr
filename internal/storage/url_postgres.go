package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type config struct {
	slots *Shards

	host     string
	port     int
	user     string
	password string
	database string
}

type Option func(s *config)

func WithDefaultPgConnect(host string, port int, user, password, database string) func(cfg *config) {
	return func(cfg *config) {
		cfg.host = host
		cfg.port = port
		cfg.user = user
		cfg.password = password
		cfg.database = database
	}
}

func WithShards(sm *Shards) func(cfg *config) {
	return func(cfg *config) {
		cfg.slots = sm
	}
}

func defaultConfig() *config {
	return &config{
		slots:    &Shards{},
		host:     "localhost",
		port:     5432,
		user:     "postgres",
		password: "",
		database: "public",
	}
}

func NewStorage(options ...Option) (*PostgreShardedStorage, error) {
	cfg := defaultConfig()
	for _, f := range options {
		f(cfg)
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.user, cfg.password, cfg.host, cfg.port, cfg.database)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("can not open postgres: %v", err)
	}

	/*
		err = db.Ping()
		if err != nil {
			return nil, fmt.Errorf("can not ping postgres: %v", err)
		}
	*/

	return &PostgreShardedStorage{
		db: db,
	}, nil
}

type PostgreShardedStorage struct {
	db *sql.DB
}

func (ps *PostgreShardedStorage) Get(id uint64) (string, error) {
	return "", fmt.Errorf("not implemented")
}

func (ps *PostgreShardedStorage) Put(url string) (uint64, error) {
	return 0, fmt.Errorf("not implemented")

}
