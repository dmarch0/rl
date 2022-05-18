package systems

import (
	"fmt"
	"rl/src/core/ecs"
)

func MovementSystem(world *ecs.World, resources *ecs.Resources) {
	for _, e := range world.Entities {
		velocity := e.Velocity
		position := e.Position
		fmt.Println("Movement: got velo and pos", *velocity, *position)
		if velocity != nil && position != nil {
			position.Vector.Add(velocity.Vector)
		}
	}
}
