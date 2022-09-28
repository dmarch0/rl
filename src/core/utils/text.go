package utils

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

var font *ttf.Font

func GetOrLoadFont(path string) *ttf.Font {
	if font == nil {
		f, err := ttf.OpenFont(path, 24)
		if err != nil {
			panic(err)
		}
		font = f
		return font
	}

	return font
}

func PrintText(text string, rect *sdl.Rect, renderer *sdl.Renderer) {
	font := GetOrLoadFont("resources/open-sans.ttf")
	surface, err := font.RenderUTF8Solid(text, sdl.Color{255, 255, 255, 255})
	if err != nil {
		fmt.Println("Error rendering text surface", err)
	}
	texture, err := renderer.CreateTextureFromSurface(surface)
	if err != nil {
		fmt.Println("Error creating texture from surface")
	}

	renderer.Copy(texture, nil, rect)
}
