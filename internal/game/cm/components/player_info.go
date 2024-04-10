package components

import (
	"online_game/internal/game/cm"
	"online_game/internal/models"
)

// TestComponent stores value about Player (models.PublicPlayer).
//
// Init: none
// Update: none
// Render: none
type PlayerInfoComponent struct {
	Info models.PublicPlayer

	obj *cm.GameObject // pointer to object that has this component
}

func (t *PlayerInfoComponent) Init(obj *cm.GameObject) {
	t.obj = obj // set pointer to object that has this component
}

func (t *PlayerInfoComponent) Update() {
	/*
	   Called in every frame by managers Update func.
	   Calculate, read input, etc.
	   Mustn't render anything
	*/
}

func (t *PlayerInfoComponent) Render() {
	/*
	   Called in every frame by managers Render func.
	   If component has some graphic it needs to be rendered here.
	*/
}
