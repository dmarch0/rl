package ecs

import "rl/src/core/utils"

type MovementIntention struct {
	Entity *Entity
	Where  utils.Vector
}

func (m MovementIntention) Lock()   {}
func (m MovementIntention) Unlock() {}

func BindMovementIntention(e *Entity, m *MovementIntention) {
	e.MovementIntention = m
	m.Entity = e
}
