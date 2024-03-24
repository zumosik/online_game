package server

import (
	"bytes"
	"log/slog"
	"math/rand"
	"net"
	"server/internal/models"
	"server/internal/utils"
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
	Pin      uint32 // like password
}

type ConnectResp struct {
	OK bool
	ID uint16
}

type PlayerPosReq struct {
	ID     uint16
	Vector models.Vector
}

type PlayerPosResp struct {
	OK bool
}

func (s *Server) handleConnectReq(req ConnectReq, conn net.Conn) ConnectResp {
	_, exists := s.playerMap[conn]
	if exists { // already connected
		return ConnectResp{OK: false}
	}

	for _, pl := range s.playerMap {
		if pl.Username == req.Username {
			return ConnectResp{OK: false} // someone is already playing
		}
	}
	// TODO find user in save

	// getting unique rnd id
	id := uint16(rand.Intn(65535))
	for !s.isUserIDUnique(id) {
		id = uint16(rand.Intn(65535))
	}

	s.playerMap[conn] = models.Player{
		Username: req.Username,
		UserID:   id,
		Pos:      models.Vector{X: 0, Y: 0},
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
	var vec models.Vector
	err = vec.Deserialize(buf)
	return err
}

func (v *ConnectReq) Serialize() []byte {
	var buf *bytes.Buffer

	utils.WriteString(buf, v.Username)
	utils.WriteUint32(buf, v.Pin)

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

func (v *PlayerPosResp) Serialize() []byte {
	var buf bytes.Buffer

	utils.WriteBool(&buf, v.OK)

	return buf.Bytes()
}

func (v *PlayerPosResp) Deserialize(b []byte) error {
	buf := bytes.NewBuffer(b)

	ok, err := utils.ReadBool(buf)

	v.OK = ok

	return err
}

func (v *ConnectResp) Serialize() []byte {
	var buf bytes.Buffer

	utils.WriteBool(&buf, v.OK)
	utils.WriteUint16(&buf, v.ID)

	return buf.Bytes()
}

func (v *ConnectResp) Deserialize(b []byte) error {
	buf := bytes.NewBuffer(b)

	ok, err := utils.ReadBool(buf)
	if err != nil {
		return err
	}
	id, err := utils.ReadUint16(buf)

	v.OK = ok
	v.ID = id

	return err
}

func (s *Server) isUserIDUnique(userID uint16) bool {
	// Iterate over the playerMap and check if the UserID already exists
	for _, player := range s.playerMap {
		if player.UserID == userID {
			return false // UserID is not unique
		}
	}
	return true // UserID is unique
}
