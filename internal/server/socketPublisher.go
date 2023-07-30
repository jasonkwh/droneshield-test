package server

import (
	"context"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/jasonkwh/droneshield-test/internal/config"
	"go.uber.org/zap"
	"nhooyr.io/websocket"
)

type socketPublisher struct {
	conn        WebsocketConn
	rcfg        config.RedisConfig
	redisPsCh   chan []byte
	done        chan struct{}
	msgInterval time.Duration
	clock       clock.Clock

	zl *zap.Logger
}

func NewSocketPublisher(conn *websocket.Conn, rcfg config.RedisConfig, redisPsCh chan []byte, done chan struct{}, zl *zap.Logger) (SocketPublisher, error) {
	sp := &socketPublisher{
		conn:        conn,
		rcfg:        rcfg,
		redisPsCh:   redisPsCh,
		done:        done,
		msgInterval: 1 * time.Second,
		clock:       clock.New(),

		zl: zl,
	}

	go sp.gracefulCloseListener()
	return sp, nil
}

func (sp *socketPublisher) PublishLoop() error {
	sp.zl.Info("start the publish_loop")

	t := sp.clock.Ticker(sp.msgInterval)

	for {
		select {
		case <-t.C:
			select {
			case msg := <-sp.redisPsCh:
				sp.zl.Info("received redis pubsub message")

				err := sp.SendMessage(msg)
				if err != nil {
					sp.zl.Error("failed to send message into websocket", zap.Error(err))
					continue
				}
			default:
				sp.zl.Info("no redis pubsub message")
			}
		case <-sp.done:
			t.Stop()
			return nil
		}
	}
}

func (sp *socketPublisher) SendMessage(msg []byte) error {
	return sp.conn.Write(context.Background(), websocket.MessageText, msg)
}

func (sp *socketPublisher) Close() error {
	close(sp.done)

	return nil
}

// gracefulCloseListener - gracefully close socket publisher if websocket connection is closed
func (sp *socketPublisher) gracefulCloseListener() {
	for {
		_, _, err := sp.conn.Read(context.Background())
		if err != nil {
			sp.Close()
			return
		}
	}
}
