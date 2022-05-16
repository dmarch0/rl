package systems

import (
	"fmt"
	"rl/src/core/ecs"
)

func PrepearSystem(world *ecs.World, resources *ecs.Resources) {
	fmt.Println("Prepear system")
	resources.Renderer.SetDrawColor(0, 0, 0, 255)
	resources.Renderer.Clear()
}
