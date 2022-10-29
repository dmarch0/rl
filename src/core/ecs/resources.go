package ecs

import "github.com/veandco/go-sdl2/sdl"

type ButtonsState struct {
	LeftMouseButton bool
}

type Resources struct {
	Running     bool
	Renderer    *sdl.Renderer
	ButtonState *ButtonsState
}
