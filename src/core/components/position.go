package components

import "rl/src/core/utils"

type Position struct {
	utils.Vector
}

func (p Position) Lock()   {}
func (p Position) Unlock() {}
