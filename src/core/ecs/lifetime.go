package ecs

type Lifetime struct {
	Entity   *Entity
	Born     int64
	Lifetime int64
}

func (p Lifetime) Lock()   {}
func (p Lifetime) Unlock() {}

func BindLifetime(e *Entity, p *Lifetime) {
	e.Lifetime = p
	p.Entity = e
}
