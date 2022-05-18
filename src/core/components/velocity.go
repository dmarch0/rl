package components

import "rl/src/core/utils"

type Velocity struct {
	utils.Vector
}

func (v Velocity) Lock()   {}
func (v Velocity) Unlock() {}
