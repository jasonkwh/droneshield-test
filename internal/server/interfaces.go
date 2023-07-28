package server

import "io"

type Subscriber interface {
	Listen() error

	io.Closer
}

type SocketPublisher interface {
	PublishLoop() error
	SendMessage(msg []byte) error

	io.Closer
}
