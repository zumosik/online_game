package components

import (
	"online_game/internal/game/cm"
	"online_game/internal/models"
	"online_game/internal/tcpclient"
)

// OtherPlayerInfoComponent stores value about Player (models.PublicPlayer). Doesn't send data to server.
//
// Init: none
// Update: none
// Render: none
//
// # Cant be used without RigidbodyComponent.
type OtherPlayerInfoComponent struct {
	Info   models.PublicPlayer
	Client *tcpclient.TCPClient

	obj *cm.GameObject // pointer to object that has this component
}

func (t *OtherPlayerInfoComponent) Init(obj *cm.GameObject) {
	t.obj = obj // set pointer to object that has this component
}

func (t *OtherPlayerInfoComponent) Update() {
}

func (t *OtherPlayerInfoComponent) Render() {
	/*
	   Called in every frame by managers Render func.
	   If component has some graphic it needs to be rendered here.
	*/
}
