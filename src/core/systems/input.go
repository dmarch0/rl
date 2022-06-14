package systems

import (
	"rl/src/core/ecs"
	"rl/src/core/utils"

	"github.com/veandco/go-sdl2/sdl"
)

func InputSystem(world *ecs.World, resources *ecs.Resources) {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			println("Quit")
			resources.Running = false
			break
		case *sdl.KeyboardEvent:

			HandleKeyboardInput(event.(*sdl.KeyboardEvent), world)
			break
		}
	}
}

func HandleKeyboardInput(event *sdl.KeyboardEvent, world *ecs.World) {
	switch event.Type {
	case sdl.KEYDOWN:
		HandleKeyDown(event, world)
		break
	case sdl.KEYUP:
		HandleKeyUp(event, world)
		break
	}
}

func HandleKeyDown(event *sdl.KeyboardEvent, world *ecs.World) {
	if event.Repeat == 0 {
		switch event.Keysym.Scancode {
		case sdl.SCANCODE_W, sdl.SCANCODE_S, sdl.SCANCODE_A, sdl.SCANCODE_D:
			HandleMovement(event.Keysym.Scancode, world)
			break
		}

	}
}

func HandleMovement(scancode sdl.Scancode, world *ecs.World) {
	player := world.Player
	if player == nil {
		return
	}
	velocity := player.Velocity
	if velocity != nil {
		switch scancode {
		case sdl.SCANCODE_S:
			velocity.Value.Add(utils.Vector{0, 1})
			break
		case sdl.SCANCODE_W:
			velocity.Value.Add(utils.Vector{0, -1})
			break
		case sdl.SCANCODE_A:
			velocity.Value.Add(utils.Vector{-1, 0})
			break
		case sdl.SCANCODE_D:
			velocity.Value.Add(utils.Vector{1, 0})
			break
		}
	}
}

func HandleKeyUp(event *sdl.KeyboardEvent, world *ecs.World) {
	if event.Repeat == 0 {
		switch event.Keysym.Scancode {
		case sdl.SCANCODE_W, sdl.SCANCODE_S, sdl.SCANCODE_A, sdl.SCANCODE_D:
			HandleStopMovement(event.Keysym.Scancode, world)
			break
		}

	}
}

func HandleStopMovement(scancode sdl.Scancode, world *ecs.World) {
	player := world.Player
	if player == nil {
		return
	}
	velocity := player.Velocity
	if velocity != nil {
		switch scancode {
		case sdl.SCANCODE_S:
			velocity.Value.Add(utils.Vector{0, -1})
			break
		case sdl.SCANCODE_W:
			velocity.Value.Add(utils.Vector{0, 1})
			break
		case sdl.SCANCODE_A:
			velocity.Value.Add(utils.Vector{1, 0})
			break
		case sdl.SCANCODE_D:
			velocity.Value.Add(utils.Vector{-1, 0})
			break
		}
	}
}
