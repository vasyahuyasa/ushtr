package main

import (
	"log"

	"github.com/vasyahuyasa/ushtr/internal/web"

	"github.com/vasyahuyasa/ushtr/internal/shortener/simple"
	"github.com/vasyahuyasa/ushtr/internal/storage"

	// autoload .env for developer purposes
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	cfg := loadConfig()
	shards := loadShards(cfg.pgPort, cfg.pgUser, cfg.pgPassword, cfg.pgDatabe)
	if len(shards) == 0 {
		log.Fatal("no shards defined, define at least one shard")
	}

	shardList := &storage.Shards{}
	for _, shard := range shards {
		err := shardList.AddShard(shard)
		if err != nil {
			log.Fatalf("can not add shard: %v", err)
		}
	}

	log.Printf("use %d shards", len(shards))
	for _, shard := range shards {
		log.Println("shard", shard)
	}

	storage, err := storage.NewStorage(
		storage.WithDefaultPgConnect(cfg.pgHost, cfg.pgPort, cfg.pgUser, cfg.pgPassword, cfg.pgDatabe),
		storage.WithShards(shardList),
	)
	if err != nil {
		log.Fatalf("can not create sorage: %v", err)
	}

	shortener, err := simple.New()
	if err != nil {
		log.Fatalf("can not create shortener: %v", err)
	}

	srv := web.NewServer(storage, shortener)
	log.Fatal(srv.Run())
}
