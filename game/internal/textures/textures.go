package textures

import rl "github.com/gen2brain/raylib-go/raylib"

type Textures struct {
	Grass  rl.Texture2D
	Player rl.Texture2D
}

func Load() *Textures {
	var textures Textures

	textures.Grass = rl.LoadTexture("./assets/sprout_lands/Tilesets/Grass.png")
	textures.Player = rl.LoadTexture("./assets/sprout_lands/Characters/Basic Charakter Spritesheet.png")

	return &textures
}

func (textures *Textures) Unload() {
	rl.UnloadTexture(textures.Grass)
	rl.UnloadTexture(textures.Player)
}
