package game

import (
	"errors"
	"log"
	"online_game/internal/game/cm"
	"online_game/internal/game/cm/components"
	"online_game/internal/game/textures"
	"online_game/internal/models"
	"online_game/internal/packets"
	"online_game/internal/tcpclient"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	player  *cm.GameObject
	players []*cm.GameObject
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
		Size:  rl.NewVector2(16, 16),
		Scale: rl.NewVector2(3, 3),
	})
	player.AddComponent(&components.SpriteComponent{
		Tex:   g.tex.Player,
		Color: rl.White,
	})
	player.AddComponent(&components.RigidbodyComponent{Velocity: rl.NewVector2(0, 0), Speed: 5})
	player.AddComponent(&components.PlayerKeyboardComponent{TypeOfInput: components.WASDInput})
	player.Layer = cm.LayerPlayer

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
			if resp.OK {
				g.connected = true
			} else {
				// we will not go further if we cant connect to server
				log.Fatal("Cant connect to server")
			}

			log.Printf("Connected to server, other players: %v", resp.Players)
			for _, p := range resp.Players {
				g.createNewPlayer(p)
			}

			// set player info
			player.AddComponent(&components.PlayerInfoComponent{
				Info:   resp.Player,
				Client: g.cl,
				Token:  resp.Token,
			})

			player.GetComponent(&components.TransformComponent{}).(*components.TransformComponent).Pos =
				rl.NewVector2(resp.Player.Pos.X, resp.Player.Pos.Y) // set player position

			log.Printf("\nPlayers: %v\n", players)

		case packets.TypeOfPacketNewPlayerConnect:
			resp := p.Payload.(packets.NewPlayerConnect)
			log.Printf("new player connect: %v", resp.Player)

			g.createNewPlayer(resp.Player)

			log.Printf("\nPlayers: %v\n", players)
		case packets.TypeOfPacketPlayerPosResp:
			resp := p.Payload.(packets.PlayerPosResp)
			for _, pl := range players {
				transform := pl.GetComponent(&components.TransformComponent{}).(*components.TransformComponent)
				transform.Pos =
					rl.NewVector2(resp.Player.Pos.X, resp.Player.Pos.Y)
			}
		case packets.TypeOfPacketPlayerDisconnect:
			log.Printf("Player disconnected: %v", p.Payload.(packets.PlayerDisconnect).Player)
			// delete player from players
			resp := p.Payload.(packets.PlayerDisconnect)
			for i, pl := range players {
				info := pl.GetComponent(&components.OtherPlayerInfoComponent{}).(*components.OtherPlayerInfoComponent)
				if info.Info.UserID == resp.Player.UserID {
					players = append(players[:i], players[i+1:]...) // deleting player from players

					g.manager.DeleteGameObjectByID(pl.ID)

					break
				}
			}
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

func (g *Game) createNewPlayer(pl models.PublicPlayer) {
	newPlayer := g.manager.CreateGameObject()
	newPlayer.AddComponent(&components.TransformComponent{
		Pos:   rl.NewVector2(pl.Pos.X, pl.Pos.Y),
		Size:  rl.NewVector2(16, 16),
		Scale: rl.NewVector2(3, 3),
	})
	newPlayer.AddComponent(&components.OtherPlayerInfoComponent{Info: pl})
	newPlayer.AddComponent(&components.SpriteComponent{
		Tex:   g.tex.OtherPlayer,
		Color: rl.White,
	})
	newPlayer.Layer = cm.LayerOtherPlayer

	players = append(players, newPlayer)
}
