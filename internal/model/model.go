package model

type Coordinate struct {
	Latitude  float32
	Longitude float32
	Altitude  float32
}

type Movement int32

const (
	MovementUnknown = iota
	MovementTakeOff
	MovementUp
	MovementDown
	MovementRight
	MovementLeft
	MovementForward
	MovementBackward
	MovementFigure8
	MovementCircle
	MovementZigzag
	MovementLanding
)
