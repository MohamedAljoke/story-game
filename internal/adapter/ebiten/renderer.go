package ebiten

import (
	"story-game/internal/adapter/ebiten/helpers"
	"story-game/internal/adapter/ebiten/sprites"
	"story-game/internal/domain"

	eb "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type EbitenRenderer struct {
	player  *sprites.PlayerSprite
	cat     *sprites.CatSprite
	tileset *Tileset
}

func NewRenderer(tileset *Tileset) *EbitenRenderer {
	return &EbitenRenderer{
		player:  sprites.NewPlayerSprite(),
		cat:     sprites.NewCatSprite(),
		tileset: tileset,
	}
}

func (r *EbitenRenderer) Draw(screen *eb.Image, world *domain.World) {
	screen.Fill(helpers.HexColor(0x1a1a2e))

	if r.tileset != nil && world.TileMap != nil {
		r.tileset.DrawMap(screen, world.TileMap)
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
