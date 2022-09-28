package loop

import (
	"fmt"
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
	ecs.BindVelocity(&e, &ecs.Velocity{
		Direction: utils.Vector{0, 0},
		Value:     2.0,
	})
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
	ecs.BindMover(&e, &ecs.Mover{})
	scheduler.World.AddEntity(&e)

	commonScale := utils.Vector{30.0, 30.0}
	for i := 0.0; i < 20.0; i++ {
		for j := 0.0; j < 20.0; j++ {
			ecs.SpawnSimpleBlock(scheduler.World, commonScale, utils.Vector{i * commonScale.X, j * commonScale.Y})
		}
	}

	fps := 0
	for resources.Running {
		start := sdl.GetTicks()
		systems.PrepearSystem(&world, &resources)
		scheduler.Execute()
		systems.RenderSystem(&world, &resources)

		fpsRect := sdl.Rect{0, 0, 100, 24}
		utils.PrintText(fmt.Sprintf("FPS: %d", fps), &fpsRect, renderer)

		systems.PresentSystem(&world, &resources)
		sdl.Delay(8)

		end := sdl.GetTicks()
		frame_time := end - start
		if frame_time > 0 {
			fps = 1000 / int(frame_time)
		}
	}
}
