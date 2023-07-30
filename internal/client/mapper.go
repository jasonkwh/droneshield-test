package client

import (
	"github.com/jasonkwh/droneshield-test-upstream/svc/dronev1"
	"github.com/jasonkwh/droneshield-test/internal/model"
)

func mapMovement(m dronev1.Movement_Enum) model.Movement {
	switch m {
	case dronev1.Movement_takeoff:
		return model.MovementTakeOff
	case dronev1.Movement_up:
		return model.MovementUp
	case dronev1.Movement_down:
		return model.MovementDown
	case dronev1.Movement_right:
		return model.MovementRight
	case dronev1.Movement_left:
		return model.MovementLeft
	case dronev1.Movement_forward:
		return model.MovementForward
	case dronev1.Movement_backward:
		return model.MovementBackward
	case dronev1.Movement_figure8:
		return model.MovementFigure8
	case dronev1.Movement_circle:
		return model.MovementCircle
	case dronev1.Movement_zigzag:
		return model.MovementZigzag
	case dronev1.Movement_landing:
		return model.MovementLanding
	default:
		return model.MovementUnknown
	}
}
