package ecs

type Health struct {
	Entity *Entity
	Value  float32
}

func (p Health) Lock()   {}
func (p Health) Unlock() {}

func BindHealth(e *Entity, p *Health) {
	e.Health = p
	p.Entity = e
}
