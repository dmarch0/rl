package systems

import (
	"rl/src/core/ecs"
)

func CollisionSystem(world *ecs.World, resources *ecs.Resources) {
	colliders := world.GetEntities(ecs.HasCollider)
	for _, collider := range colliders {
		if collider.MovementIntention != nil && collider.Transform != nil {
			newTransform := ecs.Transform{
				Position: collider.MovementIntention.Where,
				Scale:    collider.Transform.Scale,
			}
			for _, otherCollider := range colliders {
				if collider != otherCollider && otherCollider.Transform != nil && newTransform.DoIntersect(otherCollider.Transform) {
					collider.MovementIntention = nil
				}
			}
		}
	}
}
