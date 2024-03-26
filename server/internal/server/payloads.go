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
	OK            bool
	AlreadyExists bool // if player already was in save
	Player        models.Player
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

	// Player is not connected:

	var pl models.Player
	pl, playerExists := s.save.Players[req.Username]
	if !playerExists {
		pl.Pos = models.Vector{X: 1, Y: 1}
		pl.Username = req.Username

		// getting unique rnd id
		id := uint16(rand.Intn(65535))
		for !s.isUserIDUnique(id) {
			id = uint16(rand.Intn(65535))
		}

		pl.UserID = id

		// we don't need to save player here because it will be saves on shutdown
	}

	s.playerMap[conn] = pl

	s.l.Debug("New player registered", slog.String("username", req.Username), slog.Int("id", int(pl.UserID)))

	return ConnectResp{OK: true, AlreadyExists: playerExists, Player: pl}

}

func (s *Server) handlePlayerPosReq(req PlayerPosReq, conn net.Conn) PlayerPosResp {

	player, exists := s.playerMap[conn]
	if !exists && player.UserID != req.ID {
		return PlayerPosResp{
			OK: false,
		}
	}

	player.Pos = req.Vector
	s.playerMap[conn] = player

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

func (s *Server) isUserIDUnique(userID uint16) bool {
	// Iterate over the playerMap and check if the UserID already exists
	for _, player := range s.playerMap {
		if player.UserID == userID {
			return false // UserID is not unique
		}
	}
	return true // UserID is unique
}
