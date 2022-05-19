package systems

import (
	"rl/src/core/ecs"
)

func VelocitySystem(world *ecs.World, resources *ecs.Resources) {
	for _, e := range world.Entities {
		velocity := e.Velocity
		transform := e.Transform
		if velocity != nil && transform != nil && !velocity.Value.IsZero() {
			newPosition := transform.Position
			newPosition.Add(velocity.Value)
			movementIntention := ecs.MovementIntention{
				Where: newPosition,
			}
			ecs.BindMovementIntention(e, &movementIntention)
		}
	}
}
