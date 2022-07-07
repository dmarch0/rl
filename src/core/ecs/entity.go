package ecs

type Entity struct {
	Player            *Player
	Transform         *Transform
	SimpleRender      *SimpleRender
	Velocity          *Velocity
	Collider          *Collider
	Health            *Health
	MovementIntention *MovementIntention
	DeathIntention    *DeathIntention
	Damager           *Damager
	Mover             *Mover
	DiesOnImpact      *DiesOnImpact
}
