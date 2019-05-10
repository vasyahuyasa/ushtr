package main

import (
	"os"
	"strconv"
)

const defaultAddr = "0.0.0.0"
const defaultPort = 38663

type config struct {
	// web
	addr string
	port int

	// postgresql
	pgHost     string
	pgPort     int
	pgUser     string
	pgDatabe   string
	pgPassword string
}

func loadConfig() config {
	return config{
		addr:       getEnv("USHTR_ADDR", defaultAddr),
		port:       getEnvAsInt("USHTR_PORT", defaultPort),
		pgUser:     getEnv("USHTR_PG_USER", ""),
		pgPassword: getEnv("USHTR_PG_PASSWORD", ""),
	}
}

func getEnv(key, defaultVal string) string {
	val, ok := os.LookupEnv(key)
	if ok {
		return val
	}
	return defaultVal
}

func getEnvAsInt(key string, defaultVal int) int {
	str := os.Getenv(key)
	if v, err := strconv.Atoi(str); err == nil {
		return v
	}

	return defaultVal
}
