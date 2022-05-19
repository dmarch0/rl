package ecs

type Collider struct {
	Entity *Entity
}

func (c Collider) Lock()   {}
func (c Collider) Unlock() {}

func BindCollider(e *Entity, c *Collider) {
	e.Collider = c
	c.Entity = e
}
