package components

import (
	"online_game/internal/game/cm"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// RigidbodyComponent stores value about velocity and speed.
//
// Init: Gets transform component.
// Update: applies velocity to transform position (velocity * speed).
// Render: none.
//
// # Cant be used without TransformComponent.
type RigidbodyComponent struct {
	Velocity rl.Vector2
	Speed    float32

	transform *TransformComponent
	obj       *cm.GameObject // link to object that has this component
}

func (r *RigidbodyComponent) Init(obj *cm.GameObject) {
	r.obj = obj

	r.transform = r.obj.GetComponent(&TransformComponent{}).(*TransformComponent)
}

func (r *RigidbodyComponent) Update() {
	r.transform.Pos.X += r.Velocity.X * r.Speed
	r.transform.Pos.Y += r.Velocity.Y * r.Speed
}

func (r *RigidbodyComponent) Render() {
	// no render
}
