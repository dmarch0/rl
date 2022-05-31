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
	ecs.BindDamager(&e, &ecs.Damager{})
	scheduler.World.AddEntity(&e)

	commonScale := utils.Vector{20, 20}
	ecs.SpawnSimpleBlock(scheduler.World, commonScale, utils.Vector{0, 20})
	ecs.SpawnSimpleBlock(scheduler.World, commonScale, utils.Vector{20, 20})
	ecs.SpawnSimpleBlock(scheduler.World, commonScale, utils.Vector{40, 20})
	ecs.SpawnSimpleBlock(scheduler.World, commonScale, utils.Vector{60, 20})
	ecs.SpawnSimpleBlock(scheduler.World, commonScale, utils.Vector{80, 20})

	for resources.Running {
		systems.PrepearSystem(&world, &resources)
		scheduler.Execute()
		systems.RenderSystem(&world, &resources)
		systems.PresentSystem(&world, &resources)
		sdl.Delay(16)
	}
}
