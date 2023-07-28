package redis

import (
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/jasonkwh/droneshield-test/internal/config"
)

func NewRedisPool(rcfg config.RedisConfig) Pool {
	return &redis.Pool{
		MaxIdle:     rcfg.MaxIdleCons,
		MaxActive:   rcfg.MaxActiveCons,
		IdleTimeout: rcfg.IdleTimeout,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", rcfg.Endpoints, redis.DialDatabase(rcfg.Database))
			if err != nil {
				return nil, err
			}
			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}
