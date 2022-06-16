package utils

type Vector struct {
	X int
	Y int
}

func (v Vector) Add(other Vector) Vector {
	return Vector{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

func (v *Vector) IsZero() bool {
	return v.X == 0 && v.Y == 0
}

func (v Vector) Mul(factor int) Vector {
	return Vector{
		X: v.X * factor,
		Y: v.Y * factor,
	}
}
