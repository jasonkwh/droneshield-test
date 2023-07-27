package config

import "time"

type RedisConfig struct {
	PubSubChannel string
	Endpoints     string
	Database      int
	IdleTimeout   time.Duration
	MaxIdleCons   int
	MaxActiveCons int
}
