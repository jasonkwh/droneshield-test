//go:build integration
// +build integration

package integration

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/jasonkwh/droneshield-test-upstream/svc/dronev1"
	"github.com/jasonkwh/droneshield-test/internal/model"
	"github.com/onsi/gomega"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"nhooyr.io/websocket"
)

// TestIntegration - simple integration test, initialize the websocket client and verifies the result
func TestIntegration(t *testing.T) {
	// setup
	ctx := context.Background()
	zl, _ := zap.NewDevelopment()
	g := gomega.NewWithT(t)

	// initialize grpc client
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("failed to dial grpc: %v", err)
	}
	dronecl := dronev1.NewDroneServiceClient(conn)

	// take off
	resp, err := dronecl.Movement(ctx, &dronev1.MovementRequest{
		Movement: dronev1.Movement_takeoff,
	})
	if err != nil {
		t.Fatalf("failed to take off: %v", err)
	}

	expectAltitude := resp.Coordinate.Altitude

	// initialize websocket client
	wsconn, _, err := websocket.Dial(ctx, "ws://localhost:8080", &websocket.DialOptions{
		CompressionMode: websocket.CompressionDisabled,
	})
	if err != nil {
		t.Fatalf("failed to create websocket connection: %v", err)
	}

	// sleep
	time.Sleep(2 * time.Second)

	// reading status success message
	_, msg, err := wsconn.Read(ctx)
	if err != nil {
		zl.Fatal("failed to read websocket message", zap.Error(err))
	}

	coord, err := mapMessage(msg)
	if err != nil {
		zl.Fatal("failed to map websocket message", zap.Error(err))
	}

	// expect
	// if it gets expected altitude, then
	// it means we are getting the coordinates from websocket
	g.Expect(coord.Altitude).To(gomega.Equal(expectAltitude))

	// close websocket & grpc connection
	wsconn.Close(websocket.StatusNormalClosure, "")
	conn.Close()
}

func mapMessage(msg []byte) (model.Coordinate, error) {
	coord := model.Coordinate{}

	err := json.Unmarshal(msg, &coord)
	if err != nil {
		return coord, err
	}

	return coord, nil
}
