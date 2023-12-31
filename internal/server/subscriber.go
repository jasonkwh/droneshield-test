package server

import (
	"sync"

	"github.com/gomodule/redigo/redis"
	"github.com/jasonkwh/droneshield-test/internal/config"
	"go.uber.org/zap"
)

type subscriber struct {
	lock   sync.Mutex
	psConn RedisPubSubConn
	close  bool
	ch     chan []byte

	zl *zap.Logger
}

func NewSubscriber(rcfg config.RedisConfig, ch chan []byte, zl *zap.Logger) (Subscriber, error) {
	s := &subscriber{
		ch: ch,
		zl: zl,
	}

	// create redis connection
	conn, err := redis.Dial("tcp", rcfg.Endpoints, redis.DialDatabase(rcfg.Database))
	if err != nil {
		zl.Error("failed to create new redis connection", zap.Error(err))
		return nil, err
	}

	// subscribe channel
	s.psConn = &redis.PubSubConn{Conn: conn}
	if err := s.psConn.Subscribe(rcfg.PubSubChannel); err != nil {
		zl.Error("failed to subscribe redis pubsub", zap.Error(err))
		return nil, err
	}

	return s, nil
}

func (s *subscriber) Listen() error {
	for {
		s.lock.Lock()
		if s.close {
			return nil
		}
		msg := s.psConn.Receive()
		s.lock.Unlock()

		switch v := msg.(type) {
		case error:
			return v
		case redis.Message:
			//try to unmarshal it
			// coord := model.Coordinate{}
			// json.Unmarshal(v.Data, &coord)
			// s.zl.Info("sending coordinates", zap.Float64("lat", coord.Latitude), zap.Float64("lot", coord.Longitude), zap.Float64("alt", coord.Altitude))

			s.ch <- v.Data
		case redis.Subscription:
			s.zl.Debug("redis channel subscription received", zap.Any("channel", v.Channel), zap.Any("kind", v.Kind))
			continue
		}
	}
}

func (s *subscriber) Close() error {
	s.lock.Lock()
	s.close = true
	err := s.psConn.Close()
	s.lock.Unlock()

	return err
}
