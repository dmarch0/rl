package systems

import (
	"fmt"
	"rl/src/core/ecs"
)

func PresentSystem(world *ecs.World, resources *ecs.Resources) {
	fmt.Println("Present system")
	resources.Renderer.Present()
}
