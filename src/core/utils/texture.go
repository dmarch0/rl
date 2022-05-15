package utils

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func LoadTexture(renderer *sdl.Renderer, filename string) *sdl.Texture {
	sdl.LogMessage(sdl.LOG_CATEGORY_APPLICATION, sdl.LOG_PRIORITY_INFO, "loading %s", filename)
	texture, err := img.LoadTexture(renderer, filename)
	if err != nil {
		panic(err)
	}
	return texture
}

func Blit(renderer *sdl.Renderer, texture *sdl.Texture, x int, y int) {
	dest := sdl.Rect{X: int32(x), Y: int32(y), W: 50, H: 50}
	renderer.Copy(texture, nil, &dest)
}
