package components

import (
	"game/internal/cm"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// TransformComponent stores value about pos, size, rotation and scale.
//
// Init: sets scale to 1 if it was 0.
// Update: none.
// Render: none.
//
// # TransformComponent is the most important component, a lot of other components rely on it.
type TransformComponent struct {
	Pos  rl.Vector2
	Size rl.Vector2

	Rotation float32
	Scale    rl.Vector2

	obj *cm.GameObject // link to object that has this component
}

func (t *TransformComponent) Init(obj *cm.GameObject) {
	t.obj = obj

	if t.Scale.X == 0 {
		t.Scale.X = 1
	}
	if t.Scale.Y == 0 {
		t.Scale.Y = 1
	}
}

func (t *TransformComponent) Update() {
}

func (t *TransformComponent) Render() {
}
