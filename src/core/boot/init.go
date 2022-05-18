package boot

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func InitSdl() (*sdl.Window, *sdl.Renderer) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	fmt.Println("Init")

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	fmt.Println("Create window")

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	fmt.Println("Create renderer")

	if err := img.Init(img.INIT_JPG | img.INIT_PNG); err != nil {
		panic(err)
	}
	fmt.Println("Image init")

	return window, renderer
}
