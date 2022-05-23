package ecs

type DeathIntention struct {
	Entity *Entity
}

func (m DeathIntention) Lock()   {}
func (m DeathIntention) Unlock() {}

func BindDeathIntention(e *Entity, m *DeathIntention) {
	e.DeathIntention = m
	m.Entity = e
}
