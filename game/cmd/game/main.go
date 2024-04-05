package main

import (
	"game/internal/game"
	"game/internal/lib/logger/sl"
	"game/internal/lib/logger/slogpretty"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	logger := setupLogger("local")

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		logger.Error("Cant init SDL2", sl.Err(err))
		os.Exit(1)
	}
	defer sdl.Quit()

	if err := img.Init(img.INIT_PNG); err != nil {
		logger.Error("Cant init SDL2_image", sl.Err(err))
		os.Exit(1)
	}

	g := game.New(logger, 60)

	err := g.Start(800, 600)
	if err != nil {
		logger.Error("Error starting game", sl.Err(err))
	}
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = setupPrettySlog()
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
