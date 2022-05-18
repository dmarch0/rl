package ecs

import (
	"rl/src/core/components"
)

type Entity struct {
	Player       *components.Player
	Position     *components.Position
	SimpleRender *components.SimpleRender
	Velocity     *components.Velocity
}
