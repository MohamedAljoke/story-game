package ebiten

import (
	"story-game/internal/adapter/ebiten/helpers"
	"story-game/internal/adapter/ebiten/sprites"
	"story-game/internal/domain"

	eb "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type EbitenRenderer struct {
	player *sprites.PlayerSprite
	cat    *sprites.CatSprite
}

func NewRenderer() *EbitenRenderer {
	return &EbitenRenderer{
		player: sprites.NewPlayerSprite(),
		cat:    sprites.NewCatSprite(),
	}
}

func (r *EbitenRenderer) Draw(screen *eb.Image, world *domain.World) {

	// Background
	screen.Fill(helpers.HexColor(0x1a1a2e))

	// Grid
	gridColor := helpers.HexColor(0x2a2a3e)
	for x := 0; x < domain.WorldWidth; x += domain.TileSize {
		vector.StrokeLine(screen, float32(x), 0, float32(x), domain.WorldHeight, 1, gridColor, false)
	}
	for y := 0; y < domain.WorldHeight; y += domain.TileSize {
		vector.StrokeLine(screen, 0, float32(y), domain.WorldWidth, float32(y), 1, gridColor, false)
	}

	// Characters
	world.EachCharacter(func(id domain.CharacterID, char *domain.Character, pos *domain.Position) {
		switch char.Type {

		case domain.CharacterPlayer:
			r.player.Draw(screen, pos, char.Facing)

		case domain.CharacterPet:
			leader := world.Character(char.Leader)
			r.cat.Draw(screen, pos, leader.Facing)

		}
	})

	// HUD
	ebitenutil.DebugPrint(screen, "WASD / Arrow keys to move | ESC to quit")
}
