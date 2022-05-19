package ecs

type Player struct {
	Entity *Entity
}

func (p Player) Lock()   {}
func (p Player) Unlock() {}

func BindPlayer(e *Entity, p *Player) {
	e.Player = p
	p.Entity = e
}
