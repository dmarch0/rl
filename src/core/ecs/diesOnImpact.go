package ecs

type DiesOnImpact struct {
	Entity *Entity
}

func (d DiesOnImpact) Lock()   {}
func (d DiesOnImpact) Unlock() {}

func BindDiesOnImpact(e *Entity, d *DiesOnImpact) {
	e.DiesOnImpact = d
	d.Entity = e
}
