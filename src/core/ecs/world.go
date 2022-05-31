package ecs

import (
	"errors"
)

type World struct {
	Entities           []*Entity
	CollidersHashTable CollidersHashTable
}

func (w *World) AddEntity(e *Entity) {
	w.Entities = append(w.Entities, e)
	if e.Collider != nil && e.Transform != nil {
		w.CollidersHashTable.Insert(e)
	}
}

type EntityFilter func(e *Entity) bool

func (w *World) TryGetEntity(filter EntityFilter) (*Entity, error) {
	for _, e := range w.Entities {
		if filter(e) {
			return e, nil
		}
	}

	return nil, errors.New("Entity not found")
}

func (w *World) GetEntities(filter EntityFilter) []*Entity {
	result := make([]*Entity, 0)
	for _, e := range w.Entities {
		if filter(e) {
			result = append(result, e)
		}
	}
	return result
}

func IsPlayer(e *Entity) bool {
	return e.Player != nil
}

func HasCollider(e *Entity) bool {
	return e.Collider != nil
}
