package config

import "time"

type ServerConfig struct {
	Port string
}

type RedisConfig struct {
	PubSubChannel string
	Endpoints     string
	Database      int
	IdleTimeout   time.Duration
	MaxIdleCons   int
	MaxActiveCons int
}
