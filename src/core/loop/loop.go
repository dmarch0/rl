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
		Entities:           make([]*ecs.Entity, 0),
		CollidersHashTable: ecs.CreateCollidersHashTable(),
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
		Border: utils.RGBA{
			R: 255,
			G: 255,
			B: 255,
			A: 255,
		},
		Fill: utils.RGBA{
			R: 0,
			G: 0,
			B: 0,
			A: 0,
		},
	})
	ecs.BindCollider(&e, &ecs.Collider{IsSolid: true})
	ecs.BindDamager(&e, &ecs.Damager{})
	scheduler.World.AddEntity(&e)

	commonScale := utils.Vector{20, 20}
	for i := 0; i < 30; i++ {
		for j := 0; j < 30; j++ {
			ecs.SpawnSimpleBlock(scheduler.World, commonScale, utils.Vector{i * 20, j * 20})
		}
	}

	for resources.Running {
		systems.PrepearSystem(&world, &resources)
		scheduler.Execute()
		systems.RenderSystem(&world, &resources)
		systems.PresentSystem(&world, &resources)
		sdl.Delay(16)
	}
}
