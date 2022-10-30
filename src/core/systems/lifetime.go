package systems

import (
	"rl/src/core/ecs"
	"rl/src/core/utils"
)

func LifetimeSystem(world *ecs.World, resources *ecs.Resources) {
	now := utils.GetTimestamp()
	for _, e := range world.Entities {
		if e.Lifetime != nil && e.Lifetime.Born+e.Lifetime.Lifetime < now {
			deathIntention := &ecs.DeathIntention{}
			ecs.BindDeathIntention(e, deathIntention)
		}
	}
}
