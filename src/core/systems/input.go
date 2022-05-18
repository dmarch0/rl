package systems

import (
	"fmt"
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
	fmt.Println("KeyDown")
	if event.Repeat == 0 {
		switch event.Keysym.Scancode {
		case sdl.SCANCODE_UP, sdl.SCANCODE_DOWN, sdl.SCANCODE_LEFT, sdl.SCANCODE_RIGHT:
			HandleMovement(event.Keysym.Scancode, world)
			break
		}

	}
}

func HandleMovement(scancode sdl.Scancode, world *ecs.World) {
	player, error := world.TryGetEntity(ecs.IsPlayer)
	fmt.Println(player, error)
	if error != nil {
		return
	}
	velocity := player.Velocity
	fmt.Println(velocity)
	if velocity != nil {
		fmt.Println("Got player and velocity")
		switch scancode {
		case sdl.SCANCODE_DOWN:
			velocity.Vector.Add(utils.Vector{0, -1})
			break
		case sdl.SCANCODE_UP:
			velocity.Vector.Add(utils.Vector{0, 1})
			break
		case sdl.SCANCODE_LEFT:
			velocity.Vector.Add(utils.Vector{-1, 0})
			break
		case sdl.SCANCODE_RIGHT:
			velocity.Vector.Add(utils.Vector{1, 0})
			break
		}
		fmt.Println("NEW VELO", velocity)
		fmt.Println(player, velocity)
	}
}

func HandleKeyUp(event *sdl.KeyboardEvent, world *ecs.World) {
	fmt.Println("KeyUp")
	if event.Repeat == 0 {
		switch event.Keysym.Scancode {
		case sdl.SCANCODE_UP, sdl.SCANCODE_DOWN, sdl.SCANCODE_LEFT, sdl.SCANCODE_RIGHT:
			HandleStopMovement(event.Keysym.Scancode, world)
			break
		}

	}
}

func HandleStopMovement(scancode sdl.Scancode, world *ecs.World) {
	player, error := world.TryGetEntity(ecs.IsPlayer)
	if error != nil {
		return
	}
	velocity := player.Velocity
	if velocity != nil {
		switch scancode {
		case sdl.SCANCODE_DOWN:
			velocity.Vector.Add(utils.Vector{0, 1})
			break
		case sdl.SCANCODE_UP:
			velocity.Vector.Add(utils.Vector{0, -1})
			break
		case sdl.SCANCODE_LEFT:
			velocity.Vector.Add(utils.Vector{1, 0})
			break
		case sdl.SCANCODE_RIGHT:
			velocity.Vector.Add(utils.Vector{-1, 0})
			break
		}
		fmt.Println("NEW VELO", velocity)
	}
}
