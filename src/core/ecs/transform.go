package ecs

import (
	"rl/src/core/utils"
)

type Transform struct {
	Position utils.Vector
	Scale    utils.Vector
	Entity   *Entity
}

func (p Transform) Lock()   {}
func (p Transform) Unlock() {}

func BindTransform(e *Entity, p *Transform) {
	p.Entity = e
	e.Transform = p
}

func (t Transform) DoIntersect(other *Transform) bool {
	leftA := t.Position.X
	topA := t.Position.Y
	rightA := leftA + t.Scale.X
	bottomA := topA + t.Scale.Y

	leftB := other.Position.X
	topB := other.Position.Y
	rightB := leftB + other.Scale.X
	bottomB := topB + other.Scale.Y

	return leftA <= rightB && rightA >= leftB && topA <= bottomB && bottomA >= topB
}

func (t Transform) GetBoundingPoints() []utils.Vector {
	return []utils.Vector{
		utils.Vector{
			X: t.Position.X,
			Y: t.Position.Y,
		},
		utils.Vector{
			X: t.Position.X + t.Scale.X,
			Y: t.Position.Y,
		},
		utils.Vector{
			X: t.Position.X,
			Y: t.Position.Y + t.Scale.Y,
		},
		utils.Vector{
			X: t.Position.X + t.Scale.X,
			Y: t.Position.Y + t.Scale.Y,
		},
	}
}
