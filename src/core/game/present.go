package game

import "github.com/veandco/go-sdl2/sdl"

func PresentScene(renderer *sdl.Renderer) {
	renderer.Present()
}
