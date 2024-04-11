package components

import (
	"online_game/internal/game/cm"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// SpriteComponent stores value about pointer to texture and color.
//
// Init: gets transform and sets scrRect to transform size.
// Update: sets render pos to transform pos, and size to transforms size * scale.
// Render: renders texture.
//
// # Cant be used without TransformComponent.
type SpriteComponent struct {
	Tex   rl.Texture2D
	Color rl.Color

	srcRect rl.Rectangle
	dstRect rl.Rectangle

	transform *TransformComponent

	obj *cm.GameObject // link to object that has this component
}

func (s *SpriteComponent) Init(obj *cm.GameObject) {
	s.obj = obj

	s.transform = s.obj.GetComponent(&TransformComponent{}).(*TransformComponent)

	s.srcRect.X = 0
	s.srcRect.Y = 0

	s.srcRect.Width = s.transform.Size.X
	s.srcRect.Height = s.transform.Size.Y
}

func (s *SpriteComponent) Update() {
	s.dstRect.X = s.transform.Pos.X
	s.dstRect.Y = s.transform.Pos.Y

	s.dstRect.Width = s.transform.Size.X * s.transform.Scale.X
	s.dstRect.Height = s.transform.Size.Y * s.transform.Scale.Y
}

func (s *SpriteComponent) Render() {
	rl.DrawTexturePro(s.Tex, s.srcRect, s.dstRect,
		rl.NewVector2(s.dstRect.Width, s.dstRect.Height),
		s.transform.Rotation, s.Color)

}
