package ecs

import "rl/src/core/utils"

func SpawnSimpleBlock(world *World, scale utils.Vector, position utils.Vector) {
	block := Entity{}
	BindTransform(&block, &Transform{
		Position: position,
		Scale:    scale,
	})
	BindSimpleRender(&block, &SimpleRender{
		Color: utils.RGBA{
			R: 255,
			G: 0,
			B: 0,
			A: 255,
		},
	})
	BindCollider(&block, &Collider{IsSolid: true})
	BindHealth(&block, &Health{Value: 20.0})
	world.AddEntity(&block)
}
