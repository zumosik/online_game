package saver

import (
	"encoding/gob"
	"os"
	"server/internal/models"
)

type Save struct {
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
