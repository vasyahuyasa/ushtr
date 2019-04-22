package postgre

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/vasyahuyasa/ushtr/internal/hash"
)

type config struct {
	host     string
	port     int
	user     string
	password string
	database string
}

type Option func(s *config)

func WithPgConnect(host string, port int, user, password, database string) func(cfg *config) {
	return func(cfg *config) {
		cfg.host = host
		cfg.port = port
		cfg.user = user
		cfg.password = password
		cfg.database = database
	}
}

func defaultConfig() *config {
	return &config{
		host:     "localhost",
		port:     5432,
		user:     "postgres",
		password: "",
		database: "public",
	}
}

func New(options ...Option) (*PostgreShardedStorage, error) {
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

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("can not ping postgres: %v", err)
	}

	return &PostgreShardedStorage{
		db: db,
	}, nil
}

type PostgreShardedStorage struct {
	hashFunc hash.HashFunc
	db       *sql.DB
}

func (ps *PostgreShardedStorage) GetByUrl(url string, shard int) (string, error) {

}

func (ps *PostgreShardedStorage) Get(id uint64, shard int) (string, error) {

}

func (ps *PostgreShardedStorage) Put(url string, shard int) (uint64, error) {

}
