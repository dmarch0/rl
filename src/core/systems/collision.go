package systems

import (
	"rl/src/core/ecs"
	"rl/src/core/utils"
)

func CollisionSystem(world *ecs.World, resources *ecs.Resources) {
	colliders := world.GetEntities(ecs.HasCollider)
	for _, collider := range colliders {
		if collider.Transform != nil {
			otherColliders := world.CollidersHashTable.QueryVolumedPosition(collider)
			for _, otherCollider := range otherColliders {
				if collider == otherCollider {
					continue
				}
				doIntersect := false
				if collider.MovementIntention != nil {
					newTransform := ecs.Transform{
						Position: collider.MovementIntention.Where,
						Scale:    collider.Transform.Scale,
					}
					doIntersect = newTransform.DoIntersect(otherCollider.Transform)
				} else {
					doIntersect = collider.Transform.DoIntersect(otherCollider.Transform)
				}
				if otherCollider.Transform != nil && doIntersect {
					if collider.Collider.IsSolid && otherCollider.Collider.IsSolid {
						HandleSolidCollision(collider)
					}
					if otherCollider.Health != nil && collider.Damager != nil {
						HandleDamageCollision(collider, otherCollider)
					}
					if collider.DiesOnImpact != nil && otherCollider.Collider.IsSolid {
						HandleDiesOnImpactCollision(collider)
					}
				}
			}
		}
	}
}

func HandleDiesOnImpactCollision(collider *ecs.Entity) {
	ecs.BindDeathIntention(collider, &ecs.DeathIntention{})
}

func HandleSolidCollision(collider *ecs.Entity) {
	collider.MovementIntention = nil
}

func HandleDamageCollision(collider *ecs.Entity, other *ecs.Entity) {
	other.Health.Value -= 1.0
	if other.Health.Value <= 0 {
		deathIntention := &ecs.DeathIntention{}
		ecs.BindDeathIntention(other, deathIntention)
	}
	if other.SimpleRender != nil {
		other.SimpleRender.Fill = utils.RGBA{
			R: 0,
			G: 255,
			B: 0,
			A: 255,
		}
	}
}
