package main

import (
	"flag"
	"log"
	"online_game/internal/models"
	"online_game/internal/saver"
	"os"
)

func main() {
	var pathToFile string
	flag.StringVar(&pathToFile, "path-to-file", "default.111", "path to file to save the default save")
	flag.Parse()

	f, err := os.Create(pathToFile)
	if err != nil {
		panic(err)
	}

	s := saver.Save{
		Name:    "default_server",
		Players: make(map[string]models.Player),
	}

	err = s.WriteToFile(f)
	if err != nil {
		panic(err)
	}

	log.Printf("Created file - %s with data", pathToFile)
}
