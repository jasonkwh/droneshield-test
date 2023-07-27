package config

import "time"

type RedisConfig struct {
	Endpoints     string
	Database      int
	IdleTimeout   time.Duration
	MaxIdleCons   int
	MaxActiveCons int
}
