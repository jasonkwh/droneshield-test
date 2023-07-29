package server

import (
	"context"
	"io"
	"time"

	"nhooyr.io/websocket"
)

type Subscriber interface {
	Listen() error

	io.Closer
}

type SocketPublisher interface {
	PublishLoop() error
	SendMessage(msg []byte) error

	io.Closer
}

type WebsocketConn interface {
	// for test mocking purposes
	Close(code websocket.StatusCode, reason string) error
	CloseRead(ctx context.Context) context.Context
	Ping(ctx context.Context) error
	Read(ctx context.Context) (websocket.MessageType, []byte, error)
	Reader(ctx context.Context) (websocket.MessageType, io.Reader, error)
	SetReadLimit(n int64)
	Subprotocol() string
	Write(ctx context.Context, typ websocket.MessageType, p []byte) error
	Writer(ctx context.Context, typ websocket.MessageType) (io.WriteCloser, error)
}

type RedisPubSubConn interface {
	// for test mocking purposes
	PSubscribe(channel ...interface{}) error
	PUnsubscribe(channel ...interface{}) error
	Ping(data string) error
	Receive() interface{}
	ReceiveContext(ctx context.Context) interface{}
	ReceiveWithTimeout(timeout time.Duration) interface{}
	Subscribe(channel ...interface{}) error
	Unsubscribe(channel ...interface{}) error

	io.Closer
}
