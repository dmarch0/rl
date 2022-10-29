package systems

import (
	"rl/src/core/ecs"
	"rl/src/core/utils"

	"github.com/veandco/go-sdl2/sdl"
)

func ShootingSystem(world *ecs.World, resources *ecs.Resources) {
	X, Y, _ := sdl.GetMouseState()
	clickVec := utils.Vector{
		X: float64(X),
		Y: float64(Y),
	}
	if world.Player != nil && world.Player.Transform != nil && resources.ButtonState.LeftMouseButton {
		player := world.Player
		dif := clickVec.Sub(player.Transform.Position)
		normalDif := dif.Normalise()
		ecs.SpawnBullet(world, normalDif, player.Transform.Position.Add(normalDif.Mul(10)))
	}
}
