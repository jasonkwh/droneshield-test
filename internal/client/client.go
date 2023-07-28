package client

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/jasonkwh/droneshield-test/internal/config"
	"github.com/jasonkwh/droneshield-test/internal/model"
	"github.com/jasonkwh/droneshield-test/internal/redis"
	"go.uber.org/zap"
)

type client struct {
	lock        sync.Mutex // using sync.Mutex lock to avoid race condition
	rpool       redis.Pool
	coordinate  *model.Coordinate
	psChan      string
	done        chan struct{}
	msgInterval time.Duration

	zl *zap.Logger
}

func NewClient(rcfg config.RedisConfig, zl *zap.Logger) DroneClient {
	cl := &client{
		rpool:       redis.NewRedisPool(rcfg),
		psChan:      rcfg.PubSubChannel,
		done:        make(chan struct{}),
		coordinate:  &model.Coordinate{},
		msgInterval: 1 * time.Second,
		zl:          zl,
	}

	// start sending coordinate after drone intialized
	go cl.sendCoordinate()

	// take off by default
	cl.Movement(model.MovementTakeOff)
	return cl
}

func (cl *client) Movement(move model.Movement) {
	cl.lock.Lock()
	switch move {
	case model.MovementTakeOff:
		cl.coordinate.Altitude = 100.0
	case model.MovementUp:
		cl.coordinate.Altitude++
	case model.MovementDown:
		cl.coordinate.Altitude--
	case model.MovementLeft:
		cl.coordinate.Longitude++
	case model.MovementRight:
		cl.coordinate.Longitude--
	case model.MovementForward:
		cl.coordinate.Latitude++
	case model.MovementBackward:
		cl.coordinate.Latitude--
	}
	cl.lock.Unlock()
}

func (cl *client) sendCoordinate() {
	// create redis connection
	conn := cl.rpool.Get()
	defer conn.Close()

	t := time.NewTicker(cl.msgInterval)

	for {
		select {
		case <-cl.done:
			t.Stop()
			return
		case <-t.C:
			cl.lock.Lock()
			bCoor, err := json.Marshal(cl.coordinate)
			if err != nil {
				cl.zl.Error("failed to marshal coordinate into json", zap.Error(err))
				continue
			}
			cl.lock.Unlock()

			cl.zl.Info("sending coordinates", zap.Float64("lat", cl.coordinate.Latitude), zap.Float64("lot", cl.coordinate.Longitude), zap.Float64("alt", cl.coordinate.Altitude))
			err = conn.Send("PUBLISH", cl.psChan, bCoor)
			if err != nil {
				cl.zl.Error("failed to publish coordinate to redis pubsub", zap.Error(err))
				continue
			}
		}

	}
}

func (cl *client) Close() error {
	if err := cl.rpool.Close(); err != nil {
		cl.zl.Error("failed to close redis connection", zap.Error(err))
		return err
	}

	close(cl.done)
	return nil
}
