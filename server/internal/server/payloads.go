package server

import (
	"bytes"
	"encoding/binary"
	"log/slog"
	"math"
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
	s.l.Debug("handlePlayerPosReq()")
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

func writeFloat32(buf *bytes.Buffer, f float32) {
	writeUint32(buf, math.Float32bits(f))
}

func writeFloat64(buf *bytes.Buffer, f float64) {
	writeUint64(buf, math.Float64bits(f))
}

func writeBool(buf *bytes.Buffer, b bool) {
	if b == true {
		writeByte(buf, 1)
	} else {
		writeByte(buf, 0)
	}
}

// UTF-8
func writeString(buf *bytes.Buffer, s string) {
	l := uint32(len(s))
	writeUint32(buf, l)

	buf.Write([]byte(s))
}

func writeUint16(buf *bytes.Buffer, n uint16) {
	binary.Write(buf, binary.BigEndian, n)
}

func writeUint32(buf *bytes.Buffer, n uint32) {
	binary.Write(buf, binary.BigEndian, n)
}

func writeUint64(buf *bytes.Buffer, n uint64) {
	binary.Write(buf, binary.BigEndian, n)
}

func writeByte(buf *bytes.Buffer, n byte) {
	binary.Write(buf, binary.BigEndian, n)
}

func readString(buf *bytes.Buffer) (string, error) {
	n, err := readUint32(buf)
	if err != nil {
		return "", err
	}

	strBytes := make([]byte, n)
	_, err = buf.Read(strBytes)
	if err != nil {
		return "", err
	}

	str := string(strBytes)

	return str, nil
}

func readFloat32(buf *bytes.Buffer) (float32, error) {
	n, err := readUint32(buf)
	return math.Float32frombits(n), err
}

func readFloat64(buf *bytes.Buffer) (float64, error) {
	n, err := readUint64(buf)
	return math.Float64frombits(n), err
}

func readUint16(buf *bytes.Buffer) (n uint16, err error) {
	err = binary.Read(buf, binary.BigEndian, &n)
	return n, err
}

func readUint32(buf *bytes.Buffer) (n uint32, err error) {
	err = binary.Read(buf, binary.BigEndian, &n)
	return n, err
}

func readUint64(buf *bytes.Buffer) (n uint64, err error) {
	err = binary.Read(buf, binary.BigEndian, &n)
	return n, err
}

func readByte(buf *bytes.Buffer) (n byte, err error) {
	err = binary.Read(buf, binary.BigEndian, &n)
	return n, err
}

func readBool(buf *bytes.Buffer) (bool, error) {
	n, err := buf.ReadByte()
	if n == 1 {
		return true, err
	}
	return false, err

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
