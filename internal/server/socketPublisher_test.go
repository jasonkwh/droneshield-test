package server

import (
	"encoding/json"
	"runtime"
	"testing"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/golang/mock/gomock"
	"github.com/gomodule/redigo/redis"
	"github.com/jasonkwh/droneshield-test/internal/model"
	"github.com/jasonkwh/droneshield-test/test/mocks"
	"github.com/onsi/gomega"
	"go.uber.org/zap"
	"nhooyr.io/websocket"
)

func Test_SocketPublisher_PublishLoop(t *testing.T) {
	tests := []struct {
		name           string
		coordinate     model.Coordinate
		wantCoordinate model.Coordinate
	}{
		{
			name: "normal redis receive to websocker write",
			coordinate: model.Coordinate{
				Latitude:  5.0,
				Longitude: 4.0,
				Altitude:  100.0,
			},
			wantCoordinate: model.Coordinate{
				Latitude:  5.0,
				Longitude: 4.0,
				Altitude:  100.0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// setup
			g := gomega.NewWithT(t)
			ctrl := gomock.NewController(t)
			redisPsCh := make(chan []byte)
			done := make(chan struct{})
			zl, _ := zap.NewDevelopment()
			s, sp, tm := newMockSubAndSockPub(ctrl, redisPsCh, done, zl)

			// initialise mocked clock for testing
			mockedClock := clock.NewMock()
			sp.clock = mockedClock

			bCoor, err := json.Marshal(tt.coordinate)
			g.Expect(err).To(gomega.BeNil())
			wantBCoor, err := json.Marshal(tt.wantCoordinate)
			g.Expect(err).To(gomega.BeNil())

			// expects
			tm.mockRedisPubSubConn.EXPECT().Receive().Return(redis.Message{
				Data: bCoor,
			}).MinTimes(1)
			tm.mockWebsocketConn.EXPECT().Write(gomock.Any(), websocket.MessageText, wantBCoor).Return(nil).MinTimes(1)
			tm.mockRedisPubSubConn.EXPECT().Close().Return(nil)

			// listen on redis pubsub
			go s.Listen()
			go sp.PublishLoop()

			// sleep 2 seconds and close
			time.Sleep(2 * time.Second)

			// advance the clock to tick
			runtime.Gosched()
			mockedClock.Add(sp.msgInterval)

			// gratefully close the client
			err = sp.Close()
			g.Expect(err).To(gomega.BeNil())
		})
	}
}

type testMocks struct {
	mockRedisPubSubConn *mocks.MockRedisPubSubConn
	mockWebsocketConn   *mocks.MockWebsocketConn
}

func newMockSubAndSockPub(ctrl *gomock.Controller, redisPsCh chan []byte, done chan struct{}, zl *zap.Logger) (*subscriber, *socketPublisher, *testMocks) {
	tm := testMocks{
		mockRedisPubSubConn: mocks.NewMockRedisPubSubConn(ctrl),
		mockWebsocketConn:   mocks.NewMockWebsocketConn(ctrl),
	}

	s := subscriber{
		ch:     redisPsCh,
		zl:     zl,
		psConn: tm.mockRedisPubSubConn,
	}

	sp := socketPublisher{
		conn:        tm.mockWebsocketConn,
		redisPsCh:   redisPsCh,
		done:        done,
		msgInterval: 1 * time.Second,
		zl:          zl,
	}

	return &s, &sp, &tm
}
