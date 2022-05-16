package ecs

type System func(world *World, resources *Resources)
