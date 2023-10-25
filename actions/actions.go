package actions

import (
	"github.com/google/uuid"
	"time"
)

type Action interface{}

type Direction int

type MoveAction struct {
	Action
	Direction Direction
	ID        uuid.UUID
	Created   time.Time
}
