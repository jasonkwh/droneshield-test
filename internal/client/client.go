package client

import (
	"encoding/json"
	"sync"

	"github.com/gomodule/redigo/redis"
	"github.com/jasonkwh/droneshield-test/internal/config"
	"github.com/jasonkwh/droneshield-test/internal/model"
	"go.uber.org/zap"
)

type client struct {
	lock       sync.Mutex // using sync.Mutex lock to avoid race condition
	rconn      redis.Conn
	coordinate *model.Coordinate
	psChan     string
	done       chan struct{}
	zl         *zap.Logger
}

func NewClient(rcfg config.RedisConfig, zl *zap.Logger) (DroneClient, error) {
	var err error
	cl := &client{
		psChan:     rcfg.PubSubChannel,
		done:       make(chan struct{}),
		zl:         zl,
		coordinate: &model.Coordinate{},
	}

	cl.rconn, err = redis.Dial("tcp", rcfg.Endpoints, redis.DialDatabase(rcfg.Database))
	if err != nil {
		zl.Error("failed to initialize redis connection", zap.Error(err))
		return nil, err
	}

	// start sending coordinate after drone intialized
	go cl.sendCoordinate()

	// take off by default
	cl.Movement(model.MovementTakeOff)
	return cl, nil
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
	for {
		select {
		case <-cl.done:
			return
		default:
			/*
			* TODO: use timeticker, send in each seconds
			 */
			cl.lock.Lock()
			bCoor, err := json.Marshal(cl.coordinate)
			if err != nil {
				cl.zl.Error("failed to marshal coordinate into json", zap.Error(err))
				continue
			}
			cl.lock.Unlock()

			cl.zl.Info("sending coordinates", zap.Float64("lat", cl.coordinate.Latitude), zap.Float64("lot", cl.coordinate.Longitude), zap.Float64("alt", cl.coordinate.Altitude))
			err = cl.rconn.Send("PUBLISH", cl.psChan, bCoor)
			if err != nil {
				cl.zl.Error("failed to publish coordinate to redis pubsub", zap.Error(err))
				continue
			}
		}

	}
}

func (cl *client) Close() error {
	if err := cl.rconn.Close(); err != nil {
		cl.zl.Error("failed to close redis connection", zap.Error(err))
		return err
	}

	close(cl.done)
	return nil
}
