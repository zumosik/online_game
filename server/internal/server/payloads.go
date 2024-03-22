package server

import (
	"bytes"
	"log/slog"
	"math/rand"
	"net"
)

const (
	TypeOfPacketEmpty = iota // for errors
	TypeOfPacketConnectReq
	TypeOfPacketConnectResp
	TypeOfPacketPlayerPosReq
	TypeOfPacketPlayerPosResp
)

type Payload interface {
	Serialize() []byte
	Deserialize([]byte) error
}

type ConnectReq struct {
	Username string
}

type ConnectResp struct {
	OK bool
	ID uint16
}

type PlayerPosReq struct {
	ID     uint16
	Vector Vector
}

type PlayerPosResp struct {
	OK bool
}

func (s *Server) handleConnectReq(req ConnectReq, conn net.Conn) ConnectResp {
	_, exists := s.playerMap[conn]
	if exists {
		return ConnectResp{OK: false}
	}

	// getting unique rnd id
	id := uint16(rand.Intn(65535))
	for !s.isUserIDUnique(id) {
		id = uint16(rand.Intn(65535))
	}

	s.playerMap[conn] = Player{
		Username: req.Username,
		UserID:   id,
		Pos:      Vector{X: 0, Y: 0},
	}

	s.l.Debug("New player registered", slog.String("username", req.Username), slog.Int("id", int(id)))

	return ConnectResp{OK: true, ID: id}

}

func (s *Server) handlePlayerPosReq(req PlayerPosReq, conn net.Conn) PlayerPosResp {
	player, exists := s.playerMap[conn]
	if !exists && player.UserID != req.ID {
		return PlayerPosResp{
			OK: false,
		}
	}

	player.Pos = req.Vector

	return PlayerPosResp{
		OK: true,
	}
}

func (v *PlayerPosReq) Serialize() []byte {
	var buf bytes.Buffer

	writeUint16(&buf, v.ID)
	v.Vector.Serialize(&buf)

	return buf.Bytes()
}

func (v *PlayerPosReq) Deserialize(b []byte) error {
	buf := bytes.NewBuffer(b)

	id, err := readUint16(buf)
	if err != nil {
		return err
	}
	v.ID = id
	var vec Vector
	err = vec.Deserialize(buf)
	return err
}

func (v *ConnectReq) Serialize() []byte {
	var buf bytes.Buffer

	writeString(&buf, v.Username)

	return buf.Bytes()
}

func (v *ConnectReq) Deserialize(b []byte) error {
	buf := bytes.NewBuffer(b)

	s, err := readString(buf)
	v.Username = s
	return err
}

func (v *PlayerPosResp) Serialize() []byte {
	var buf bytes.Buffer

	writeBool(&buf, v.OK)

	return buf.Bytes()
}

func (v *PlayerPosResp) Deserialize(b []byte) error {
	buf := bytes.NewBuffer(b)

	ok, err := readBool(buf)

	v.OK = ok

	return err
}

func (v *ConnectResp) Serialize() []byte {
	var buf bytes.Buffer

	writeBool(&buf, v.OK)
	writeUint16(&buf, v.ID)

	return buf.Bytes()
}

func (v *ConnectResp) Deserialize(b []byte) error {
	buf := bytes.NewBuffer(b)

	ok, err := readBool(buf)
	if err != nil {
		return err
	}
	id, err := readUint16(buf)

	v.OK = ok
	v.ID = id

	return err
}
