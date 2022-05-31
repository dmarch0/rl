package ecs

import "rl/src/core/utils"

type SimpleRender struct {
	Border utils.RGBA
	Fill   utils.RGBA
	Entity *Entity
}

func (c SimpleRender) Lock()   {}
func (c SimpleRender) Unlock() {}

func BindSimpleRender(e *Entity, r *SimpleRender) {
	e.SimpleRender = r
	r.Entity = e
}
