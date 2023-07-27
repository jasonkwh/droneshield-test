package server

type Subscriber interface {
	Listen() error
	Close() error
}
