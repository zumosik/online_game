package main

import "game/internal/game"

func main() {
	g := game.New(720, 480, 60, "Game")

	g.Init()
	g.Start()
	g.Quit()
}
