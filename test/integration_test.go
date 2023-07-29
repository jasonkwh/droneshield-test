//go:build integration
// +build integration

package integration

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/jasonkwh/droneshield-test/internal/model"
	"github.com/onsi/gomega"
	"go.uber.org/zap"
	"nhooyr.io/websocket"
)

// TestIntegration - simple integration test, initialize the websocket client and verifies the result
func TestIntegration(t *testing.T) {
	// setup
	ctx := context.Background()
	zl, _ := zap.NewDevelopment()
	g := gomega.NewWithT(t)

	conn, _, err := websocket.Dial(ctx, "ws://localhost:8080", &websocket.DialOptions{
		CompressionMode: websocket.CompressionDisabled,
	})
	if err != nil {
		t.Fatalf("failed to create websocket connection: %v", err)
	}

	// sleep
	time.Sleep(2 * time.Second)

	// reading status success message
	_, msg, err := conn.Read(ctx)
	if err != nil {
		zl.Fatal("failed to read websocket message", zap.Error(err))
	}

	coord, err := mapMessage(msg)
	if err != nil {
		zl.Fatal("failed to map websocket message", zap.Error(err))
	}

	// expect
	// if it gets altitude 100.0, then
	// it means we are getting the coordinates
	g.Expect(coord.Altitude).To(gomega.Equal(100.0))

	// close websocket connection
	conn.Close(websocket.StatusNormalClosure, "")
}

func mapMessage(msg []byte) (model.Coordinate, error) {
	coord := model.Coordinate{}

	err := json.Unmarshal(msg, &coord)
	if err != nil {
		return coord, err
	}

	return coord, nil
}
