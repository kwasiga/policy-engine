package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	DatabaseURL        string
	CedarAgentURL      string
	CedarAgentToken    string
	RESTListenAddr     string
	GRPCListenAddr     string
	PolicyCacheTTL     time.Duration
	LogLevel           string
}

func Load() Config {
	return Config{
		DatabaseURL:     getenv("DATABASE_URL", "postgres://policy_engine:policy_engine@localhost:5432/policy_engine?sslmode=disable"),
		CedarAgentURL:   getenv("CEDAR_AGENT_URL", "http://localhost:8180"),
		CedarAgentToken: getenv("CEDAR_AGENT_AUTH_TOKEN", ""),
		RESTListenAddr:  getenv("REST_LISTEN_ADDR", "0.0.0.0:8080"),
		GRPCListenAddr:  getenv("GRPC_LISTEN_ADDR", "0.0.0.0:9090"),
		PolicyCacheTTL:  time.Duration(getenvInt("POLICY_CACHE_TTL_SECONDS", 300)) * time.Second,
		LogLevel:        getenv("LOG_LEVEL", "info"),
	}
}

func getenv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func getenvInt(key string, fallback int) int {
	if v := os.Getenv(key); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			return n
		}
	}
	return fallback
}
