package systems

import (
	"fmt"
	"rl/src/core/ecs"

	"github.com/veandco/go-sdl2/sdl"
)

func WaitSystem(world *ecs.World, resources *ecs.Resources) {
	fmt.Println("Wait")
	sdl.Delay(16)
}
