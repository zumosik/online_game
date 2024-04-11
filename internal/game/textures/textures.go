package textures

import rl "github.com/gen2brain/raylib-go/raylib"

type Textures struct {
	OtherPlayer rl.Texture2D
	Player      rl.Texture2D
}

func Load() *Textures {
	var textures Textures

	textures.OtherPlayer = rl.LoadTexture("./assets/character/vampire_v2_1.png")
	textures.Player = rl.LoadTexture("./assets/character2/skeleton_v2_1.png")

	return &textures
}

func (textures *Textures) Unload() {
	rl.UnloadTexture(textures.OtherPlayer)
	rl.UnloadTexture(textures.Player)
}
