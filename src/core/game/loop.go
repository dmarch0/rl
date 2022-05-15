package game

import (
	"fmt"
	"rl/src/core/ecs"

	"github.com/veandco/go-sdl2/sdl"
)

func repeat(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(s)
	}
}

func GameLoop(window *sdl.Window, renderer *sdl.Renderer) {
	running := true

	world := ecs.World{
		Entities: make([]ecs.Entity, 0),
	}
	builder := ecs.CreateScheduleBuilder(&world)
	builder.AddSystem(func(world *ecs.World) {
		repeat("System A")
	})
	builder.AddSystem(func(world *ecs.World) {
		repeat("System B")
	})
	builder.Flush()
	builder.AddSystem(func(world *ecs.World) {
		repeat("System C")
	})
	builder.AddSystem(func(world *ecs.World) {
		repeat("System D")
	})
	builder.Flush()
	scheduler := builder.Build()

	for running {
		PrepearScene(renderer)
		PollEvents(&running)
		scheduler.Execute()
		Render(renderer)
		PresentScene(renderer)
		sdl.Delay(16)
	}
}
