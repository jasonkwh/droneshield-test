package redis

import (
	"io"

	"github.com/gomodule/redigo/redis"
)

type Pool interface {
	Get() redis.Conn
	Stats() redis.PoolStats
	ActiveCount() int
	IdleCount() int

	io.Closer
}
