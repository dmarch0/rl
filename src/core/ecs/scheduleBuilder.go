package ecs

type ScheduleBuilder struct {
	stages       []Stage
	currentStage []SystemContainer
	world        *World
}

func CreateScheduleBuilder() *ScheduleBuilder {
	b := new(ScheduleBuilder)
	b.stages = make([]Stage, 0)
	b.currentStage = make([]SystemContainer, 0)
	return b
}

func (b *ScheduleBuilder) AddSystem(system System) *ScheduleBuilder {
	c := make(chan int)
	systemContainer := SystemContainer{
		system: func(world *World, resources *Resources) {
			system(world, resources)
			c <- 1
		},
		c: c,
	}
	b.currentStage = append(b.currentStage, systemContainer)
	return b
}

func (b *ScheduleBuilder) Flush() *ScheduleBuilder {
	b.stages = append(b.stages, b.currentStage)
	b.currentStage = make([]SystemContainer, 0)
	return b
}

func (b *ScheduleBuilder) Build() *Scheduler {
	scheduler := new(Scheduler)
	scheduler.stages = b.stages
	return scheduler
}
