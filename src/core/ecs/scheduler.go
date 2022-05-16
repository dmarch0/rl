package ecs

type SystemContainer struct {
	system func(world *World, resources *Resources)
	c      chan int
}

type Stage []SystemContainer

type Scheduler struct {
	World     *World
	Resources *Resources
	stages    []Stage
}

func (s *Scheduler) Execute() {
	for _, stage := range s.stages {
		for _, sc := range stage {
			go sc.system(s.World, s.Resources)
		}
		for _, sc := range stage {
			<-sc.c
		}
	}
}
