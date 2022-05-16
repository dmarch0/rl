package components

import "rl/src/core/utils"

type SimpleRender struct {
	W     uint
	H     uint
	Color utils.RGBA
}

func (c SimpleRender) Lock()   {}
func (c SimpleRender) Unlock() {}
