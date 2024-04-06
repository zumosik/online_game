package components

import (
	"game/internal/cm"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type TransformComponent struct { // this component only store values
	Pos  rl.Vector2
	Size rl.Vector2

	Rotation float32
	Scale    rl.Vector2

	obj *cm.GameObject // link to object that has this component
}

func (t *TransformComponent) SetGameObject(obj *cm.GameObject) {
	t.obj = obj
}

func (t *TransformComponent) Init() {
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
