package models

type Player struct {
	Username string
	UserID   uint16
	Pos      Vector
	Pin      uint32
	Token    string `gob:"-"` // token is generated each time player connects, we dont need to save it
}

// PublicPlayer is a player that is visible to other players (doesnt have Pin field)
type PublicPlayer struct {
	Username string
	UserID   uint16
	Pos      Vector
	Pin      uint32
}

type Vector struct {
	X float32
	Y float32
}

func (v Vector) Multiply(multiplier float32) {
	v.X *= multiplier
	v.Y *= multiplier
}

func (v Vector) Add(v2 Vector) {
	v.X += v2.X
	v.Y += v2.Y
}
