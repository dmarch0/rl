package ecs

type World struct {
	Entities []Entity
}

func (w *World) AddEntity(e Entity) {
	w.Entities = append(w.Entities, e)
}
