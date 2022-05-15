package game

import "github.com/veandco/go-sdl2/sdl"

func PollEvents(running *bool) {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			println("Quit")
			*running = false
			break
		}
	}
}
