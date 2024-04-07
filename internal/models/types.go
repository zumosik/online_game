package models

type Player struct {
	Username string
	UserID   uint16
	Pos      Vector
	Pin      uint32
}

type Vector struct {
	X float64
	Y float64
}

func (v Vector) Multiply(multiplier float64) {
	v.X *= multiplier
	v.Y *= multiplier
}

func (v Vector) Add(v2 Vector) {
	v.X += v2.X
	v.Y += v2.Y
}
