package server

import "io"

type Subscriber interface {
	Listen() error

	io.Closer
}
