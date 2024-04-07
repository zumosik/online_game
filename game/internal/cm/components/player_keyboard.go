package components

import (
	"game/internal/cm"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	WASDInput = iota
	ARROWInput
)

// PlayerKeyboardComponent stores info about type of input (see consts)
//
// Init: checks if type of input is correct (see consts)
// Update: handle input and set rigidbody velocity.
// Render: none.
//
// # Cant be used without RigidbodyComponent.
type PlayerKeyboardComponent struct {
	TypeOfInput int

	keyUp     int32 // W
	keyDown   int32 // S
	keyLeft   int32 // A
	keyRight  int32 // D
	rigidbody *RigidbodyComponent

	obj *cm.GameObject
}

func (p *PlayerKeyboardComponent) Init(obj *cm.GameObject) {
	p.obj = obj
	p.rigidbody = p.obj.GetComponent(&RigidbodyComponent{}).(*RigidbodyComponent)

	switch p.TypeOfInput {
	case WASDInput:
		p.keyUp = rl.KeyW
		p.keyDown = rl.KeyS
		p.keyLeft = rl.KeyA
		p.keyRight = rl.KeyD
	case ARROWInput:
		p.keyUp = rl.KeyUp
		p.keyDown = rl.KeyDown
		p.keyLeft = rl.KeyLeft
		p.keyRight = rl.KeyRight
	default: // invalid type of input
		return
	}
}

func (p *PlayerKeyboardComponent) Update() {
	if rl.IsKeyUp(p.keyUp) && rl.IsKeyUp(p.keyDown) { // set y vel 0 if no buttons pressed
		p.rigidbody.Velocity.Y = 0
	}
	if rl.IsKeyUp(p.keyLeft) && rl.IsKeyUp(p.keyRight) { // set x vel 0 if no buttons pressed
		p.rigidbody.Velocity.X = 0
	}

	if rl.IsKeyDown(p.keyUp) {
		p.rigidbody.Velocity.Y = -1.0
	}

	if rl.IsKeyDown(p.keyDown) {
		p.rigidbody.Velocity.Y = 1.0
	}

	if rl.IsKeyDown(p.keyLeft) {
		p.rigidbody.Velocity.X = -1.0
	}

	if rl.IsKeyDown(p.keyRight) {
		p.rigidbody.Velocity.X = 1.0
	}
}

func (p *PlayerKeyboardComponent) Render() {
}
