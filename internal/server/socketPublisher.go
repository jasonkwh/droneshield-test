package server

import (
	"context"

	"github.com/jasonkwh/droneshield-test/internal/config"
	"go.uber.org/zap"
	"nhooyr.io/websocket"
)

type socketPublisher struct {
	conn      *websocket.Conn
	rcfg      config.RedisConfig
	redisPsCh chan []byte
	done      chan struct{}
	zl        *zap.Logger
}

func NewSocketPublisher(conn *websocket.Conn, rcfg config.RedisConfig, redisPsCh chan []byte, done chan struct{}, zl *zap.Logger) (SocketPublisher, error) {
	sp := &socketPublisher{
		conn:      conn,
		rcfg:      rcfg,
		redisPsCh: redisPsCh,
		done:      done,
		zl:        zl,
	}

	go sp.gratefulCloseListener()
	return sp, nil
}

func (sp *socketPublisher) PublishLoop() error {
	sp.zl.Info("start the publish_loop")

	for {
		select {
		case msg := <-sp.redisPsCh:
			sp.zl.Info("received redis pubsub message")

			err := sp.SendMessage(msg)
			if err != nil {
				sp.zl.Error("failed to send message into websocket", zap.Error(err))
				continue
			}
		case <-sp.done:
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

// gratefulCloseListener - gratefully close socket publisher if websocket connection is closed
func (sp *socketPublisher) gratefulCloseListener() {
	for {
		_, _, err := sp.conn.Read(context.Background())
		if err != nil {
			sp.Close()
			return
		}
	}
}
