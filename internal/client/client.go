package client

import (
	"encoding/json"
	"sync"

	"github.com/gomodule/redigo/redis"
	"github.com/jasonkwh/droneshield-test/internal/model"
	"go.uber.org/zap"
)

type client struct {
	lock       sync.Mutex
	rpool      *redis.Pool
	coordinate *model.Coordinate
	psChan     string
	done       chan struct{}
	zl         *zap.Logger
}

func NewClient(psChan string, zl *zap.Logger) DroneClient {
	cl := &client{
		psChan: psChan,
		done:   make(chan struct{}),
		zl:     zl,
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

	for {
		select {
		case <-cl.done:
			return
		default:
			cl.lock.Lock()
			bCoor, err := json.Marshal(cl.coordinate)
			if err != nil {
				cl.zl.Error("failed to marshal coordinate into json", zap.Error(err))
				continue
			}
			cl.lock.Unlock()

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
		cl.zl.Error("failed to close redis pool", zap.Error(err))
		return err
	}

	close(cl.done)
	return nil
}
