package systems

import (
	"rl/src/core/ecs"
)

func PresentSystem(world *ecs.World, resources *ecs.Resources) {
	resources.Renderer.Present()
}
