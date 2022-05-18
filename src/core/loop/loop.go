package loop

import (
	"rl/src/core/components"
	"rl/src/core/ecs"
	"rl/src/core/systems"
	"rl/src/core/utils"

	"github.com/veandco/go-sdl2/sdl"
)

func GameLoop(window *sdl.Window, renderer *sdl.Renderer) {
	resources := ecs.Resources{
		Running:  true,
		Renderer: renderer,
	}
	world := ecs.World{
		Entities: make([]ecs.Entity, 0),
	}

	builder := ecs.CreateScheduleBuilder()
	builder.AddSystem(systems.InputSystem)
	builder.Flush()
	builder.AddSystem(systems.MovementSystem)
	builder.Flush()
	scheduler := builder.Build()

	scheduler.World = &world
	scheduler.Resources = &resources

	e := ecs.Entity{
		Player: &components.Player{},
		SimpleRender: &components.SimpleRender{
			Color: utils.RGBA{
				R: 255,
				G: 0,
				B: 0,
				A: 255,
			},
			W: 100,
			H: 100,
		},
		Position: &components.Position{
			Vector: utils.Vector{10, 10},
		},
		Velocity: &components.Velocity{Vector: utils.Vector{0, 0}},
	}
	scheduler.World.AddEntity(e)

	for resources.Running {
		systems.PrepearSystem(&world, &resources)
		scheduler.Execute()
		systems.RenderSystem(&world, &resources)
		systems.PresentSystem(&world, &resources)
		sdl.Delay(16)
	}
}
