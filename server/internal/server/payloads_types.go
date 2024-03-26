package server

import "server/internal/models"

const (
	TypeOfPacketEmpty            = iota // for errors
	TypeOfPacketConnectReq              // to connect
	TypeOfPacketConnectResp             // connect resp
	TypeOfPacketNewPlayerConnect        // notify other players about new conn
	TypeOfPacketDisconnectReq           // safe disconnect
	TypeOfPacketDisconnectResp          // if disconnected safe
	TypeOfPacketPlayerPosReq            // send your pos
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

type NewPlayerConnect struct {
	models.Player
}

type DisconnectReq struct {
	ID uint16
}

type DisconnectResp struct {
	OK bool
}
