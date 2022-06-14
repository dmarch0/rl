package ecs

import (
	"rl/src/core/utils"
)

type CollidersHashTable struct {
	Table map[int]*[]*Entity
}

func CreateCollidersHashTable() CollidersHashTable {
	table := make(map[int]*[]*Entity)

	return CollidersHashTable{
		Table: table,
	}
}

func (c *CollidersHashTable) Insert(e *Entity) {
	if e.Transform != nil {
		points := e.Transform.GetBoundingPoints()
		for _, point := range points {
			c.InsertForPoint(e, &point)
		}
	}
}

func (c *CollidersHashTable) InsertForPoint(e *Entity, v *utils.Vector) {
	hash := HashPosition(v)
	if entities, ok := c.Table[hash]; ok {
		*entities = append(*entities, e)
	} else {
		c.Table[hash] = &[]*Entity{e}
	}
}

func (c *CollidersHashTable) UpdateObject(e *Entity) {
	if e.Transform != nil {
		c.Remove(e)
		c.Insert(e)
	}
}

func (c *CollidersHashTable) Remove(e *Entity) {
	if e.Transform != nil {
		points := e.Transform.GetBoundingPoints()
		for _, point := range points {
			c.RemoveForPoint(e, &point)
		}
	}
}

func (c *CollidersHashTable) RemoveForPoint(e *Entity, v *utils.Vector) {
	hash := HashPosition(v)
	if entities, ok := c.Table[hash]; ok {
		index := -1
		for i, entity := range *entities {
			if entity == e {
				index = i
			}
		}
		if index != -1 {
			(*entities)[index] = (*entities)[len(*entities)-1]
			(*entities) = (*entities)[:len(*entities)-1]
		}
	}
}

func (c *CollidersHashTable) QueryPosition(e *Entity) []*Entity {
	if e.Transform != nil {
		hash := HashPosition(&e.Transform.Position)
		if entities, ok := c.Table[hash]; ok {
			return *entities
		}
	}
	return []*Entity{}
}

func (c *CollidersHashTable) QueryVolumedPosition(e *Entity) []*Entity {
	result := []*Entity{}
	if e.Transform != nil {
		points := e.Transform.GetBoundingPoints()
		for _, point := range points {
			hash := HashPosition(&point)
			if entities, ok := c.Table[hash]; ok {
				result = append(result, *entities...)
			}
		}
	}
	return result
}

//50 is cell size
//800 is width
func HashPosition(position *utils.Vector) int {
	opA := (position.X / 50) * 73856093
	opB := (position.Y / 50) * 19349663
	return opA ^ opB
}
