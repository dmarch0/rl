package utils

type Vector struct {
	X int
	Y int
}

func (v *Vector) Add(other Vector) {
	v.X += other.X
	v.Y += other.Y
}

func (v *Vector) IsZero() bool {
	return v.X == 0 && v.Y == 0
}
