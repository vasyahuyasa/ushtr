package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/vasyahuyasa/ushtr/internal/storage"
)

func isShardConfig(key string) bool {
	parts := strings.Split(key, "_")
	return len(parts) == 5 && parts[0] == "USHTR" && parts[3] == "PG" && parts[4] == "HOST"
}

func parseShardConfig(key string) (from int, to int, err error) {
	parts := strings.Split(key, "_")
	from, err = strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, err
	}

	to, err = strconv.Atoi(parts[2])
	if err != nil {
		return 0, 0, err
	}

	return from, to, nil
}

func loadShards(port int, user, password, database string) []storage.Shard {
	shards := []storage.Shard{}

	for _, env := range os.Environ() {
		parts := strings.Split(env, "=")
		if len(parts) != 2 || !isShardConfig(parts[0]) {
			continue
		}

		from, to, err := parseShardConfig(parts[0])
		if err != nil {
			log.Printf("can not parse %q: %v", parts[0], err)
			continue
		}

		shard, err := storage.MakeShard(from, to, parts[1], port, user, password, database)
		if err != nil {
			log.Printf("can not create shard %q: %v", parts[1], err)
			continue
		}

		shards = append(shards, shard)
	}
	return shards
}
