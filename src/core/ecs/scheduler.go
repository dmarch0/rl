package ecs

import "fmt"

type SystemContainer struct {
	system func(world *World)
	c      chan int
}

type Stage []SystemContainer

type Scheduler struct {
	world  *World
	stages []Stage
}

func (s *Scheduler) Execute() {
	for _, stage := range s.stages {
		for _, sc := range stage {
			go sc.system(s.world)
		}
		fmt.Println("Scheduled Stage")
		for _, sc := range stage {
			<-sc.c
		}
		fmt.Println("Stage done")
	}
}
