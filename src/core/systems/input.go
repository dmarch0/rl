package systems

import (
	"fmt"
	"rl/src/core/ecs"

	"github.com/veandco/go-sdl2/sdl"
)

func InputSystem(world *ecs.World, resources *ecs.Resources) {
	fmt.Println("Input system")
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			println("Quit")
			resources.Running = false
			break
		}
	}
}
