package main

import (
	"rl/src/core/boot"
	"rl/src/core/loop"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	window, renderer := boot.InitSdl()

	loop.GameLoop(window, renderer)

	sdl.Quit()
	window.Destroy()
}
