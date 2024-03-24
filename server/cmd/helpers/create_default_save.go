package main

import (
	"log"
	"os"
	"server/internal/models"
	"server/internal/saver"
)

func main() {
	const pathToFile = "default.111"

	f, err := os.Create(pathToFile)
	if err != nil {
		panic(err)
	}

	s := saver.Save{
		Players: make(map[string]models.Player),
	}

	err = s.WriteToFile(f)
	if err != nil {
		panic(err)
	}

	log.Printf("Created file - %s with data", pathToFile)
}
