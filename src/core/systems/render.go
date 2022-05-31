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
			rect := sdl.Rect{
				X: int32(transform.Position.X),
				Y: int32(transform.Position.Y),
				W: int32(transform.Scale.X),
				H: int32(transform.Scale.Y),
			}
			resources.Renderer.SetDrawColor(render.Fill.R, render.Fill.G, render.Fill.B, render.Fill.A)
			resources.Renderer.FillRect(&rect)
			resources.Renderer.SetDrawColor(render.Border.R, render.Border.G, render.Border.B, render.Border.A)
			resources.Renderer.DrawRect(&rect)
		}
	}
}
