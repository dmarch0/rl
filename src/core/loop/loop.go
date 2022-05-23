package loop

import (
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
		Entities: make([]*ecs.Entity, 0),
	}

	builder := ecs.CreateScheduleBuilder()
	builder.AddSystem(systems.InputSystem)
	builder.Flush()
	builder.AddSystem(systems.VelocitySystem)
	builder.Flush()
	builder.AddSystem(systems.CollisionSystem)
	builder.Flush()
	builder.AddSystem(systems.MovementSystem)
	builder.Flush()
	builder.AddSystem(systems.DeathSystem)
	builder.Flush()
	scheduler := builder.Build()

	scheduler.World = &world
	scheduler.Resources = &resources

	e := ecs.Entity{}
	ecs.BindPlayer(&e, &ecs.Player{})
	ecs.BindTransform(&e, &ecs.Transform{
		Position: utils.Vector{10, 10},
		Scale:    utils.Vector{10, 10},
	})
	ecs.BindVelocity(&e, &ecs.Velocity{Value: utils.Vector{0, 0}})
	ecs.BindSimpleRender(&e, &ecs.SimpleRender{
		Color: utils.RGBA{
			R: 255,
			G: 255,
			B: 255,
			A: 255,
		},
	})
	ecs.BindCollider(&e, &ecs.Collider{IsSolid: true})
	scheduler.World.AddEntity(&e)

	block := ecs.Entity{}
	ecs.BindTransform(&block, &ecs.Transform{
		Position: utils.Vector{30, 30},
		Scale:    utils.Vector{30, 30},
	})
	ecs.BindSimpleRender(&block, &ecs.SimpleRender{
		Color: utils.RGBA{
			R: 255,
			G: 0,
			B: 0,
			A: 255,
		},
	})
	ecs.BindCollider(&block, &ecs.Collider{IsSolid: true})
	ecs.BindHealth(&block, &ecs.Health{Value: 20.0})
	scheduler.World.AddEntity(&block)

	for resources.Running {
		systems.PrepearSystem(&world, &resources)
		scheduler.Execute()
		systems.RenderSystem(&world, &resources)
		systems.PresentSystem(&world, &resources)
		sdl.Delay(16)
	}
}
