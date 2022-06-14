package ecs

type Mover struct {
	Entity  *Entity
	IsSolid bool
}

func (m Mover) Lock()   {}
func (m Mover) Unlock() {}

func BindMover(e *Entity, m *Mover) {
	e.Mover = m
	m.Entity = e
}
