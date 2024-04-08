package main

import (
	"fmt"
	config_game "online_game/internal/config/game"
	"online_game/internal/game"
	"online_game/internal/tcpclient"
)

func main() {
	cfg, err := config_game.ReadConfig()
	if err != nil {
		panic(err) // cant go further without config
	}

	client, err := tcpclient.New(
		fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		tcpclient.User{
			Username: cfg.Username,
			Pin:      cfg.Pin,
		},
	)
	if err != nil {
		panic(err) // cant go further without server connection
	}

	g := game.New(client, 720, 480, 60, "Game")

	g.Init()
	g.Start()
	if err := g.Quit(); err != nil {
		panic(err)
	}
}
