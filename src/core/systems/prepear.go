package systems

import (
	"rl/src/core/ecs"
)

func PrepearSystem(world *ecs.World, resources *ecs.Resources) {
	resources.Renderer.SetDrawColor(0, 0, 0, 255)
	resources.Renderer.Clear()
}
