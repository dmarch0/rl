package components

type Position struct {
	X int
	Y int
}

func (p Position) Lock()   {}
func (p Position) Unlock() {}
