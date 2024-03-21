package server

import "bytes"

type Player struct {
	Username string
	UserID   uint16
	Pos      Vector
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

func (v Vector) Serialize(buf *bytes.Buffer) {
	writeFloat64(buf, v.X)
	writeFloat64(buf, v.Y)
}

func (v Vector) Deserialize(buf *bytes.Buffer) error {

	x, err := readFloat64(buf)
	if err != nil {
		return err
	}
	y, err := readFloat64(buf)

	v.X = x
	v.Y = y

	return err
}
