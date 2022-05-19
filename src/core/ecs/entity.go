package ecs

type Entity struct {
	Player            *Player
	Transform         *Transform
	SimpleRender      *SimpleRender
	Velocity          *Velocity
	MovementIntention *MovementIntention
	Collider          *Collider
}
