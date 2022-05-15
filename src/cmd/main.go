package main

import (
	"rl/src/core/boot"
	"rl/src/core/game"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	window, renderer := boot.InitSdl()

	game.GameLoop(window, renderer)

	sdl.Quit()
	window.Destroy()
}
