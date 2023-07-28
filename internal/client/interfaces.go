package client

import (
	"io"

	"github.com/jasonkwh/droneshield-test/internal/model"
)

type DroneClient interface {
	Movement(model.Movement)

	io.Closer
}
