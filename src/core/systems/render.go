package systems

import (
	"fmt"
	"rl/src/core/components"
	"rl/src/core/ecs"

	"github.com/veandco/go-sdl2/sdl"
)

func RenderSystem(world *ecs.World, resources *ecs.Resources) {
	fmt.Println("Render system")
	for _, e := range world.Entities {
		position := ecs.TryGetComponent[components.Position](&e)
		render := ecs.TryGetComponent[components.SimpleRender](&e)
		fmt.Println("Position: ", position)
		fmt.Println("Render: ", render)
		if render != nil && position != nil {
			resources.Renderer.SetDrawColor(render.Color.R, render.Color.G, render.Color.B, render.Color.A)
			rect := sdl.Rect{
				X: int32(position.X),
				Y: int32(position.Y),
				W: int32(render.W),
				H: int32(render.H),
			}
			resources.Renderer.DrawRect(&rect)
		}
	}
}
