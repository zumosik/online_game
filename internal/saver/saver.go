package saver

import (
	"encoding/gob"
	"fmt"
	"log/slog"
	"online_game/internal/lib/logger/sl"
	"online_game/internal/models"
	"os"
)

type Save struct {
	Name    string
	Players map[string]models.Player
}

func (data *Save) WriteToFile(f *os.File) error {
	gob.Register(data)

	if err := gob.NewEncoder(f).Encode(data); err != nil {
		return err
	}

	return nil
}

func (data *Save) ReadFromFile(f *os.File) error {
	gob.Register(data)
	if err := gob.NewDecoder(f).Decode(data); err != nil {
		return err
	}

	return nil
}

func (data *Save) Print(l *slog.Logger) {
	var playersStr string
	var i int
	for _, player := range data.Players {
		if i > 10 {
			break
		}

		playersStr += fmt.Sprintf("%d: (id: %d, username: %s, pos: (%3f, %3f))", i, player.UserID, player.Username, player.Pos.X, player.Pos.Y)
		i++
	}
	l.Info("Save data", sl.Attr("name", data.Name), sl.Attr("players", playersStr))
}
