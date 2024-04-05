package game

import (
	"game/internal/lib/logger/sl"
	"github.com/veandco/go-sdl2/sdl"
	"log/slog"
)

type Game struct {
	isRunning bool
	l         *slog.Logger
}

func New(logger *slog.Logger) *Game {
	return &Game{
		l: logger,
	}
}

func (g *Game) Start() error {
	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		return err
	}
	defer func(window *sdl.Window) {
		err := window.Destroy()
		if err != nil {
			g.l.Error("cant destroy window", sl.Err(err))
		}
	}(window)

	surface, err := window.GetSurface()
	if err != nil {
		return err
	}
	err = surface.FillRect(nil, 0)
	if err != nil {
		return err
	}

	rect := sdl.Rect{0, 0, 200, 200}
	colour := sdl.Color{R: 255, G: 0, B: 255, A: 255} // purple
	pixel := sdl.MapRGBA(surface.Format, colour.R, colour.G, colour.B, colour.A)
	err = surface.FillRect(&rect, pixel)
	if err != nil {
		return err
	}
	err = window.UpdateSurface()
	if err != nil {
		return err
	}

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				g.l.Debug("Quit event...")
				running = false
				break
			}
		}

		sdl.Delay(33)
	}

	return nil
}
