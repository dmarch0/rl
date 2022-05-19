package ecs

import "rl/src/core/utils"

type Velocity struct {
	Value  utils.Vector
	Entity *Entity
}

func (v Velocity) Lock()   {}
func (v Velocity) Unlock() {}

func BindVelocity(e *Entity, v *Velocity) {
	e.Velocity = v
	v.Entity = e
}
