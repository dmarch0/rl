package ecs

import "rl/src/core/utils"

func SpawnSimpleBlock(world *World, scale utils.Vector, position utils.Vector) {
	block := Entity{}
	BindTransform(&block, &Transform{
		Position: position,
		Scale:    scale,
	})
	BindSimpleRender(&block, &SimpleRender{
		Border: utils.RGBA{
			R: 255,
			G: 0,
			B: 0,
			A: 255,
		},
		Fill: utils.RGBA{
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

func SpawnBullet(world *World, velocity utils.Vector, position utils.Vector) {
	bullet := Entity{}
	BindTransform(&bullet, &Transform{
		Position: position,
		Scale: utils.Vector{
			X: 5,
			Y: 5,
		},
	})
	BindSimpleRender(&bullet, &SimpleRender{
		Border: utils.RGBA{
			R: 0,
			G: 255,
			B: 0,
			A: 255,
		},
		Fill: utils.RGBA{
			R: 255,
			G: 255,
			B: 255,
			A: 255,
		},
	})
	BindCollider(&bullet, &Collider{IsSolid: false})
	BindDamager(&bullet, &Damager{})
	BindVelocity(&bullet, &Velocity{
		Direction: velocity,
		Value:     5.0,
	})
	BindMover(&bullet, &Mover{})
	BindDiesOnImpact(&bullet, &DiesOnImpact{})
	world.AddEntity(&bullet)
}
