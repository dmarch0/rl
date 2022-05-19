package systems

import (
	"rl/src/core/ecs"

	"github.com/veandco/go-sdl2/sdl"
)

func RenderSystem(world *ecs.World, resources *ecs.Resources) {
	for _, e := range world.Entities {
		transform := e.Transform
		render := e.SimpleRender

		if render != nil && transform != nil {
			resources.Renderer.SetDrawColor(render.Color.R, render.Color.G, render.Color.B, render.Color.A)
			rect := sdl.Rect{
				X: int32(transform.Position.X),
				Y: int32(transform.Position.Y),
				W: int32(transform.Scale.X),
				H: int32(transform.Scale.Y),
			}
			resources.Renderer.DrawRect(&rect)
		}
	}
}
