package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/vasyahuyasa/ushtr/internal/storage/pg"
)

func isShardConfig(key string) bool {
	parts := strings.Split(key, "_")
	return len(parts) == 4 && parts[0] == "USHTR" && parts[2] == "PG" && parts[3] == "HOST"
}

func parseShardConfig(key string) (int, error) {
	parts := strings.Split(key, "_")
	return strconv.Atoi(parts[1])
}

func getShardPort(shard int) (int, bool) {
	env, ok := os.LookupEnv("USHTR_" + strconv.Itoa(shard) + "_PG_PORT")
	if !ok {
		return 0, false
	}

	port, err := strconv.Atoi(env)
	if err != nil {
		log.Printf("can not parse port %q for shard %d, default port used", env, shard)
		return 0, false
	}

	return port, true
}

func getShardDatabase(shard int) string {
	dbName, ok := os.LookupEnv("USHTR_" + strconv.Itoa(shard) + "_PG_DATABASE")
	if !ok {
		return "public"
	}

	return dbName
}

func loadShards(port int, user, password string) []pg.Shard {
	shards := []pg.Shard{}

	for _, env := range os.Environ() {
		parts := strings.Split(env, "=")
		if len(parts) != 2 || !isShardConfig(parts[0]) {
			continue
		}

		key, host := parts[0], parts[1]

		// get the shard id
		id, err := parseShardConfig(key)
		if err != nil {
			log.Printf("can not parse %q: %v", key, err)
			continue
		}

		// check if there specified the shard port or use default port
		shardPort, ok := getShardPort(id)
		if !ok {
			shardPort = port
		}

		//dbName := fmt.Sprintf("ushtr_%d", id)
		dbName := getShardDatabase(id)
		shard, err := pg.MakeShard(id, host, shardPort, user, password, dbName)
		if err != nil {
			log.Printf("can not create shard %q: %v", host, err)
			continue
		}

		shards = append(shards, shard)
	}
	return shards
}
