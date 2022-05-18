package systems

import (
	"fmt"
	"rl/src/core/ecs"

	"github.com/veandco/go-sdl2/sdl"
)

func RenderSystem(world *ecs.World, resources *ecs.Resources) {
	for _, e := range world.Entities {
		fmt.Println(e)
		position := e.Position
		render := e.SimpleRender

		if render != nil && position != nil {
			resources.Renderer.SetDrawColor(render.Color.R, render.Color.G, render.Color.B, render.Color.A)
			rect := sdl.Rect{
				X: int32(position.Vector.X),
				Y: int32(position.Vector.Y),
				W: int32(render.W),
				H: int32(render.H),
			}
			resources.Renderer.DrawRect(&rect)
		}
	}
}
