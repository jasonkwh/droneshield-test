package model

type Coordinate struct {
	Latitude  float64
	Longitude float64
	Altitude  float64
}

type Movement int64

const (
	MovementTakeOff = iota
	MovementUp
	MovementDown
	MovementRight
	MovementLeft
	MovementForward
	MovementBackward
)
