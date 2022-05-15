package ecs

type Component interface {
	Lock()
	Unlock()
}
