package components

import (
	"online_game/internal/game/cm"
	"online_game/internal/models"
	"online_game/internal/packets"
	"online_game/internal/tcpclient"
)

// PlayerInfoComponent stores value about Player (models.PublicPlayer) and Token (to work with server). It is used to send data to server.
//
// Init: gets components.
// Update: Updates info, sends position to server.
// Render: none
//
// # Cant be used without TransformComponent.
type PlayerInfoComponent struct {
	Info   models.PublicPlayer
	Client *tcpclient.TCPClient
	Token  string // token to authenticate user

	transform *TransformComponent
	obj       *cm.GameObject // pointer to object that has this component
}

func (t *PlayerInfoComponent) Init(obj *cm.GameObject) {
	t.transform = obj.GetComponent(&TransformComponent{}).(*TransformComponent)

	t.obj = obj // set pointer to object that has this component
}

func (t *PlayerInfoComponent) Update() {
	// update info
	t.Info.Pos = models.Vector{
		X: t.transform.Pos.X,
		Y: t.transform.Pos.Y,
	}

	// send player position to server
	err := t.Client.Send(
		packets.Packet{
			TypeOfPacket: packets.TypeOfPacketPlayerPosReq,
			Payload: packets.PlayerPosReq{
				ID:     t.Info.UserID,
				Vector: t.Info.Pos,
				Token:  t.Token,
			},
		},
	)
	if err != nil {
		return // TODO handle error
	}

}

func (t *PlayerInfoComponent) Render() {
	/*
	   Called in every frame by managers Render func.
	   If component has some graphic it needs to be rendered here.
	*/
}
