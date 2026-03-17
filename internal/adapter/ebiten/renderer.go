package ebiten

import (
	"story-game/internal/adapter/ebiten/helpers"
	"story-game/internal/adapter/ebiten/sprites"
	"story-game/internal/domain"

	eb "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type EbitenRenderer struct {
	player *sprites.PlayerSprite
	cat    *sprites.CatSprite
	tiles  *TileRenderer
}

func NewRenderer(tiles *TileRenderer) *EbitenRenderer {
	return &EbitenRenderer{
		player: sprites.NewPlayerSprite(),
		cat:    sprites.NewCatSprite(),
		tiles:  tiles,
	}
}

func (r *EbitenRenderer) Draw(screen *eb.Image, world *domain.World) {
	screen.Fill(helpers.HexColor(0x1a1a2e))

	cam := world.Camera

	if r.tiles != nil && world.TileMap != nil {
		r.tiles.DrawMap(screen, world.TileMap, cam)
	}

	world.EachCharacter(func(id domain.CharacterID, char *domain.Character, pos *domain.Position) {
		screenPos := &domain.Position{X: pos.X - cam.X, Y: pos.Y - cam.Y}

		switch char.Type {
		case domain.CharacterPlayer:
			r.player.Draw(screen, screenPos, char.Facing)
		case domain.CharacterPet:
			leader := world.Character(char.Leader)
			r.cat.Draw(screen, screenPos, leader.Facing)
		}
	})

	ebitenutil.DebugPrint(screen, "WASD / Arrow keys to move | ESC to quit")
}
