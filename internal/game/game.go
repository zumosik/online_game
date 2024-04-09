package game

import (
	"errors"
	"log"
	"online_game/internal/game/cm"
	"online_game/internal/game/cm/components"
	"online_game/internal/game/textures"
	"online_game/internal/packets"
	"online_game/internal/tcpclient"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	player *cm.GameObject
)

type Game struct {
	cl *tcpclient.TCPClient

	screenWidth  int32
	screenHeight int32
	fps          int32
	title        string

	running   bool // is game loop running
	connected bool // is connected to server

	tex     *textures.Textures
	manager *cm.Manager
}

func New(cl *tcpclient.TCPClient, w, h, fps int32, title string) *Game {
	return &Game{
		screenHeight: h,
		screenWidth:  w,
		fps:          fps,
		title:        title,
		running:      true,

		cl: cl,
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
	player = g.manager.CreateGameObject()
	player.AddComponent(&components.TransformComponent{
		Pos:   rl.NewVector2(500, 200),
		Size:  rl.NewVector2(48, 48),
		Scale: rl.NewVector2(5, 5),
	})
	player.AddComponent(&components.SpriteComponent{
		Tex:   g.tex.Player,
		Color: rl.White,
	})
	player.AddComponent(&components.RigidbodyComponent{Velocity: rl.NewVector2(0, 0), Speed: 5})
	player.AddComponent(&components.PlayerKeyboardComponent{TypeOfInput: components.WASDInput})

	err := g.cl.Send(packets.Packet{
		TypeOfPacket: packets.TypeOfPacketConnectReq,
		Payload: packets.ConnectReq{
			Username: g.cl.User.Username,
			Pin:      g.cl.User.Pin,
		},
	})
	if err != nil {
		return // cant go further if we cant send server about new player
	}

	go g.TCPLoopRead()

	for g.running { // game loop
		g.update()
		g.render()
	}
}

func (g *Game) TCPLoopRead() {
	for g.running {
		p, err := g.cl.Receive()
		if err != nil {
			if errors.Is(err, tcpclient.ErrNoDataRead) {
				continue
			}
			log.Printf("Error receiving packet: %v", err)
			continue
		}

		switch p.TypeOfPacket {
		case packets.TypeOfPacketConnectResp:
			resp := p.Payload.(packets.ConnectResp)
			g.connected = resp.OK
		default:
			continue
		}
	}
}

func (g *Game) Quit() error {
	rl.CloseWindow()
	err := g.cl.Close()
	if err != nil {
		return err
	}

	return nil
}

func (g *Game) update() {
	g.running = !rl.WindowShouldClose()

	g.manager.Update()
}

func (g *Game) render() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.Color{R: 147, G: 211, B: 139, A: 255})

	if !g.connected { // if not connected draw menu
		g.drawWaitingMenu()
	} else { // if connected draw game
		g.drawScene()
	}

	rl.EndDrawing()
}

func (g *Game) drawScene() {
	g.manager.Render()
}

func (g *Game) drawWaitingMenu() {
	rl.DrawText("Waiting for connection...", 10, 10, 20, rl.Black)
}
