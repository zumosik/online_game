package packets

import "online_game/internal/models"

const (
	TypeOfPacketEmpty            = iota // for errors
	TypeOfPacketConnectReq              // to connect
	TypeOfPacketConnectResp             // connect resp
	TypeOfPacketNewPlayerConnect        // notify other players about new conn
	TypeOfPacketPlayerPosReq            // send your pos
	TypeOfPacketPlayerPosResp           // send other player pos
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
	AlreadyExists bool                // if player already was in save
	Player        models.PublicPlayer // return the player its info
	Token         string              // token for future requests
	Players       []models.PublicPlayer
}

type PlayerPosReq struct {
	Token string

	ID     uint16
	Vector models.Vector
}

type NewPlayerConnect struct {
	Player models.PublicPlayer
}

type PlayerPosResp struct {
	Player models.PublicPlayer // TODO
}
