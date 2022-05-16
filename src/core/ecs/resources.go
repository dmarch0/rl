package ecs

import "github.com/veandco/go-sdl2/sdl"

type Resources struct {
	Running  bool
	Renderer *sdl.Renderer
}
