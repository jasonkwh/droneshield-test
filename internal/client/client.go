package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/gomodule/redigo/redis"
	"github.com/jasonkwh/droneshield-test-upstream/svc/dronev1"
	"github.com/jasonkwh/droneshield-test/internal/config"
	"github.com/jasonkwh/droneshield-test/internal/model"
	"go.uber.org/multierr"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type client struct {
	lock        sync.Mutex // using sync.Mutex lock to avoid race condition
	rconn       redis.Conn
	coordinate  model.Coordinate
	takeOff     bool
	psChan      string
	done        chan struct{}
	msgInterval time.Duration
	clock       clock.Clock

	//grpc
	listener net.Listener
	server   *grpc.Server

	// please refer to
	// # https://stackoverflow.com/questions/65079032/grpc-with-mustembedunimplemented-method
	dronev1.UnimplementedDroneServiceServer

	zl *zap.Logger
}

func NewClient(rcfg config.RedisConfig, scfg config.ServerConfig, windSimulation bool, zl *zap.Logger) (*client, error) {
	var err error

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", scfg.Port))
	if err != nil {
		return nil, fmt.Errorf("unable to listen: %v", err)
	}

	cl := &client{
		listener:    listener,
		psChan:      rcfg.PubSubChannel,
		done:        make(chan struct{}),
		msgInterval: 1 * time.Second,
		clock:       clock.New(),
		zl:          zl,
	}

	// register grpc server
	cl.server = grpc.NewServer()
	dronev1.RegisterDroneServiceServer(cl.server, cl)

	// create redis connection
	cl.rconn, err = redis.Dial("tcp", rcfg.Endpoints, redis.DialDatabase(rcfg.Database))
	if err != nil {
		zl.Error("failed to initialize redis connection", zap.Error(err))
		return nil, err
	}

	// simulate the wind effect, just for fun :)
	if windSimulation {
		go cl.windSimulation()
	}

	// start sending coordinate after drone intialized
	go cl.sendCoordinate()

	cl.zl.Info("drone client server is running")
	return cl, nil
}

func (cl *client) GetCoordinate(ctx context.Context, r *dronev1.GetCoordinateRequest) (*dronev1.GetCoordinateResponse, error) {
	cl.zl.Info("received GetCoordinate() call")

	cl.lock.Lock()
	coord := cl.coordinate
	cl.lock.Unlock()

	return &dronev1.GetCoordinateResponse{
		Coordinate: &dronev1.Coordinate{
			Latitude:  coord.Latitude,
			Longitude: coord.Longitude,
			Altitude:  coord.Altitude,
		},
		Time: timestamppb.Now(),
	}, nil
}

func (cl *client) Movement(ctx context.Context, r *dronev1.MovementRequest) (*dronev1.GetCoordinateResponse, error) {
	cl.zl.Info("received Movement() call")

	// set drone movement
	cl.setMovement(mapMovement(r.Movement))

	// call GetCoordinate()
	return cl.GetCoordinate(ctx, &dronev1.GetCoordinateRequest{})
}

func (cl *client) setMovement(move model.Movement) {
	cl.zl.Info("setting movement")

	cl.lock.Lock()
	switch move {
	case model.MovementTakeOff:
		if !cl.takeOff {
			cl.zl.Info("drone take off")
			cl.takeOff = true
			cl.coordinate.Altitude = 100
		}
	case model.MovementUp:
		if cl.takeOff {
			cl.zl.Info("drone go up")
			cl.coordinate.Altitude++
		}
	case model.MovementDown:
		if cl.takeOff {
			cl.zl.Info("drone go down")
			cl.coordinate.Altitude--
		}
	case model.MovementLeft:
		if cl.takeOff {
			cl.zl.Info("drone go left")
			cl.coordinate.Longitude++
		}
	case model.MovementRight:
		if cl.takeOff {
			cl.zl.Info("drone go right")
			cl.coordinate.Longitude--
		}
	case model.MovementForward:
		if cl.takeOff {
			cl.zl.Info("drone go forward")
			cl.coordinate.Latitude++
		}
	case model.MovementBackward:
		if cl.takeOff {
			cl.zl.Info("drone go backward")
			cl.coordinate.Latitude--
		}
	case model.MovementCircle:
		if cl.takeOff {
			cl.zl.Info("drone doing circle")
		}
	case model.MovementZigzag:
		if cl.takeOff {
			cl.zl.Info("drone doing zigzag")
		}
	case model.MovementFigure8:
		if cl.takeOff {
			cl.zl.Info("drone doing figure8")
		}
	case model.MovementLanding:
		if cl.takeOff {
			cl.zl.Info("drone landing")
			cl.coordinate.Altitude = 0
			cl.takeOff = false
		}
	default:
		cl.zl.Info("drone remains stable")
	}
	cl.lock.Unlock()
}

func (cl *client) sendCoordinate() {
	t := cl.clock.Ticker(cl.msgInterval)

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

			cl.zl.Info("sending coordinates", zap.Float32("lat", cl.coordinate.Latitude), zap.Float32("lot", cl.coordinate.Longitude), zap.Float32("alt", cl.coordinate.Altitude))
			err = cl.rconn.Send("PUBLISH", cl.psChan, bCoor)
			if err != nil {
				cl.zl.Error("failed to publish coordinate to redis pubsub", zap.Error(err))
				continue
			}
		}

	}
}

// Run starts the server run
func (cl *client) Run() error {
	return cl.server.Serve(cl.listener)
}

func (cl *client) Close() error {
	// landing the drone
	cl.setMovement(model.MovementLanding)

	if cl.server != nil {
		cl.server.GracefulStop()
	}

	var errs error
	if cl.listener != nil {
		err := cl.listener.Close()
		if err != nil {
			errs = multierr.Append(errs, err)
		}
	}

	if err := cl.rconn.Close(); err != nil {
		errs = multierr.Append(errs, err)
	}

	close(cl.done)
	return errs
}
