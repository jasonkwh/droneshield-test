package client

import (
	"encoding/json"
	"runtime"
	"testing"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/golang/mock/gomock"
	"github.com/jasonkwh/droneshield-test/internal/model"
	"github.com/jasonkwh/droneshield-test/test/mocks"
	"github.com/onsi/gomega"
	"go.uber.org/zap"
)

const psChan = "test_drone_channel"

func Test_DroneClient_SendCoordinate(t *testing.T) {
	tests := []struct {
		name           string
		movements      []model.Movement
		wantCoordinate *model.Coordinate
	}{
		{
			name: "normal send coordinate test",
			movements: []model.Movement{
				model.MovementUp,
				model.MovementForward,
				model.MovementRight,
			},
			wantCoordinate: &model.Coordinate{
				Altitude:  101.0,
				Latitude:  1.0,
				Longitude: -1.0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := gomega.NewWithT(t)
			ctrl := gomock.NewController(t)
			cl, tm := newMockClient(ctrl)

			// initialise mocked clock for testing
			mockedClock := clock.NewMock()
			cl.clock = mockedClock

			if len(tt.movements) != 0 {
				for _, movement := range tt.movements {
					cl.Movement(movement)
				}
			}

			bCoor, err := json.Marshal(tt.wantCoordinate)
			g.Expect(err).To(gomega.BeNil())

			// expects the right coordinate
			tm.mockRedisConn.EXPECT().Send("PUBLISH", psChan, bCoor)
			tm.mockRedisConn.EXPECT().Close().Return(nil)

			// start sending coordinate after drone intialized
			go cl.SendCoordinate()

			// advance the clock to tick
			runtime.Gosched()
			mockedClock.Add(cl.msgInterval)

			// gratefully close the client
			err = cl.Close()
			g.Expect(err).To(gomega.BeNil())
		})
	}
}

type testMocks struct {
	mockRedisConn *mocks.MockConn
}

func newMockClient(ctrl *gomock.Controller) (*client, *testMocks) {
	tm := testMocks{
		mockRedisConn: mocks.NewMockConn(ctrl),
	}

	zl, _ := zap.NewDevelopment()

	cl := client{
		psChan:      psChan,
		done:        make(chan struct{}),
		coordinate:  &model.Coordinate{},
		msgInterval: 1 * time.Second,
		rconn:       tm.mockRedisConn,
		zl:          zl,
	}

	// take off by default
	cl.Movement(model.MovementTakeOff)

	return &cl, &tm
}
