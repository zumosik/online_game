package game

import (
	"game/internal/cm"
	"game/internal/cm/components"
	"game/internal/textures"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	screenWidth  int32
	screenHeight int32
	fps          int32
	title        string
	running      bool

	tex     *textures.Textures
	manager *cm.Manager
}

func New(w, h, fps int32, title string) *Game {
	return &Game{
		screenHeight: h,
		screenWidth:  w,
		fps:          fps,
		title:        title,
		running:      true,
	}
}

func (g *Game) Init() {
	rl.InitWindow(g.screenWidth, g.screenHeight, g.title)
	rl.SetTargetFPS(g.fps)

	g.tex = textures.Load()

	rl.SetExitKey(rl.KeyNull)

	g.manager = cm.NewManager()
}

func (g *Game) Start() {
	player := g.manager.CreateGameObject()
	player.AddComponent(&components.TransformComponent{
		Pos:   rl.NewVector2(500, 200),
		Size:  rl.NewVector2(48, 48),
		Scale: rl.NewVector2(5, 5),
	})
	player.AddComponent(&components.SpriteComponent{
		Tex:   g.tex.Player,
		Color: rl.White,
	})

	for g.running {
		g.handleInput()
		g.update()
		g.render()
	}
}

func (g *Game) Quit() {
	rl.CloseWindow()
}

func (g *Game) handleInput() {
	// TODO
}

func (g *Game) update() {
	g.running = !rl.WindowShouldClose()

	g.manager.Update()
}

func (g *Game) render() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.Color{R: 147, G: 211, B: 139, A: 255})
	g.drawScene()

	rl.EndDrawing()
}

func (g *Game) drawScene() {
	g.manager.Render()
}
