package systems

import (
	"rl/src/core/ecs"
)

func MovementSystem(world *ecs.World, resources *ecs.Resources) {
	for _, e := range world.Movers {
		velocity := e.Velocity
		transform := e.Transform
		movementIntention := e.MovementIntention
		if velocity != nil && transform != nil && movementIntention != nil {
			world.CollidersHashTable.Remove(e)
			transform.Position = movementIntention.Where
			world.CollidersHashTable.Insert(e)
			e.MovementIntention = nil
		}
	}
}
