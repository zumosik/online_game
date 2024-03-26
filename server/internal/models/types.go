package models

import (
	"bytes"
	"server/internal/utils"
)

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

func (v *Vector) Serialize(buf *bytes.Buffer) {
	utils.WriteFloat64(buf, v.X)
	utils.WriteFloat64(buf, v.Y)
}

func (v *Vector) Deserialize(buf *bytes.Buffer) error {

	x, err := utils.ReadFloat64(buf)
	if err != nil {
		return err
	}
	y, err := utils.ReadFloat64(buf)

	v.X = x
	v.Y = y

	return err
}

func (p *Player) Serialize(buf *bytes.Buffer) {
	utils.WriteString(buf, p.Username)
	utils.WriteUint16(buf, p.UserID)
	p.Pos.Serialize(buf)
}

func (p *Player) Deserialize(buf *bytes.Buffer) error {
	s, err := utils.ReadString(buf)
	if err != nil {
		return err
	}
	n, err := utils.ReadUint16(buf)
	if err != nil {
		return err
	}
	err = p.Pos.Deserialize(buf)
	if err != nil {
		return err
	}

	p.Username = s
	p.UserID = n

	return nil
}
