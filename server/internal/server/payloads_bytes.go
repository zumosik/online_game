package server

import (
	"bytes"
	"server/internal/utils"
)

func (v *PlayerPosReq) Serialize() []byte {
	var buf bytes.Buffer

	utils.WriteUint16(&buf, v.ID)
	v.Vector.Serialize(&buf)

	return buf.Bytes()
}

func (v *PlayerPosReq) Deserialize(b []byte) error {
	buf := bytes.NewBuffer(b)

	id, err := utils.ReadUint16(buf)
	if err != nil {
		return err
	}
	v.ID = id

	err = v.Vector.Deserialize(buf)
	return err
}

func (v *ConnectReq) Serialize() []byte {
	var buf bytes.Buffer

	utils.WriteString(&buf, v.Username)
	utils.WriteUint32(&buf, v.Pin)

	return buf.Bytes()
}

func (v *ConnectReq) Deserialize(b []byte) error {
	buf := bytes.NewBuffer(b)

	s, err := utils.ReadString(buf)
	if err != nil {
		return err
	}
	n, err := utils.ReadUint32(buf)

	v.Username = s
	v.Pin = n

	return err
}

func (v *ConnectResp) Serialize() []byte {
	var buf bytes.Buffer

	utils.WriteBool(&buf, v.OK)
	utils.WriteBool(&buf, v.AlreadyExists)
	v.Player.Serialize(&buf)

	return buf.Bytes()
}

func (v *ConnectResp) Deserialize(b []byte) error {
	buf := bytes.NewBuffer(b)

	ok, err := utils.ReadBool(buf)
	if err != nil {
		return err
	}
	exists, err := utils.ReadBool(buf)
	if err != nil {
		return err
	}

	v.OK = ok
	v.AlreadyExists = exists
	err = v.Player.Deserialize(buf)
	if err != nil {
		return err
	}

	return nil
}

func (v *NewPlayerConnect) Serialize() []byte {
	var buf bytes.Buffer

	v.Player.Serialize(&buf)

	return buf.Bytes()
}

func (v *NewPlayerConnect) Deserialize(b []byte) error {
	buf := bytes.NewBuffer(b)
	err := v.Player.Deserialize(buf)
	return err
}

func (v *DisconnectReq) Serialize() []byte {
	var buf bytes.Buffer

	utils.WriteUint16(&buf, v.ID)

	return buf.Bytes()
}

func (v *DisconnectReq) Deserialize(b []byte) error {
	buf := bytes.NewBuffer(b)
	n, err := utils.ReadUint16(buf)
	v.ID = n
	return err
}

func (v *DisconnectResp) Serialize() []byte {
	var buf bytes.Buffer

	utils.WriteBool(&buf, v.OK)

	return buf.Bytes()
}

func (v *DisconnectResp) Deserialize(b []byte) error {
	buf := bytes.NewBuffer(b)
	ok, err := utils.ReadBool(buf)
	v.OK = ok
	return err
}
