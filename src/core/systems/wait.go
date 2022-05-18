package systems

import (
	"rl/src/core/ecs"

	"github.com/veandco/go-sdl2/sdl"
)

func WaitSystem(world *ecs.World, resources *ecs.Resources) {
	sdl.Delay(16)
}
