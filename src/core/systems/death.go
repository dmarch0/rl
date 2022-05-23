package systems

import (
	"rl/src/core/ecs"
)

func DeathSystem(world *ecs.World, resources *ecs.Resources) {
	newEntities := make([]*ecs.Entity, 0)
	for _, e := range world.Entities {
		if e.DeathIntention != nil {
			continue
		}
		newEntities = append(newEntities, e)
	}
	world.Entities = newEntities
}
