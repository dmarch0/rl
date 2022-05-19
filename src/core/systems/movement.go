package systems

import (
	"rl/src/core/ecs"
)

func MovementSystem(world *ecs.World, resources *ecs.Resources) {
	for _, e := range world.Entities {
		velocity := e.Velocity
		transform := e.Transform
		movementIntention := e.MovementIntention
		if velocity != nil && transform != nil && movementIntention != nil {
			transform.Position = movementIntention.Where
			e.MovementIntention = nil
		}
	}
}