package utils

import "math"

type Vector struct {
	X float64
	Y float64
}

func (v Vector) Add(other Vector) Vector {
	return Vector{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

func (v Vector) Sub(other Vector) Vector {
	return Vector{
		X: v.X - other.X,
		Y: v.Y - other.Y,
	}
}

func (v *Vector) IsZero() bool {
	return v.X == 0 && v.Y == 0
}

func (v Vector) Mul(factor float64) Vector {
	return Vector{
		X: v.X * factor,
		Y: v.Y * factor,
	}
}

func (v Vector) Magnitude() float64 {
	return math.Pow(math.Pow(v.X, 2)+math.Pow(v.Y, 2), 0.5)
}

func (v Vector) Normalise() Vector {
	mag := v.Magnitude()
	return Vector{
		X: v.X / mag,
		Y: v.Y / mag,
	}
}
