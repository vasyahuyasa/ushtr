package main

import (
	"fmt"
	"log"

	"github.com/vasyahuyasa/ushtr/internal/web"

	"github.com/vasyahuyasa/ushtr/internal/shortener/simple"
	"github.com/vasyahuyasa/ushtr/internal/storage/pg"

	// autoload .env for developer purposes
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	cfg := loadConfig()
	shards := loadShards(cfg.pgPort, cfg.pgUser, cfg.pgPassword)
	if len(shards) == 0 {
		log.Fatal("no shards defined, define at least one shard")
	}

	shardList := &pg.Shards{}
	for _, shard := range shards {
		err := shardList.AddShard(shard)
		if err != nil {
			log.Fatalf("can not add shard: %v", err)
		}
	}

	log.Printf("use %d shard(s)", len(shards))
	for _, shard := range shards {
		log.Println("shard", shard)
	}

	storage := pg.NewStorage(shardList)

	shortener, err := simple.New()
	if err != nil {
		log.Fatalf("can not create shortener: %v", err)
	}

	addr := fmt.Sprintf("%s:%d", cfg.addr, cfg.port)
	log.Printf("Listen on http://%s", addr)
	srv := web.NewServer(storage, shortener, cfg.url)
	log.Fatal(srv.Run(addr))
}
