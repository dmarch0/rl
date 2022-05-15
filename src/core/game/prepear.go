package game

import "github.com/veandco/go-sdl2/sdl"

func PrepearScene(renderer *sdl.Renderer) {
	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.Clear()
}
