package ecs

type Damager struct {
	Entity  *Entity
	IsSolid bool
}

func (c Damager) Lock()   {}
func (c Damager) Unlock() {}

func BindDamager(e *Entity, c *Damager) {
	e.Damager = c
	c.Entity = e
}
