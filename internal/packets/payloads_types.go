package packets

import "online_game/internal/models"

const (
	TypeOfPacketEmpty            = iota // for errors
	TypeOfPacketConnectReq              // to connect
	TypeOfPacketConnectResp             // connect resp
	TypeOfPacketNewPlayerConnect        // notify other players about new conn
	TypeOfPacketDisconnectReq           // safe disconnect
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
	Username      string
	UserID        uint16
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
