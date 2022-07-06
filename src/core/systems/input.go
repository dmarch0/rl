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
		case *sdl.MouseButtonEvent:
			HandleMouseButtonEvent(event.(*sdl.MouseButtonEvent), world)
			break
		}
	}
}

func HandleMouseButtonEvent(event *sdl.MouseButtonEvent, world *ecs.World) {
	switch event.Type {
	case sdl.MOUSEBUTTONDOWN:
		HandleMouseButtonDown(event, world)
		break
	}
}

func HandleMouseButtonDown(event *sdl.MouseButtonEvent, world *ecs.World) {
	switch event.Button {
	case sdl.BUTTON_LEFT:
		HandleLeftButtonDown(event, world)
	}
}

func HandleLeftButtonDown(event *sdl.MouseButtonEvent, world *ecs.World) {
	fmt.Println("Click at X=", event.X, " Y= ", event.Y)
	clickVec := utils.Vector{
		X: float64(event.X),
		Y: float64(event.Y),
	}
	if world.Player != nil && world.Player.Transform != nil {
		player := world.Player
		dif := clickVec.Sub(player.Transform.Position)
		normalDif := dif.Normalise()
		ecs.SpawnBullet(world, normalDif, player.Transform.Position)
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
			player.Velocity.Direction = velocity.Direction.Add(utils.Vector{0, 1})
			break
		case sdl.SCANCODE_W:
			player.Velocity.Direction = velocity.Direction.Add(utils.Vector{0, -1})
			break
		case sdl.SCANCODE_A:
			player.Velocity.Direction = velocity.Direction.Add(utils.Vector{-1, 0})
			break
		case sdl.SCANCODE_D:
			player.Velocity.Direction = velocity.Direction.Add(utils.Vector{1, 0})
			break
		}
	}
	fmt.Println("Velo: ", player.Velocity)
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
			player.Velocity.Direction = velocity.Direction.Add(utils.Vector{0, -1})
			break
		case sdl.SCANCODE_W:
			player.Velocity.Direction = velocity.Direction.Add(utils.Vector{0, 1})
			break
		case sdl.SCANCODE_A:
			player.Velocity.Direction = velocity.Direction.Add(utils.Vector{1, 0})
			break
		case sdl.SCANCODE_D:
			player.Velocity.Direction = velocity.Direction.Add(utils.Vector{-1, 0})
			break
		}
	}
}
