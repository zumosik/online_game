package saver

import (
	"encoding/gob"
	"os"
	"server/internal/models"
)

type DataToSave struct {
	players []models.Player
	// map or smth
}

func (data *DataToSave) WriteToFile(f *os.File) error {
	gob.Register(data)

	if err := gob.NewEncoder(f).Encode(data); err != nil {
		return err
	}

	return nil
}

func (data *DataToSave) ReadFromFile(f *os.File) error {
	gob.Register(data)
	if err := gob.NewDecoder(f).Decode(data); err != nil {
		return err
	}

	return nil
}
