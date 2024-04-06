package components

import (
	"game/internal/cm"
	rl "github.com/gen2brain/raylib-go/raylib"
	"log"
)

type SpriteComponent struct {
	Tex   rl.Texture2D
	Color rl.Color

	srcRect rl.Rectangle
	dstRect rl.Rectangle

	transform *TransformComponent

	obj *cm.GameObject // link to object that has this component
}

func (s *SpriteComponent) SetGameObject(obj *cm.GameObject) {
	s.obj = obj
}

func (s *SpriteComponent) Init() {
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
	log.Printf("Render: %v, %v", s.srcRect, s.dstRect)

	rl.DrawTexturePro(s.Tex, s.srcRect, s.dstRect,
		rl.NewVector2(s.dstRect.Width, s.dstRect.Height),
		s.transform.Rotation, s.Color)

}
