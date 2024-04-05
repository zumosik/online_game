package game

import (
	"game/internal/game/texturemanager"
	"game/internal/lib/logger/sl"
	"github.com/veandco/go-sdl2/sdl"
	"log/slog"
)

type Game struct {
	isRunning bool
	l         *slog.Logger
	FPS       uint32

	renderer *sdl.Renderer

	// TMP
	plTex *sdl.Texture // TODO: delete this after GameObject implementation
	c     int
	cc    int
}

func New(logger *slog.Logger, FPS uint32) *Game {
	return &Game{
		l:   logger,
		FPS: FPS,
	}
}

func (g *Game) Start(w, h int32) error {
	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, w, h, sdl.WINDOW_SHOWN)
	if err != nil {
		return err
	}
	defer func(window *sdl.Window) {
		err := window.Destroy()
		if err != nil {
			g.l.Error("cant destroy window", sl.Err(err))
		}
	}(window)

	g.renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return err
	}
	defer func(renderer *sdl.Renderer) {
		err := renderer.Destroy()
		if err != nil {
			g.l.Error("cant destroy renderer", sl.Err(err))
		}
	}(g.renderer)

	tm := texturemanager.New(g.renderer)

	g.plTex, err = tm.LoadTexture("assets/imgs/player.png")
	if err != nil {
		g.l.Error("cant load image", sl.Attr("path", "assets/imgs/player.png"), sl.Err(err))
	}

	defer func(texture *sdl.Texture) {
		err := texture.Destroy()
		if err != nil {
			g.l.Error("cant destory image", sl.Attr("path", "assets/imgs/player.png"), sl.Err(err))
		}
	}(g.plTex)

	g.isRunning = true

	g.cc = 1
	
	for g.isRunning {

		g.handleEvents()
		g.render()

		sdl.Delay(1000 / g.FPS) // wait some time before next frame
	}

	return nil
}

func (g *Game) render() {

	if g.c > 50 {
		g.cc = -1
	}
	if g.c < -50 {
		g.cc = 1
	}

	g.c += g.cc

	if err := g.renderer.Clear(); err != nil {
		g.l.Error("cant clear renderer")
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if err := g.renderer.Copy(g.plTex, &sdl.Rect{
				X: 0,
				Y: 0,
				W: 200,
				H: 231,
			}, &sdl.Rect{
				X: (100 * int32(i)) + int32(g.c),
				Y: (50 * int32(j)) + int32(g.c)*2,
				W: 50,
				H: 50,
			}); err != nil {
				g.l.Error("cant copy texture into renderer")
			}
		}
	}

	g.renderer.Present()

}

func (g *Game) handleEvents() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			g.l.Debug("Quit event...")
			g.isRunning = false
			break
		}
	}

}
